package mtp

import (
	"bytes"
	"compress/gzip"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/cjongseok/slog"
	"math"
	"math/big"
	"reflect"
	"runtime"
	"strings"
	"time"
)

const (
	layer = 71

	// https://core.telegram.org/schema/mtproto
	//crc_vector                     = 0x1cb5c415
	crc_resPQ                      = 0x05162463
	crc_p_q_inner_data             = 0x83c95aec
	crc_server_DH_params_fail      = 0x79cb045d
	crc_server_DH_params_ok        = 0xd0e8075c
	crc_server_DH_inner_data       = 0xb5890dba
	crc_client_DH_inner_data       = 0x6643b654
	crc_dh_gen_ok                  = 0x3bcbf734
	crc_dh_gen_retry               = 0x46dc1fb9
	crc_dh_gen_fail                = 0xa69dae02
	crc_rpc_result                 = 0xf35c6d01
	crc_rpc_error                  = 0x2144ca19
	crc_rpc_answer_unknown         = 0x5e2ad36e
	crc_rpc_answer_dropped_running = 0xcd78e586
	crc_rpc_answer_dropped         = 0xa43ad8b7
	crc_future_salt                = 0x0949d9dc
	crc_future_salts               = 0xae500895
	crc_pong                       = 0x347773c5
	crc_destroy_session_ok         = 0xe22045fc
	crc_destroy_session_none       = 0x62d350c9
	crc_new_session_created        = 0x9ec20908
	crc_msg_container              = 0x73f1f8dc
	crc_msg_copy                   = 0xe06046b2
	crc_gzip_packed                = 0x3072cfa1
	crc_msgs_ack                   = 0x62d6b459
	crc_bad_msg_notification       = 0xa7eff811
	crc_bad_server_salt            = 0xedab447b
	crc_msg_resend_req             = 0x7d861a08
	crc_msgs_state_req             = 0xda69fb52
	crc_msgs_state_info            = 0x04deb57d
	crc_msgs_all_info              = 0x8cc0d131
	crc_msg_detailed_info          = 0x276d3ec6
	crc_msg_new_detailed_info      = 0x809db6df
	crc_req_pq                     = 0x60469778
	crc_req_DH_params              = 0xd712e4be
	crc_set_client_DH_params       = 0xf5045f1f
	crc_rpc_drop_answer            = 0x58e4a740
	crc_get_future_salts           = 0xb921bd04
	crc_ping                       = 0x7abe77ec
	crc_ping_delay_disconnect      = 0xf3427b8c
	crc_destroy_session            = 0xe7512126
	crc_http_wait                  = 0x9299359f
)

type TL interface {
	encode() []byte
}

type RemoteProcedureCall interface {
	InvokeBlocked(msg TL) (interface{}, error)
}

type Predicate interface {
	ToType() TL
}

type RPCaller struct {
	RPC RemoteProcedureCall
}

type EncodeBuf struct {
	buf []byte
}

type DecodeBuf struct {
	buf  []byte
	off  int
	size int
	err  error
}

type TL_msg_container struct {
	Items []TL_MT_message
}

type TL_MT_message struct {
	Msg_id int64
	Seq_no int32
	Size   int32
	Data   interface{}
}

type TL_req_pq struct {
	nonce []byte
}

type TL_p_q_inner_data struct {
	pq           *big.Int
	p            *big.Int
	q            *big.Int
	nonce        []byte
	server_nonce []byte
	new_nonce    []byte
}
type TL_req_DH_params struct {
	nonce        []byte
	server_nonce []byte
	p            *big.Int
	q            *big.Int
	fp           uint64
	encdata      []byte
}
type TL_client_DH_inner_data struct {
	nonce        []byte
	server_nonce []byte
	retry        int64
	g_b          *big.Int
}
type TL_set_client_DH_params struct {
	nonce        []byte
	server_nonce []byte
	encdata      []byte
}
type TL_resPQ struct {
	nonce        []byte
	server_nonce []byte
	pq           *big.Int
	fingerprints []int64
}

type TL_server_DH_params_ok struct {
	nonce            []byte
	server_nonce     []byte
	encrypted_answer []byte
}

type TL_server_DH_inner_data struct {
	nonce        []byte
	server_nonce []byte
	g            int32
	dh_prime     *big.Int
	g_a          *big.Int
	server_time  int32
}

type TL_new_session_created struct {
	first_msg_id int64
	unique_id    int64
	server_salt  []byte
}

type TL_bad_server_salt struct {
	bad_msg_id      int64
	bad_msg_seqno   int32
	error_code      int32
	new_server_salt []byte
}

type TL_crc_bad_msg_notification struct {
	bad_msg_id    int64
	bad_msg_seqno int32
	error_code    int32
}

type TL_msgs_ack struct {
	msgIds []int64
}

type TL_rpc_result struct {
	req_msg_id int64
	Obj        interface{}
}

type TL_rpc_error struct {
	error_code    int32
	error_message string
}

type TL_dh_gen_ok struct {
	nonce           []byte
	server_nonce    []byte
	new_nonce_hash1 []byte
}

type TL_ping struct {
	ping_id int64
}

type TL_pong struct {
	msg_id  int64
	ping_id int64
}

// Encoders
func GenerateNonce(size int) []byte {
	b := make([]byte, size)
	_, _ = rand.Read(b)
	return b
}

func GenerateMessageId() int64 {
	const nano = 1000 * 1000 * 1000
	//FIXME: Windows system clock has time resolution issue. https://github.com/golang/go/issues/17696
	//Remove the sleep when the issue is resolved.
	if strings.Contains(runtime.GOOS, "windows") {
		time.Sleep(2 * time.Millisecond)
	}
	unixnano := time.Now().UnixNano()

	return ((unixnano / nano) << 32) | ((unixnano % nano) & -4)
}

func NewEncodeBuf(cap int) *EncodeBuf {
	if __debug&DEBUG_LEVEL_ENCODE_DETAILS != 0 {
		slog.Logln("Encode::NewBuf::", "cap=", cap)
	}
	return &EncodeBuf{make([]byte, 0, cap)}
}

func (e *EncodeBuf) Int(s int32) {
	e.buf = append(e.buf, 0, 0, 0, 0)
	binary.LittleEndian.PutUint32(e.buf[len(e.buf)-4:], uint32(s))
	if __debug&DEBUG_LEVEL_ENCODE_DETAILS != 0 {
		slog.Logln("Encode::Int::", s)
	}
}

func (e *EncodeBuf) UInt(s uint32) {
	e.buf = append(e.buf, 0, 0, 0, 0)
	binary.LittleEndian.PutUint32(e.buf[len(e.buf)-4:], s)
	if __debug&DEBUG_LEVEL_ENCODE_DETAILS != 0 {
		slog.Logf("Encode::UInt::", "%d(0x%x)", s, s)
	}
}

func (e *EncodeBuf) Long(s int64) {
	e.buf = append(e.buf, 0, 0, 0, 0, 0, 0, 0, 0)
	binary.LittleEndian.PutUint64(e.buf[len(e.buf)-8:], uint64(s))
	if __debug&DEBUG_LEVEL_ENCODE_DETAILS != 0 {
		slog.Logln("Encode::Long::", s)
	}
}

func (e *EncodeBuf) Double(s float64) {
	e.buf = append(e.buf, 0, 0, 0, 0, 0, 0, 0, 0)
	binary.LittleEndian.PutUint64(e.buf[len(e.buf)-8:], math.Float64bits(s))
	if __debug&DEBUG_LEVEL_ENCODE_DETAILS != 0 {
		slog.Logln("Encode::Double::", s)
	}
}

func (e *EncodeBuf) String(s string) {
	e.StringBytes([]byte(s))
	if __debug&DEBUG_LEVEL_ENCODE_DETAILS != 0 {
		slog.Logln("Encode::String::", s)
	}
}

func (e *EncodeBuf) BigInt(s *big.Int) {
	e.StringBytes(s.Bytes())
	if __debug&DEBUG_LEVEL_ENCODE_DETAILS != 0 {
		slog.Logln("Encode::BigInt::", s)
	}
}

func (e *EncodeBuf) StringBytes(s []byte) {
	var res []byte
	size := len(s)
	if size < 254 {
		nl := 1 + size + (4-(size+1)%4)&3
		res = make([]byte, nl)
		res[0] = byte(size)
		copy(res[1:], s)

	} else {
		nl := 4 + size + (4-size%4)&3
		res = make([]byte, nl)
		binary.LittleEndian.PutUint32(res, uint32(size<<8|254))
		copy(res[4:], s)

	}
	e.buf = append(e.buf, res...)
	if __debug&DEBUG_LEVEL_ENCODE_DETAILS != 0 {
		slog.Logln("Encode::StringBytes::", s)
	}
}

func (e *EncodeBuf) Bytes(s []byte) {
	e.buf = append(e.buf, s...)
	if __debug&DEBUG_LEVEL_ENCODE_DETAILS != 0 {
		slog.Logln("Encode::Bytes::", s)
	}
}

func (e *EncodeBuf) VectorInt(v []int32) {
	x := make([]byte, 4+4+len(v)*4)
	binary.LittleEndian.PutUint32(x, crc_vector)
	binary.LittleEndian.PutUint32(x[4:], uint32(len(v)))
	i := 8
	for _, v := range v {
		binary.LittleEndian.PutUint32(x[i:], uint32(v))
		i += 4
	}
	e.buf = append(e.buf, x...)
	if __debug&DEBUG_LEVEL_ENCODE_DETAILS != 0 {
		slog.Logln("Encode::VectorInt::", v)
	}
}

func (e *EncodeBuf) VectorLong(v []int64) {
	x := make([]byte, 4+4+len(v)*8)
	binary.LittleEndian.PutUint32(x, crc_vector)
	binary.LittleEndian.PutUint32(x[4:], uint32(len(v)))
	i := 8
	for _, v := range v {
		binary.LittleEndian.PutUint64(x[i:], uint64(v))
		i += 8
	}
	e.buf = append(e.buf, x...)
	if __debug&DEBUG_LEVEL_ENCODE_DETAILS != 0 {
		slog.Logln("Encode::VectorLong::", v)
	}
}

func (e *EncodeBuf) VectorString(v []string) {
	x := make([]byte, 8)
	binary.LittleEndian.PutUint32(x, crc_vector)
	binary.LittleEndian.PutUint32(x[4:], uint32(len(v)))
	e.buf = append(e.buf, x...)
	for _, v := range v {
		e.String(v)
	}
	if __debug&DEBUG_LEVEL_ENCODE_DETAILS != 0 {
		slog.Logln("Encode::VectorString::", v)
	}
}

func (e *EncodeBuf) Vector(v []TL) {
	x := make([]byte, 8)
	binary.LittleEndian.PutUint32(x, crc_vector)
	binary.LittleEndian.PutUint32(x[4:], uint32(len(v)))
	e.buf = append(e.buf, x...)
	for _, v := range v {
		e.buf = append(e.buf, v.encode()...)
	}
	if __debug&DEBUG_LEVEL_ENCODE_DETAILS != 0 {
		slog.Logln("Encode::Vector::", v)
	}
}

func (e TL_msg_container) encode() []byte            { return nil }
func (e TL_resPQ) encode() []byte                    { return nil }
func (e TL_server_DH_params_ok) encode() []byte      { return nil }
func (e TL_server_DH_inner_data) encode() []byte     { return nil }
func (e TL_dh_gen_ok) encode() []byte                { return nil }
func (e TL_rpc_result) encode() []byte               { return nil }
func (e TL_rpc_error) encode() []byte                { return nil }
func (e TL_new_session_created) encode() []byte      { return nil }
func (e TL_bad_server_salt) encode() []byte          { return nil }
func (e TL_crc_bad_msg_notification) encode() []byte { return nil }

func (e TL_req_pq) encode() []byte {
	x := NewEncodeBuf(20)
	x.UInt(crc_req_pq)
	x.Bytes(e.nonce)
	return x.buf
}

func (e TL_p_q_inner_data) encode() []byte {
	x := NewEncodeBuf(256)
	x.UInt(crc_p_q_inner_data)
	x.BigInt(e.pq)
	x.BigInt(e.p)
	x.BigInt(e.q)
	x.Bytes(e.nonce)
	x.Bytes(e.server_nonce)
	x.Bytes(e.new_nonce)
	return x.buf
}

func (e TL_req_DH_params) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_req_DH_params)
	x.Bytes(e.nonce)
	x.Bytes(e.server_nonce)
	x.BigInt(e.p)
	x.BigInt(e.q)
	x.Long(int64(e.fp))
	x.StringBytes(e.encdata)
	return x.buf
}

func (e TL_client_DH_inner_data) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_client_DH_inner_data)
	x.Bytes(e.nonce)
	x.Bytes(e.server_nonce)
	x.Long(e.retry)
	x.BigInt(e.g_b)
	return x.buf
}

func (e TL_set_client_DH_params) encode() []byte {
	x := NewEncodeBuf(256)
	x.UInt(crc_set_client_DH_params)
	x.Bytes(e.nonce)
	x.Bytes(e.server_nonce)
	x.StringBytes(e.encdata)
	return x.buf
}

func (e TL_ping) encode() []byte {
	x := NewEncodeBuf(32)
	x.UInt(crc_ping)
	x.Long(e.ping_id)
	return x.buf
}

func (e TL_pong) encode() []byte {
	x := NewEncodeBuf(32)
	x.UInt(crc_pong)
	x.Long(e.msg_id)
	x.Long(e.ping_id)
	return x.buf
}

func (e TL_msgs_ack) encode() []byte {
	x := NewEncodeBuf(64)
	x.UInt(crc_msgs_ack)
	x.VectorLong(e.msgIds)
	return x.buf
}

func (e *EncodeBuf) FlaggedLong(flags, f int32, s int64) {
	bit := int32(1 << uint(f))
	if flags&bit == 0 {
		return
	}
	e.Long(s)
}
func (e *EncodeBuf) FlaggedDouble(flags, f int32, s float64) {
	bit := int32(1 << uint(f))
	if flags&bit == 0 {
		return
	}
	e.Double(s)
}
func (e *EncodeBuf) FlaggedInt(flags, f int32, s int32) {
	bit := int32(1 << uint(f))
	if flags&bit == 0 {
		return
	}
	e.Int(s)
}
func (e *EncodeBuf) FlaggedString(flags, f int32, s string) {
	bit := int32(1 << uint(f))
	if flags&bit == 0 {
		return
	}
	e.String(s)
}
func (e *EncodeBuf) FlaggedVector(flags, f int32, v []TL) {
	bit := int32(1 << uint(f))
	if flags&bit == 0 {
		return
	}
	e.Vector(v)
}
func (e *EncodeBuf) FlaggedObject(flags, f int32, o TL) {
	bit := int32(1 << uint(f))
	if flags&bit == 0 {
		return
	}
	e.Bytes(o.encode())
}
func (e *EncodeBuf) FlaggedStringBytes(flags, f int32, s []byte) {
	bit := int32(1 << uint(f))
	if flags&bit == 0 {
		return
	}
	e.StringBytes(s)
}
func (e *EncodeBuf) FlaggedVectorInt(flags, f int32, v []int32) {
	bit := int32(1 << uint(f))
	if flags&bit == 0 {
		return
	}
	e.VectorInt(v)
}
func (e *EncodeBuf) FlaggedVectorLong(flags, f int32, v []int64) {
	bit := int32(1 << uint(f))
	if flags&bit == 0 {
		return
	}
	e.VectorLong(v)
}
func (e *EncodeBuf) FlaggedVectorString(flags, f int32, v []string) {
	bit := int32(1 << uint(f))
	if flags&bit == 0 {
		return
	}
	e.VectorString(v)
}

func toTLslice(slice interface{}) []TL {
	if reflect.TypeOf(slice).Kind() != reflect.Slice {
		return nil
	}
	s := reflect.ValueOf(slice)
	if s.Len() < 1 {
		return nil
	}
	switch s.Index(0).Interface().(type) {
	case TL:
		tlslice := make([]TL, s.Len())
		for i := 0; i < s.Len(); i++ {
			tlslice[i] = s.Index(i).Interface().(TL)
		}
		return tlslice
	default:
		return nil
	}
}

// Decoders
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
