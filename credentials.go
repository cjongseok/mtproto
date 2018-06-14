package mtproto

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"os"
)

type Credentials struct {
	Phone       string
	ApiID       int32
	ApiHash     string
	IP          string
	Port        int
	Salt        []byte
	AuthKey     []byte
	AuthKeyHash []byte
}

type credentialsJSON struct {
	Phone   string `json:"phone"`
	ApiID   int32  `json:"api_id"`
	ApiHash string `json:"api_hash"`
	IP      string `json:"ip"`
	Port    int    `json:"port"`
	Salt    string `json:"salt"`
	AuthKey string `json:"auth_key"`
}

// SaveToFile stores the Credentials as a JSON file in the format of credentialsJson.
// It creates the file, if not exists, and overwrite the file, if exists.

// JSON turns the given credentials into JSON binary in the format of credentialsJSON
func (c *Credentials) JSON() ([]byte, error) {
	return json.Marshal(credentialsJSON{
		c.Phone,
		c.ApiID,
		c.ApiHash,
		c.IP,
		c.Port,
		//string(c.Salt),
		base64.StdEncoding.EncodeToString(c.Salt),
		base64.StdEncoding.EncodeToString(c.AuthKey),
	})
}
func NewCredentials(jsonInBytes []byte) (c *Credentials, err error) {
	unmarshaled := &credentialsJSON{}
	err = json.Unmarshal(jsonInBytes, unmarshaled)
	if err != nil {
		return
	}
	var salt, authKey, authKeyHash []byte
	salt, err = base64.StdEncoding.DecodeString(unmarshaled.Salt)
	if err != nil {
		return
	}
	authKey, err = base64.StdEncoding.DecodeString(unmarshaled.AuthKey)
	if err != nil {
		return
	}
	authKeyHash = sha1(authKey)[12:20]
	c = &Credentials{
		unmarshaled.Phone,
		unmarshaled.ApiID,
		unmarshaled.ApiHash,
		unmarshaled.IP,
		unmarshaled.Port,
		//[]byte(unmarshaled.Salt),
		salt,
		authKey,
		authKeyHash,
	}
	return
}

// Save session
//TODO: save channel and datacenter information
func (c *Credentials) Save(f *os.File) (err error) {
	/*session.encrypted = true
	b := NewEncodeBuf(1024)
	b.StringBytes(session.authKey)
	b.StringBytes(session.authKeyHash)
	b.StringBytes(session.serverSalt)
	b.String(session.addr)
	var useIPv6UInt uint32
	if session.useIPv6 {
		useIPv6UInt = 1
	}
	b.UInt(useIPv6UInt)
	*/

	var jsonBytes []byte
	jsonBytes, err = c.JSON()
	if err != nil {
		return err
	}

	err = f.Truncate(0)
	if err != nil {
		return err
	}

	_, err = f.WriteAt(jsonBytes, 0)
	if err != nil {
		return err
	}

	return nil
}

func NewCredentialsFromFile(f *os.File) (*Credentials, error) {
	// Decode session file
	b := make([]byte, 1024*4)
	n, err := f.ReadAt(b, 0)
	if n <= 0 || (err != nil && err.Error() != "EOF") {
		return nil, errors.New("nothing in the file")
	}
	return NewCredentials(b[:n])

	//d := NewDecodeBuf(b)
	//session.authKey = d.StringBytes()
	//session.authKeyHash = d.StringBytes()
	//session.serverSalt = d.StringBytes()
	//session.addr = d.String()
	//session.useIPv6 = false
	//if d.UInt() == 1 {
	//	session.useIPv6 = true
	//}
	//
	//if d.err != nil {
	//	// Failed to load session
	//	return d.err
	//}
	//
	//session.encrypted = true
	//return nil
}
