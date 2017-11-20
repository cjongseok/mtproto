package mtproto

import (
	"math/big"
	"fmt"
)

type TL interface {
	encode() []byte
}

const crc_vector = 0x1cb5c415 // Processed manually

const crc_msg_container = 0x73f1f8dc

type TL_msg_container struct {
	Items []TL_MT_message
}

type TL_MT_message struct {
	Msg_id int64
	Seq_no int32
	Size   int32
	Data   interface{}
}

const crc_req_pq = 0x60469778

type TL_req_pq struct {
	Nonce []byte
}

const crc_p_q_inner_data = 0x83c95aec

type TL_p_q_inner_data struct {
	Pq           *big.Int
	P            *big.Int
	Q            *big.Int
	Nonce        []byte
	Server_nonce []byte
	New_nonce    []byte
}

const crc_req_DH_params = 0xd712e4be

type TL_req_DH_params struct {
	Nonce        []byte
	Server_nonce []byte
	P            *big.Int
	Q            *big.Int
	Fp           uint64
	Encdata      []byte
}

const crc_client_DH_inner_data = 0x6643b654

type TL_client_DH_inner_data struct {
	Nonce        []byte
	Server_nonce []byte
	Retry        int64
	G_b          *big.Int
}

const crc_set_client_DH_params = 0xf5045f1f

type TL_set_client_DH_params struct {
	Nonce        []byte
	Server_nonce []byte
	Encdata      []byte
}

const crc_resPQ = 0x05162463

type TL_resPQ struct {
	Nonce        []byte
	Server_nonce []byte
	Pq           *big.Int
	Fingerprints []int64
}

const crc_server_DH_params_ok = 0xd0e8075c

type TL_server_DH_params_ok struct {
	Nonce            []byte
	Server_nonce     []byte
	Encrypted_answer []byte
}

const crc_server_DH_params_fail = 0x79cb045d

type TL_server_DH_params_fail struct {
	Nonce          []byte
	Server_nonce   []byte
	New_nonce_hash []byte
}

const crc_server_DH_inner_data = 0xb5890dba

type TL_server_DH_inner_data struct {
	Nonce        []byte
	Server_nonce []byte
	G            int32
	Dh_prime     *big.Int
	G_a          *big.Int
	Server_time  int32
}

const crc_new_session_created = 0x9ec20908

type TL_new_session_created struct {
	First_msg_id int64
	Unique_id    int64
	Server_salt  []byte
}

const crc_bad_server_salt = 0xedab447b

type TL_bad_server_salt struct {
	Bad_msg_id      int64
	Bad_msg_seqno   int32
	Error_code      int32
	New_server_salt []byte
}

const crc_bad_msg_notification = 0xa7eff811

type TL_bad_msg_notification struct {
	Bad_msg_id    int64
	Bad_msg_seqno int32
	Error_code    int32
}

const crc_msgs_ack = 0x62d6b459

type TL_msgs_ack struct {
	MsgIds []int64
}

const crc_rpc_result = 0xf35c6d01

type TL_rpc_result struct {
	Req_msg_id int64
	Obj        interface{}
}

const crc_rpc_error = 0x2144ca19

type TL_rpc_error struct {
	Error_code    int32
	Error_message string
}

const crc_dh_gen_ok = 0x3bcbf734

type TL_dh_gen_ok struct {
	Nonce           []byte
	Server_nonce    []byte
	New_nonce_hash1 []byte
}

const crc_ping = 0x7abe77ec

type TL_ping struct {
	Ping_id int64
}

const crc_pong = 0x347773c5

type TL_pong struct {
	Msg_id  int64
	Ping_id int64
}

const crc_gzip_packed = 0x3072cfa1

// Implements interface error
func (e TL_rpc_error) Error() string {
	switch e.Error_code {
	case errorSeeOther:
		return fmt.Sprintf("mtproto RPC error: %d %s", e.Error_code, e.Error_message)
	case errorBadRequest, errorUnauthorized, errorFlood, errorInternal:
		return fmt.Sprintf("mtproto RPC error: %d %s", e.Error_code, e.Error_message)
	default:
		return fmt.Sprintf("mtproto unknow RPC error: %d %s", e.Error_code, e.Error_message)
	}
}