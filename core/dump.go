package core

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/cjongseok/slog"
	"io"
	"os"
	"strings"
	"sync"
)

type dumpCallback struct {
	out chan interface{}
}

func (cb dumpCallback) OnUpdate(u Update) {
	cb.out <- u
}

type Dump struct {
	updatesState    *PredUpdatesState
	updateCallback  dumpCallback
	readWaitGroup   sync.WaitGroup
	readInterrupter chan interface{}
	authKey         []byte
	authKeyHash     []byte
	serverSalt      []byte
	reader          io.Reader
}

func NewDump(authFileName, dumpFilename string, out chan interface{}) (*Dump, error) {
	mdump := new(Dump)
	authf, err := os.OpenFile(authFileName, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	dumpf, err := os.OpenFile(dumpFilename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	mdump.updatesState = &PredUpdatesState{}
	mdump.updateCallback = dumpCallback{out}
	mdump.readWaitGroup = sync.WaitGroup{}
	mdump.readInterrupter = make(chan interface{})

	mdump.readSessionFile(authf)
	reader, err := slog.DumpReader(dumpf)
	if err != nil {
		return nil, err
	}
	mdump.reader = reader
	return mdump, nil
}

func (md *Dump) Play() {
	md.readWaitGroup.Add(1)
	go md.readRoutine()
}

func (d *Dump) Wait() {
	d.readWaitGroup.Wait()
}

func (md *Dump) read() (interface{}, error) {
	var err error
	var n int
	var size int
	var data interface{}

	// Read packet size
	b := make([]byte, 1)
	n, err = md.reader.Read(b) // Wait for an incoming byte
	if err != nil {
		return nil, err
	}
	slog.Record(b)

	if b[0] < 127 {
		size = int(b[0]) << 2
	} else {
		b := make([]byte, 3)
		n, err = md.reader.Read(b)
		slog.Record(b)
		if err != nil {
			return nil, err
		}
		size = (int(b[0]) | int(b[1])<<8 | int(b[2])<<16) << 2
	}

	// Read packet
	left := size
	buf := make([]byte, size)
	for left > 0 {
		n, err = md.reader.Read(buf[size-left:])
		if err != nil {
			return nil, err
		}
		left -= n
	}
	slog.Record(buf)

	if size == 4 {
		return nil, fmt.Errorf("Server response error: %d", int32(binary.LittleEndian.Uint32(buf)))
	}

	// decrypt incoming packet
	data, _, _, err = decryptMtproto(buf, md.authKey)
	if err != nil {
		return nil, err
	}
	return data, nil

}

func (md *Dump) process( /*msgId int64, seqNo int32, */ data interface{}) interface{} {
	switch data := data.(type) {
	case TL_msg_container:
		//data := data.(TL_msg_container).Items
		for _, v := range data.Items {
			md.process(v.Data)
		}

	case TL_bad_server_salt:
	case TL_new_session_created:
	case TL_ping:
	case TL_pong:
	case TL_msgs_ack:
	case TL_rpc_result:
		//data := data.(TL_rpc_result)
		md.process(data.Obj)
	case *PredUpdatesState:
		//data := data.(*PredUpdatesState)
		md.updatesState.Pts = data.Pts
		md.updatesState.Qts = data.Qts
		md.updatesState.Date = data.Date
		md.updatesState.Seq = data.Seq
		//marshaled, err := json.Marshal(data)
		//if err == nil {
		//slog.Logf(md.process, "updatesState: %s\n", marshaled)
		//} else {
		//slog.Logf(md.process, "updatesState: %v\n", data)
		//}
		return data

		// Date updates
	case *PredUpdates:
		//data := data.(*PredUpdates)
		md.updatesState.Date = data.Date
		md.updatesState.Seq = data.Seq
		md.updateCallback.OnUpdate(data)
		return data
	case *PredUpdateShort:
		//data := data.(*PredUpdateShort)
		md.updatesState.Date = data.Date
		md.updateCallback.OnUpdate(data)
		return data

		// Pts updates
	case *PredUpdateNewMessage:
		//data := data.(*PredUpdateNewMessage)
		md.updatesState.Pts = data.Pts
		md.updateCallback.OnUpdate(data)
		return data
	case *PredUpdateReadMessagesContents:
		//data := data.(TL_updateReadMessagesContents)
		md.updatesState.Pts = data.Pts
		md.updateCallback.OnUpdate(data)
		return data
	case *PredUpdateDeleteMessages:
		//data := data.(TL_updateDeleteMessages)
		md.updatesState.Pts = data.Pts
		md.updateCallback.OnUpdate(data)
		return data

		// Pts and Date updates
	case *PredUpdateShortMessage:
		//data := data.(TL_updateShortMessage)
		md.updatesState.Pts = data.Pts
		md.updatesState.Date = data.Date
		md.updateCallback.OnUpdate(data)
		return data
	case *PredUpdateShortChatMessage:
		//data := data.(TL_updateShortChatMessage)
		md.updatesState.Pts = data.Pts
		md.updatesState.Date = data.Date
		md.updateCallback.OnUpdate(data)
		return data
	case *PredUpdateShortSentMessage:
		//data := data.(TL_updateShortSentMessage)
		md.updatesState.Pts = data.Pts
		md.updatesState.Date = data.Date
		md.updateCallback.OnUpdate(data)
		return data

		// Qts updates
	case *PredUpdateNewEncryptedMessage:
		//data := data.(TL_updateNewEncryptedMessage)
		md.updatesState.Qts = data.Qts
		md.updateCallback.OnUpdate(data)
		return data

		// Channel updates
	case *PredUpdateChannel:
		//data := data.(TL_updateChannel)
		md.updateCallback.OnUpdate(data)
		return data
	case *PredUpdateChannelMessageViews:
		//data := data.(TL_updateChannelMessageViews)
		md.updateCallback.OnUpdate(data)
		return data
	case *PredUpdateChannelTooLong:
		//data := data.(TL_updateChannelTooLong)
		md.updatesState.Pts = data.Pts
		md.updateCallback.OnUpdate(data)
		return data
	case *PredUpdateReadChannelInbox:
		//data := data.(TL_updateReadChannelInbox)
		md.updateCallback.OnUpdate(data)
		return data
	case *PredUpdateReadChannelOutbox:
		//data := data.(TL_updateReadChannelOutbox)
		md.updateCallback.OnUpdate(data)
		return data
	case *PredUpdateNewChannelMessage:
		//data := data.(TL_updateNewChannelMessage)
		md.updatesState.Pts = data.Pts
		md.updateCallback.OnUpdate(data)
		return data

	default:
		//marshaled, err := json.Marshal(data)
		//if err == nil {
		//slog.Logf(md.process, "process: unknown data type %T {%s}\n", data, marshaled)
		//} else {
		//slog.Logf(md.process, "process: unknown data type %T {%v}\n", data, data)
		//}
		return data
	}

	return nil
}

func (md *Dump) readRoutine() {
	//slog.Logln(md.readRoutine, "read: start")
	defer md.readWaitGroup.Done()

	innerRoutineWG := sync.WaitGroup{}

	ch := make(chan interface{}, 1)
	innerRoutineWG.Add(1)
	go func(ch chan<- interface{}) {
		defer innerRoutineWG.Done()

		for {
			data, err := md.read()
			//slog.Logf(md.readRoutine, "read: type: %v, data: %v, err: %v\n", reflect.TypeOf(data), data, err)
			if err == io.EOF {
				// Connection closed by server, trying to reconnect
				//slog.Logln(md.readRoutine, "read: lost connection (captured EOF). reconnect")
				//TODO: figure out the end of dump file, and stop the loop
				close(ch)
				return
			} else if err != nil {
				if strings.Contains(err.Error(), "use of closed network connection") {
					//slog.Logf(md.readRoutine, "read: TCP connection closed (%s)\n", err)
				} else if strings.Contains(err.Error(), "connection reset by peer") {
					//slog.Logf(md.readRoutine, "read: lost connection (%s). reconnect\n", err)
				} else {
					//slog.Logf(md.readRoutine, "read: unknown error, %s. reconnect\n", err)
				}
			} else {
				ch <- data
			}
		}
	}(ch)

	for {
		// Run async wait for data from server

		select {
		case <-md.readInterrupter:
			//slog.Logln(md.readRoutine, "read: wait for inner routine ...")
			innerRoutineWG.Wait()
			close(ch)
			//slog.Logln(md.readRoutine, "read: stop")
			return
		case data := <-ch:
			if data == nil {
				//slog.Logln(md.readRoutine, "data is nil")
				md.updateCallback.OnUpdate(nil)
				return
			}
			md.process(data)
		}
	}
}

// Implements interface error
func (md *Dump) readSessionFile(f *os.File) error {
	// Decode session file
	b := make([]byte, 1024*4)
	n, err := f.ReadAt(b, 0)
	if n <= 0 || (err != nil && err.Error() != "EOF") {
		return errors.New("New session")
	}

	d := NewDecodeBuf(b)
	md.authKey = d.StringBytes()
	md.authKeyHash = d.StringBytes()
	md.serverSalt = d.StringBytes()
	return nil
}
