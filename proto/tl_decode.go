package mtp

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/cjongseok/slog"
	"math"
	"math/big"
)

type DecodeBuf struct {
	buf  []byte
	off  int
	size int
	err  error
}

func NewDecodeBuf(b []byte) *DecodeBuf {
	if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
		slog.Logln("Decode::NewBuf::", "bytes = ", b)
	}
	return &DecodeBuf{b, 0, len(b), nil}
}

func (m *DecodeBuf) Long() int64 {
	if m.err != nil {
		return 0
	}
	if m.off+8 > m.size {
		m.err = errors.New("DecodeLong")
		return 0
	}
	x := int64(binary.LittleEndian.Uint64(m.buf[m.off : m.off+8]))
	m.off += 8
	if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
		slog.Logln("Decode::Long::", x)
	}
	return x
}

func (m *DecodeBuf) Double() float64 {
	if m.err != nil {
		return 0
	}
	if m.off+8 > m.size {
		m.err = errors.New("DecodeDouble")
		return 0
	}
	x := math.Float64frombits(binary.LittleEndian.Uint64(m.buf[m.off : m.off+8]))
	m.off += 8
	if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
		slog.Logln("Decode::Double::", x)
	}
	return x
}

func (m *DecodeBuf) Int() int32 {
	if m.err != nil {
		return 0
	}
	if m.off+4 > m.size {
		m.err = errors.New("DecodeInt")
		return 0
	}
	x := binary.LittleEndian.Uint32(m.buf[m.off : m.off+4])
	m.off += 4
	if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
		slog.Logln("Decode::Int::", x)
	}
	return int32(x)
}

func (m *DecodeBuf) UInt() uint32 {
	if m.err != nil {
		return 0
	}
	if m.off+4 > m.size {
		m.err = errors.New("DecodeUInt")
		return 0
	}
	x := binary.LittleEndian.Uint32(m.buf[m.off : m.off+4])
	m.off += 4
	if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
		slog.Logf("Decode::UInt::", "%d(0x%x)", x, x)
	}
	return x
}

func (m *DecodeBuf) Bytes(size int) []byte {
	if m.err != nil {
		return nil
	}
	if m.off+size > m.size {
		m.err = errors.New("DecodeBytes")
		return nil
	}
	x := make([]byte, size)
	copy(x, m.buf[m.off:m.off+size])
	m.off += size
	if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
		if len(x) > 10 {
			slog.Logln("Decode::Bytes::", len(x), x[:10], " ...")
		} else {
			slog.Logln("Decode::Bytes::", len(x), x)
		}

	}
	return x
}

func (m *DecodeBuf) StringBytes() []byte {
	if m.err != nil {
		return nil
	}
	var size, padding int

	if m.off+1 > m.size {
		m.err = errors.New("DecodeStringBytes")
		return nil
	}
	size = int(m.buf[m.off])
	m.off++
	padding = (4 - ((size + 1) % 4)) & 3
	if size == 254 {
		if m.off+3 > m.size {
			m.err = errors.New("DecodeStringBytes")
			return nil
		}
		size = int(m.buf[m.off]) | int(m.buf[m.off+1])<<8 | int(m.buf[m.off+2])<<16
		m.off += 3
		padding = (4 - size%4) & 3
	}

	if m.off+size > m.size {
		m.err = errors.New("DecodeStringBytes: Wrong size")
		return nil
	}
	x := make([]byte, size)
	copy(x, m.buf[m.off:m.off+size])
	m.off += size

	if m.off+padding > m.size {
		m.err = errors.New("DecodeStringBytes: Wrong padding")
		return nil
	}
	m.off += padding
	if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
		if len(x) > 10 {
			slog.Logln("Decode::StringBytes::", len(x), x[:10], " ...")
		} else {
			slog.Logln("Decode::StringBytes::", len(x), x)
		}

	}
	return x
}

func (m *DecodeBuf) String() string {
	b := m.StringBytes()
	if m.err != nil {
		return ""
	}
	x := string(b)
	if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
		slog.Logln("Decode::String::", x)
	}
	return x
}

func (m *DecodeBuf) BigInt() *big.Int {
	b := m.StringBytes()
	if m.err != nil {
		return nil
	}
	y := make([]byte, len(b)+1)
	y[0] = 0
	copy(y[1:], b)
	x := new(big.Int).SetBytes(y)
	if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
		slog.Logln("Decode::BigInt::", x)
	}
	return x
}

func (m *DecodeBuf) VectorInt() []int32 {
	constructor := m.UInt()
	if m.err != nil {
		return nil
	}
	if constructor != crc_vector {
		m.err = fmt.Errorf("DecodeVectorInt: Wrong constructor (0x%08x)", constructor)
		return nil
	}
	size := m.Int()
	if m.err != nil {
		return nil
	}
	if size < 0 {
		m.err = errors.New("DecodeVectorInt: Wrong size")
		return nil
	}
	x := make([]int32, size)
	i := int32(0)
	for i < size {
		y := m.Int()
		if m.err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
		slog.Logln("Decode::VectorInt::", x)
	}
	return x
}

func (m *DecodeBuf) VectorLong() []int64 {
	constructor := m.UInt()
	if m.err != nil {
		return nil
	}
	if constructor != crc_vector {
		m.err = fmt.Errorf("DecodeVectorLong: Wrong constructor (0x%08x)", constructor)
		return nil
	}
	size := m.Int()
	if m.err != nil {
		return nil
	}
	if size < 0 {
		m.err = errors.New("DecodeVectorLong: Wrong size")
		return nil
	}
	x := make([]int64, size)
	i := int32(0)
	for i < size {
		y := m.Long()
		if m.err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
		slog.Logln("Decode::VectorLong::", x)
	}
	return x
}

func (m *DecodeBuf) VectorString() []string {
	constructor := m.UInt()
	if m.err != nil {
		return nil
	}
	if constructor != crc_vector {
		m.err = fmt.Errorf("DecodeVectorString: Wrong constructor (0x%08x)", constructor)
		return nil
	}
	size := m.Int()
	if m.err != nil {
		return nil
	}
	if size < 0 {
		m.err = errors.New("DecodeVectorString: Wrong size")
		return nil
	}
	x := make([]string, size)
	i := int32(0)
	for i < size {
		y := m.String()
		if m.err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
		slog.Logln("Decode::VectorString::", x)
	}
	return x
}

func (m *DecodeBuf) Bool() bool {
	constructor := m.UInt()
	if m.err != nil {
		return false
	}
	switch constructor {
	case crc_boolFalse:
		if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
			slog.Logln("Decode::Bool::", false)
		}
		return false
	case crc_boolTrue:
		if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
			slog.Logln("Decode::Bool::", true)
		}
		return true
	}
	return false
}

func (m *DecodeBuf) TL_Vector() []TL {
	//constructor := m.UInt()
	//if m.err != nil {
	//	return nil
	//}
	//if constructor != crc_vector {
	//	m.err = fmt.Errorf("DecodeTL_Vector: Wrong constructor (0x%08x)", constructor)
	//	return nil
	//}
	//slog.Logln("byte len=", len(m.buf))
	//slog.Logln("m.size=", m.size)
	//slog.Logln(hex.EncodeToString(m.buf))
	//xx := m.UInt()
	//slog.Logln("x=", xx)
	//size := m.Int()
	//slog.Logln("size=", size)
	//if m.err != nil {
	//	return nil
	//}
	//if size < 0 {
	//	m.err = errors.New("DecodeTL_Vector: Wrong size")
	//	return nil
	//}
	//x := make([]TL, size)
	//i := int32(0)
	//for i < size {
	//	y := m.Object()
	//	if m.err != nil {
	//		return nil
	//	}
	//	x[i] = y
	//	i++
	//}
	//if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
	//	slog.Logln("Decode::Vector::", x)
	//}
	//return x
	m.err = fmt.Errorf("DecodeTL_Vector: NOT SUPPORTED YET")
	return nil
}

func (m *DecodeBuf) Vector() []TL {
	constructor := m.UInt()
	if m.err != nil {
		return nil
	}
	if constructor != crc_vector {
		m.err = fmt.Errorf("DecodeVector: Wrong constructor (0x%08x)", constructor)
		return nil
	}
	size := m.Int()
	if m.err != nil {
		return nil
	}
	if size < 0 {
		m.err = errors.New("DecodeVector: Wrong size")
		return nil
	}
	x := make([]TL, size)
	i := int32(0)
	for i < size {
		y := m.Object()
		if m.err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
		slog.Logln("Decode::Vector::", x)
	}
	return x
}

func (m *DecodeBuf) Object() (r TL) {
	constructor := m.UInt()
	if m.err != nil {
		return nil
	}
	switch constructor {

	case crc_resPQ:
		if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
			slog.Logln("reqPQ", constructor)
		}
		r = TL_resPQ{m.Bytes(16), m.Bytes(16), m.BigInt(), m.VectorLong()}

	case crc_server_DH_params_ok:
		if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
			slog.Logln("server_DH_params_ok", constructor)
		}
		r = TL_server_DH_params_ok{m.Bytes(16), m.Bytes(16), m.StringBytes()}

	case crc_server_DH_inner_data:
		if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
			slog.Logln("server_DH_inner_data", constructor)
		}
		r = TL_server_DH_inner_data{
			m.Bytes(16), m.Bytes(16), m.Int(),
			m.BigInt(), m.BigInt(), m.Int(),
		}

	case crc_dh_gen_ok:
		if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
			slog.Logln("dh_gen_ok", constructor)
		}
		r = TL_dh_gen_ok{m.Bytes(16), m.Bytes(16), m.Bytes(16)}

	case crc_ping:
		if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
			slog.Logln("ping", constructor)
		}
		r = TL_ping{m.Long()}

	case crc_pong:
		if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
			slog.Logln("pong", constructor)
		}
		r = TL_pong{m.Long(), m.Long()}

	case crc_msg_container:
		if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
			slog.Logln("msg_container", constructor)
		}
		size := m.Int()
		arr := make([]TL_MT_message, size)
		for i := int32(0); i < size; i++ {
			arr[i] = TL_MT_message{m.Long(), m.Int(), m.Int(), m.Object()}
			//slog.Logln(constructor, arr[i])
			if m.err != nil {
				slog.Logln(m.err.Error())
				return nil
			}
		}
		r = TL_msg_container{arr}

	case crc_rpc_result:
		if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
			slog.Logln("rpc_result", constructor)
		}
		r = TL_rpc_result{m.Long(), m.Object()}

	case crc_rpc_error:
		if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
			slog.Logln("rpc_error", constructor)
		}
		r = TL_rpc_error{m.Int(), m.String()}

	case crc_new_session_created:
		if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
			slog.Logln("new_session_created", constructor)
		}
		r = TL_new_session_created{m.Long(), m.Long(), m.Bytes(8)}

	case crc_bad_server_salt:
		if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
			slog.Logln("bad_server_salt", constructor)
		}
		r = TL_bad_server_salt{m.Long(), m.Int(), m.Int(), m.Bytes(8)}

	case crc_bad_msg_notification:
		if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
			slog.Logln("bad_msg_notification", constructor)
		}
		r = TL_crc_bad_msg_notification{m.Long(), m.Int(), m.Int()}

	case crc_msgs_ack:
		if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
			slog.Logln("msgs_ack", constructor)
		}
		r = TL_msgs_ack{m.VectorLong()}

	case crc_gzip_packed:
		if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
			slog.Logln("gzip_packed", constructor)
		}
		obj := make([]byte, 0, 4096)

		var buf bytes.Buffer
		_, _ = buf.Write(m.StringBytes())
		gz, _ := gzip.NewReader(&buf)
		b := make([]byte, 4096)
		for true {
			n, _ := gz.Read(b)
			obj = append(obj, b[0:n]...)
			if n <= 0 {
				break
			}
		}
		d := NewDecodeBuf(obj)
		r = d.Object()

	default:
		if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
			slog.Logln(fmt.Sprintf("default %x", constructor))
		}
		r = m.ObjectGenerated(constructor)

	}

	if m.err != nil {
		slog.Logln(m.err.Error())
		return nil
	}
	return
}

func (m *DecodeBuf) Flags() int32 {
	if m.err != nil {
		return 0
	}
	if m.off+4 > m.size {
		m.err = errors.New("DecodeInt")
		return 0
	}
	x := binary.LittleEndian.Uint32(m.buf[m.off : m.off+4])
	m.off += 4
	if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
		slog.Logln("Decode::Flags::", x)
	}
	return int32(x)
}

func (m *DecodeBuf) FlaggedLong(flags, f int32) int64 {
	bit := int32(1 << uint(f))
	if flags&bit == 0 {
		return 0
	}
	return m.Long()
}

func (m *DecodeBuf) FlaggedDouble(flags, f int32) float64 {
	bit := int32(1 << uint(f))
	if flags&bit == 0 {
		return 0
	}
	return m.Double()
}

func (m *DecodeBuf) FlaggedInt(flags, f int32) int32 {
	if m.err != nil {
		return 0
	}
	bit := int32(1 << uint(f))
	if flags&bit == 0 {
		return 0
	}
	return m.Int()
}

func (m *DecodeBuf) FlaggedString(flags, f int32) string {
	bit := int32(1 << uint(f))
	if flags&bit == 0 {
		return ""
	}
	return m.String()
}

func (m *DecodeBuf) FlaggedVector(flags, f int32) []TL {
	bit := int32(1 << uint(f))
	if flags&bit == 0 {
		return nil
	}
	return m.Vector()
}

func (m *DecodeBuf) FlaggedObject(flags, f int32) (r TL) {
	bit := int32(1 << uint(f))
	if flags&bit == 0 {
		return nil
	}
	return m.Object()
}

func (m *DecodeBuf) FlaggedStringBytes(flags, f int32) []byte {
	bit := int32(1 << uint(f))
	if flags&bit == 0 {
		return nil
	}
	return m.StringBytes()
}

func (d *DecodeBuf) dump() {
	slog.Logln(hex.Dump(d.buf[d.off:d.size]))
}

func toBool(x TL) bool {
	_, ok := x.(*PredBoolTrue)
	return ok
}
