package mtproto

import (
	"github.com/cjongseok/slog"
	"fmt"
	"encoding/binary"
	"io"
	//"encoding/json"
	"sync"
	"strings"
	"errors"
	"os"
)

type dumpCallback struct {
	out chan interface{}
}
func (cb dumpCallback) OnUpdate(u MUpdate) {
	cb.out <- u
}

type MDump struct {
	updatesState    *TL_updates_state
	updateCallback  dumpCallback
	readWaitGroup   sync.WaitGroup
	readInterrupter chan interface{}
	authKey     []byte
	authKeyHash []byte
	serverSalt  []byte
	reader      io.Reader
}

func NewMdump(authFileName, dumpFilename string, out chan interface{}) (*MDump, error) {
	mdump := new(MDump)
	authf, err := os.OpenFile(authFileName, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	dumpf, err := os.OpenFile(dumpFilename, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	mdump.updatesState = new(TL_updates_state)
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

func (md *MDump) Play() {
	md.readWaitGroup.Add(1)
	go md.readRoutine()
}

func (d *MDump) Wait() {
	d.readWaitGroup.Wait()
}

func (md *MDump) read() (interface{}, error) {
	var err error
	var n int
	var size int
	var data interface{}

	// Read packet size
	b := make([]byte, 1)
	n, err = md.reader.Read(b)	// Wait for an incoming byte
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

	// Deserialize incoming packet
	data, _, _, err = Deserialize(buf, md.authKey)
	if err != nil {
		return nil, err
	}
	return data, nil

}

func (md *MDump) process(/*msgId int64, seqNo int32, */data interface{}) interface{} {
	switch data.(type) {
	case TL_msg_container:
		data := data.(TL_msg_container).Items
		for _, v := range data {
			md.process(v.Data)
		}

	case TL_bad_server_salt:
	case TL_new_session_created:
	case TL_ping:
	case TL_pong:
	case TL_msgs_ack:
	case TL_rpc_result:
		data := data.(TL_rpc_result)
		md.process(data.Obj)
	case TL_updates_state:
		data := data.(TL_updates_state)
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
	case TL_updates:
		data := data.(TL_updates)
		md.updatesState.Date = data.Date
		md.updatesState.Seq = data.Seq
		md.updateCallback.OnUpdate(data)
		return data
	case TL_updateShort:
		data := data.(TL_updateShort)
		md.updatesState.Date = data.Date
		md.updateCallback.OnUpdate(data)
		return data

		// Pts updates
	case TL_updateNewMessage:
		data := data.(TL_updateNewMessage)
		md.updatesState.Pts = data.Pts
		md.updateCallback.OnUpdate(data)
		return data
	case TL_updateReadMessagesContents:
		data := data.(TL_updateReadMessagesContents)
		md.updatesState.Pts = data.Pts
		md.updateCallback.OnUpdate(data)
		return data
	case TL_updateDeleteMessages:
		data := data.(TL_updateDeleteMessages)
		md.updatesState.Pts = data.Pts
		md.updateCallback.OnUpdate(data)
		return data

		// Pts and Date updates
	case TL_updateShortMessage:
		data := data.(TL_updateShortMessage)
		md.updatesState.Pts = data.Pts
		md.updatesState.Date = data.Date
		md.updateCallback.OnUpdate(data)
		return data
	case TL_updateShortChatMessage:
		data := data.(TL_updateShortChatMessage)
		md.updatesState.Pts = data.Pts
		md.updatesState.Date = data.Date
		md.updateCallback.OnUpdate(data)
		return data
	case TL_updateShortSentMessage:
		data := data.(TL_updateShortSentMessage)
		md.updatesState.Pts = data.Pts
		md.updatesState.Date = data.Date
		md.updateCallback.OnUpdate(data)
		return data

		// Qts updates
	case TL_updateNewEncryptedMessage:
		data := data.(TL_updateNewEncryptedMessage)
		md.updatesState.Qts = data.Qts
		md.updateCallback.OnUpdate(data)
		return data

		// Channel updates
	case TL_updateChannel:
		data := data.(TL_updateChannel)
		md.updateCallback.OnUpdate(data)
		return data
	case TL_updateChannelMessageViews:
		data := data.(TL_updateChannelMessageViews)
		md.updateCallback.OnUpdate(data)
		return data
	case TL_updateChannelTooLong:
		data := data.(TL_updateChannelTooLong)
		md.updatesState.Pts = data.Pts
		md.updateCallback.OnUpdate(data)
		return data
	case TL_updateReadChannelInbox:
		data := data.(TL_updateReadChannelInbox)
		md.updateCallback.OnUpdate(data)
		return data
	case TL_updateReadChannelOutbox:
		data := data.(TL_updateReadChannelOutbox)
		md.updateCallback.OnUpdate(data)
		return data
	case TL_updateNewChannelMessage:
		data := data.(TL_updateNewChannelMessage)
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

func (md *MDump) readRoutine() {
	//slog.Logln(md.readRoutine, "read: start")
	defer md.readWaitGroup.Done()

	innerRoutineWG := sync.WaitGroup{}

	for {
		// Run async wait for data from server
		ch := make(chan interface{}, 1)
		go func(ch chan<- interface{}) {
			innerRoutineWG.Add(1)
			defer innerRoutineWG.Done()

			data, err := md.read()
			//slog.Logf(md.readRoutine, "read: type: %v, data: %v, err: %v\n", reflect.TypeOf(data), data, err)
			if err == io.EOF {
				// Connection closed by server, trying to reconnect
				//slog.Logln(md.readRoutine, "read: lost connection (captured EOF). reconnect")
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
		}(ch)

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
				return
			}
			md.process(data)
		}
	}
}

// Implements interface error
func (md *MDump) readSessionFile(f *os.File) error {
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
