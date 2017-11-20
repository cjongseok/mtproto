package mtproto

import "fmt"

// boolFalse#bc799737 = Bool;

const crc_boolFalse = 0xbc799737

type TL_boolFalse struct {
}

// boolTrue#997275b5 = Bool;

const crc_boolTrue = 0x997275b5

type TL_boolTrue struct {
}

// true#3fedd339 = True;

const crc_true = 0x3fedd339

type TL_true struct {
}

// Encoding TL_true
func (e TL_true) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_true)
	return x.buf
}

// error#c4b9f9bb code:int text:string = Error;

const crc_error = 0xc4b9f9bb

type TL_error struct {
	Code int32  // code:int
	Text string // text:string
}

// Encoding TL_error
func (e TL_error) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_error)
	x.Int(e.Code)
	x.String(e.Text)
	return x.buf
}

// null#56730bcc = Null;

const crc_null = 0x56730bcc

type TL_null struct {
}

// inputPeerEmpty#7f3b18ea = InputPeer;

const crc_inputPeerEmpty = 0x7f3b18ea

type TL_inputPeerEmpty struct {
}

// Encoding TL_inputPeerEmpty
func (e TL_inputPeerEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerEmpty)
	return x.buf
}

// inputPeerSelf#7da07ec9 = InputPeer;

const crc_inputPeerSelf = 0x7da07ec9

type TL_inputPeerSelf struct {
}

// Encoding TL_inputPeerSelf
func (e TL_inputPeerSelf) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerSelf)
	return x.buf
}

// inputPeerChat#179be863 chat_id:int = InputPeer;

const crc_inputPeerChat = 0x179be863

type TL_inputPeerChat struct {
	Chat_id int32 // chat_id:int
}

// Encoding TL_inputPeerChat
func (e TL_inputPeerChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerChat)
	x.Int(e.Chat_id)
	return x.buf
}

// inputPeerUser#7b8e7de6 user_id:int access_hash:long = InputPeer;

const crc_inputPeerUser = 0x7b8e7de6

type TL_inputPeerUser struct {
	User_id     int32 // user_id:int
	Access_hash int64 // access_hash:long
}

// Encoding TL_inputPeerUser
func (e TL_inputPeerUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerUser)
	x.Int(e.User_id)
	x.Long(e.Access_hash)
	return x.buf
}

// inputPeerChannel#20adaef8 channel_id:int access_hash:long = InputPeer;

const crc_inputPeerChannel = 0x20adaef8

type TL_inputPeerChannel struct {
	Channel_id  int32 // channel_id:int
	Access_hash int64 // access_hash:long
}

// Encoding TL_inputPeerChannel
func (e TL_inputPeerChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerChannel)
	x.Int(e.Channel_id)
	x.Long(e.Access_hash)
	return x.buf
}

// inputUserEmpty#b98886cf = InputUser;

const crc_inputUserEmpty = 0xb98886cf

type TL_inputUserEmpty struct {
}

// Encoding TL_inputUserEmpty
func (e TL_inputUserEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputUserEmpty)
	return x.buf
}

// inputUserSelf#f7c1b13f = InputUser;

const crc_inputUserSelf = 0xf7c1b13f

type TL_inputUserSelf struct {
}

// Encoding TL_inputUserSelf
func (e TL_inputUserSelf) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputUserSelf)
	return x.buf
}

// inputUser#d8292816 user_id:int access_hash:long = InputUser;

const crc_inputUser = 0xd8292816

type TL_inputUser struct {
	User_id     int32 // user_id:int
	Access_hash int64 // access_hash:long
}

// Encoding TL_inputUser
func (e TL_inputUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputUser)
	x.Int(e.User_id)
	x.Long(e.Access_hash)
	return x.buf
}

// inputPhoneContact#f392b7f4 client_id:long phone:string first_name:string last_name:string = InputContact;

const crc_inputPhoneContact = 0xf392b7f4

type TL_inputPhoneContact struct {
	Client_id  int64  // client_id:long
	Phone      string // phone:string
	First_name string // first_name:string
	Last_name  string // last_name:string
}

// Encoding TL_inputPhoneContact
func (e TL_inputPhoneContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPhoneContact)
	x.Long(e.Client_id)
	x.String(e.Phone)
	x.String(e.First_name)
	x.String(e.Last_name)
	return x.buf
}

// inputFile#f52ff27f Id:long parts:int name:string md5_checksum:string = InputFile;

const crc_inputFile = 0xf52ff27f

type TL_inputFile struct {
	Id           int64  // Id:long
	Parts        int32  // parts:int
	Name         string // name:string
	Md5_checksum string // md5_checksum:string
}

// Encoding TL_inputFile
func (e TL_inputFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputFile)
	x.Long(e.Id)
	x.Int(e.Parts)
	x.String(e.Name)
	x.String(e.Md5_checksum)
	return x.buf
}

// inputFileBig#fa4f0bb5 Id:long parts:int name:string = InputFile;

const crc_inputFileBig = 0xfa4f0bb5

type TL_inputFileBig struct {
	Id    int64  // Id:long
	Parts int32  // parts:int
	Name  string // name:string
}

// Encoding TL_inputFileBig
func (e TL_inputFileBig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputFileBig)
	x.Long(e.Id)
	x.Int(e.Parts)
	x.String(e.Name)
	return x.buf
}

// inputMediaEmpty#9664f57f = InputMedia;

const crc_inputMediaEmpty = 0x9664f57f

type TL_inputMediaEmpty struct {
}

// Encoding TL_inputMediaEmpty
func (e TL_inputMediaEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaEmpty)
	return x.buf
}

// inputMediaUploadedPhoto#630c9af1 flags:# file:InputFile caption:string stickers:flags.0?Vector<InputDocument> = InputMedia;

const crc_inputMediaUploadedPhoto = 0x630c9af1

type TL_inputMediaUploadedPhoto struct {
	Flags    int32
	File     TL     // file:InputFile
	Caption  string // caption:string
	Stickers []TL   // stickers:flags.0?Vector<InputDocument>
}

// Encoding TL_inputMediaUploadedPhoto
func (e TL_inputMediaUploadedPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaUploadedPhoto)
	var flags int32
	if len(e.Stickers) != 0 {
		flags |= (1 << 0)
	}
	x.Int(flags)
	x.Bytes(e.File.encode())
	x.String(e.Caption)
	if flags&(1<<0) != 0 {
		x.Vector(e.Stickers)
	}
	return x.buf
}

// inputMediaPhoto#e9bfb4f3 Id:InputPhoto caption:string = InputMedia;

const crc_inputMediaPhoto = 0xe9bfb4f3

type TL_inputMediaPhoto struct {
	Id      TL     // Id:InputPhoto
	Caption string // caption:string
}

// Encoding TL_inputMediaPhoto
func (e TL_inputMediaPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaPhoto)
	x.Bytes(e.Id.encode())
	x.String(e.Caption)
	return x.buf
}

// inputMediaGeoPoint#f9c44144 geo_point:InputGeoPoint = InputMedia;

const crc_inputMediaGeoPoint = 0xf9c44144

type TL_inputMediaGeoPoint struct {
	Geo_point TL // geo_point:InputGeoPoint
}

// Encoding TL_inputMediaGeoPoint
func (e TL_inputMediaGeoPoint) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaGeoPoint)
	x.Bytes(e.Geo_point.encode())
	return x.buf
}

// inputMediaContact#a6e45987 phone_number:string first_name:string last_name:string = InputMedia;

const crc_inputMediaContact = 0xa6e45987

type TL_inputMediaContact struct {
	Phone_number string // phone_number:string
	First_name   string // first_name:string
	Last_name    string // last_name:string
}

// Encoding TL_inputMediaContact
func (e TL_inputMediaContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaContact)
	x.String(e.Phone_number)
	x.String(e.First_name)
	x.String(e.Last_name)
	return x.buf
}

// inputMediaUploadedDocument#d070f1e9 flags:# file:InputFile mime_type:string attributes:Vector<DocumentAttribute> caption:string stickers:flags.0?Vector<InputDocument> = InputMedia;

const crc_inputMediaUploadedDocument = 0xd070f1e9

type TL_inputMediaUploadedDocument struct {
	Flags      int32
	File       TL     // file:InputFile
	Mime_type  string // mime_type:string
	Attributes []TL   // attributes:Vector<DocumentAttribute>
	Caption    string // caption:string
	Stickers   []TL   // stickers:flags.0?Vector<InputDocument>
}

// Encoding TL_inputMediaUploadedDocument
func (e TL_inputMediaUploadedDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaUploadedDocument)
	var flags int32
	if len(e.Stickers) != 0 {
		flags |= (1 << 0)
	}
	x.Int(flags)
	x.Bytes(e.File.encode())
	x.String(e.Mime_type)
	x.Vector(e.Attributes)
	x.String(e.Caption)
	if flags&(1<<0) != 0 {
		x.Vector(e.Stickers)
	}
	return x.buf
}

// inputMediaUploadedThumbDocument#50d88cae flags:# file:InputFile thumb:InputFile mime_type:string attributes:Vector<DocumentAttribute> caption:string stickers:flags.0?Vector<InputDocument> = InputMedia;

const crc_inputMediaUploadedThumbDocument = 0x50d88cae

type TL_inputMediaUploadedThumbDocument struct {
	Flags      int32
	File       TL     // file:InputFile
	Thumb      TL     // thumb:InputFile
	Mime_type  string // mime_type:string
	Attributes []TL   // attributes:Vector<DocumentAttribute>
	Caption    string // caption:string
	Stickers   []TL   // stickers:flags.0?Vector<InputDocument>
}

// Encoding TL_inputMediaUploadedThumbDocument
func (e TL_inputMediaUploadedThumbDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaUploadedThumbDocument)
	var flags int32
	if len(e.Stickers) != 0 {
		flags |= (1 << 0)
	}
	x.Int(flags)
	x.Bytes(e.File.encode())
	x.Bytes(e.Thumb.encode())
	x.String(e.Mime_type)
	x.Vector(e.Attributes)
	x.String(e.Caption)
	if flags&(1<<0) != 0 {
		x.Vector(e.Stickers)
	}
	return x.buf
}

// inputMediaDocument#1a77f29c Id:InputDocument caption:string = InputMedia;

const crc_inputMediaDocument = 0x1a77f29c

type TL_inputMediaDocument struct {
	Id      TL     // Id:InputDocument
	Caption string // caption:string
}

// Encoding TL_inputMediaDocument
func (e TL_inputMediaDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaDocument)
	x.Bytes(e.Id.encode())
	x.String(e.Caption)
	return x.buf
}

// inputMediaVenue#2827a81a geo_point:InputGeoPoint title:string address:string provider:string venue_id:string = InputMedia;

const crc_inputMediaVenue = 0x2827a81a

type TL_inputMediaVenue struct {
	Geo_point TL     // geo_point:InputGeoPoint
	Title     string // title:string
	Address   string // address:string
	Provider  string // provider:string
	Venue_id  string // venue_id:string
}

// Encoding TL_inputMediaVenue
func (e TL_inputMediaVenue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaVenue)
	x.Bytes(e.Geo_point.encode())
	x.String(e.Title)
	x.String(e.Address)
	x.String(e.Provider)
	x.String(e.Venue_id)
	return x.buf
}

// inputMediaGifExternal#4843b0fd url:string q:string = InputMedia;

const crc_inputMediaGifExternal = 0x4843b0fd

type TL_inputMediaGifExternal struct {
	Url string // url:string
	Q   string // q:string
}

// Encoding TL_inputMediaGifExternal
func (e TL_inputMediaGifExternal) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaGifExternal)
	x.String(e.Url)
	x.String(e.Q)
	return x.buf
}

// inputMediaPhotoExternal#b55f4f18 url:string caption:string = InputMedia;

const crc_inputMediaPhotoExternal = 0xb55f4f18

type TL_inputMediaPhotoExternal struct {
	Url     string // url:string
	Caption string // caption:string
}

// Encoding TL_inputMediaPhotoExternal
func (e TL_inputMediaPhotoExternal) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaPhotoExternal)
	x.String(e.Url)
	x.String(e.Caption)
	return x.buf
}

// inputMediaDocumentExternal#e5e9607c url:string caption:string = InputMedia;

const crc_inputMediaDocumentExternal = 0xe5e9607c

type TL_inputMediaDocumentExternal struct {
	Url     string // url:string
	Caption string // caption:string
}

// Encoding TL_inputMediaDocumentExternal
func (e TL_inputMediaDocumentExternal) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaDocumentExternal)
	x.String(e.Url)
	x.String(e.Caption)
	return x.buf
}

// inputMediaGame#d33f43f3 Id:InputGame = InputMedia;

const crc_inputMediaGame = 0xd33f43f3

type TL_inputMediaGame struct {
	Id TL // Id:InputGame
}

// Encoding TL_inputMediaGame
func (e TL_inputMediaGame) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaGame)
	x.Bytes(e.Id.encode())
	return x.buf
}

// inputMediaInvoice#92153685 flags:# title:string description:string photo:flags.0?InputWebDocument invoice:Invoice payload:bytes provider:string start_param:string = InputMedia;

const crc_inputMediaInvoice = 0x92153685

type TL_inputMediaInvoice struct {
	Flags       int32
	Title       string // title:string
	Description string // description:string
	Photo       TL     // photo:flags.0?InputWebDocument
	Invoice     TL     // invoice:Invoice
	Payload     []byte // payload:bytes
	Provider    string // provider:string
	Start_param string // start_param:string
}

// Encoding TL_inputMediaInvoice
func (e TL_inputMediaInvoice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaInvoice)
	var flags int32
	if _, ok := (e.Photo).(TL_null); !ok {
		flags |= (1 << 0)
	}
	x.Int(flags)
	x.String(e.Title)
	x.String(e.Description)
	if flags&(1<<0) != 0 {
		x.Bytes(e.Photo.encode())
	}
	x.Bytes(e.Invoice.encode())
	x.Bytes(e.Payload)
	x.String(e.Provider)
	x.String(e.Start_param)
	return x.buf
}

// inputChatPhotoEmpty#1ca48f57 = InputChatPhoto;

const crc_inputChatPhotoEmpty = 0x1ca48f57

type TL_inputChatPhotoEmpty struct {
}

// Encoding TL_inputChatPhotoEmpty
func (e TL_inputChatPhotoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputChatPhotoEmpty)
	return x.buf
}

// inputChatUploadedPhoto#927c55b4 file:InputFile = InputChatPhoto;

const crc_inputChatUploadedPhoto = 0x927c55b4

type TL_inputChatUploadedPhoto struct {
	File TL // file:InputFile
}

// Encoding TL_inputChatUploadedPhoto
func (e TL_inputChatUploadedPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputChatUploadedPhoto)
	x.Bytes(e.File.encode())
	return x.buf
}

// inputChatPhoto#8953ad37 Id:InputPhoto = InputChatPhoto;

const crc_inputChatPhoto = 0x8953ad37

type TL_inputChatPhoto struct {
	Id TL // Id:InputPhoto
}

// Encoding TL_inputChatPhoto
func (e TL_inputChatPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputChatPhoto)
	x.Bytes(e.Id.encode())
	return x.buf
}

// inputGeoPointEmpty#e4c123d6 = InputGeoPoint;

const crc_inputGeoPointEmpty = 0xe4c123d6

type TL_inputGeoPointEmpty struct {
}

// Encoding TL_inputGeoPointEmpty
func (e TL_inputGeoPointEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputGeoPointEmpty)
	return x.buf
}

// inputGeoPoint#f3b7acc9 lat:double long:double = InputGeoPoint;

const crc_inputGeoPoint = 0xf3b7acc9

type TL_inputGeoPoint struct {
	Lat  float64 // lat:double
	Long float64 // long:double
}

// Encoding TL_inputGeoPoint
func (e TL_inputGeoPoint) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputGeoPoint)
	x.Double(e.Lat)
	x.Double(e.Long)
	return x.buf
}

// inputPhotoEmpty#1cd7bf0d = InputPhoto;

const crc_inputPhotoEmpty = 0x1cd7bf0d

type TL_inputPhotoEmpty struct {
}

// Encoding TL_inputPhotoEmpty
func (e TL_inputPhotoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPhotoEmpty)
	return x.buf
}

// inputPhoto#fb95c6c4 Id:long access_hash:long = InputPhoto;

const crc_inputPhoto = 0xfb95c6c4

type TL_inputPhoto struct {
	Id          int64 // Id:long
	Access_hash int64 // access_hash:long
}

// Encoding TL_inputPhoto
func (e TL_inputPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPhoto)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	return x.buf
}

// inputFileLocation#14637196 volume_id:long local_id:int secret:long = InputFileLocation;

const crc_inputFileLocation = 0x14637196

type TL_inputFileLocation struct {
	Volume_id int64 // volume_id:long
	Local_id  int32 // local_id:int
	Secret    int64 // secret:long
}

// Encoding TL_inputFileLocation
func (e TL_inputFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputFileLocation)
	x.Long(e.Volume_id)
	x.Int(e.Local_id)
	x.Long(e.Secret)
	return x.buf
}

// inputEncryptedFileLocation#f5235d55 Id:long access_hash:long = InputFileLocation;

const crc_inputEncryptedFileLocation = 0xf5235d55

type TL_inputEncryptedFileLocation struct {
	Id          int64 // Id:long
	Access_hash int64 // access_hash:long
}

// Encoding TL_inputEncryptedFileLocation
func (e TL_inputEncryptedFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputEncryptedFileLocation)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	return x.buf
}

// inputDocumentFileLocation#430f0724 Id:long access_hash:long Version:int = InputFileLocation;

const crc_inputDocumentFileLocation = 0x430f0724

type TL_inputDocumentFileLocation struct {
	Id          int64 // Id:long
	Access_hash int64 // access_hash:long
	Version     int32 // Version:int
}

// Encoding TL_inputDocumentFileLocation
func (e TL_inputDocumentFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputDocumentFileLocation)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	x.Int(e.Version)
	return x.buf
}

// inputAppEvent#770656a8 time:double type:string peer:long data:string = InputAppEvent;

const crc_inputAppEvent = 0x770656a8

type TL_inputAppEvent struct {
	Time      float64 // time:double
	Code_type string  // type:string
	Peer      int64   // peer:long
	Data      string  // data:string
}

// Encoding TL_inputAppEvent
func (e TL_inputAppEvent) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputAppEvent)
	x.Double(e.Time)
	x.String(e.Code_type)
	x.Long(e.Peer)
	x.String(e.Data)
	return x.buf
}

// peerUser#9db1bc6d user_id:int = Peer;

const crc_peerUser = 0x9db1bc6d

type TL_peerUser struct {
	User_id int32 // user_id:int
}

// Encoding TL_peerUser
func (e TL_peerUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerUser)
	x.Int(e.User_id)
	return x.buf
}

// peerChat#bad0e5bb chat_id:int = Peer;

const crc_peerChat = 0xbad0e5bb

type TL_peerChat struct {
	Chat_id int32 // chat_id:int
}

// Encoding TL_peerChat
func (e TL_peerChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerChat)
	x.Int(e.Chat_id)
	return x.buf
}

// peerChannel#bddde532 channel_id:int = Peer;

const crc_peerChannel = 0xbddde532

type TL_peerChannel struct {
	Channel_id int32 // channel_id:int
}

// Encoding TL_peerChannel
func (e TL_peerChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerChannel)
	x.Int(e.Channel_id)
	return x.buf
}

// storage.fileUnknown#aa963b05 = storage.FileType;

const crc_storage_fileUnknown = 0xaa963b05

type TL_storage_fileUnknown struct {
}

// Encoding TL_storage_fileUnknown
func (e TL_storage_fileUnknown) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_fileUnknown)
	return x.buf
}

// storage.filePartial#40bc6f52 = storage.FileType;

const crc_storage_filePartial = 0x40bc6f52

type TL_storage_filePartial struct {
}

// Encoding TL_storage_filePartial
func (e TL_storage_filePartial) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_filePartial)
	return x.buf
}

// storage.fileJpeg#7efe0e = storage.FileType;

const crc_storage_fileJpeg = 0x7efe0e

type TL_storage_fileJpeg struct {
}

// Encoding TL_storage_fileJpeg
func (e TL_storage_fileJpeg) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_fileJpeg)
	return x.buf
}

// storage.fileGif#cae1aadf = storage.FileType;

const crc_storage_fileGif = 0xcae1aadf

type TL_storage_fileGif struct {
}

// Encoding TL_storage_fileGif
func (e TL_storage_fileGif) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_fileGif)
	return x.buf
}

// storage.filePng#a4f63c0 = storage.FileType;

const crc_storage_filePng = 0xa4f63c0

type TL_storage_filePng struct {
}

// Encoding TL_storage_filePng
func (e TL_storage_filePng) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_filePng)
	return x.buf
}

// storage.filePdf#ae1e508d = storage.FileType;

const crc_storage_filePdf = 0xae1e508d

type TL_storage_filePdf struct {
}

// Encoding TL_storage_filePdf
func (e TL_storage_filePdf) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_filePdf)
	return x.buf
}

// storage.fileMp3#528a0677 = storage.FileType;

const crc_storage_fileMp3 = 0x528a0677

type TL_storage_fileMp3 struct {
}

// Encoding TL_storage_fileMp3
func (e TL_storage_fileMp3) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_fileMp3)
	return x.buf
}

// storage.fileMov#4b09ebbc = storage.FileType;

const crc_storage_fileMov = 0x4b09ebbc

type TL_storage_fileMov struct {
}

// Encoding TL_storage_fileMov
func (e TL_storage_fileMov) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_fileMov)
	return x.buf
}

// storage.fileMp4#b3cea0e4 = storage.FileType;

const crc_storage_fileMp4 = 0xb3cea0e4

type TL_storage_fileMp4 struct {
}

// Encoding TL_storage_fileMp4
func (e TL_storage_fileMp4) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_fileMp4)
	return x.buf
}

// storage.fileWebp#1081464c = storage.FileType;

const crc_storage_fileWebp = 0x1081464c

type TL_storage_fileWebp struct {
}

// Encoding TL_storage_fileWebp
func (e TL_storage_fileWebp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_fileWebp)
	return x.buf
}

// fileLocationUnavailable#7c596b46 volume_id:long local_id:int secret:long = FileLocation;

const crc_fileLocationUnavailable = 0x7c596b46

type TL_fileLocationUnavailable struct {
	Volume_id int64 // volume_id:long
	Local_id  int32 // local_id:int
	Secret    int64 // secret:long
}

// Encoding TL_fileLocationUnavailable
func (e TL_fileLocationUnavailable) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_fileLocationUnavailable)
	x.Long(e.Volume_id)
	x.Int(e.Local_id)
	x.Long(e.Secret)
	return x.buf
}

// fileLocation#53d69076 dc_id:int volume_id:long local_id:int secret:long = FileLocation;

const crc_fileLocation = 0x53d69076

type TL_fileLocation struct {
	Dc_id     int32 // dc_id:int
	Volume_id int64 // volume_id:long
	Local_id  int32 // local_id:int
	Secret    int64 // secret:long
}

// Encoding TL_fileLocation
func (e TL_fileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_fileLocation)
	x.Int(e.Dc_id)
	x.Long(e.Volume_id)
	x.Int(e.Local_id)
	x.Long(e.Secret)
	return x.buf
}

// userEmpty#200250ba Id:int = User;

const crc_userEmpty = 0x200250ba

type TL_userEmpty struct {
	Id int32 // Id:int
}

// Encoding TL_userEmpty
func (e TL_userEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userEmpty)
	x.Int(e.Id)
	return x.buf
}

// user#d10d979a flags:# self:flags.10?true contact:flags.11?true mutual_contact:flags.12?true deleted:flags.13?true bot:flags.14?true bot_chat_history:flags.15?true bot_nochats:flags.16?true verified:flags.17?true restricted:flags.18?true min:flags.20?true bot_inline_geo:flags.21?true Id:int access_hash:flags.0?long first_name:flags.1?string last_name:flags.2?string username:flags.3?string phone:flags.4?string photo:flags.5?UserProfilePhoto status:flags.6?UserStatus bot_info_version:flags.14?int restriction_reason:flags.18?string bot_inline_placeholder:flags.19?string = User;

const crc_user = 0xd10d979a

type TL_user struct {
	Flags                  int32
	Self                   bool   // self:flags.10?true
	Contact                bool   // contact:flags.11?true
	Mutual_contact         bool   // mutual_contact:flags.12?true
	Deleted                bool   // deleted:flags.13?true
	Bot                    bool   // bot:flags.14?true
	Bot_chat_history       bool   // bot_chat_history:flags.15?true
	Bot_nochats            bool   // bot_nochats:flags.16?true
	Verified               bool   // verified:flags.17?true
	Restricted             bool   // restricted:flags.18?true
	Min                    bool   // min:flags.20?true
	Bot_inline_geo         bool   // bot_inline_geo:flags.21?true
	Id                     int32  // Id:int
	Access_hash            int64  // access_hash:flags.0?long
	First_name             string // first_name:flags.1?string
	Last_name              string // last_name:flags.2?string
	Username               string // username:flags.3?string
	Phone                  string // phone:flags.4?string
	Photo                  TL     // photo:flags.5?UserProfilePhoto
	Status                 TL     // status:flags.6?UserStatus
	Bot_info_version       int32  // bot_info_version:flags.14?int
	Restriction_reason     string // restriction_reason:flags.18?string
	Bot_inline_placeholder string // bot_inline_placeholder:flags.19?string
}

// Encoding TL_user
func (e TL_user) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_user)
	var flags int32
	if e.Self {
		flags |= (1 << 10)
	}
	if e.Contact {
		flags |= (1 << 11)
	}
	if e.Mutual_contact {
		flags |= (1 << 12)
	}
	if e.Deleted {
		flags |= (1 << 13)
	}
	if e.Bot {
		flags |= (1 << 14)
	}
	if e.Bot_chat_history {
		flags |= (1 << 15)
	}
	if e.Bot_nochats {
		flags |= (1 << 16)
	}
	if e.Verified {
		flags |= (1 << 17)
	}
	if e.Restricted {
		flags |= (1 << 18)
	}
	if e.Min {
		flags |= (1 << 20)
	}
	if e.Bot_inline_geo {
		flags |= (1 << 21)
	}
	if e.Access_hash > 0 {
		flags |= (1 << 0)
	}
	if e.First_name != "" {
		flags |= (1 << 1)
	}
	if e.Last_name != "" {
		flags |= (1 << 2)
	}
	if e.Username != "" {
		flags |= (1 << 3)
	}
	if e.Phone != "" {
		flags |= (1 << 4)
	}
	if _, ok := (e.Photo).(TL_null); !ok {
		flags |= (1 << 5)
	}
	if _, ok := (e.Status).(TL_null); !ok {
		flags |= (1 << 6)
	}
	if e.Bot_info_version > 0 {
		flags |= (1 << 14)
	}
	if e.Restriction_reason != "" {
		flags |= (1 << 18)
	}
	if e.Bot_inline_placeholder != "" {
		flags |= (1 << 19)
	}
	x.Int(flags)
	x.Int(e.Id)
	if flags&(1<<0) != 0 {
		x.Long(e.Access_hash)
	}
	if flags&(1<<1) != 0 {
		x.String(e.First_name)
	}
	if flags&(1<<2) != 0 {
		x.String(e.Last_name)
	}
	if flags&(1<<3) != 0 {
		x.String(e.Username)
	}
	if flags&(1<<4) != 0 {
		x.String(e.Phone)
	}
	if flags&(1<<5) != 0 {
		x.Bytes(e.Photo.encode())
	}
	if flags&(1<<6) != 0 {
		x.Bytes(e.Status.encode())
	}
	if flags&(1<<14) != 0 {
		x.Int(e.Bot_info_version)
	}
	if flags&(1<<18) != 0 {
		x.String(e.Restriction_reason)
	}
	if flags&(1<<19) != 0 {
		x.String(e.Bot_inline_placeholder)
	}
	return x.buf
}

// userProfilePhotoEmpty#4f11bae1 = UserProfilePhoto;

const crc_userProfilePhotoEmpty = 0x4f11bae1

type TL_userProfilePhotoEmpty struct {
}

// Encoding TL_userProfilePhotoEmpty
func (e TL_userProfilePhotoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userProfilePhotoEmpty)
	return x.buf
}

// userProfilePhoto#d559d8c8 photo_id:long photo_small:FileLocation photo_big:FileLocation = UserProfilePhoto;

const crc_userProfilePhoto = 0xd559d8c8

type TL_userProfilePhoto struct {
	Photo_id    int64 // photo_id:long
	Photo_small TL    // photo_small:FileLocation
	Photo_big   TL    // photo_big:FileLocation
}

// Encoding TL_userProfilePhoto
func (e TL_userProfilePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userProfilePhoto)
	x.Long(e.Photo_id)
	x.Bytes(e.Photo_small.encode())
	x.Bytes(e.Photo_big.encode())
	return x.buf
}

// userStatusEmpty#9d05049 = UserStatus;

const crc_userStatusEmpty = 0x9d05049

type TL_userStatusEmpty struct {
}

// Encoding TL_userStatusEmpty
func (e TL_userStatusEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userStatusEmpty)
	return x.buf
}

// userStatusOnline#edb93949 expires:int = UserStatus;

const crc_userStatusOnline = 0xedb93949

type TL_userStatusOnline struct {
	Expires int32 // expires:int
}

// Encoding TL_userStatusOnline
func (e TL_userStatusOnline) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userStatusOnline)
	x.Int(e.Expires)
	return x.buf
}

// userStatusOffline#8c703f was_online:int = UserStatus;

const crc_userStatusOffline = 0x8c703f

type TL_userStatusOffline struct {
	Was_online int32 // was_online:int
}

// Encoding TL_userStatusOffline
func (e TL_userStatusOffline) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userStatusOffline)
	x.Int(e.Was_online)
	return x.buf
}

// userStatusRecently#e26f42f1 = UserStatus;

const crc_userStatusRecently = 0xe26f42f1

type TL_userStatusRecently struct {
}

// Encoding TL_userStatusRecently
func (e TL_userStatusRecently) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userStatusRecently)
	return x.buf
}

// userStatusLastWeek#7bf09fc = UserStatus;

const crc_userStatusLastWeek = 0x7bf09fc

type TL_userStatusLastWeek struct {
}

// Encoding TL_userStatusLastWeek
func (e TL_userStatusLastWeek) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userStatusLastWeek)
	return x.buf
}

// userStatusLastMonth#77ebc742 = UserStatus;

const crc_userStatusLastMonth = 0x77ebc742

type TL_userStatusLastMonth struct {
}

// Encoding TL_userStatusLastMonth
func (e TL_userStatusLastMonth) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userStatusLastMonth)
	return x.buf
}

// chatEmpty#9ba2d800 Id:int = Chat;

const crc_chatEmpty = 0x9ba2d800

type TL_chatEmpty struct {
	Id int32 // Id:int
}

// Encoding TL_chatEmpty
func (e TL_chatEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatEmpty)
	x.Int(e.Id)
	return x.buf
}

// chat#d91cdd54 flags:# creator:flags.0?true kicked:flags.1?true left:flags.2?true admins_enabled:flags.3?true admin:flags.4?true deactivated:flags.5?true Id:int title:string photo:ChatPhoto participants_count:int date:int Version:int migrated_to:flags.6?InputChannel = Chat;

const crc_chat = 0xd91cdd54

type TL_chat struct {
	Flags              int32
	Creator            bool   // creator:flags.0?true
	Kicked             bool   // kicked:flags.1?true
	Left               bool   // left:flags.2?true
	Admins_enabled     bool   // admins_enabled:flags.3?true
	Admin              bool   // admin:flags.4?true
	Deactivated        bool   // deactivated:flags.5?true
	Id                 int32  // Id:int
	Title              string // title:string
	Photo              TL     // photo:ChatPhoto
	Participants_count int32  // participants_count:int
	Date               int32  // date:int
	Version            int32  // Version:int
	Migrated_to        TL     // migrated_to:flags.6?InputChannel
}

// Encoding TL_chat
func (e TL_chat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chat)
	var flags int32
	if e.Creator {
		flags |= (1 << 0)
	}
	if e.Kicked {
		flags |= (1 << 1)
	}
	if e.Left {
		flags |= (1 << 2)
	}
	if e.Admins_enabled {
		flags |= (1 << 3)
	}
	if e.Admin {
		flags |= (1 << 4)
	}
	if e.Deactivated {
		flags |= (1 << 5)
	}
	if _, ok := (e.Migrated_to).(TL_null); !ok {
		flags |= (1 << 6)
	}
	x.Int(flags)
	x.Int(e.Id)
	x.String(e.Title)
	x.Bytes(e.Photo.encode())
	x.Int(e.Participants_count)
	x.Int(e.Date)
	x.Int(e.Version)
	if flags&(1<<6) != 0 {
		x.Bytes(e.Migrated_to.encode())
	}
	return x.buf
}

// chatForbidden#7328bdb Id:int title:string = Chat;

const crc_chatForbidden = 0x7328bdb

type TL_chatForbidden struct {
	Id    int32  // Id:int
	Title string // title:string
}

// Encoding TL_chatForbidden
func (e TL_chatForbidden) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatForbidden)
	x.Int(e.Id)
	x.String(e.Title)
	return x.buf
}

// channel#a14dca52 flags:# creator:flags.0?true kicked:flags.1?true left:flags.2?true editor:flags.3?true moderator:flags.4?true broadcast:flags.5?true verified:flags.7?true megagroup:flags.8?true restricted:flags.9?true democracy:flags.10?true signatures:flags.11?true min:flags.12?true Id:int access_hash:flags.13?long title:string username:flags.6?string photo:ChatPhoto date:int Version:int restriction_reason:flags.9?string = Chat;

const crc_channel = 0xa14dca52

type TL_channel struct {
	Flags              int32
	Creator            bool   // creator:flags.0?true
	Kicked             bool   // kicked:flags.1?true
	Left               bool   // left:flags.2?true
	Editor             bool   // editor:flags.3?true
	Moderator          bool   // moderator:flags.4?true
	Broadcast          bool   // broadcast:flags.5?true
	Verified           bool   // verified:flags.7?true
	Megagroup          bool   // megagroup:flags.8?true
	Restricted         bool   // restricted:flags.9?true
	Democracy          bool   // democracy:flags.10?true
	Signatures         bool   // signatures:flags.11?true
	Min                bool   // min:flags.12?true
	Id                 int32  // Id:int
	Access_hash        int64  // access_hash:flags.13?long
	Title              string // title:string
	Username           string // username:flags.6?string
	Photo              TL     // photo:ChatPhoto
	Date               int32  // date:int
	Version            int32  // Version:int
	Restriction_reason string // restriction_reason:flags.9?string
}

// Encoding TL_channel
func (e TL_channel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channel)
	var flags int32
	if e.Creator {
		flags |= (1 << 0)
	}
	if e.Kicked {
		flags |= (1 << 1)
	}
	if e.Left {
		flags |= (1 << 2)
	}
	if e.Editor {
		flags |= (1 << 3)
	}
	if e.Moderator {
		flags |= (1 << 4)
	}
	if e.Broadcast {
		flags |= (1 << 5)
	}
	if e.Verified {
		flags |= (1 << 7)
	}
	if e.Megagroup {
		flags |= (1 << 8)
	}
	if e.Restricted {
		flags |= (1 << 9)
	}
	if e.Democracy {
		flags |= (1 << 10)
	}
	if e.Signatures {
		flags |= (1 << 11)
	}
	if e.Min {
		flags |= (1 << 12)
	}
	if e.Access_hash > 0 {
		flags |= (1 << 13)
	}
	if e.Username != "" {
		flags |= (1 << 6)
	}
	if e.Restriction_reason != "" {
		flags |= (1 << 9)
	}
	x.Int(flags)
	x.Int(e.Id)
	if flags&(1<<13) != 0 {
		x.Long(e.Access_hash)
	}
	x.String(e.Title)
	if flags&(1<<6) != 0 {
		x.String(e.Username)
	}
	x.Bytes(e.Photo.encode())
	x.Int(e.Date)
	x.Int(e.Version)
	if flags&(1<<9) != 0 {
		x.String(e.Restriction_reason)
	}
	return x.buf
}

// channelForbidden#8537784f flags:# broadcast:flags.5?true megagroup:flags.8?true Id:int access_hash:long title:string = Chat;

const crc_channelForbidden = 0x8537784f

type TL_channelForbidden struct {
	Flags       int32
	Broadcast   bool   // broadcast:flags.5?true
	Megagroup   bool   // megagroup:flags.8?true
	Id          int32  // Id:int
	Access_hash int64  // access_hash:long
	Title       string // title:string
}

// Encoding TL_channelForbidden
func (e TL_channelForbidden) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelForbidden)
	var flags int32
	if e.Broadcast {
		flags |= (1 << 5)
	}
	if e.Megagroup {
		flags |= (1 << 8)
	}
	x.Int(flags)
	x.Int(e.Id)
	x.Long(e.Access_hash)
	x.String(e.Title)
	return x.buf
}

// chatFull#2e02a614 Id:int participants:ChatParticipants chat_photo:Photo notify_settings:PeerNotifySettings exported_invite:ExportedChatInvite bot_info:Vector<BotInfo> = ChatFull;

const crc_chatFull = 0x2e02a614

type TL_chatFull struct {
	Id              int32 // Id:int
	Participants    TL    // participants:ChatParticipants
	Chat_photo      TL    // chat_photo:Photo
	Notify_settings TL    // notify_settings:PeerNotifySettings
	Exported_invite TL    // exported_invite:ExportedChatInvite
	Bot_info        []TL  // bot_info:Vector<BotInfo>
}

// Encoding TL_chatFull
func (e TL_chatFull) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatFull)
	x.Int(e.Id)
	x.Bytes(e.Participants.encode())
	x.Bytes(e.Chat_photo.encode())
	x.Bytes(e.Notify_settings.encode())
	x.Bytes(e.Exported_invite.encode())
	x.Vector(e.Bot_info)
	return x.buf
}

// channelFull#c3d5512f flags:# can_view_participants:flags.3?true can_set_username:flags.6?true Id:int about:string participants_count:flags.0?int admins_count:flags.1?int kicked_count:flags.2?int read_inbox_max_id:int read_outbox_max_id:int unread_count:int chat_photo:Photo notify_settings:PeerNotifySettings exported_invite:ExportedChatInvite bot_info:Vector<BotInfo> migrated_from_chat_id:flags.4?int migrated_from_max_id:flags.4?int pinned_msg_id:flags.5?int = ChatFull;

const crc_channelFull = 0xc3d5512f

type TL_channelFull struct {
	Flags                 int32
	Can_view_participants bool   // can_view_participants:flags.3?true
	Can_set_username      bool   // can_set_username:flags.6?true
	Id                    int32  // Id:int
	About                 string // about:string
	Participants_count    int32  // participants_count:flags.0?int
	Admins_count          int32  // admins_count:flags.1?int
	Kicked_count          int32  // kicked_count:flags.2?int
	Read_inbox_max_id     int32  // read_inbox_max_id:int
	Read_outbox_max_id    int32  // read_outbox_max_id:int
	Unread_count          int32  // unread_count:int
	Chat_photo            TL     // chat_photo:Photo
	Notify_settings       TL     // notify_settings:PeerNotifySettings
	Exported_invite       TL     // exported_invite:ExportedChatInvite
	Bot_info              []TL   // bot_info:Vector<BotInfo>
	Migrated_from_chat_id int32  // migrated_from_chat_id:flags.4?int
	Migrated_from_max_id  int32  // migrated_from_max_id:flags.4?int
	Pinned_msg_id         int32  // pinned_msg_id:flags.5?int
}

// Encoding TL_channelFull
func (e TL_channelFull) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelFull)
	var flags int32
	if e.Can_view_participants {
		flags |= (1 << 3)
	}
	if e.Can_set_username {
		flags |= (1 << 6)
	}
	if e.Participants_count > 0 {
		flags |= (1 << 0)
	}
	if e.Admins_count > 0 {
		flags |= (1 << 1)
	}
	if e.Kicked_count > 0 {
		flags |= (1 << 2)
	}
	if e.Migrated_from_chat_id > 0 {
		flags |= (1 << 4)
	}
	if e.Migrated_from_max_id > 0 {
		flags |= (1 << 4)
	}
	if e.Pinned_msg_id > 0 {
		flags |= (1 << 5)
	}
	x.Int(flags)
	x.Int(e.Id)
	x.String(e.About)
	if flags&(1<<0) != 0 {
		x.Int(e.Participants_count)
	}
	if flags&(1<<1) != 0 {
		x.Int(e.Admins_count)
	}
	if flags&(1<<2) != 0 {
		x.Int(e.Kicked_count)
	}
	x.Int(e.Read_inbox_max_id)
	x.Int(e.Read_outbox_max_id)
	x.Int(e.Unread_count)
	x.Bytes(e.Chat_photo.encode())
	x.Bytes(e.Notify_settings.encode())
	x.Bytes(e.Exported_invite.encode())
	x.Vector(e.Bot_info)
	if flags&(1<<4) != 0 {
		x.Int(e.Migrated_from_chat_id)
	}
	if flags&(1<<4) != 0 {
		x.Int(e.Migrated_from_max_id)
	}
	if flags&(1<<5) != 0 {
		x.Int(e.Pinned_msg_id)
	}
	return x.buf
}

// chatParticipant#c8d7493e user_id:int inviter_id:int date:int = ChatParticipant;

const crc_chatParticipant = 0xc8d7493e

type TL_chatParticipant struct {
	User_id    int32 // user_id:int
	Inviter_id int32 // inviter_id:int
	Date       int32 // date:int
}

// Encoding TL_chatParticipant
func (e TL_chatParticipant) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatParticipant)
	x.Int(e.User_id)
	x.Int(e.Inviter_id)
	x.Int(e.Date)
	return x.buf
}

// chatParticipantCreator#da13538a user_id:int = ChatParticipant;

const crc_chatParticipantCreator = 0xda13538a

type TL_chatParticipantCreator struct {
	User_id int32 // user_id:int
}

// Encoding TL_chatParticipantCreator
func (e TL_chatParticipantCreator) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatParticipantCreator)
	x.Int(e.User_id)
	return x.buf
}

// chatParticipantAdmin#e2d6e436 user_id:int inviter_id:int date:int = ChatParticipant;

const crc_chatParticipantAdmin = 0xe2d6e436

type TL_chatParticipantAdmin struct {
	User_id    int32 // user_id:int
	Inviter_id int32 // inviter_id:int
	Date       int32 // date:int
}

// Encoding TL_chatParticipantAdmin
func (e TL_chatParticipantAdmin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatParticipantAdmin)
	x.Int(e.User_id)
	x.Int(e.Inviter_id)
	x.Int(e.Date)
	return x.buf
}

// chatParticipantsForbidden#fc900c2b flags:# chat_id:int self_participant:flags.0?ChatParticipant = ChatParticipants;

const crc_chatParticipantsForbidden = 0xfc900c2b

type TL_chatParticipantsForbidden struct {
	Flags            int32
	Chat_id          int32 // chat_id:int
	Self_participant TL    // self_participant:flags.0?ChatParticipant
}

// Encoding TL_chatParticipantsForbidden
func (e TL_chatParticipantsForbidden) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatParticipantsForbidden)
	var flags int32
	if _, ok := (e.Self_participant).(TL_null); !ok {
		flags |= (1 << 0)
	}
	x.Int(flags)
	x.Int(e.Chat_id)
	if flags&(1<<0) != 0 {
		x.Bytes(e.Self_participant.encode())
	}
	return x.buf
}

// chatParticipants#3f460fed chat_id:int participants:Vector<ChatParticipant> Version:int = ChatParticipants;

const crc_chatParticipants = 0x3f460fed

type TL_chatParticipants struct {
	Chat_id      int32 // chat_id:int
	Participants []TL  // participants:Vector<ChatParticipant>
	Version      int32 // Version:int
}

// Encoding TL_chatParticipants
func (e TL_chatParticipants) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatParticipants)
	x.Int(e.Chat_id)
	x.Vector(e.Participants)
	x.Int(e.Version)
	return x.buf
}

// chatPhotoEmpty#37c1011c = ChatPhoto;

const crc_chatPhotoEmpty = 0x37c1011c

type TL_chatPhotoEmpty struct {
}

// Encoding TL_chatPhotoEmpty
func (e TL_chatPhotoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatPhotoEmpty)
	return x.buf
}

// chatPhoto#6153276a photo_small:FileLocation photo_big:FileLocation = ChatPhoto;

const crc_chatPhoto = 0x6153276a

type TL_chatPhoto struct {
	Photo_small TL // photo_small:FileLocation
	Photo_big   TL // photo_big:FileLocation
}

// Encoding TL_chatPhoto
func (e TL_chatPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatPhoto)
	x.Bytes(e.Photo_small.encode())
	x.Bytes(e.Photo_big.encode())
	return x.buf
}

// messageEmpty#83e5de54 Id:int = Message;

const crc_messageEmpty = 0x83e5de54

type TL_messageEmpty struct {
	Id int32 // Id:int
}

// Encoding TL_messageEmpty
func (e TL_messageEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEmpty)
	x.Int(e.Id)
	return x.buf
}

// message#c09be45f flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true post:flags.14?true Id:int from_id:flags.8?int to_id:Peer fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to_msg_id:flags.3?int date:int message:string media:flags.9?MessageMedia reply_markup:flags.6?ReplyMarkup entities:flags.7?Vector<MessageEntity> views:flags.10?int edit_date:flags.15?int = Message;

const crc_message = 0xc09be45f

type TL_message struct {
	Flags           int32
	Out             bool   // out:flags.1?true
	Mentioned       bool   // mentioned:flags.4?true
	Media_unread    bool   // media_unread:flags.5?true
	Silent          bool   // silent:flags.13?true
	Post            bool   // post:flags.14?true
	Id              int32  // Id:int
	From_id         int32  // from_id:flags.8?int
	To_id           TL     // to_id:Peer
	Fwd_from        TL     // fwd_from:flags.2?MessageFwdHeader
	Via_bot_id      int32  // via_bot_id:flags.11?int
	Reply_to_msg_id int32  // reply_to_msg_id:flags.3?int
	Date            int32  // date:int
	Message         string // message:string
	Media           TL     // media:flags.9?MessageMedia
	Reply_markup    TL     // reply_markup:flags.6?ReplyMarkup
	Entities        []TL   // entities:flags.7?Vector<MessageEntity>
	Views           int32  // views:flags.10?int
	Edit_date       int32  // edit_date:flags.15?int
}

// Encoding TL_message
func (e TL_message) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_message)
	var flags int32
	if e.Out {
		flags |= (1 << 1)
	}
	if e.Mentioned {
		flags |= (1 << 4)
	}
	if e.Media_unread {
		flags |= (1 << 5)
	}
	if e.Silent {
		flags |= (1 << 13)
	}
	if e.Post {
		flags |= (1 << 14)
	}
	if e.From_id > 0 {
		flags |= (1 << 8)
	}
	if _, ok := (e.Fwd_from).(TL_null); !ok {
		flags |= (1 << 2)
	}
	if e.Via_bot_id > 0 {
		flags |= (1 << 11)
	}
	if e.Reply_to_msg_id > 0 {
		flags |= (1 << 3)
	}
	if _, ok := (e.Media).(TL_null); !ok {
		flags |= (1 << 9)
	}
	if _, ok := (e.Reply_markup).(TL_null); !ok {
		flags |= (1 << 6)
	}
	if len(e.Entities) != 0 {
		flags |= (1 << 7)
	}
	if e.Views > 0 {
		flags |= (1 << 10)
	}
	if e.Edit_date > 0 {
		flags |= (1 << 15)
	}
	x.Int(flags)
	x.Int(e.Id)
	if flags&(1<<8) != 0 {
		x.Int(e.From_id)
	}
	x.Bytes(e.To_id.encode())
	if flags&(1<<2) != 0 {
		x.Bytes(e.Fwd_from.encode())
	}
	if flags&(1<<11) != 0 {
		x.Int(e.Via_bot_id)
	}
	if flags&(1<<3) != 0 {
		x.Int(e.Reply_to_msg_id)
	}
	x.Int(e.Date)
	x.String(e.Message)
	if flags&(1<<9) != 0 {
		x.Bytes(e.Media.encode())
	}
	if flags&(1<<6) != 0 {
		x.Bytes(e.Reply_markup.encode())
	}
	if flags&(1<<7) != 0 {
		x.Vector(e.Entities)
	}
	if flags&(1<<10) != 0 {
		x.Int(e.Views)
	}
	if flags&(1<<15) != 0 {
		x.Int(e.Edit_date)
	}
	return x.buf
}

// messageService#9e19a1f6 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true post:flags.14?true Id:int from_id:flags.8?int to_id:Peer reply_to_msg_id:flags.3?int date:int action:MessageAction = Message;

const crc_messageService = 0x9e19a1f6

type TL_messageService struct {
	Flags           int32
	Out             bool  // out:flags.1?true
	Mentioned       bool  // mentioned:flags.4?true
	Media_unread    bool  // media_unread:flags.5?true
	Silent          bool  // silent:flags.13?true
	Post            bool  // post:flags.14?true
	Id              int32 // Id:int
	From_id         int32 // from_id:flags.8?int
	To_id           TL    // to_id:Peer
	Reply_to_msg_id int32 // reply_to_msg_id:flags.3?int
	Date            int32 // date:int
	Action          TL    // action:MessageAction
}

// Encoding TL_messageService
func (e TL_messageService) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageService)
	var flags int32
	if e.Out {
		flags |= (1 << 1)
	}
	if e.Mentioned {
		flags |= (1 << 4)
	}
	if e.Media_unread {
		flags |= (1 << 5)
	}
	if e.Silent {
		flags |= (1 << 13)
	}
	if e.Post {
		flags |= (1 << 14)
	}
	if e.From_id > 0 {
		flags |= (1 << 8)
	}
	if e.Reply_to_msg_id > 0 {
		flags |= (1 << 3)
	}
	x.Int(flags)
	x.Int(e.Id)
	if flags&(1<<8) != 0 {
		x.Int(e.From_id)
	}
	x.Bytes(e.To_id.encode())
	if flags&(1<<3) != 0 {
		x.Int(e.Reply_to_msg_id)
	}
	x.Int(e.Date)
	x.Bytes(e.Action.encode())
	return x.buf
}

// messageMediaEmpty#3ded6320 = MessageMedia;

const crc_messageMediaEmpty = 0x3ded6320

type TL_messageMediaEmpty struct {
}

// Encoding TL_messageMediaEmpty
func (e TL_messageMediaEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaEmpty)
	return x.buf
}

// messageMediaPhoto#3d8ce53d photo:Photo caption:string = MessageMedia;

const crc_messageMediaPhoto = 0x3d8ce53d

type TL_messageMediaPhoto struct {
	Photo   TL     // photo:Photo
	Caption string // caption:string
}

// Encoding TL_messageMediaPhoto
func (e TL_messageMediaPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaPhoto)
	x.Bytes(e.Photo.encode())
	x.String(e.Caption)
	return x.buf
}

// messageMediaGeo#56e0d474 geo:GeoPoint = MessageMedia;

const crc_messageMediaGeo = 0x56e0d474

type TL_messageMediaGeo struct {
	Geo TL // geo:GeoPoint
}

// Encoding TL_messageMediaGeo
func (e TL_messageMediaGeo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaGeo)
	x.Bytes(e.Geo.encode())
	return x.buf
}

// messageMediaContact#5e7d2f39 phone_number:string first_name:string last_name:string user_id:int = MessageMedia;

const crc_messageMediaContact = 0x5e7d2f39

type TL_messageMediaContact struct {
	Phone_number string // phone_number:string
	First_name   string // first_name:string
	Last_name    string // last_name:string
	User_id      int32  // user_id:int
}

// Encoding TL_messageMediaContact
func (e TL_messageMediaContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaContact)
	x.String(e.Phone_number)
	x.String(e.First_name)
	x.String(e.Last_name)
	x.Int(e.User_id)
	return x.buf
}

// messageMediaUnsupported#9f84f49e = MessageMedia;

const crc_messageMediaUnsupported = 0x9f84f49e

type TL_messageMediaUnsupported struct {
}

// Encoding TL_messageMediaUnsupported
func (e TL_messageMediaUnsupported) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaUnsupported)
	return x.buf
}

// messageMediaDocument#f3e02ea8 document:Document caption:string = MessageMedia;

const crc_messageMediaDocument = 0xf3e02ea8

type TL_messageMediaDocument struct {
	Document TL     // document:Document
	Caption  string // caption:string
}

// Encoding TL_messageMediaDocument
func (e TL_messageMediaDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaDocument)
	x.Bytes(e.Document.encode())
	x.String(e.Caption)
	return x.buf
}

// messageMediaWebPage#a32dd600 webpage:WebPage = MessageMedia;

const crc_messageMediaWebPage = 0xa32dd600

type TL_messageMediaWebPage struct {
	Webpage TL // webpage:WebPage
}

// Encoding TL_messageMediaWebPage
func (e TL_messageMediaWebPage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaWebPage)
	x.Bytes(e.Webpage.encode())
	return x.buf
}

// messageMediaVenue#7912b71f geo:GeoPoint title:string address:string provider:string venue_id:string = MessageMedia;

const crc_messageMediaVenue = 0x7912b71f

type TL_messageMediaVenue struct {
	Geo      TL     // geo:GeoPoint
	Title    string // title:string
	Address  string // address:string
	Provider string // provider:string
	Venue_id string // venue_id:string
}

// Encoding TL_messageMediaVenue
func (e TL_messageMediaVenue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaVenue)
	x.Bytes(e.Geo.encode())
	x.String(e.Title)
	x.String(e.Address)
	x.String(e.Provider)
	x.String(e.Venue_id)
	return x.buf
}

// messageMediaGame#fdb19008 game:Game = MessageMedia;

const crc_messageMediaGame = 0xfdb19008

type TL_messageMediaGame struct {
	Game TL // game:Game
}

// Encoding TL_messageMediaGame
func (e TL_messageMediaGame) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaGame)
	x.Bytes(e.Game.encode())
	return x.buf
}

// messageMediaInvoice#84551347 flags:# shipping_address_requested:flags.1?true test:flags.3?true title:string description:string photo:flags.0?WebDocument receipt_msg_id:flags.2?int currency:string total_amount:long start_param:string = MessageMedia;

const crc_messageMediaInvoice = 0x84551347

type TL_messageMediaInvoice struct {
	Flags                      int32
	Shipping_address_requested bool   // shipping_address_requested:flags.1?true
	Test                       bool   // test:flags.3?true
	Title                      string // title:string
	Description                string // description:string
	Photo                      TL     // photo:flags.0?WebDocument
	Receipt_msg_id             int32  // receipt_msg_id:flags.2?int
	Currency                   string // currency:string
	Total_amount               int64  // total_amount:long
	Start_param                string // start_param:string
}

// Encoding TL_messageMediaInvoice
func (e TL_messageMediaInvoice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaInvoice)
	var flags int32
	if e.Shipping_address_requested {
		flags |= (1 << 1)
	}
	if e.Test {
		flags |= (1 << 3)
	}
	if _, ok := (e.Photo).(TL_null); !ok {
		flags |= (1 << 0)
	}
	if e.Receipt_msg_id > 0 {
		flags |= (1 << 2)
	}
	x.Int(flags)
	x.String(e.Title)
	x.String(e.Description)
	if flags&(1<<0) != 0 {
		x.Bytes(e.Photo.encode())
	}
	if flags&(1<<2) != 0 {
		x.Int(e.Receipt_msg_id)
	}
	x.String(e.Currency)
	x.Long(e.Total_amount)
	x.String(e.Start_param)
	return x.buf
}

// messageActionEmpty#b6aef7b0 = MessageAction;

const crc_messageActionEmpty = 0xb6aef7b0

type TL_messageActionEmpty struct {
}

// Encoding TL_messageActionEmpty
func (e TL_messageActionEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionEmpty)
	return x.buf
}

// messageActionChatCreate#a6638b9a title:string users:Vector<int> = MessageAction;

const crc_messageActionChatCreate = 0xa6638b9a

type TL_messageActionChatCreate struct {
	Title string  // title:string
	Users []int32 // users:Vector<int>
}

// Encoding TL_messageActionChatCreate
func (e TL_messageActionChatCreate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatCreate)
	x.String(e.Title)
	x.VectorInt(e.Users)
	return x.buf
}

// messageActionChatEditTitle#b5a1ce5a title:string = MessageAction;

const crc_messageActionChatEditTitle = 0xb5a1ce5a

type TL_messageActionChatEditTitle struct {
	Title string // title:string
}

// Encoding TL_messageActionChatEditTitle
func (e TL_messageActionChatEditTitle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatEditTitle)
	x.String(e.Title)
	return x.buf
}

// messageActionChatEditPhoto#7fcb13a8 photo:Photo = MessageAction;

const crc_messageActionChatEditPhoto = 0x7fcb13a8

type TL_messageActionChatEditPhoto struct {
	Photo TL // photo:Photo
}

// Encoding TL_messageActionChatEditPhoto
func (e TL_messageActionChatEditPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatEditPhoto)
	x.Bytes(e.Photo.encode())
	return x.buf
}

// messageActionChatDeletePhoto#95e3fbef = MessageAction;

const crc_messageActionChatDeletePhoto = 0x95e3fbef

type TL_messageActionChatDeletePhoto struct {
}

// Encoding TL_messageActionChatDeletePhoto
func (e TL_messageActionChatDeletePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatDeletePhoto)
	return x.buf
}

// messageActionChatAddUser#488a7337 users:Vector<int> = MessageAction;

const crc_messageActionChatAddUser = 0x488a7337

type TL_messageActionChatAddUser struct {
	Users []int32 // users:Vector<int>
}

// Encoding TL_messageActionChatAddUser
func (e TL_messageActionChatAddUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatAddUser)
	x.VectorInt(e.Users)
	return x.buf
}

// messageActionChatDeleteUser#b2ae9b0c user_id:int = MessageAction;

const crc_messageActionChatDeleteUser = 0xb2ae9b0c

type TL_messageActionChatDeleteUser struct {
	User_id int32 // user_id:int
}

// Encoding TL_messageActionChatDeleteUser
func (e TL_messageActionChatDeleteUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatDeleteUser)
	x.Int(e.User_id)
	return x.buf
}

// messageActionChatJoinedByLink#f89cf5e8 inviter_id:int = MessageAction;

const crc_messageActionChatJoinedByLink = 0xf89cf5e8

type TL_messageActionChatJoinedByLink struct {
	Inviter_id int32 // inviter_id:int
}

// Encoding TL_messageActionChatJoinedByLink
func (e TL_messageActionChatJoinedByLink) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatJoinedByLink)
	x.Int(e.Inviter_id)
	return x.buf
}

// messageActionChannelCreate#95d2ac92 title:string = MessageAction;

const crc_messageActionChannelCreate = 0x95d2ac92

type TL_messageActionChannelCreate struct {
	Title string // title:string
}

// Encoding TL_messageActionChannelCreate
func (e TL_messageActionChannelCreate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChannelCreate)
	x.String(e.Title)
	return x.buf
}

// messageActionChatMigrateTo#51bdb021 channel_id:int = MessageAction;

const crc_messageActionChatMigrateTo = 0x51bdb021

type TL_messageActionChatMigrateTo struct {
	Channel_id int32 // channel_id:int
}

// Encoding TL_messageActionChatMigrateTo
func (e TL_messageActionChatMigrateTo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatMigrateTo)
	x.Int(e.Channel_id)
	return x.buf
}

// messageActionChannelMigrateFrom#b055eaee title:string chat_id:int = MessageAction;

const crc_messageActionChannelMigrateFrom = 0xb055eaee

type TL_messageActionChannelMigrateFrom struct {
	Title   string // title:string
	Chat_id int32  // chat_id:int
}

// Encoding TL_messageActionChannelMigrateFrom
func (e TL_messageActionChannelMigrateFrom) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChannelMigrateFrom)
	x.String(e.Title)
	x.Int(e.Chat_id)
	return x.buf
}

// messageActionPinMessage#94bd38ed = MessageAction;

const crc_messageActionPinMessage = 0x94bd38ed

type TL_messageActionPinMessage struct {
}

// Encoding TL_messageActionPinMessage
func (e TL_messageActionPinMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionPinMessage)
	return x.buf
}

// messageActionHistoryClear#9fbab604 = MessageAction;

const crc_messageActionHistoryClear = 0x9fbab604

type TL_messageActionHistoryClear struct {
}

// Encoding TL_messageActionHistoryClear
func (e TL_messageActionHistoryClear) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionHistoryClear)
	return x.buf
}

// messageActionGameScore#92a72876 game_id:long score:int = MessageAction;

const crc_messageActionGameScore = 0x92a72876

type TL_messageActionGameScore struct {
	Game_id int64 // game_id:long
	Score   int32 // score:int
}

// Encoding TL_messageActionGameScore
func (e TL_messageActionGameScore) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionGameScore)
	x.Long(e.Game_id)
	x.Int(e.Score)
	return x.buf
}

// messageActionPaymentSentMe#8f31b327 flags:# currency:string total_amount:long payload:bytes info:flags.0?PaymentRequestedInfo shipping_option_id:flags.1?string charge:PaymentCharge = MessageAction;

const crc_messageActionPaymentSentMe = 0x8f31b327

type TL_messageActionPaymentSentMe struct {
	Flags              int32
	Currency           string // currency:string
	Total_amount       int64  // total_amount:long
	Payload            []byte // payload:bytes
	Info               TL     // info:flags.0?PaymentRequestedInfo
	Shipping_option_id string // shipping_option_id:flags.1?string
	Charge             TL     // charge:PaymentCharge
}

// Encoding TL_messageActionPaymentSentMe
func (e TL_messageActionPaymentSentMe) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionPaymentSentMe)
	var flags int32
	if _, ok := (e.Info).(TL_null); !ok {
		flags |= (1 << 0)
	}
	if e.Shipping_option_id != "" {
		flags |= (1 << 1)
	}
	x.Int(flags)
	x.String(e.Currency)
	x.Long(e.Total_amount)
	x.Bytes(e.Payload)
	if flags&(1<<0) != 0 {
		x.Bytes(e.Info.encode())
	}
	if flags&(1<<1) != 0 {
		x.String(e.Shipping_option_id)
	}
	x.Bytes(e.Charge.encode())
	return x.buf
}

// messageActionPaymentSent#40699cd0 currency:string total_amount:long = MessageAction;

const crc_messageActionPaymentSent = 0x40699cd0

type TL_messageActionPaymentSent struct {
	Currency     string // currency:string
	Total_amount int64  // total_amount:long
}

// Encoding TL_messageActionPaymentSent
func (e TL_messageActionPaymentSent) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionPaymentSent)
	x.String(e.Currency)
	x.Long(e.Total_amount)
	return x.buf
}

// messageActionPhoneCall#80e11a7f flags:# call_id:long reason:flags.0?PhoneCallDiscardReason duration:flags.1?int = MessageAction;

const crc_messageActionPhoneCall = 0x80e11a7f

type TL_messageActionPhoneCall struct {
	Flags    int32
	Call_id  int64 // call_id:long
	Reason   TL    // reason:flags.0?PhoneCallDiscardReason
	Duration int32 // duration:flags.1?int
}

// Encoding TL_messageActionPhoneCall
func (e TL_messageActionPhoneCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionPhoneCall)
	var flags int32
	if _, ok := (e.Reason).(TL_null); !ok {
		flags |= (1 << 0)
	}
	if e.Duration > 0 {
		flags |= (1 << 1)
	}
	x.Int(flags)
	x.Long(e.Call_id)
	if flags&(1<<0) != 0 {
		x.Bytes(e.Reason.encode())
	}
	if flags&(1<<1) != 0 {
		x.Int(e.Duration)
	}
	return x.buf
}

// dialog#66ffba14 flags:# pinned:flags.2?true peer:Peer top_message:int read_inbox_max_id:int read_outbox_max_id:int unread_count:int notify_settings:PeerNotifySettings pts:flags.0?int draft:flags.1?DraftMessage = Dialog;

const crc_dialog = 0x66ffba14

type TL_dialog struct {
	Flags              int32
	Pinned             bool  // pinned:flags.2?true
	Peer               TL    // peer:Peer
	Top_message        int32 // top_message:int
	Read_inbox_max_id  int32 // read_inbox_max_id:int
	Read_outbox_max_id int32 // read_outbox_max_id:int
	Unread_count       int32 // unread_count:int
	Notify_settings    TL    // notify_settings:PeerNotifySettings
	Pts                int32 // pts:flags.0?int
	Draft              TL    // draft:flags.1?DraftMessage
}

// Encoding TL_dialog
func (e TL_dialog) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_dialog)
	var flags int32
	if e.Pinned {
		flags |= (1 << 2)
	}
	if e.Pts > 0 {
		flags |= (1 << 0)
	}
	if _, ok := (e.Draft).(TL_null); !ok {
		flags |= (1 << 1)
	}
	x.Int(flags)
	x.Bytes(e.Peer.encode())
	x.Int(e.Top_message)
	x.Int(e.Read_inbox_max_id)
	x.Int(e.Read_outbox_max_id)
	x.Int(e.Unread_count)
	x.Bytes(e.Notify_settings.encode())
	if flags&(1<<0) != 0 {
		x.Int(e.Pts)
	}
	if flags&(1<<1) != 0 {
		x.Bytes(e.Draft.encode())
	}
	return x.buf
}

// photoEmpty#2331b22d Id:long = Photo;

const crc_photoEmpty = 0x2331b22d

type TL_photoEmpty struct {
	Id int64 // Id:long
}

// Encoding TL_photoEmpty
func (e TL_photoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photoEmpty)
	x.Long(e.Id)
	return x.buf
}

// photo#9288dd29 flags:# has_stickers:flags.0?true Id:long access_hash:long date:int sizes:Vector<PhotoSize> = Photo;

const crc_photo = 0x9288dd29

type TL_photo struct {
	Flags        int32
	Has_stickers bool  // has_stickers:flags.0?true
	Id           int64 // Id:long
	Access_hash  int64 // access_hash:long
	Date         int32 // date:int
	Sizes        []TL  // sizes:Vector<PhotoSize>
}

// Encoding TL_photo
func (e TL_photo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photo)
	var flags int32
	if e.Has_stickers {
		flags |= (1 << 0)
	}
	x.Int(flags)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	x.Int(e.Date)
	x.Vector(e.Sizes)
	return x.buf
}

// photoSizeEmpty#e17e23c type:string = PhotoSize;

const crc_photoSizeEmpty = 0xe17e23c

type TL_photoSizeEmpty struct {
	Code_type string // type:string
}

// Encoding TL_photoSizeEmpty
func (e TL_photoSizeEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photoSizeEmpty)
	x.String(e.Code_type)
	return x.buf
}

// photoSize#77bfb61b type:string location:FileLocation w:int h:int size:int = PhotoSize;

const crc_photoSize = 0x77bfb61b

type TL_photoSize struct {
	Code_type string // type:string
	Location  TL     // location:FileLocation
	W         int32  // w:int
	H         int32  // h:int
	Size      int32  // size:int
}

// Encoding TL_photoSize
func (e TL_photoSize) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photoSize)
	x.String(e.Code_type)
	x.Bytes(e.Location.encode())
	x.Int(e.W)
	x.Int(e.H)
	x.Int(e.Size)
	return x.buf
}

// photoCachedSize#e9a734fa type:string location:FileLocation w:int h:int bytes:bytes = PhotoSize;

const crc_photoCachedSize = 0xe9a734fa

type TL_photoCachedSize struct {
	Code_type string // type:string
	Location  TL     // location:FileLocation
	W         int32  // w:int
	H         int32  // h:int
	Bytes     []byte // bytes:bytes
}

// Encoding TL_photoCachedSize
func (e TL_photoCachedSize) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photoCachedSize)
	x.String(e.Code_type)
	x.Bytes(e.Location.encode())
	x.Int(e.W)
	x.Int(e.H)
	x.Bytes(e.Bytes)
	return x.buf
}

// geoPointEmpty#1117dd5f = GeoPoint;

const crc_geoPointEmpty = 0x1117dd5f

type TL_geoPointEmpty struct {
}

// Encoding TL_geoPointEmpty
func (e TL_geoPointEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_geoPointEmpty)
	return x.buf
}

// geoPoint#2049d70c long:double lat:double = GeoPoint;

const crc_geoPoint = 0x2049d70c

type TL_geoPoint struct {
	Long float64 // long:double
	Lat  float64 // lat:double
}

// Encoding TL_geoPoint
func (e TL_geoPoint) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_geoPoint)
	x.Double(e.Long)
	x.Double(e.Lat)
	return x.buf
}

// auth.checkedPhone#811ea28e phone_registered:Bool = auth.CheckedPhone;

const crc_auth_checkedPhone = 0x811ea28e

type TL_auth_checkedPhone struct {
	Phone_registered TL // phone_registered:Bool
}

// Encoding TL_auth_checkedPhone
func (e TL_auth_checkedPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_checkedPhone)
	x.Bytes(e.Phone_registered.encode())
	return x.buf
}

// auth.sentCode#5e002502 flags:# phone_registered:flags.0?true type:auth.SentCodeType phone_code_hash:string next_type:flags.1?auth.CodeType timeout:flags.2?int = auth.SentCode;

const crc_auth_sentCode = 0x5e002502

type TL_auth_sentCode struct {
	Flags            int32
	Phone_registered bool   // phone_registered:flags.0?true
	Code_type        TL     // type:auth.SentCodeType
	Phone_code_hash  string // phone_code_hash:string
	Next_type        TL     // next_type:flags.1?auth.CodeType
	Timeout          int32  // timeout:flags.2?int
}

// Encoding TL_auth_sentCode
func (e TL_auth_sentCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_sentCode)
	var flags int32
	if e.Phone_registered {
		flags |= (1 << 0)
	}
	if _, ok := (e.Next_type).(TL_null); !ok {
		flags |= (1 << 1)
	}
	if e.Timeout > 0 {
		flags |= (1 << 2)
	}
	x.Int(flags)
	x.Bytes(e.Code_type.encode())
	x.String(e.Phone_code_hash)
	if flags&(1<<1) != 0 {
		x.Bytes(e.Next_type.encode())
	}
	if flags&(1<<2) != 0 {
		x.Int(e.Timeout)
	}
	return x.buf
}

// auth.authorization#cd050916 flags:# tmp_sessions:flags.0?int user:User = auth.Authorization;

const crc_auth_authorization = 0xcd050916

type TL_auth_authorization struct {
	Flags        int32
	Tmp_sessions int32 // tmp_sessions:flags.0?int
	User         TL    // user:User
}

// Encoding TL_auth_authorization
func (e TL_auth_authorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_authorization)
	var flags int32
	if e.Tmp_sessions > 0 {
		flags |= (1 << 0)
	}
	x.Int(flags)
	if flags&(1<<0) != 0 {
		x.Int(e.Tmp_sessions)
	}
	x.Bytes(e.User.encode())
	return x.buf
}

// auth.exportedAuthorization#df969c2d Id:int bytes:bytes = auth.ExportedAuthorization;

const crc_auth_exportedAuthorization = 0xdf969c2d

type TL_auth_exportedAuthorization struct {
	Id    int32  // Id:int
	Bytes []byte // bytes:bytes
}

// Encoding TL_auth_exportedAuthorization
func (e TL_auth_exportedAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_exportedAuthorization)
	x.Int(e.Id)
	x.Bytes(e.Bytes)
	return x.buf
}

// inputNotifyPeer#b8bc5b0c peer:InputPeer = InputNotifyPeer;

const crc_inputNotifyPeer = 0xb8bc5b0c

type TL_inputNotifyPeer struct {
	Peer TL // peer:InputPeer
}

// Encoding TL_inputNotifyPeer
func (e TL_inputNotifyPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputNotifyPeer)
	x.Bytes(e.Peer.encode())
	return x.buf
}

// inputNotifyUsers#193b4417 = InputNotifyPeer;

const crc_inputNotifyUsers = 0x193b4417

type TL_inputNotifyUsers struct {
}

// Encoding TL_inputNotifyUsers
func (e TL_inputNotifyUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputNotifyUsers)
	return x.buf
}

// inputNotifyChats#4a95e84e = InputNotifyPeer;

const crc_inputNotifyChats = 0x4a95e84e

type TL_inputNotifyChats struct {
}

// Encoding TL_inputNotifyChats
func (e TL_inputNotifyChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputNotifyChats)
	return x.buf
}

// inputNotifyAll#a429b886 = InputNotifyPeer;

const crc_inputNotifyAll = 0xa429b886

type TL_inputNotifyAll struct {
}

// Encoding TL_inputNotifyAll
func (e TL_inputNotifyAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputNotifyAll)
	return x.buf
}

// inputPeerNotifyEventsEmpty#f03064d8 = InputPeerNotifyEvents;

const crc_inputPeerNotifyEventsEmpty = 0xf03064d8

type TL_inputPeerNotifyEventsEmpty struct {
}

// Encoding TL_inputPeerNotifyEventsEmpty
func (e TL_inputPeerNotifyEventsEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerNotifyEventsEmpty)
	return x.buf
}

// inputPeerNotifyEventsAll#e86a2c74 = InputPeerNotifyEvents;

const crc_inputPeerNotifyEventsAll = 0xe86a2c74

type TL_inputPeerNotifyEventsAll struct {
}

// Encoding TL_inputPeerNotifyEventsAll
func (e TL_inputPeerNotifyEventsAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerNotifyEventsAll)
	return x.buf
}

// inputPeerNotifySettings#38935eb2 flags:# show_previews:flags.0?true silent:flags.1?true mute_until:int sound:string = InputPeerNotifySettings;

const crc_inputPeerNotifySettings = 0x38935eb2

type TL_inputPeerNotifySettings struct {
	Flags         int32
	Show_previews bool   // show_previews:flags.0?true
	Silent        bool   // silent:flags.1?true
	Mute_until    int32  // mute_until:int
	Sound         string // sound:string
}

// Encoding TL_inputPeerNotifySettings
func (e TL_inputPeerNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerNotifySettings)
	var flags int32
	if e.Show_previews {
		flags |= (1 << 0)
	}
	if e.Silent {
		flags |= (1 << 1)
	}
	x.Int(flags)
	x.Int(e.Mute_until)
	x.String(e.Sound)
	return x.buf
}

// peerNotifyEventsEmpty#add53cb3 = PeerNotifyEvents;

const crc_peerNotifyEventsEmpty = 0xadd53cb3

type TL_peerNotifyEventsEmpty struct {
}

// Encoding TL_peerNotifyEventsEmpty
func (e TL_peerNotifyEventsEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerNotifyEventsEmpty)
	return x.buf
}

// peerNotifyEventsAll#6d1ded88 = PeerNotifyEvents;

const crc_peerNotifyEventsAll = 0x6d1ded88

type TL_peerNotifyEventsAll struct {
}

// Encoding TL_peerNotifyEventsAll
func (e TL_peerNotifyEventsAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerNotifyEventsAll)
	return x.buf
}

// peerNotifySettingsEmpty#70a68512 = PeerNotifySettings;

const crc_peerNotifySettingsEmpty = 0x70a68512

type TL_peerNotifySettingsEmpty struct {
}

// Encoding TL_peerNotifySettingsEmpty
func (e TL_peerNotifySettingsEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerNotifySettingsEmpty)
	return x.buf
}

// peerNotifySettings#9acda4c0 flags:# show_previews:flags.0?true silent:flags.1?true mute_until:int sound:string = PeerNotifySettings;

const crc_peerNotifySettings = 0x9acda4c0

type TL_peerNotifySettings struct {
	Flags         int32
	Show_previews bool   // show_previews:flags.0?true
	Silent        bool   // silent:flags.1?true
	Mute_until    int32  // mute_until:int
	Sound         string // sound:string
}

// Encoding TL_peerNotifySettings
func (e TL_peerNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerNotifySettings)
	var flags int32
	if e.Show_previews {
		flags |= (1 << 0)
	}
	if e.Silent {
		flags |= (1 << 1)
	}
	x.Int(flags)
	x.Int(e.Mute_until)
	x.String(e.Sound)
	return x.buf
}

// peerSettings#818426cd flags:# report_spam:flags.0?true = PeerSettings;

const crc_peerSettings = 0x818426cd

type TL_peerSettings struct {
	Flags       int32
	Report_spam bool // report_spam:flags.0?true
}

// Encoding TL_peerSettings
func (e TL_peerSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerSettings)
	var flags int32
	if e.Report_spam {
		flags |= (1 << 0)
	}
	x.Int(flags)
	return x.buf
}

// wallPaper#ccb03657 Id:int title:string sizes:Vector<PhotoSize> color:int = WallPaper;

const crc_wallPaper = 0xccb03657

type TL_wallPaper struct {
	Id    int32  // Id:int
	Title string // title:string
	Sizes []TL   // sizes:Vector<PhotoSize>
	Color int32  // color:int
}

// Encoding TL_wallPaper
func (e TL_wallPaper) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_wallPaper)
	x.Int(e.Id)
	x.String(e.Title)
	x.Vector(e.Sizes)
	x.Int(e.Color)
	return x.buf
}

// wallPaperSolid#63117f24 Id:int title:string bg_color:int color:int = WallPaper;

const crc_wallPaperSolid = 0x63117f24

type TL_wallPaperSolid struct {
	Id       int32  // Id:int
	Title    string // title:string
	Bg_color int32  // bg_color:int
	Color    int32  // color:int
}

// Encoding TL_wallPaperSolid
func (e TL_wallPaperSolid) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_wallPaperSolid)
	x.Int(e.Id)
	x.String(e.Title)
	x.Int(e.Bg_color)
	x.Int(e.Color)
	return x.buf
}

// inputReportReasonSpam#58dbcab8 = ReportReason;

const crc_inputReportReasonSpam = 0x58dbcab8

type TL_inputReportReasonSpam struct {
}

// Encoding TL_inputReportReasonSpam
func (e TL_inputReportReasonSpam) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputReportReasonSpam)
	return x.buf
}

// inputReportReasonViolence#1e22c78d = ReportReason;

const crc_inputReportReasonViolence = 0x1e22c78d

type TL_inputReportReasonViolence struct {
}

// Encoding TL_inputReportReasonViolence
func (e TL_inputReportReasonViolence) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputReportReasonViolence)
	return x.buf
}

// inputReportReasonPornography#2e59d922 = ReportReason;

const crc_inputReportReasonPornography = 0x2e59d922

type TL_inputReportReasonPornography struct {
}

// Encoding TL_inputReportReasonPornography
func (e TL_inputReportReasonPornography) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputReportReasonPornography)
	return x.buf
}

// inputReportReasonOther#e1746d0a text:string = ReportReason;

const crc_inputReportReasonOther = 0xe1746d0a

type TL_inputReportReasonOther struct {
	Text string // text:string
}

// Encoding TL_inputReportReasonOther
func (e TL_inputReportReasonOther) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputReportReasonOther)
	x.String(e.Text)
	return x.buf
}

// userFull#f220f3f flags:# blocked:flags.0?true phone_calls_available:flags.4?true phone_calls_private:flags.5?true user:User about:flags.1?string link:contacts.Link profile_photo:flags.2?Photo notify_settings:PeerNotifySettings bot_info:flags.3?BotInfo common_chats_count:int = UserFull;

const crc_userFull = 0xf220f3f

type TL_userFull struct {
	Flags                 int32
	Blocked               bool   // blocked:flags.0?true
	Phone_calls_available bool   // phone_calls_available:flags.4?true
	Phone_calls_private   bool   // phone_calls_private:flags.5?true
	User                  TL     // user:User
	About                 string // about:flags.1?string
	Link                  TL     // link:contacts.Link
	Profile_photo         TL     // profile_photo:flags.2?Photo
	Notify_settings       TL     // notify_settings:PeerNotifySettings
	Bot_info              TL     // bot_info:flags.3?BotInfo
	Common_chats_count    int32  // common_chats_count:int
}

// Encoding TL_userFull
func (e TL_userFull) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userFull)
	var flags int32
	if e.Blocked {
		flags |= (1 << 0)
	}
	if e.Phone_calls_available {
		flags |= (1 << 4)
	}
	if e.Phone_calls_private {
		flags |= (1 << 5)
	}
	if e.About != "" {
		flags |= (1 << 1)
	}
	if _, ok := (e.Profile_photo).(TL_null); !ok {
		flags |= (1 << 2)
	}
	if _, ok := (e.Bot_info).(TL_null); !ok {
		flags |= (1 << 3)
	}
	x.Int(flags)
	x.Bytes(e.User.encode())
	if flags&(1<<1) != 0 {
		x.String(e.About)
	}
	x.Bytes(e.Link.encode())
	if flags&(1<<2) != 0 {
		x.Bytes(e.Profile_photo.encode())
	}
	x.Bytes(e.Notify_settings.encode())
	if flags&(1<<3) != 0 {
		x.Bytes(e.Bot_info.encode())
	}
	x.Int(e.Common_chats_count)
	return x.buf
}

// contact#f911c994 user_id:int mutual:Bool = Contact;

const crc_contact = 0xf911c994

type TL_contact struct {
	User_id int32 // user_id:int
	Mutual  TL    // mutual:Bool
}

// Encoding TL_contact
func (e TL_contact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contact)
	x.Int(e.User_id)
	x.Bytes(e.Mutual.encode())
	return x.buf
}

// importedContact#d0028438 user_id:int client_id:long = ImportedContact;

const crc_importedContact = 0xd0028438

type TL_importedContact struct {
	User_id   int32 // user_id:int
	Client_id int64 // client_id:long
}

// Encoding TL_importedContact
func (e TL_importedContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_importedContact)
	x.Int(e.User_id)
	x.Long(e.Client_id)
	return x.buf
}

// contactBlocked#561bc879 user_id:int date:int = ContactBlocked;

const crc_contactBlocked = 0x561bc879

type TL_contactBlocked struct {
	User_id int32 // user_id:int
	Date    int32 // date:int
}

// Encoding TL_contactBlocked
func (e TL_contactBlocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactBlocked)
	x.Int(e.User_id)
	x.Int(e.Date)
	return x.buf
}

// contactStatus#d3680c61 user_id:int status:UserStatus = ContactStatus;

const crc_contactStatus = 0xd3680c61

type TL_contactStatus struct {
	User_id int32 // user_id:int
	Status  TL    // status:UserStatus
}

// Encoding TL_contactStatus
func (e TL_contactStatus) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactStatus)
	x.Int(e.User_id)
	x.Bytes(e.Status.encode())
	return x.buf
}

// contacts.link#3ace484c my_link:ContactLink foreign_link:ContactLink user:User = contacts.Link;

const crc_contacts_link = 0x3ace484c

type TL_contacts_link struct {
	My_link      TL // my_link:ContactLink
	Foreign_link TL // foreign_link:ContactLink
	User         TL // user:User
}

// Encoding TL_contacts_link
func (e TL_contacts_link) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_link)
	x.Bytes(e.My_link.encode())
	x.Bytes(e.Foreign_link.encode())
	x.Bytes(e.User.encode())
	return x.buf
}

// contacts.contactsNotModified#b74ba9d2 = contacts.Contacts;

const crc_contacts_contactsNotModified = 0xb74ba9d2

type TL_contacts_contactsNotModified struct {
}

// Encoding TL_contacts_contactsNotModified
func (e TL_contacts_contactsNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_contactsNotModified)
	return x.buf
}

// contacts.contacts#6f8b8cb2 contacts:Vector<Contact> users:Vector<User> = contacts.Contacts;

const crc_contacts_contacts = 0x6f8b8cb2

type TL_contacts_contacts struct {
	Contacts []TL // contacts:Vector<Contact>
	Users    []TL // users:Vector<User>
}

// Encoding TL_contacts_contacts
func (e TL_contacts_contacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_contacts)
	x.Vector(e.Contacts)
	x.Vector(e.Users)
	return x.buf
}

// contacts.importedContacts#ad524315 imported:Vector<ImportedContact> retry_contacts:Vector<long> users:Vector<User> = contacts.ImportedContacts;

const crc_contacts_importedContacts = 0xad524315

type TL_contacts_importedContacts struct {
	Imported       []TL    // imported:Vector<ImportedContact>
	Retry_contacts []int64 // retry_contacts:Vector<long>
	Users          []TL    // users:Vector<User>
}

// Encoding TL_contacts_importedContacts
func (e TL_contacts_importedContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_importedContacts)
	x.Vector(e.Imported)
	x.VectorLong(e.Retry_contacts)
	x.Vector(e.Users)
	return x.buf
}

// contacts.blocked#1c138d15 blocked:Vector<ContactBlocked> users:Vector<User> = contacts.Blocked;

const crc_contacts_blocked = 0x1c138d15

type TL_contacts_blocked struct {
	Blocked []TL // blocked:Vector<ContactBlocked>
	Users   []TL // users:Vector<User>
}

// Encoding TL_contacts_blocked
func (e TL_contacts_blocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_blocked)
	x.Vector(e.Blocked)
	x.Vector(e.Users)
	return x.buf
}

// contacts.blockedSlice#900802a1 count:int blocked:Vector<ContactBlocked> users:Vector<User> = contacts.Blocked;

const crc_contacts_blockedSlice = 0x900802a1

type TL_contacts_blockedSlice struct {
	Count   int32 // count:int
	Blocked []TL  // blocked:Vector<ContactBlocked>
	Users   []TL  // users:Vector<User>
}

// Encoding TL_contacts_blockedSlice
func (e TL_contacts_blockedSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_blockedSlice)
	x.Int(e.Count)
	x.Vector(e.Blocked)
	x.Vector(e.Users)
	return x.buf
}

// messages.dialogs#15ba6c40 dialogs:Vector<Dialog> messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Dialogs;

const crc_messages_dialogs = 0x15ba6c40

type TL_messages_dialogs struct {
	Dialogs  []TL // dialogs:Vector<Dialog>
	Messages []TL // messages:Vector<Message>
	Chats    []TL // chats:Vector<Chat>
	Users    []TL // users:Vector<User>
}

// Encoding TL_messages_dialogs
func (e TL_messages_dialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_dialogs)
	x.Vector(e.Dialogs)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

// messages.dialogsSlice#71e094f3 count:int dialogs:Vector<Dialog> messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Dialogs;

const crc_messages_dialogsSlice = 0x71e094f3

type TL_messages_dialogsSlice struct {
	Count    int32 // count:int
	Dialogs  []TL  // dialogs:Vector<Dialog>
	Messages []TL  // messages:Vector<Message>
	Chats    []TL  // chats:Vector<Chat>
	Users    []TL  // users:Vector<User>
}

// Encoding TL_messages_dialogsSlice
func (e TL_messages_dialogsSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_dialogsSlice)
	x.Int(e.Count)
	x.Vector(e.Dialogs)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

// messages.messages#8c718e87 messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;

const crc_messages_messages = 0x8c718e87

type TL_messages_messages struct {
	Messages []TL // messages:Vector<Message>
	Chats    []TL // chats:Vector<Chat>
	Users    []TL // users:Vector<User>
}

// Encoding TL_messages_messages
func (e TL_messages_messages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_messages)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

// messages.messagesSlice#b446ae3 count:int messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;

const crc_messages_messagesSlice = 0xb446ae3

type TL_messages_messagesSlice struct {
	Count    int32 // count:int
	Messages []TL  // messages:Vector<Message>
	Chats    []TL  // chats:Vector<Chat>
	Users    []TL  // users:Vector<User>
}

// Encoding TL_messages_messagesSlice
func (e TL_messages_messagesSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_messagesSlice)
	x.Int(e.Count)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

// messages.channelMessages#99262e37 flags:# pts:int count:int messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;

const crc_messages_channelMessages = 0x99262e37

type TL_messages_channelMessages struct {
	Flags    int32
	Pts      int32 // pts:int
	Count    int32 // count:int
	Messages []TL  // messages:Vector<Message>
	Chats    []TL  // chats:Vector<Chat>
	Users    []TL  // users:Vector<User>
}

// Encoding TL_messages_channelMessages
func (e TL_messages_channelMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_channelMessages)
	var flags int32
	x.Int(flags)
	x.Int(e.Pts)
	x.Int(e.Count)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

// messages.chats#64ff9fd5 chats:Vector<Chat> = messages.Chats;

const crc_messages_chats = 0x64ff9fd5

type TL_messages_chats struct {
	Chats []TL // chats:Vector<Chat>
}

// Encoding TL_messages_chats
func (e TL_messages_chats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_chats)
	x.Vector(e.Chats)
	return x.buf
}

// messages.chatsSlice#9cd81144 count:int chats:Vector<Chat> = messages.Chats;

const crc_messages_chatsSlice = 0x9cd81144

type TL_messages_chatsSlice struct {
	Count int32 // count:int
	Chats []TL  // chats:Vector<Chat>
}

// Encoding TL_messages_chatsSlice
func (e TL_messages_chatsSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_chatsSlice)
	x.Int(e.Count)
	x.Vector(e.Chats)
	return x.buf
}

// messages.chatFull#e5d7d19c full_chat:ChatFull chats:Vector<Chat> users:Vector<User> = messages.ChatFull;

const crc_messages_chatFull = 0xe5d7d19c

type TL_messages_chatFull struct {
	Full_chat TL   // full_chat:ChatFull
	Chats     []TL // chats:Vector<Chat>
	Users     []TL // users:Vector<User>
}

// Encoding TL_messages_chatFull
func (e TL_messages_chatFull) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_chatFull)
	x.Bytes(e.Full_chat.encode())
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

// messages.affectedHistory#b45c69d1 pts:int pts_count:int offset:int = messages.AffectedHistory;

const crc_messages_affectedHistory = 0xb45c69d1

type TL_messages_affectedHistory struct {
	Pts       int32 // pts:int
	Pts_count int32 // pts_count:int
	Offset    int32 // offset:int
}

// Encoding TL_messages_affectedHistory
func (e TL_messages_affectedHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_affectedHistory)
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	x.Int(e.Offset)
	return x.buf
}

// inputMessagesFilterEmpty#57e2f66c = MessagesFilter;

const crc_inputMessagesFilterEmpty = 0x57e2f66c

type TL_inputMessagesFilterEmpty struct {
}

// Encoding TL_inputMessagesFilterEmpty
func (e TL_inputMessagesFilterEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterEmpty)
	return x.buf
}

// inputMessagesFilterPhotos#9609a51c = MessagesFilter;

const crc_inputMessagesFilterPhotos = 0x9609a51c

type TL_inputMessagesFilterPhotos struct {
}

// Encoding TL_inputMessagesFilterPhotos
func (e TL_inputMessagesFilterPhotos) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterPhotos)
	return x.buf
}

// inputMessagesFilterVideo#9fc00e65 = MessagesFilter;

const crc_inputMessagesFilterVideo = 0x9fc00e65

type TL_inputMessagesFilterVideo struct {
}

// Encoding TL_inputMessagesFilterVideo
func (e TL_inputMessagesFilterVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterVideo)
	return x.buf
}

// inputMessagesFilterPhotoVideo#56e9f0e4 = MessagesFilter;

const crc_inputMessagesFilterPhotoVideo = 0x56e9f0e4

type TL_inputMessagesFilterPhotoVideo struct {
}

// Encoding TL_inputMessagesFilterPhotoVideo
func (e TL_inputMessagesFilterPhotoVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterPhotoVideo)
	return x.buf
}

// inputMessagesFilterPhotoVideoDocuments#d95e73bb = MessagesFilter;

const crc_inputMessagesFilterPhotoVideoDocuments = 0xd95e73bb

type TL_inputMessagesFilterPhotoVideoDocuments struct {
}

// Encoding TL_inputMessagesFilterPhotoVideoDocuments
func (e TL_inputMessagesFilterPhotoVideoDocuments) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterPhotoVideoDocuments)
	return x.buf
}

// inputMessagesFilterDocument#9eddf188 = MessagesFilter;

const crc_inputMessagesFilterDocument = 0x9eddf188

type TL_inputMessagesFilterDocument struct {
}

// Encoding TL_inputMessagesFilterDocument
func (e TL_inputMessagesFilterDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterDocument)
	return x.buf
}

// inputMessagesFilterUrl#7ef0dd87 = MessagesFilter;

const crc_inputMessagesFilterUrl = 0x7ef0dd87

type TL_inputMessagesFilterUrl struct {
}

// Encoding TL_inputMessagesFilterUrl
func (e TL_inputMessagesFilterUrl) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterUrl)
	return x.buf
}

// inputMessagesFilterGif#ffc86587 = MessagesFilter;

const crc_inputMessagesFilterGif = 0xffc86587

type TL_inputMessagesFilterGif struct {
}

// Encoding TL_inputMessagesFilterGif
func (e TL_inputMessagesFilterGif) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterGif)
	return x.buf
}

// inputMessagesFilterVoice#50f5c392 = MessagesFilter;

const crc_inputMessagesFilterVoice = 0x50f5c392

type TL_inputMessagesFilterVoice struct {
}

// Encoding TL_inputMessagesFilterVoice
func (e TL_inputMessagesFilterVoice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterVoice)
	return x.buf
}

// inputMessagesFilterMusic#3751b49e = MessagesFilter;

const crc_inputMessagesFilterMusic = 0x3751b49e

type TL_inputMessagesFilterMusic struct {
}

// Encoding TL_inputMessagesFilterMusic
func (e TL_inputMessagesFilterMusic) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterMusic)
	return x.buf
}

// inputMessagesFilterChatPhotos#3a20ecb8 = MessagesFilter;

const crc_inputMessagesFilterChatPhotos = 0x3a20ecb8

type TL_inputMessagesFilterChatPhotos struct {
}

// Encoding TL_inputMessagesFilterChatPhotos
func (e TL_inputMessagesFilterChatPhotos) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterChatPhotos)
	return x.buf
}

// inputMessagesFilterPhoneCalls#80c99768 flags:# missed:flags.0?true = MessagesFilter;

const crc_inputMessagesFilterPhoneCalls = 0x80c99768

type TL_inputMessagesFilterPhoneCalls struct {
	Flags  int32
	Missed bool // missed:flags.0?true
}

// Encoding TL_inputMessagesFilterPhoneCalls
func (e TL_inputMessagesFilterPhoneCalls) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterPhoneCalls)
	var flags int32
	if e.Missed {
		flags |= (1 << 0)
	}
	x.Int(flags)
	return x.buf
}

// updateNewMessage#1f2b0afd message:Message pts:int pts_count:int = Update;

const crc_updateNewMessage = 0x1f2b0afd

type TL_updateNewMessage struct {
	Message   TL    // message:Message
	Pts       int32 // pts:int
	Pts_count int32 // pts_count:int
}

// Encoding TL_updateNewMessage
func (e TL_updateNewMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateNewMessage)
	x.Bytes(e.Message.encode())
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	return x.buf
}

// updateMessageID#4e90bfd6 Id:int random_id:long = Update;

const crc_updateMessageID = 0x4e90bfd6

type TL_updateMessageID struct {
	Id        int32 // Id:int
	Random_id int64 // random_id:long
}

// Encoding TL_updateMessageID
func (e TL_updateMessageID) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateMessageID)
	x.Int(e.Id)
	x.Long(e.Random_id)
	return x.buf
}

// updateDeleteMessages#a20db0e5 messages:Vector<int> pts:int pts_count:int = Update;

const crc_updateDeleteMessages = 0xa20db0e5

type TL_updateDeleteMessages struct {
	Messages  []int32 // messages:Vector<int>
	Pts       int32   // pts:int
	Pts_count int32   // pts_count:int
}

// Encoding TL_updateDeleteMessages
func (e TL_updateDeleteMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateDeleteMessages)
	x.VectorInt(e.Messages)
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	return x.buf
}

// updateUserTyping#5c486927 user_id:int action:SendMessageAction = Update;

const crc_updateUserTyping = 0x5c486927

type TL_updateUserTyping struct {
	User_id int32 // user_id:int
	Action  TL    // action:SendMessageAction
}

// Encoding TL_updateUserTyping
func (e TL_updateUserTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateUserTyping)
	x.Int(e.User_id)
	x.Bytes(e.Action.encode())
	return x.buf
}

// updateChatUserTyping#9a65ea1f chat_id:int user_id:int action:SendMessageAction = Update;

const crc_updateChatUserTyping = 0x9a65ea1f

type TL_updateChatUserTyping struct {
	Chat_id int32 // chat_id:int
	User_id int32 // user_id:int
	Action  TL    // action:SendMessageAction
}

// Encoding TL_updateChatUserTyping
func (e TL_updateChatUserTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChatUserTyping)
	x.Int(e.Chat_id)
	x.Int(e.User_id)
	x.Bytes(e.Action.encode())
	return x.buf
}

// updateChatParticipants#7761198 participants:ChatParticipants = Update;

const crc_updateChatParticipants = 0x7761198

type TL_updateChatParticipants struct {
	Participants TL // participants:ChatParticipants
}

// Encoding TL_updateChatParticipants
func (e TL_updateChatParticipants) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChatParticipants)
	x.Bytes(e.Participants.encode())
	return x.buf
}

// updateUserStatus#1bfbd823 user_id:int status:UserStatus = Update;

const crc_updateUserStatus = 0x1bfbd823

type TL_updateUserStatus struct {
	User_id int32 // user_id:int
	Status  TL    // status:UserStatus
}

// Encoding TL_updateUserStatus
func (e TL_updateUserStatus) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateUserStatus)
	x.Int(e.User_id)
	x.Bytes(e.Status.encode())
	return x.buf
}

// updateUserName#a7332b73 user_id:int first_name:string last_name:string username:string = Update;

const crc_updateUserName = 0xa7332b73

type TL_updateUserName struct {
	User_id    int32  // user_id:int
	First_name string // first_name:string
	Last_name  string // last_name:string
	Username   string // username:string
}

// Encoding TL_updateUserName
func (e TL_updateUserName) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateUserName)
	x.Int(e.User_id)
	x.String(e.First_name)
	x.String(e.Last_name)
	x.String(e.Username)
	return x.buf
}

// updateUserPhoto#95313b0c user_id:int date:int photo:UserProfilePhoto previous:Bool = Update;

const crc_updateUserPhoto = 0x95313b0c

type TL_updateUserPhoto struct {
	User_id  int32 // user_id:int
	Date     int32 // date:int
	Photo    TL    // photo:UserProfilePhoto
	Previous TL    // previous:Bool
}

// Encoding TL_updateUserPhoto
func (e TL_updateUserPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateUserPhoto)
	x.Int(e.User_id)
	x.Int(e.Date)
	x.Bytes(e.Photo.encode())
	x.Bytes(e.Previous.encode())
	return x.buf
}

// updateContactRegistered#2575bbb9 user_id:int date:int = Update;

const crc_updateContactRegistered = 0x2575bbb9

type TL_updateContactRegistered struct {
	User_id int32 // user_id:int
	Date    int32 // date:int
}

// Encoding TL_updateContactRegistered
func (e TL_updateContactRegistered) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateContactRegistered)
	x.Int(e.User_id)
	x.Int(e.Date)
	return x.buf
}

// updateContactLink#9d2e67c5 user_id:int my_link:ContactLink foreign_link:ContactLink = Update;

const crc_updateContactLink = 0x9d2e67c5

type TL_updateContactLink struct {
	User_id      int32 // user_id:int
	My_link      TL    // my_link:ContactLink
	Foreign_link TL    // foreign_link:ContactLink
}

// Encoding TL_updateContactLink
func (e TL_updateContactLink) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateContactLink)
	x.Int(e.User_id)
	x.Bytes(e.My_link.encode())
	x.Bytes(e.Foreign_link.encode())
	return x.buf
}

// updateNewEncryptedMessage#12bcbd9a message:EncryptedMessage qts:int = Update;

const crc_updateNewEncryptedMessage = 0x12bcbd9a

type TL_updateNewEncryptedMessage struct {
	Message TL    // message:EncryptedMessage
	Qts     int32 // qts:int
}

// Encoding TL_updateNewEncryptedMessage
func (e TL_updateNewEncryptedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateNewEncryptedMessage)
	x.Bytes(e.Message.encode())
	x.Int(e.Qts)
	return x.buf
}

// updateEncryptedChatTyping#1710f156 chat_id:int = Update;

const crc_updateEncryptedChatTyping = 0x1710f156

type TL_updateEncryptedChatTyping struct {
	Chat_id int32 // chat_id:int
}

// Encoding TL_updateEncryptedChatTyping
func (e TL_updateEncryptedChatTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateEncryptedChatTyping)
	x.Int(e.Chat_id)
	return x.buf
}

// updateEncryption#b4a2e88d chat:EncryptedChat date:int = Update;

const crc_updateEncryption = 0xb4a2e88d

type TL_updateEncryption struct {
	Chat TL    // chat:EncryptedChat
	Date int32 // date:int
}

// Encoding TL_updateEncryption
func (e TL_updateEncryption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateEncryption)
	x.Bytes(e.Chat.encode())
	x.Int(e.Date)
	return x.buf
}

// updateEncryptedMessagesRead#38fe25b7 chat_id:int max_date:int date:int = Update;

const crc_updateEncryptedMessagesRead = 0x38fe25b7

type TL_updateEncryptedMessagesRead struct {
	Chat_id  int32 // chat_id:int
	Max_date int32 // max_date:int
	Date     int32 // date:int
}

// Encoding TL_updateEncryptedMessagesRead
func (e TL_updateEncryptedMessagesRead) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateEncryptedMessagesRead)
	x.Int(e.Chat_id)
	x.Int(e.Max_date)
	x.Int(e.Date)
	return x.buf
}

// updateChatParticipantAdd#ea4b0e5c chat_id:int user_id:int inviter_id:int date:int Version:int = Update;

const crc_updateChatParticipantAdd = 0xea4b0e5c

type TL_updateChatParticipantAdd struct {
	Chat_id    int32 // chat_id:int
	User_id    int32 // user_id:int
	Inviter_id int32 // inviter_id:int
	Date       int32 // date:int
	Version    int32 // Version:int
}

// Encoding TL_updateChatParticipantAdd
func (e TL_updateChatParticipantAdd) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChatParticipantAdd)
	x.Int(e.Chat_id)
	x.Int(e.User_id)
	x.Int(e.Inviter_id)
	x.Int(e.Date)
	x.Int(e.Version)
	return x.buf
}

// updateChatParticipantDelete#6e5f8c22 chat_id:int user_id:int Version:int = Update;

const crc_updateChatParticipantDelete = 0x6e5f8c22

type TL_updateChatParticipantDelete struct {
	Chat_id int32 // chat_id:int
	User_id int32 // user_id:int
	Version int32 // Version:int
}

// Encoding TL_updateChatParticipantDelete
func (e TL_updateChatParticipantDelete) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChatParticipantDelete)
	x.Int(e.Chat_id)
	x.Int(e.User_id)
	x.Int(e.Version)
	return x.buf
}

// updateDcOptions#8e5e9873 dc_options:Vector<DcOption> = Update;

const crc_updateDcOptions = 0x8e5e9873

type TL_updateDcOptions struct {
	Dc_options []TL // dc_options:Vector<DcOption>
}

// Encoding TL_updateDcOptions
func (e TL_updateDcOptions) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateDcOptions)
	x.Vector(e.Dc_options)
	return x.buf
}

// updateUserBlocked#80ece81a user_id:int blocked:Bool = Update;

const crc_updateUserBlocked = 0x80ece81a

type TL_updateUserBlocked struct {
	User_id int32 // user_id:int
	Blocked TL    // blocked:Bool
}

// Encoding TL_updateUserBlocked
func (e TL_updateUserBlocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateUserBlocked)
	x.Int(e.User_id)
	x.Bytes(e.Blocked.encode())
	return x.buf
}

// updateNotifySettings#bec268ef peer:NotifyPeer notify_settings:PeerNotifySettings = Update;

const crc_updateNotifySettings = 0xbec268ef

type TL_updateNotifySettings struct {
	Peer            TL // peer:NotifyPeer
	Notify_settings TL // notify_settings:PeerNotifySettings
}

// Encoding TL_updateNotifySettings
func (e TL_updateNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateNotifySettings)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Notify_settings.encode())
	return x.buf
}

// updateServiceNotification#ebe46819 flags:# popup:flags.0?true inbox_date:flags.1?int type:string message:string media:MessageMedia entities:Vector<MessageEntity> = Update;

const crc_updateServiceNotification = 0xebe46819

type TL_updateServiceNotification struct {
	Flags      int32
	Popup      bool   // popup:flags.0?true
	Inbox_date int32  // inbox_date:flags.1?int
	Code_type  string // type:string
	Message    string // message:string
	Media      TL     // media:MessageMedia
	Entities   []TL   // entities:Vector<MessageEntity>
}

// Encoding TL_updateServiceNotification
func (e TL_updateServiceNotification) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateServiceNotification)
	var flags int32
	if e.Popup {
		flags |= (1 << 0)
	}
	if e.Inbox_date > 0 {
		flags |= (1 << 1)
	}
	x.Int(flags)
	if flags&(1<<1) != 0 {
		x.Int(e.Inbox_date)
	}
	x.String(e.Code_type)
	x.String(e.Message)
	x.Bytes(e.Media.encode())
	x.Vector(e.Entities)
	return x.buf
}

// updatePrivacy#ee3b272a key:PrivacyKey rules:Vector<PrivacyRule> = Update;

const crc_updatePrivacy = 0xee3b272a

type TL_updatePrivacy struct {
	Key   TL   // key:PrivacyKey
	Rules []TL // rules:Vector<PrivacyRule>
}

// Encoding TL_updatePrivacy
func (e TL_updatePrivacy) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatePrivacy)
	x.Bytes(e.Key.encode())
	x.Vector(e.Rules)
	return x.buf
}

// updateUserPhone#12b9417b user_id:int phone:string = Update;

const crc_updateUserPhone = 0x12b9417b

type TL_updateUserPhone struct {
	User_id int32  // user_id:int
	Phone   string // phone:string
}

// Encoding TL_updateUserPhone
func (e TL_updateUserPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateUserPhone)
	x.Int(e.User_id)
	x.String(e.Phone)
	return x.buf
}

// updateReadHistoryInbox#9961fd5c peer:Peer max_id:int pts:int pts_count:int = Update;

const crc_updateReadHistoryInbox = 0x9961fd5c

type TL_updateReadHistoryInbox struct {
	Peer      TL    // peer:Peer
	Max_id    int32 // max_id:int
	Pts       int32 // pts:int
	Pts_count int32 // pts_count:int
}

// Encoding TL_updateReadHistoryInbox
func (e TL_updateReadHistoryInbox) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateReadHistoryInbox)
	x.Bytes(e.Peer.encode())
	x.Int(e.Max_id)
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	return x.buf
}

// updateReadHistoryOutbox#2f2f21bf peer:Peer max_id:int pts:int pts_count:int = Update;

const crc_updateReadHistoryOutbox = 0x2f2f21bf

type TL_updateReadHistoryOutbox struct {
	Peer      TL    // peer:Peer
	Max_id    int32 // max_id:int
	Pts       int32 // pts:int
	Pts_count int32 // pts_count:int
}

// Encoding TL_updateReadHistoryOutbox
func (e TL_updateReadHistoryOutbox) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateReadHistoryOutbox)
	x.Bytes(e.Peer.encode())
	x.Int(e.Max_id)
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	return x.buf
}

// updateWebPage#7f891213 webpage:WebPage pts:int pts_count:int = Update;

const crc_updateWebPage = 0x7f891213

type TL_updateWebPage struct {
	Webpage   TL    // webpage:WebPage
	Pts       int32 // pts:int
	Pts_count int32 // pts_count:int
}

// Encoding TL_updateWebPage
func (e TL_updateWebPage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateWebPage)
	x.Bytes(e.Webpage.encode())
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	return x.buf
}

// updateReadMessagesContents#68c13933 messages:Vector<int> pts:int pts_count:int = Update;

const crc_updateReadMessagesContents = 0x68c13933

type TL_updateReadMessagesContents struct {
	Messages  []int32 // messages:Vector<int>
	Pts       int32   // pts:int
	Pts_count int32   // pts_count:int
}

// Encoding TL_updateReadMessagesContents
func (e TL_updateReadMessagesContents) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateReadMessagesContents)
	x.VectorInt(e.Messages)
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	return x.buf
}

// updateChannelTooLong#eb0467fb flags:# channel_id:int pts:flags.0?int = Update;

const crc_updateChannelTooLong = 0xeb0467fb

type TL_updateChannelTooLong struct {
	Flags      int32
	Channel_id int32 // channel_id:int
	Pts        int32 // pts:flags.0?int
}

// Encoding TL_updateChannelTooLong
func (e TL_updateChannelTooLong) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChannelTooLong)
	var flags int32
	if e.Pts > 0 {
		flags |= (1 << 0)
	}
	x.Int(flags)
	x.Int(e.Channel_id)
	if flags&(1<<0) != 0 {
		x.Int(e.Pts)
	}
	return x.buf
}

// updateChannel#b6d45656 channel_id:int = Update;

const crc_updateChannel = 0xb6d45656

type TL_updateChannel struct {
	Channel_id int32 // channel_id:int
}

// Encoding TL_updateChannel
func (e TL_updateChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChannel)
	x.Int(e.Channel_id)
	return x.buf
}

// updateNewChannelMessage#62ba04d9 message:Message pts:int pts_count:int = Update;

const crc_updateNewChannelMessage = 0x62ba04d9

type TL_updateNewChannelMessage struct {
	Message   TL    // message:Message
	Pts       int32 // pts:int
	Pts_count int32 // pts_count:int
}

// Encoding TL_updateNewChannelMessage
func (e TL_updateNewChannelMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateNewChannelMessage)
	x.Bytes(e.Message.encode())
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	return x.buf
}

// updateReadChannelInbox#4214f37f channel_id:int max_id:int = Update;

const crc_updateReadChannelInbox = 0x4214f37f

type TL_updateReadChannelInbox struct {
	Channel_id int32 // channel_id:int
	Max_id     int32 // max_id:int
}

// Encoding TL_updateReadChannelInbox
func (e TL_updateReadChannelInbox) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateReadChannelInbox)
	x.Int(e.Channel_id)
	x.Int(e.Max_id)
	return x.buf
}

// updateDeleteChannelMessages#c37521c9 channel_id:int messages:Vector<int> pts:int pts_count:int = Update;

const crc_updateDeleteChannelMessages = 0xc37521c9

type TL_updateDeleteChannelMessages struct {
	Channel_id int32   // channel_id:int
	Messages   []int32 // messages:Vector<int>
	Pts        int32   // pts:int
	Pts_count  int32   // pts_count:int
}

// Encoding TL_updateDeleteChannelMessages
func (e TL_updateDeleteChannelMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateDeleteChannelMessages)
	x.Int(e.Channel_id)
	x.VectorInt(e.Messages)
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	return x.buf
}

// updateChannelMessageViews#98a12b4b channel_id:int Id:int views:int = Update;

const crc_updateChannelMessageViews = 0x98a12b4b

type TL_updateChannelMessageViews struct {
	Channel_id int32 // channel_id:int
	Id         int32 // Id:int
	Views      int32 // views:int
}

// Encoding TL_updateChannelMessageViews
func (e TL_updateChannelMessageViews) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChannelMessageViews)
	x.Int(e.Channel_id)
	x.Int(e.Id)
	x.Int(e.Views)
	return x.buf
}

// updateChatAdmins#6e947941 chat_id:int enabled:Bool Version:int = Update;

const crc_updateChatAdmins = 0x6e947941

type TL_updateChatAdmins struct {
	Chat_id int32 // chat_id:int
	Enabled TL    // enabled:Bool
	Version int32 // Version:int
}

// Encoding TL_updateChatAdmins
func (e TL_updateChatAdmins) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChatAdmins)
	x.Int(e.Chat_id)
	x.Bytes(e.Enabled.encode())
	x.Int(e.Version)
	return x.buf
}

// updateChatParticipantAdmin#b6901959 chat_id:int user_id:int is_admin:Bool Version:int = Update;

const crc_updateChatParticipantAdmin = 0xb6901959

type TL_updateChatParticipantAdmin struct {
	Chat_id  int32 // chat_id:int
	User_id  int32 // user_id:int
	Is_admin TL    // is_admin:Bool
	Version  int32 // Version:int
}

// Encoding TL_updateChatParticipantAdmin
func (e TL_updateChatParticipantAdmin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChatParticipantAdmin)
	x.Int(e.Chat_id)
	x.Int(e.User_id)
	x.Bytes(e.Is_admin.encode())
	x.Int(e.Version)
	return x.buf
}

// updateNewStickerSet#688a30aa stickerset:messages.StickerSet = Update;

const crc_updateNewStickerSet = 0x688a30aa

type TL_updateNewStickerSet struct {
	Stickerset TL // stickerset:messages.StickerSet
}

// Encoding TL_updateNewStickerSet
func (e TL_updateNewStickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateNewStickerSet)
	x.Bytes(e.Stickerset.encode())
	return x.buf
}

// updateStickerSetsOrder#bb2d201 flags:# masks:flags.0?true order:Vector<long> = Update;

const crc_updateStickerSetsOrder = 0xbb2d201

type TL_updateStickerSetsOrder struct {
	Flags int32
	Masks bool    // masks:flags.0?true
	Order []int64 // order:Vector<long>
}

// Encoding TL_updateStickerSetsOrder
func (e TL_updateStickerSetsOrder) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateStickerSetsOrder)
	var flags int32
	if e.Masks {
		flags |= (1 << 0)
	}
	x.Int(flags)
	x.VectorLong(e.Order)
	return x.buf
}

// updateStickerSets#43ae3dec = Update;

const crc_updateStickerSets = 0x43ae3dec

type TL_updateStickerSets struct {
}

// Encoding TL_updateStickerSets
func (e TL_updateStickerSets) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateStickerSets)
	return x.buf
}

// updateSavedGifs#9375341e = Update;

const crc_updateSavedGifs = 0x9375341e

type TL_updateSavedGifs struct {
}

// Encoding TL_updateSavedGifs
func (e TL_updateSavedGifs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateSavedGifs)
	return x.buf
}

// updateBotInlineQuery#54826690 flags:# query_id:long user_id:int query:string geo:flags.0?GeoPoint offset:string = Update;

const crc_updateBotInlineQuery = 0x54826690

type TL_updateBotInlineQuery struct {
	Flags    int32
	Query_id int64  // query_id:long
	User_id  int32  // user_id:int
	Query    string // query:string
	Geo      TL     // geo:flags.0?GeoPoint
	Offset   string // offset:string
}

// Encoding TL_updateBotInlineQuery
func (e TL_updateBotInlineQuery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateBotInlineQuery)
	var flags int32
	if _, ok := (e.Geo).(TL_null); !ok {
		flags |= (1 << 0)
	}
	x.Int(flags)
	x.Long(e.Query_id)
	x.Int(e.User_id)
	x.String(e.Query)
	if flags&(1<<0) != 0 {
		x.Bytes(e.Geo.encode())
	}
	x.String(e.Offset)
	return x.buf
}

// updateBotInlineSend#e48f964 flags:# user_id:int query:string geo:flags.0?GeoPoint Id:string msg_id:flags.1?InputBotInlineMessageID = Update;

const crc_updateBotInlineSend = 0xe48f964

type TL_updateBotInlineSend struct {
	Flags   int32
	User_id int32  // user_id:int
	Query   string // query:string
	Geo     TL     // geo:flags.0?GeoPoint
	Id      string // Id:string
	Msg_id  TL     // msg_id:flags.1?InputBotInlineMessageID
}

// Encoding TL_updateBotInlineSend
func (e TL_updateBotInlineSend) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateBotInlineSend)
	var flags int32
	if _, ok := (e.Geo).(TL_null); !ok {
		flags |= (1 << 0)
	}
	if _, ok := (e.Msg_id).(TL_null); !ok {
		flags |= (1 << 1)
	}
	x.Int(flags)
	x.Int(e.User_id)
	x.String(e.Query)
	if flags&(1<<0) != 0 {
		x.Bytes(e.Geo.encode())
	}
	x.String(e.Id)
	if flags&(1<<1) != 0 {
		x.Bytes(e.Msg_id.encode())
	}
	return x.buf
}

// updateEditChannelMessage#1b3f4df7 message:Message pts:int pts_count:int = Update;

const crc_updateEditChannelMessage = 0x1b3f4df7

type TL_updateEditChannelMessage struct {
	Message   TL    // message:Message
	Pts       int32 // pts:int
	Pts_count int32 // pts_count:int
}

// Encoding TL_updateEditChannelMessage
func (e TL_updateEditChannelMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateEditChannelMessage)
	x.Bytes(e.Message.encode())
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	return x.buf
}

// updateChannelPinnedMessage#98592475 channel_id:int Id:int = Update;

const crc_updateChannelPinnedMessage = 0x98592475

type TL_updateChannelPinnedMessage struct {
	Channel_id int32 // channel_id:int
	Id         int32 // Id:int
}

// Encoding TL_updateChannelPinnedMessage
func (e TL_updateChannelPinnedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChannelPinnedMessage)
	x.Int(e.Channel_id)
	x.Int(e.Id)
	return x.buf
}

// updateBotCallbackQuery#e73547e1 flags:# query_id:long user_id:int peer:Peer msg_id:int chat_instance:long data:flags.0?bytes game_short_name:flags.1?string = Update;

const crc_updateBotCallbackQuery = 0xe73547e1

type TL_updateBotCallbackQuery struct {
	Flags           int32
	Query_id        int64  // query_id:long
	User_id         int32  // user_id:int
	Peer            TL     // peer:Peer
	Msg_id          int32  // msg_id:int
	Chat_instance   int64  // chat_instance:long
	Data            []byte // data:flags.0?bytes
	Game_short_name string // game_short_name:flags.1?string
}

// Encoding TL_updateBotCallbackQuery
func (e TL_updateBotCallbackQuery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateBotCallbackQuery)
	var flags int32
	if len(e.Data) != 0 {
		flags |= (1 << 0)
	}
	if e.Game_short_name != "" {
		flags |= (1 << 1)
	}
	x.Int(flags)
	x.Long(e.Query_id)
	x.Int(e.User_id)
	x.Bytes(e.Peer.encode())
	x.Int(e.Msg_id)
	x.Long(e.Chat_instance)
	if flags&(1<<0) != 0 {
		x.Bytes(e.Data)
	}
	if flags&(1<<1) != 0 {
		x.String(e.Game_short_name)
	}
	return x.buf
}

// updateEditMessage#e40370a3 message:Message pts:int pts_count:int = Update;

const crc_updateEditMessage = 0xe40370a3

type TL_updateEditMessage struct {
	Message   TL    // message:Message
	Pts       int32 // pts:int
	Pts_count int32 // pts_count:int
}

// Encoding TL_updateEditMessage
func (e TL_updateEditMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateEditMessage)
	x.Bytes(e.Message.encode())
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	return x.buf
}

// updateInlineBotCallbackQuery#f9d27a5a flags:# query_id:long user_id:int msg_id:InputBotInlineMessageID chat_instance:long data:flags.0?bytes game_short_name:flags.1?string = Update;

const crc_updateInlineBotCallbackQuery = 0xf9d27a5a

type TL_updateInlineBotCallbackQuery struct {
	Flags           int32
	Query_id        int64  // query_id:long
	User_id         int32  // user_id:int
	Msg_id          TL     // msg_id:InputBotInlineMessageID
	Chat_instance   int64  // chat_instance:long
	Data            []byte // data:flags.0?bytes
	Game_short_name string // game_short_name:flags.1?string
}

// Encoding TL_updateInlineBotCallbackQuery
func (e TL_updateInlineBotCallbackQuery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateInlineBotCallbackQuery)
	var flags int32
	if len(e.Data) != 0 {
		flags |= (1 << 0)
	}
	if e.Game_short_name != "" {
		flags |= (1 << 1)
	}
	x.Int(flags)
	x.Long(e.Query_id)
	x.Int(e.User_id)
	x.Bytes(e.Msg_id.encode())
	x.Long(e.Chat_instance)
	if flags&(1<<0) != 0 {
		x.Bytes(e.Data)
	}
	if flags&(1<<1) != 0 {
		x.String(e.Game_short_name)
	}
	return x.buf
}

// updateReadChannelOutbox#25d6c9c7 channel_id:int max_id:int = Update;

const crc_updateReadChannelOutbox = 0x25d6c9c7

type TL_updateReadChannelOutbox struct {
	Channel_id int32 // channel_id:int
	Max_id     int32 // max_id:int
}

// Encoding TL_updateReadChannelOutbox
func (e TL_updateReadChannelOutbox) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateReadChannelOutbox)
	x.Int(e.Channel_id)
	x.Int(e.Max_id)
	return x.buf
}

// updateDraftMessage#ee2bb969 peer:Peer draft:DraftMessage = Update;

const crc_updateDraftMessage = 0xee2bb969

type TL_updateDraftMessage struct {
	Peer  TL // peer:Peer
	Draft TL // draft:DraftMessage
}

// Encoding TL_updateDraftMessage
func (e TL_updateDraftMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateDraftMessage)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Draft.encode())
	return x.buf
}

// updateReadFeaturedStickers#571d2742 = Update;

const crc_updateReadFeaturedStickers = 0x571d2742

type TL_updateReadFeaturedStickers struct {
}

// Encoding TL_updateReadFeaturedStickers
func (e TL_updateReadFeaturedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateReadFeaturedStickers)
	return x.buf
}

// updateRecentStickers#9a422c20 = Update;

const crc_updateRecentStickers = 0x9a422c20

type TL_updateRecentStickers struct {
}

// Encoding TL_updateRecentStickers
func (e TL_updateRecentStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateRecentStickers)
	return x.buf
}

// updateConfig#a229dd06 = Update;

const crc_updateConfig = 0xa229dd06

type TL_updateConfig struct {
}

// Encoding TL_updateConfig
func (e TL_updateConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateConfig)
	return x.buf
}

// updatePtsChanged#3354678f = Update;

const crc_updatePtsChanged = 0x3354678f

type TL_updatePtsChanged struct {
}

// Encoding TL_updatePtsChanged
func (e TL_updatePtsChanged) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatePtsChanged)
	return x.buf
}

// updateChannelWebPage#40771900 channel_id:int webpage:WebPage pts:int pts_count:int = Update;

const crc_updateChannelWebPage = 0x40771900

type TL_updateChannelWebPage struct {
	Channel_id int32 // channel_id:int
	Webpage    TL    // webpage:WebPage
	Pts        int32 // pts:int
	Pts_count  int32 // pts_count:int
}

// Encoding TL_updateChannelWebPage
func (e TL_updateChannelWebPage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChannelWebPage)
	x.Int(e.Channel_id)
	x.Bytes(e.Webpage.encode())
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	return x.buf
}

// updateDialogPinned#d711a2cc flags:# pinned:flags.0?true peer:Peer = Update;

const crc_updateDialogPinned = 0xd711a2cc

type TL_updateDialogPinned struct {
	Flags  int32
	Pinned bool // pinned:flags.0?true
	Peer   TL   // peer:Peer
}

// Encoding TL_updateDialogPinned
func (e TL_updateDialogPinned) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateDialogPinned)
	var flags int32
	if e.Pinned {
		flags |= (1 << 0)
	}
	x.Int(flags)
	x.Bytes(e.Peer.encode())
	return x.buf
}

// updatePinnedDialogs#d8caf68d flags:# order:flags.0?Vector<Peer> = Update;

const crc_updatePinnedDialogs = 0xd8caf68d

type TL_updatePinnedDialogs struct {
	Flags int32
	Order []TL // order:flags.0?Vector<Peer>
}

// Encoding TL_updatePinnedDialogs
func (e TL_updatePinnedDialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatePinnedDialogs)
	var flags int32
	if len(e.Order) != 0 {
		flags |= (1 << 0)
	}
	x.Int(flags)
	if flags&(1<<0) != 0 {
		x.Vector(e.Order)
	}
	return x.buf
}

// updateBotWebhookJSON#8317c0c3 data:DataJSON = Update;

const crc_updateBotWebhookJSON = 0x8317c0c3

type TL_updateBotWebhookJSON struct {
	Data TL // data:DataJSON
}

// Encoding TL_updateBotWebhookJSON
func (e TL_updateBotWebhookJSON) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateBotWebhookJSON)
	x.Bytes(e.Data.encode())
	return x.buf
}

// updateBotWebhookJSONQuery#9b9240a6 query_id:long data:DataJSON timeout:int = Update;

const crc_updateBotWebhookJSONQuery = 0x9b9240a6

type TL_updateBotWebhookJSONQuery struct {
	Query_id int64 // query_id:long
	Data     TL    // data:DataJSON
	Timeout  int32 // timeout:int
}

// Encoding TL_updateBotWebhookJSONQuery
func (e TL_updateBotWebhookJSONQuery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateBotWebhookJSONQuery)
	x.Long(e.Query_id)
	x.Bytes(e.Data.encode())
	x.Int(e.Timeout)
	return x.buf
}

// updateBotShippingQuery#e0cdc940 query_id:long user_id:int payload:bytes shipping_address:PostAddress = Update;

const crc_updateBotShippingQuery = 0xe0cdc940

type TL_updateBotShippingQuery struct {
	Query_id         int64  // query_id:long
	User_id          int32  // user_id:int
	Payload          []byte // payload:bytes
	Shipping_address TL     // shipping_address:PostAddress
}

// Encoding TL_updateBotShippingQuery
func (e TL_updateBotShippingQuery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateBotShippingQuery)
	x.Long(e.Query_id)
	x.Int(e.User_id)
	x.Bytes(e.Payload)
	x.Bytes(e.Shipping_address.encode())
	return x.buf
}

// updateBotPrecheckoutQuery#5d2f3aa9 flags:# query_id:long user_id:int payload:bytes info:flags.0?PaymentRequestedInfo shipping_option_id:flags.1?string currency:string total_amount:long = Update;

const crc_updateBotPrecheckoutQuery = 0x5d2f3aa9

type TL_updateBotPrecheckoutQuery struct {
	Flags              int32
	Query_id           int64  // query_id:long
	User_id            int32  // user_id:int
	Payload            []byte // payload:bytes
	Info               TL     // info:flags.0?PaymentRequestedInfo
	Shipping_option_id string // shipping_option_id:flags.1?string
	Currency           string // currency:string
	Total_amount       int64  // total_amount:long
}

// Encoding TL_updateBotPrecheckoutQuery
func (e TL_updateBotPrecheckoutQuery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateBotPrecheckoutQuery)
	var flags int32
	if _, ok := (e.Info).(TL_null); !ok {
		flags |= (1 << 0)
	}
	if e.Shipping_option_id != "" {
		flags |= (1 << 1)
	}
	x.Int(flags)
	x.Long(e.Query_id)
	x.Int(e.User_id)
	x.Bytes(e.Payload)
	if flags&(1<<0) != 0 {
		x.Bytes(e.Info.encode())
	}
	if flags&(1<<1) != 0 {
		x.String(e.Shipping_option_id)
	}
	x.String(e.Currency)
	x.Long(e.Total_amount)
	return x.buf
}

// updatePhoneCall#ab0f6b1e phone_call:PhoneCall = Update;

const crc_updatePhoneCall = 0xab0f6b1e

type TL_updatePhoneCall struct {
	Phone_call TL // phone_call:PhoneCall
}

// Encoding TL_updatePhoneCall
func (e TL_updatePhoneCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatePhoneCall)
	x.Bytes(e.Phone_call.encode())
	return x.buf
}

// updates.state#a56c2a3e pts:int qts:int date:int seq:int unread_count:int = updates.State;

const crc_updates_state = 0xa56c2a3e

type TL_updates_state struct {
	Pts          int32 // pts:int
	Qts          int32 // qts:int
	Date         int32 // date:int
	Seq          int32 // seq:int
	Unread_count int32 // unread_count:int
}

// Encoding TL_updates_state
func (e TL_updates_state) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_state)
	x.Int(e.Pts)
	x.Int(e.Qts)
	x.Int(e.Date)
	x.Int(e.Seq)
	x.Int(e.Unread_count)
	return x.buf
}

// updates.differenceEmpty#5d75a138 date:int seq:int = updates.Difference;

const crc_updates_differenceEmpty = 0x5d75a138

type TL_updates_differenceEmpty struct {
	Date int32 // date:int
	Seq  int32 // seq:int
}

// Encoding TL_updates_differenceEmpty
func (e TL_updates_differenceEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_differenceEmpty)
	x.Int(e.Date)
	x.Int(e.Seq)
	return x.buf
}

// updates.difference#f49ca0 new_messages:Vector<Message> new_encrypted_messages:Vector<EncryptedMessage> other_updates:Vector<Update> chats:Vector<Chat> users:Vector<User> state:updates.State = updates.Difference;

const crc_updates_difference = 0xf49ca0

type TL_updates_difference struct {
	New_messages           []TL // new_messages:Vector<Message>
	New_encrypted_messages []TL // new_encrypted_messages:Vector<EncryptedMessage>
	Other_updates          []TL // other_updates:Vector<Update>
	Chats                  []TL // chats:Vector<Chat>
	Users                  []TL // users:Vector<User>
	State                  TL   // state:updates.State
}

// Encoding TL_updates_difference
func (e TL_updates_difference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_difference)
	x.Vector(e.New_messages)
	x.Vector(e.New_encrypted_messages)
	x.Vector(e.Other_updates)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	x.Bytes(e.State.encode())
	return x.buf
}

// updates.differenceSlice#a8fb1981 new_messages:Vector<Message> new_encrypted_messages:Vector<EncryptedMessage> other_updates:Vector<Update> chats:Vector<Chat> users:Vector<User> intermediate_state:updates.State = updates.Difference;

const crc_updates_differenceSlice = 0xa8fb1981

type TL_updates_differenceSlice struct {
	New_messages           []TL // new_messages:Vector<Message>
	New_encrypted_messages []TL // new_encrypted_messages:Vector<EncryptedMessage>
	Other_updates          []TL // other_updates:Vector<Update>
	Chats                  []TL // chats:Vector<Chat>
	Users                  []TL // users:Vector<User>
	Intermediate_state     TL   // intermediate_state:updates.State
}

// Encoding TL_updates_differenceSlice
func (e TL_updates_differenceSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_differenceSlice)
	x.Vector(e.New_messages)
	x.Vector(e.New_encrypted_messages)
	x.Vector(e.Other_updates)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	x.Bytes(e.Intermediate_state.encode())
	return x.buf
}

// updates.differenceTooLong#4afe8f6d pts:int = updates.Difference;

const crc_updates_differenceTooLong = 0x4afe8f6d

type TL_updates_differenceTooLong struct {
	Pts int32 // pts:int
}

// Encoding TL_updates_differenceTooLong
func (e TL_updates_differenceTooLong) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_differenceTooLong)
	x.Int(e.Pts)
	return x.buf
}

// updatesTooLong#e317af7e = Updates;

const crc_updatesTooLong = 0xe317af7e

type TL_updatesTooLong struct {
}

// Encoding TL_updatesTooLong
func (e TL_updatesTooLong) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatesTooLong)
	return x.buf
}

// updateShortMessage#914fbf11 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true Id:int user_id:int message:string pts:int pts_count:int date:int fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to_msg_id:flags.3?int entities:flags.7?Vector<MessageEntity> = Updates;

const crc_updateShortMessage = 0x914fbf11

type TL_updateShortMessage struct {
	Flags           int32
	Out             bool   // out:flags.1?true
	Mentioned       bool   // mentioned:flags.4?true
	Media_unread    bool   // media_unread:flags.5?true
	Silent          bool   // silent:flags.13?true
	Id              int32  // Id:int
	User_id         int32  // user_id:int
	Message         string // message:string
	Pts             int32  // pts:int
	Pts_count       int32  // pts_count:int
	Date            int32  // date:int
	Fwd_from        TL     // fwd_from:flags.2?MessageFwdHeader
	Via_bot_id      int32  // via_bot_id:flags.11?int
	Reply_to_msg_id int32  // reply_to_msg_id:flags.3?int
	Entities        []TL   // entities:flags.7?Vector<MessageEntity>
}

// Encoding TL_updateShortMessage
func (e TL_updateShortMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateShortMessage)
	var flags int32
	if e.Out {
		flags |= (1 << 1)
	}
	if e.Mentioned {
		flags |= (1 << 4)
	}
	if e.Media_unread {
		flags |= (1 << 5)
	}
	if e.Silent {
		flags |= (1 << 13)
	}
	if _, ok := (e.Fwd_from).(TL_null); !ok {
		flags |= (1 << 2)
	}
	if e.Via_bot_id > 0 {
		flags |= (1 << 11)
	}
	if e.Reply_to_msg_id > 0 {
		flags |= (1 << 3)
	}
	if len(e.Entities) != 0 {
		flags |= (1 << 7)
	}
	x.Int(flags)
	x.Int(e.Id)
	x.Int(e.User_id)
	x.String(e.Message)
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	x.Int(e.Date)
	if flags&(1<<2) != 0 {
		x.Bytes(e.Fwd_from.encode())
	}
	if flags&(1<<11) != 0 {
		x.Int(e.Via_bot_id)
	}
	if flags&(1<<3) != 0 {
		x.Int(e.Reply_to_msg_id)
	}
	if flags&(1<<7) != 0 {
		x.Vector(e.Entities)
	}
	return x.buf
}

// updateShortChatMessage#16812688 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true Id:int from_id:int chat_id:int message:string pts:int pts_count:int date:int fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to_msg_id:flags.3?int entities:flags.7?Vector<MessageEntity> = Updates;

const crc_updateShortChatMessage = 0x16812688

type TL_updateShortChatMessage struct {
	Flags           int32
	Out             bool   // out:flags.1?true
	Mentioned       bool   // mentioned:flags.4?true
	Media_unread    bool   // media_unread:flags.5?true
	Silent          bool   // silent:flags.13?true
	Id              int32  // Id:int
	From_id         int32  // from_id:int
	Chat_id         int32  // chat_id:int
	Message         string // message:string
	Pts             int32  // pts:int
	Pts_count       int32  // pts_count:int
	Date            int32  // date:int
	Fwd_from        TL     // fwd_from:flags.2?MessageFwdHeader
	Via_bot_id      int32  // via_bot_id:flags.11?int
	Reply_to_msg_id int32  // reply_to_msg_id:flags.3?int
	Entities        []TL   // entities:flags.7?Vector<MessageEntity>
}

// Encoding TL_updateShortChatMessage
func (e TL_updateShortChatMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateShortChatMessage)
	var flags int32
	if e.Out {
		flags |= (1 << 1)
	}
	if e.Mentioned {
		flags |= (1 << 4)
	}
	if e.Media_unread {
		flags |= (1 << 5)
	}
	if e.Silent {
		flags |= (1 << 13)
	}
	if _, ok := (e.Fwd_from).(TL_null); !ok {
		flags |= (1 << 2)
	}
	if e.Via_bot_id > 0 {
		flags |= (1 << 11)
	}
	if e.Reply_to_msg_id > 0 {
		flags |= (1 << 3)
	}
	if len(e.Entities) != 0 {
		flags |= (1 << 7)
	}
	x.Int(flags)
	x.Int(e.Id)
	x.Int(e.From_id)
	x.Int(e.Chat_id)
	x.String(e.Message)
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	x.Int(e.Date)
	if flags&(1<<2) != 0 {
		x.Bytes(e.Fwd_from.encode())
	}
	if flags&(1<<11) != 0 {
		x.Int(e.Via_bot_id)
	}
	if flags&(1<<3) != 0 {
		x.Int(e.Reply_to_msg_id)
	}
	if flags&(1<<7) != 0 {
		x.Vector(e.Entities)
	}
	return x.buf
}

// updateShort#78d4dec1 update:Update date:int = Updates;

const crc_updateShort = 0x78d4dec1

type TL_updateShort struct {
	Update TL    // update:Update
	Date   int32 // date:int
}

// Encoding TL_updateShort
func (e TL_updateShort) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateShort)
	x.Bytes(e.Update.encode())
	x.Int(e.Date)
	return x.buf
}

// updatesCombined#725b04c3 updates:Vector<Update> users:Vector<User> chats:Vector<Chat> date:int seq_start:int seq:int = Updates;

const crc_updatesCombined = 0x725b04c3

type TL_updatesCombined struct {
	Updates   []TL  // updates:Vector<Update>
	Users     []TL  // users:Vector<User>
	Chats     []TL  // chats:Vector<Chat>
	Date      int32 // date:int
	Seq_start int32 // seq_start:int
	Seq       int32 // seq:int
}

// Encoding TL_updatesCombined
func (e TL_updatesCombined) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatesCombined)
	x.Vector(e.Updates)
	x.Vector(e.Users)
	x.Vector(e.Chats)
	x.Int(e.Date)
	x.Int(e.Seq_start)
	x.Int(e.Seq)
	return x.buf
}

// updates#74ae4240 updates:Vector<Update> users:Vector<User> chats:Vector<Chat> date:int seq:int = Updates;

const crc_updates = 0x74ae4240

type TL_updates struct {
	Updates []TL  // updates:Vector<Update>
	Users   []TL  // users:Vector<User>
	Chats   []TL  // chats:Vector<Chat>
	Date    int32 // date:int
	Seq     int32 // seq:int
}

// Encoding TL_updates
func (e TL_updates) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates)
	x.Vector(e.Updates)
	x.Vector(e.Users)
	x.Vector(e.Chats)
	x.Int(e.Date)
	x.Int(e.Seq)
	return x.buf
}

// updateShortSentMessage#11f1331c flags:# out:flags.1?true Id:int pts:int pts_count:int date:int media:flags.9?MessageMedia entities:flags.7?Vector<MessageEntity> = Updates;

const crc_updateShortSentMessage = 0x11f1331c

type TL_updateShortSentMessage struct {
	Flags     int32
	Out       bool  // out:flags.1?true
	Id        int32 // Id:int
	Pts       int32 // pts:int
	Pts_count int32 // pts_count:int
	Date      int32 // date:int
	Media     TL    // media:flags.9?MessageMedia
	Entities  []TL  // entities:flags.7?Vector<MessageEntity>
}

// Encoding TL_updateShortSentMessage
func (e TL_updateShortSentMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateShortSentMessage)
	var flags int32
	if e.Out {
		flags |= (1 << 1)
	}
	if _, ok := (e.Media).(TL_null); !ok {
		flags |= (1 << 9)
	}
	if len(e.Entities) != 0 {
		flags |= (1 << 7)
	}
	x.Int(flags)
	x.Int(e.Id)
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	x.Int(e.Date)
	if flags&(1<<9) != 0 {
		x.Bytes(e.Media.encode())
	}
	if flags&(1<<7) != 0 {
		x.Vector(e.Entities)
	}
	return x.buf
}

// photos.photos#8dca6aa5 photos:Vector<Photo> users:Vector<User> = photos.Photos;

const crc_photos_photos = 0x8dca6aa5

type TL_photos_photos struct {
	Photos []TL // photos:Vector<Photo>
	Users  []TL // users:Vector<User>
}

// Encoding TL_photos_photos
func (e TL_photos_photos) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photos_photos)
	x.Vector(e.Photos)
	x.Vector(e.Users)
	return x.buf
}

// photos.photosSlice#15051f54 count:int photos:Vector<Photo> users:Vector<User> = photos.Photos;

const crc_photos_photosSlice = 0x15051f54

type TL_photos_photosSlice struct {
	Count  int32 // count:int
	Photos []TL  // photos:Vector<Photo>
	Users  []TL  // users:Vector<User>
}

// Encoding TL_photos_photosSlice
func (e TL_photos_photosSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photos_photosSlice)
	x.Int(e.Count)
	x.Vector(e.Photos)
	x.Vector(e.Users)
	return x.buf
}

// photos.photo#20212ca8 photo:Photo users:Vector<User> = photos.Photo;

const crc_photos_photo = 0x20212ca8

type TL_photos_photo struct {
	Photo TL   // photo:Photo
	Users []TL // users:Vector<User>
}

// Encoding TL_photos_photo
func (e TL_photos_photo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photos_photo)
	x.Bytes(e.Photo.encode())
	x.Vector(e.Users)
	return x.buf
}

// upload.file#96a18d5 type:storage.FileType mtime:int bytes:bytes = upload.File;

const crc_upload_file = 0x96a18d5

type TL_upload_file struct {
	Code_type TL     // type:storage.FileType
	Mtime     int32  // mtime:int
	Bytes     []byte // bytes:bytes
}

// Encoding TL_upload_file
func (e TL_upload_file) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_file)
	x.Bytes(e.Code_type.encode())
	x.Int(e.Mtime)
	x.Bytes(e.Bytes)
	return x.buf
}

// dcOption#5d8c6cc flags:# ipv6:flags.0?true media_only:flags.1?true tcpo_only:flags.2?true Id:int ip_address:string port:int = DcOption;

const crc_dcOption = 0x5d8c6cc

type TL_dcOption struct {
	Flags      int32
	Ipv6       bool   // ipv6:flags.0?true
	Media_only bool   // media_only:flags.1?true
	Tcpo_only  bool   // tcpo_only:flags.2?true
	Id         int32  // Id:int
	Ip_address string // ip_address:string
	Port       int32  // port:int
}

// Encoding TL_dcOption
func (e TL_dcOption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_dcOption)
	var flags int32
	if e.Ipv6 {
		flags |= (1 << 0)
	}
	if e.Media_only {
		flags |= (1 << 1)
	}
	if e.Tcpo_only {
		flags |= (1 << 2)
	}
	x.Int(flags)
	x.Int(e.Id)
	x.String(e.Ip_address)
	x.Int(e.Port)
	return x.buf
}

// config#cb601684 flags:# phonecalls_enabled:flags.1?true date:int expires:int test_mode:Bool this_dc:int dc_options:Vector<DcOption> chat_size_max:int megagroup_size_max:int forwarded_count_max:int online_update_period_ms:int offline_blur_timeout_ms:int offline_idle_timeout_ms:int online_cloud_timeout_ms:int notify_cloud_delay_ms:int notify_default_delay_ms:int chat_big_size:int push_chat_period_ms:int push_chat_limit:int saved_gifs_limit:int edit_time_limit:int rating_e_decay:int stickers_recent_limit:int tmp_sessions:flags.0?int pinned_dialogs_count_max:int call_receive_timeout_ms:int call_ring_timeout_ms:int call_connect_timeout_ms:int call_packet_timeout_ms:int me_url_prefix:string disabled_features:Vector<DisabledFeature> = Config;

const crc_config = 0xcb601684

type TL_config struct {
	Flags                    int32
	Phonecalls_enabled       bool   // phonecalls_enabled:flags.1?true
	Date                     int32  // date:int
	Expires                  int32  // expires:int
	Test_mode                TL     // test_mode:Bool
	This_dc                  int32  // this_dc:int
	Dc_options               []TL   // dc_options:Vector<DcOption>
	Chat_size_max            int32  // chat_size_max:int
	Megagroup_size_max       int32  // megagroup_size_max:int
	Forwarded_count_max      int32  // forwarded_count_max:int
	Online_update_period_ms  int32  // online_update_period_ms:int
	Offline_blur_timeout_ms  int32  // offline_blur_timeout_ms:int
	Offline_idle_timeout_ms  int32  // offline_idle_timeout_ms:int
	Online_cloud_timeout_ms  int32  // online_cloud_timeout_ms:int
	Notify_cloud_delay_ms    int32  // notify_cloud_delay_ms:int
	Notify_default_delay_ms  int32  // notify_default_delay_ms:int
	Chat_big_size            int32  // chat_big_size:int
	Push_chat_period_ms      int32  // push_chat_period_ms:int
	Push_chat_limit          int32  // push_chat_limit:int
	Saved_gifs_limit         int32  // saved_gifs_limit:int
	Edit_time_limit          int32  // edit_time_limit:int
	Rating_e_decay           int32  // rating_e_decay:int
	Stickers_recent_limit    int32  // stickers_recent_limit:int
	Tmp_sessions             int32  // tmp_sessions:flags.0?int
	Pinned_dialogs_count_max int32  // pinned_dialogs_count_max:int
	Call_receive_timeout_ms  int32  // call_receive_timeout_ms:int
	Call_ring_timeout_ms     int32  // call_ring_timeout_ms:int
	Call_connect_timeout_ms  int32  // call_connect_timeout_ms:int
	Call_packet_timeout_ms   int32  // call_packet_timeout_ms:int
	Me_url_prefix            string // me_url_prefix:string
	Disabled_features        []TL   // disabled_features:Vector<DisabledFeature>
}

// Encoding TL_config
func (e TL_config) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_config)
	var flags int32
	if e.Phonecalls_enabled {
		flags |= (1 << 1)
	}
	if e.Tmp_sessions > 0 {
		flags |= (1 << 0)
	}
	x.Int(flags)
	x.Int(e.Date)
	x.Int(e.Expires)
	x.Bytes(e.Test_mode.encode())
	x.Int(e.This_dc)
	x.Vector(e.Dc_options)
	x.Int(e.Chat_size_max)
	x.Int(e.Megagroup_size_max)
	x.Int(e.Forwarded_count_max)
	x.Int(e.Online_update_period_ms)
	x.Int(e.Offline_blur_timeout_ms)
	x.Int(e.Offline_idle_timeout_ms)
	x.Int(e.Online_cloud_timeout_ms)
	x.Int(e.Notify_cloud_delay_ms)
	x.Int(e.Notify_default_delay_ms)
	x.Int(e.Chat_big_size)
	x.Int(e.Push_chat_period_ms)
	x.Int(e.Push_chat_limit)
	x.Int(e.Saved_gifs_limit)
	x.Int(e.Edit_time_limit)
	x.Int(e.Rating_e_decay)
	x.Int(e.Stickers_recent_limit)
	if flags&(1<<0) != 0 {
		x.Int(e.Tmp_sessions)
	}
	x.Int(e.Pinned_dialogs_count_max)
	x.Int(e.Call_receive_timeout_ms)
	x.Int(e.Call_ring_timeout_ms)
	x.Int(e.Call_connect_timeout_ms)
	x.Int(e.Call_packet_timeout_ms)
	x.String(e.Me_url_prefix)
	x.Vector(e.Disabled_features)
	return x.buf
}

// nearestDc#8e1a1775 country:string this_dc:int nearest_dc:int = NearestDc;

const crc_nearestDc = 0x8e1a1775

type TL_nearestDc struct {
	Country    string // country:string
	This_dc    int32  // this_dc:int
	Nearest_dc int32  // nearest_dc:int
}

// Encoding TL_nearestDc
func (e TL_nearestDc) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_nearestDc)
	x.String(e.Country)
	x.Int(e.This_dc)
	x.Int(e.Nearest_dc)
	return x.buf
}

// help.appUpdate#8987f311 Id:int critical:Bool url:string text:string = help.AppUpdate;

const crc_help_appUpdate = 0x8987f311

type TL_help_appUpdate struct {
	Id       int32  // Id:int
	Critical TL     // critical:Bool
	Url      string // url:string
	Text     string // text:string
}

// Encoding TL_help_appUpdate
func (e TL_help_appUpdate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_appUpdate)
	x.Int(e.Id)
	x.Bytes(e.Critical.encode())
	x.String(e.Url)
	x.String(e.Text)
	return x.buf
}

// help.noAppUpdate#c45a6536 = help.AppUpdate;

const crc_help_noAppUpdate = 0xc45a6536

type TL_help_noAppUpdate struct {
}

// Encoding TL_help_noAppUpdate
func (e TL_help_noAppUpdate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_noAppUpdate)
	return x.buf
}

// help.inviteText#18cb9f78 message:string = help.InviteText;

const crc_help_inviteText = 0x18cb9f78

type TL_help_inviteText struct {
	Message string // message:string
}

// Encoding TL_help_inviteText
func (e TL_help_inviteText) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_inviteText)
	x.String(e.Message)
	return x.buf
}

// encryptedChatEmpty#ab7ec0a0 Id:int = EncryptedChat;

const crc_encryptedChatEmpty = 0xab7ec0a0

type TL_encryptedChatEmpty struct {
	Id int32 // Id:int
}

// Encoding TL_encryptedChatEmpty
func (e TL_encryptedChatEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedChatEmpty)
	x.Int(e.Id)
	return x.buf
}

// encryptedChatWaiting#3bf703dc Id:int access_hash:long date:int admin_id:int participant_id:int = EncryptedChat;

const crc_encryptedChatWaiting = 0x3bf703dc

type TL_encryptedChatWaiting struct {
	Id             int32 // Id:int
	Access_hash    int64 // access_hash:long
	Date           int32 // date:int
	Admin_id       int32 // admin_id:int
	Participant_id int32 // participant_id:int
}

// Encoding TL_encryptedChatWaiting
func (e TL_encryptedChatWaiting) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedChatWaiting)
	x.Int(e.Id)
	x.Long(e.Access_hash)
	x.Int(e.Date)
	x.Int(e.Admin_id)
	x.Int(e.Participant_id)
	return x.buf
}

// encryptedChatRequested#c878527e Id:int access_hash:long date:int admin_id:int participant_id:int g_a:bytes = EncryptedChat;

const crc_encryptedChatRequested = 0xc878527e

type TL_encryptedChatRequested struct {
	Id             int32  // Id:int
	Access_hash    int64  // access_hash:long
	Date           int32  // date:int
	Admin_id       int32  // admin_id:int
	Participant_id int32  // participant_id:int
	G_a            []byte // g_a:bytes
}

// Encoding TL_encryptedChatRequested
func (e TL_encryptedChatRequested) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedChatRequested)
	x.Int(e.Id)
	x.Long(e.Access_hash)
	x.Int(e.Date)
	x.Int(e.Admin_id)
	x.Int(e.Participant_id)
	x.Bytes(e.G_a)
	return x.buf
}

// encryptedChat#fa56ce36 Id:int access_hash:long date:int admin_id:int participant_id:int g_a_or_b:bytes key_fingerprint:long = EncryptedChat;

const crc_encryptedChat = 0xfa56ce36

type TL_encryptedChat struct {
	Id              int32  // Id:int
	Access_hash     int64  // access_hash:long
	Date            int32  // date:int
	Admin_id        int32  // admin_id:int
	Participant_id  int32  // participant_id:int
	G_a_or_b        []byte // g_a_or_b:bytes
	Key_fingerprint int64  // key_fingerprint:long
}

// Encoding TL_encryptedChat
func (e TL_encryptedChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedChat)
	x.Int(e.Id)
	x.Long(e.Access_hash)
	x.Int(e.Date)
	x.Int(e.Admin_id)
	x.Int(e.Participant_id)
	x.Bytes(e.G_a_or_b)
	x.Long(e.Key_fingerprint)
	return x.buf
}

// encryptedChatDiscarded#13d6dd27 Id:int = EncryptedChat;

const crc_encryptedChatDiscarded = 0x13d6dd27

type TL_encryptedChatDiscarded struct {
	Id int32 // Id:int
}

// Encoding TL_encryptedChatDiscarded
func (e TL_encryptedChatDiscarded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedChatDiscarded)
	x.Int(e.Id)
	return x.buf
}

// inputEncryptedChat#f141b5e1 chat_id:int access_hash:long = InputEncryptedChat;

const crc_inputEncryptedChat = 0xf141b5e1

type TL_inputEncryptedChat struct {
	Chat_id     int32 // chat_id:int
	Access_hash int64 // access_hash:long
}

// Encoding TL_inputEncryptedChat
func (e TL_inputEncryptedChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputEncryptedChat)
	x.Int(e.Chat_id)
	x.Long(e.Access_hash)
	return x.buf
}

// encryptedFileEmpty#c21f497e = EncryptedFile;

const crc_encryptedFileEmpty = 0xc21f497e

type TL_encryptedFileEmpty struct {
}

// Encoding TL_encryptedFileEmpty
func (e TL_encryptedFileEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedFileEmpty)
	return x.buf
}

// encryptedFile#4a70994c Id:long access_hash:long size:int dc_id:int key_fingerprint:int = EncryptedFile;

const crc_encryptedFile = 0x4a70994c

type TL_encryptedFile struct {
	Id              int64 // Id:long
	Access_hash     int64 // access_hash:long
	Size            int32 // size:int
	Dc_id           int32 // dc_id:int
	Key_fingerprint int32 // key_fingerprint:int
}

// Encoding TL_encryptedFile
func (e TL_encryptedFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedFile)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	x.Int(e.Size)
	x.Int(e.Dc_id)
	x.Int(e.Key_fingerprint)
	return x.buf
}

// inputEncryptedFileEmpty#1837c364 = InputEncryptedFile;

const crc_inputEncryptedFileEmpty = 0x1837c364

type TL_inputEncryptedFileEmpty struct {
}

// Encoding TL_inputEncryptedFileEmpty
func (e TL_inputEncryptedFileEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputEncryptedFileEmpty)
	return x.buf
}

// inputEncryptedFileUploaded#64bd0306 Id:long parts:int md5_checksum:string key_fingerprint:int = InputEncryptedFile;

const crc_inputEncryptedFileUploaded = 0x64bd0306

type TL_inputEncryptedFileUploaded struct {
	Id              int64  // Id:long
	Parts           int32  // parts:int
	Md5_checksum    string // md5_checksum:string
	Key_fingerprint int32  // key_fingerprint:int
}

// Encoding TL_inputEncryptedFileUploaded
func (e TL_inputEncryptedFileUploaded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputEncryptedFileUploaded)
	x.Long(e.Id)
	x.Int(e.Parts)
	x.String(e.Md5_checksum)
	x.Int(e.Key_fingerprint)
	return x.buf
}

// inputEncryptedFile#5a17b5e5 Id:long access_hash:long = InputEncryptedFile;

const crc_inputEncryptedFile = 0x5a17b5e5

type TL_inputEncryptedFile struct {
	Id          int64 // Id:long
	Access_hash int64 // access_hash:long
}

// Encoding TL_inputEncryptedFile
func (e TL_inputEncryptedFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputEncryptedFile)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	return x.buf
}

// inputEncryptedFileBigUploaded#2dc173c8 Id:long parts:int key_fingerprint:int = InputEncryptedFile;

const crc_inputEncryptedFileBigUploaded = 0x2dc173c8

type TL_inputEncryptedFileBigUploaded struct {
	Id              int64 // Id:long
	Parts           int32 // parts:int
	Key_fingerprint int32 // key_fingerprint:int
}

// Encoding TL_inputEncryptedFileBigUploaded
func (e TL_inputEncryptedFileBigUploaded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputEncryptedFileBigUploaded)
	x.Long(e.Id)
	x.Int(e.Parts)
	x.Int(e.Key_fingerprint)
	return x.buf
}

// encryptedMessage#ed18c118 random_id:long chat_id:int date:int bytes:bytes file:EncryptedFile = EncryptedMessage;

const crc_encryptedMessage = 0xed18c118

type TL_encryptedMessage struct {
	Random_id int64  // random_id:long
	Chat_id   int32  // chat_id:int
	Date      int32  // date:int
	Bytes     []byte // bytes:bytes
	File      TL     // file:EncryptedFile
}

// Encoding TL_encryptedMessage
func (e TL_encryptedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedMessage)
	x.Long(e.Random_id)
	x.Int(e.Chat_id)
	x.Int(e.Date)
	x.Bytes(e.Bytes)
	x.Bytes(e.File.encode())
	return x.buf
}

// encryptedMessageService#23734b06 random_id:long chat_id:int date:int bytes:bytes = EncryptedMessage;

const crc_encryptedMessageService = 0x23734b06

type TL_encryptedMessageService struct {
	Random_id int64  // random_id:long
	Chat_id   int32  // chat_id:int
	Date      int32  // date:int
	Bytes     []byte // bytes:bytes
}

// Encoding TL_encryptedMessageService
func (e TL_encryptedMessageService) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedMessageService)
	x.Long(e.Random_id)
	x.Int(e.Chat_id)
	x.Int(e.Date)
	x.Bytes(e.Bytes)
	return x.buf
}

// messages.dhConfigNotModified#c0e24635 random:bytes = messages.DhConfig;

const crc_messages_dhConfigNotModified = 0xc0e24635

type TL_messages_dhConfigNotModified struct {
	Random []byte // random:bytes
}

// Encoding TL_messages_dhConfigNotModified
func (e TL_messages_dhConfigNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_dhConfigNotModified)
	x.Bytes(e.Random)
	return x.buf
}

// messages.dhConfig#2c221edd g:int p:bytes Version:int random:bytes = messages.DhConfig;

const crc_messages_dhConfig = 0x2c221edd

type TL_messages_dhConfig struct {
	G       int32  // g:int
	P       []byte // p:bytes
	Version int32  // Version:int
	Random  []byte // random:bytes
}

// Encoding TL_messages_dhConfig
func (e TL_messages_dhConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_dhConfig)
	x.Int(e.G)
	x.Bytes(e.P)
	x.Int(e.Version)
	x.Bytes(e.Random)
	return x.buf
}

// messages.sentEncryptedMessage#560f8935 date:int = messages.SentEncryptedMessage;

const crc_messages_sentEncryptedMessage = 0x560f8935

type TL_messages_sentEncryptedMessage struct {
	Date int32 // date:int
}

// Encoding TL_messages_sentEncryptedMessage
func (e TL_messages_sentEncryptedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sentEncryptedMessage)
	x.Int(e.Date)
	return x.buf
}

// messages.sentEncryptedFile#9493ff32 date:int file:EncryptedFile = messages.SentEncryptedMessage;

const crc_messages_sentEncryptedFile = 0x9493ff32

type TL_messages_sentEncryptedFile struct {
	Date int32 // date:int
	File TL    // file:EncryptedFile
}

// Encoding TL_messages_sentEncryptedFile
func (e TL_messages_sentEncryptedFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sentEncryptedFile)
	x.Int(e.Date)
	x.Bytes(e.File.encode())
	return x.buf
}

// inputDocumentEmpty#72f0eaae = InputDocument;

const crc_inputDocumentEmpty = 0x72f0eaae

type TL_inputDocumentEmpty struct {
}

// Encoding TL_inputDocumentEmpty
func (e TL_inputDocumentEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputDocumentEmpty)
	return x.buf
}

// inputDocument#18798952 Id:long access_hash:long = InputDocument;

const crc_inputDocument = 0x18798952

type TL_inputDocument struct {
	Id          int64 // Id:long
	Access_hash int64 // access_hash:long
}

// Encoding TL_inputDocument
func (e TL_inputDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputDocument)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	return x.buf
}

// documentEmpty#36f8c871 Id:long = Document;

const crc_documentEmpty = 0x36f8c871

type TL_documentEmpty struct {
	Id int64 // Id:long
}

// Encoding TL_documentEmpty
func (e TL_documentEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentEmpty)
	x.Long(e.Id)
	return x.buf
}

// document#87232bc7 Id:long access_hash:long date:int mime_type:string size:int thumb:PhotoSize dc_id:int Version:int attributes:Vector<DocumentAttribute> = Document;

const crc_document = 0x87232bc7

type TL_document struct {
	Id          int64  // Id:long
	Access_hash int64  // access_hash:long
	Date        int32  // date:int
	Mime_type   string // mime_type:string
	Size        int32  // size:int
	Thumb       TL     // thumb:PhotoSize
	Dc_id       int32  // dc_id:int
	Version     int32  // Version:int
	Attributes  []TL   // attributes:Vector<DocumentAttribute>
}

// Encoding TL_document
func (e TL_document) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_document)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	x.Int(e.Date)
	x.String(e.Mime_type)
	x.Int(e.Size)
	x.Bytes(e.Thumb.encode())
	x.Int(e.Dc_id)
	x.Int(e.Version)
	x.Vector(e.Attributes)
	return x.buf
}

// help.support#17c6b5f6 phone_number:string user:User = help.Support;

const crc_help_support = 0x17c6b5f6

type TL_help_support struct {
	Phone_number string // phone_number:string
	User         TL     // user:User
}

// Encoding TL_help_support
func (e TL_help_support) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_support)
	x.String(e.Phone_number)
	x.Bytes(e.User.encode())
	return x.buf
}

// notifyPeer#9fd40bd8 peer:Peer = NotifyPeer;

const crc_notifyPeer = 0x9fd40bd8

type TL_notifyPeer struct {
	Peer TL // peer:Peer
}

// Encoding TL_notifyPeer
func (e TL_notifyPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_notifyPeer)
	x.Bytes(e.Peer.encode())
	return x.buf
}

// notifyUsers#b4c83b4c = NotifyPeer;

const crc_notifyUsers = 0xb4c83b4c

type TL_notifyUsers struct {
}

// Encoding TL_notifyUsers
func (e TL_notifyUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_notifyUsers)
	return x.buf
}

// notifyChats#c007cec3 = NotifyPeer;

const crc_notifyChats = 0xc007cec3

type TL_notifyChats struct {
}

// Encoding TL_notifyChats
func (e TL_notifyChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_notifyChats)
	return x.buf
}

// notifyAll#74d07c60 = NotifyPeer;

const crc_notifyAll = 0x74d07c60

type TL_notifyAll struct {
}

// Encoding TL_notifyAll
func (e TL_notifyAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_notifyAll)
	return x.buf
}

// sendMessageTypingAction#16bf744e = SendMessageAction;

const crc_sendMessageTypingAction = 0x16bf744e

type TL_sendMessageTypingAction struct {
}

// Encoding TL_sendMessageTypingAction
func (e TL_sendMessageTypingAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageTypingAction)
	return x.buf
}

// sendMessageCancelAction#fd5ec8f5 = SendMessageAction;

const crc_sendMessageCancelAction = 0xfd5ec8f5

type TL_sendMessageCancelAction struct {
}

// Encoding TL_sendMessageCancelAction
func (e TL_sendMessageCancelAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageCancelAction)
	return x.buf
}

// sendMessageRecordVideoAction#a187d66f = SendMessageAction;

const crc_sendMessageRecordVideoAction = 0xa187d66f

type TL_sendMessageRecordVideoAction struct {
}

// Encoding TL_sendMessageRecordVideoAction
func (e TL_sendMessageRecordVideoAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageRecordVideoAction)
	return x.buf
}

// sendMessageUploadVideoAction#e9763aec progress:int = SendMessageAction;

const crc_sendMessageUploadVideoAction = 0xe9763aec

type TL_sendMessageUploadVideoAction struct {
	Progress int32 // progress:int
}

// Encoding TL_sendMessageUploadVideoAction
func (e TL_sendMessageUploadVideoAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageUploadVideoAction)
	x.Int(e.Progress)
	return x.buf
}

// sendMessageRecordAudioAction#d52f73f7 = SendMessageAction;

const crc_sendMessageRecordAudioAction = 0xd52f73f7

type TL_sendMessageRecordAudioAction struct {
}

// Encoding TL_sendMessageRecordAudioAction
func (e TL_sendMessageRecordAudioAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageRecordAudioAction)
	return x.buf
}

// sendMessageUploadAudioAction#f351d7ab progress:int = SendMessageAction;

const crc_sendMessageUploadAudioAction = 0xf351d7ab

type TL_sendMessageUploadAudioAction struct {
	Progress int32 // progress:int
}

// Encoding TL_sendMessageUploadAudioAction
func (e TL_sendMessageUploadAudioAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageUploadAudioAction)
	x.Int(e.Progress)
	return x.buf
}

// sendMessageUploadPhotoAction#d1d34a26 progress:int = SendMessageAction;

const crc_sendMessageUploadPhotoAction = 0xd1d34a26

type TL_sendMessageUploadPhotoAction struct {
	Progress int32 // progress:int
}

// Encoding TL_sendMessageUploadPhotoAction
func (e TL_sendMessageUploadPhotoAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageUploadPhotoAction)
	x.Int(e.Progress)
	return x.buf
}

// sendMessageUploadDocumentAction#aa0cd9e4 progress:int = SendMessageAction;

const crc_sendMessageUploadDocumentAction = 0xaa0cd9e4

type TL_sendMessageUploadDocumentAction struct {
	Progress int32 // progress:int
}

// Encoding TL_sendMessageUploadDocumentAction
func (e TL_sendMessageUploadDocumentAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageUploadDocumentAction)
	x.Int(e.Progress)
	return x.buf
}

// sendMessageGeoLocationAction#176f8ba1 = SendMessageAction;

const crc_sendMessageGeoLocationAction = 0x176f8ba1

type TL_sendMessageGeoLocationAction struct {
}

// Encoding TL_sendMessageGeoLocationAction
func (e TL_sendMessageGeoLocationAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageGeoLocationAction)
	return x.buf
}

// sendMessageChooseContactAction#628cbc6f = SendMessageAction;

const crc_sendMessageChooseContactAction = 0x628cbc6f

type TL_sendMessageChooseContactAction struct {
}

// Encoding TL_sendMessageChooseContactAction
func (e TL_sendMessageChooseContactAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageChooseContactAction)
	return x.buf
}

// sendMessageGamePlayAction#dd6a8f48 = SendMessageAction;

const crc_sendMessageGamePlayAction = 0xdd6a8f48

type TL_sendMessageGamePlayAction struct {
}

// Encoding TL_sendMessageGamePlayAction
func (e TL_sendMessageGamePlayAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageGamePlayAction)
	return x.buf
}

// contacts.found#1aa1f784 results:Vector<Peer> chats:Vector<Chat> users:Vector<User> = contacts.Found;

const crc_contacts_found = 0x1aa1f784

type TL_contacts_found struct {
	Results []TL // results:Vector<Peer>
	Chats   []TL // chats:Vector<Chat>
	Users   []TL // users:Vector<User>
}

// Encoding TL_contacts_found
func (e TL_contacts_found) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_found)
	x.Vector(e.Results)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

// inputPrivacyKeyStatusTimestamp#4f96cb18 = InputPrivacyKey;

const crc_inputPrivacyKeyStatusTimestamp = 0x4f96cb18

type TL_inputPrivacyKeyStatusTimestamp struct {
}

// Encoding TL_inputPrivacyKeyStatusTimestamp
func (e TL_inputPrivacyKeyStatusTimestamp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyKeyStatusTimestamp)
	return x.buf
}

// inputPrivacyKeyChatInvite#bdfb0426 = InputPrivacyKey;

const crc_inputPrivacyKeyChatInvite = 0xbdfb0426

type TL_inputPrivacyKeyChatInvite struct {
}

// Encoding TL_inputPrivacyKeyChatInvite
func (e TL_inputPrivacyKeyChatInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyKeyChatInvite)
	return x.buf
}

// inputPrivacyKeyPhoneCall#fabadc5f = InputPrivacyKey;

const crc_inputPrivacyKeyPhoneCall = 0xfabadc5f

type TL_inputPrivacyKeyPhoneCall struct {
}

// Encoding TL_inputPrivacyKeyPhoneCall
func (e TL_inputPrivacyKeyPhoneCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyKeyPhoneCall)
	return x.buf
}

// privacyKeyStatusTimestamp#bc2eab30 = PrivacyKey;

const crc_privacyKeyStatusTimestamp = 0xbc2eab30

type TL_privacyKeyStatusTimestamp struct {
}

// Encoding TL_privacyKeyStatusTimestamp
func (e TL_privacyKeyStatusTimestamp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyKeyStatusTimestamp)
	return x.buf
}

// privacyKeyChatInvite#500e6dfa = PrivacyKey;

const crc_privacyKeyChatInvite = 0x500e6dfa

type TL_privacyKeyChatInvite struct {
}

// Encoding TL_privacyKeyChatInvite
func (e TL_privacyKeyChatInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyKeyChatInvite)
	return x.buf
}

// privacyKeyPhoneCall#3d662b7b = PrivacyKey;

const crc_privacyKeyPhoneCall = 0x3d662b7b

type TL_privacyKeyPhoneCall struct {
}

// Encoding TL_privacyKeyPhoneCall
func (e TL_privacyKeyPhoneCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyKeyPhoneCall)
	return x.buf
}

// inputPrivacyValueAllowContacts#d09e07b = InputPrivacyRule;

const crc_inputPrivacyValueAllowContacts = 0xd09e07b

type TL_inputPrivacyValueAllowContacts struct {
}

// Encoding TL_inputPrivacyValueAllowContacts
func (e TL_inputPrivacyValueAllowContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyValueAllowContacts)
	return x.buf
}

// inputPrivacyValueAllowAll#184b35ce = InputPrivacyRule;

const crc_inputPrivacyValueAllowAll = 0x184b35ce

type TL_inputPrivacyValueAllowAll struct {
}

// Encoding TL_inputPrivacyValueAllowAll
func (e TL_inputPrivacyValueAllowAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyValueAllowAll)
	return x.buf
}

// inputPrivacyValueAllowUsers#131cc67f users:Vector<InputUser> = InputPrivacyRule;

const crc_inputPrivacyValueAllowUsers = 0x131cc67f

type TL_inputPrivacyValueAllowUsers struct {
	Users []TL // users:Vector<InputUser>
}

// Encoding TL_inputPrivacyValueAllowUsers
func (e TL_inputPrivacyValueAllowUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyValueAllowUsers)
	x.Vector(e.Users)
	return x.buf
}

// inputPrivacyValueDisallowContacts#ba52007 = InputPrivacyRule;

const crc_inputPrivacyValueDisallowContacts = 0xba52007

type TL_inputPrivacyValueDisallowContacts struct {
}

// Encoding TL_inputPrivacyValueDisallowContacts
func (e TL_inputPrivacyValueDisallowContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyValueDisallowContacts)
	return x.buf
}

// inputPrivacyValueDisallowAll#d66b66c9 = InputPrivacyRule;

const crc_inputPrivacyValueDisallowAll = 0xd66b66c9

type TL_inputPrivacyValueDisallowAll struct {
}

// Encoding TL_inputPrivacyValueDisallowAll
func (e TL_inputPrivacyValueDisallowAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyValueDisallowAll)
	return x.buf
}

// inputPrivacyValueDisallowUsers#90110467 users:Vector<InputUser> = InputPrivacyRule;

const crc_inputPrivacyValueDisallowUsers = 0x90110467

type TL_inputPrivacyValueDisallowUsers struct {
	Users []TL // users:Vector<InputUser>
}

// Encoding TL_inputPrivacyValueDisallowUsers
func (e TL_inputPrivacyValueDisallowUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyValueDisallowUsers)
	x.Vector(e.Users)
	return x.buf
}

// privacyValueAllowContacts#fffe1bac = PrivacyRule;

const crc_privacyValueAllowContacts = 0xfffe1bac

type TL_privacyValueAllowContacts struct {
}

// Encoding TL_privacyValueAllowContacts
func (e TL_privacyValueAllowContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyValueAllowContacts)
	return x.buf
}

// privacyValueAllowAll#65427b82 = PrivacyRule;

const crc_privacyValueAllowAll = 0x65427b82

type TL_privacyValueAllowAll struct {
}

// Encoding TL_privacyValueAllowAll
func (e TL_privacyValueAllowAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyValueAllowAll)
	return x.buf
}

// privacyValueAllowUsers#4d5bbe0c users:Vector<int> = PrivacyRule;

const crc_privacyValueAllowUsers = 0x4d5bbe0c

type TL_privacyValueAllowUsers struct {
	Users []int32 // users:Vector<int>
}

// Encoding TL_privacyValueAllowUsers
func (e TL_privacyValueAllowUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyValueAllowUsers)
	x.VectorInt(e.Users)
	return x.buf
}

// privacyValueDisallowContacts#f888fa1a = PrivacyRule;

const crc_privacyValueDisallowContacts = 0xf888fa1a

type TL_privacyValueDisallowContacts struct {
}

// Encoding TL_privacyValueDisallowContacts
func (e TL_privacyValueDisallowContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyValueDisallowContacts)
	return x.buf
}

// privacyValueDisallowAll#8b73e763 = PrivacyRule;

const crc_privacyValueDisallowAll = 0x8b73e763

type TL_privacyValueDisallowAll struct {
}

// Encoding TL_privacyValueDisallowAll
func (e TL_privacyValueDisallowAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyValueDisallowAll)
	return x.buf
}

// privacyValueDisallowUsers#c7f49b7 users:Vector<int> = PrivacyRule;

const crc_privacyValueDisallowUsers = 0xc7f49b7

type TL_privacyValueDisallowUsers struct {
	Users []int32 // users:Vector<int>
}

// Encoding TL_privacyValueDisallowUsers
func (e TL_privacyValueDisallowUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyValueDisallowUsers)
	x.VectorInt(e.Users)
	return x.buf
}

// account.privacyRules#554abb6f rules:Vector<PrivacyRule> users:Vector<User> = account.PrivacyRules;

const crc_account_privacyRules = 0x554abb6f

type TL_account_privacyRules struct {
	Rules []TL // rules:Vector<PrivacyRule>
	Users []TL // users:Vector<User>
}

// Encoding TL_account_privacyRules
func (e TL_account_privacyRules) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_privacyRules)
	x.Vector(e.Rules)
	x.Vector(e.Users)
	return x.buf
}

// accountDaysTTL#b8d0afdf days:int = AccountDaysTTL;

const crc_accountDaysTTL = 0xb8d0afdf

type TL_accountDaysTTL struct {
	Days int32 // days:int
}

// Encoding TL_accountDaysTTL
func (e TL_accountDaysTTL) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountDaysTTL)
	x.Int(e.Days)
	return x.buf
}

// documentAttributeImageSize#6c37c15c w:int h:int = DocumentAttribute;

const crc_documentAttributeImageSize = 0x6c37c15c

type TL_documentAttributeImageSize struct {
	W int32 // w:int
	H int32 // h:int
}

// Encoding TL_documentAttributeImageSize
func (e TL_documentAttributeImageSize) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeImageSize)
	x.Int(e.W)
	x.Int(e.H)
	return x.buf
}

// documentAttributeAnimated#11b58939 = DocumentAttribute;

const crc_documentAttributeAnimated = 0x11b58939

type TL_documentAttributeAnimated struct {
}

// Encoding TL_documentAttributeAnimated
func (e TL_documentAttributeAnimated) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeAnimated)
	return x.buf
}

// documentAttributeSticker#6319d612 flags:# mask:flags.1?true alt:string stickerset:InputStickerSet mask_coords:flags.0?MaskCoords = DocumentAttribute;

const crc_documentAttributeSticker = 0x6319d612

type TL_documentAttributeSticker struct {
	Flags       int32
	Mask        bool   // mask:flags.1?true
	Alt         string // alt:string
	Stickerset  TL     // stickerset:InputStickerSet
	Mask_coords TL     // mask_coords:flags.0?MaskCoords
}

// Encoding TL_documentAttributeSticker
func (e TL_documentAttributeSticker) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeSticker)
	var flags int32
	if e.Mask {
		flags |= (1 << 1)
	}
	if _, ok := (e.Mask_coords).(TL_null); !ok {
		flags |= (1 << 0)
	}
	x.Int(flags)
	x.String(e.Alt)
	x.Bytes(e.Stickerset.encode())
	if flags&(1<<0) != 0 {
		x.Bytes(e.Mask_coords.encode())
	}
	return x.buf
}

// documentAttributeVideo#5910cccb duration:int w:int h:int = DocumentAttribute;

const crc_documentAttributeVideo = 0x5910cccb

type TL_documentAttributeVideo struct {
	Duration int32 // duration:int
	W        int32 // w:int
	H        int32 // h:int
}

// Encoding TL_documentAttributeVideo
func (e TL_documentAttributeVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeVideo)
	x.Int(e.Duration)
	x.Int(e.W)
	x.Int(e.H)
	return x.buf
}

// documentAttributeAudio#9852f9c6 flags:# voice:flags.10?true duration:int title:flags.0?string performer:flags.1?string waveform:flags.2?bytes = DocumentAttribute;

const crc_documentAttributeAudio = 0x9852f9c6

type TL_documentAttributeAudio struct {
	Flags     int32
	Voice     bool   // voice:flags.10?true
	Duration  int32  // duration:int
	Title     string // title:flags.0?string
	Performer string // performer:flags.1?string
	Waveform  []byte // waveform:flags.2?bytes
}

// Encoding TL_documentAttributeAudio
func (e TL_documentAttributeAudio) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeAudio)
	var flags int32
	if e.Voice {
		flags |= (1 << 10)
	}
	if e.Title != "" {
		flags |= (1 << 0)
	}
	if e.Performer != "" {
		flags |= (1 << 1)
	}
	if len(e.Waveform) != 0 {
		flags |= (1 << 2)
	}
	x.Int(flags)
	x.Int(e.Duration)
	if flags&(1<<0) != 0 {
		x.String(e.Title)
	}
	if flags&(1<<1) != 0 {
		x.String(e.Performer)
	}
	if flags&(1<<2) != 0 {
		x.Bytes(e.Waveform)
	}
	return x.buf
}

// documentAttributeFilename#15590068 file_name:string = DocumentAttribute;

const crc_documentAttributeFilename = 0x15590068

type TL_documentAttributeFilename struct {
	File_name string // file_name:string
}

// Encoding TL_documentAttributeFilename
func (e TL_documentAttributeFilename) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeFilename)
	x.String(e.File_name)
	return x.buf
}

// documentAttributeHasStickers#9801d2f7 = DocumentAttribute;

const crc_documentAttributeHasStickers = 0x9801d2f7

type TL_documentAttributeHasStickers struct {
}

// Encoding TL_documentAttributeHasStickers
func (e TL_documentAttributeHasStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeHasStickers)
	return x.buf
}

// messages.stickersNotModified#f1749a22 = messages.Stickers;

const crc_messages_stickersNotModified = 0xf1749a22

type TL_messages_stickersNotModified struct {
}

// Encoding TL_messages_stickersNotModified
func (e TL_messages_stickersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_stickersNotModified)
	return x.buf
}

// messages.stickers#8a8ecd32 Hash:string stickers:Vector<Document> = messages.Stickers;

const crc_messages_stickers = 0x8a8ecd32

type TL_messages_stickers struct {
	Hash     string // Hash:string
	Stickers []TL   // stickers:Vector<Document>
}

// Encoding TL_messages_stickers
func (e TL_messages_stickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_stickers)
	x.String(e.Hash)
	x.Vector(e.Stickers)
	return x.buf
}

// stickerPack#12b299d4 emoticon:string documents:Vector<long> = StickerPack;

const crc_stickerPack = 0x12b299d4

type TL_stickerPack struct {
	Emoticon  string  // emoticon:string
	Documents []int64 // documents:Vector<long>
}

// Encoding TL_stickerPack
func (e TL_stickerPack) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_stickerPack)
	x.String(e.Emoticon)
	x.VectorLong(e.Documents)
	return x.buf
}

// messages.allStickersNotModified#e86602c3 = messages.AllStickers;

const crc_messages_allStickersNotModified = 0xe86602c3

type TL_messages_allStickersNotModified struct {
}

// Encoding TL_messages_allStickersNotModified
func (e TL_messages_allStickersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_allStickersNotModified)
	return x.buf
}

// messages.allStickers#edfd405f Hash:int sets:Vector<StickerSet> = messages.AllStickers;

const crc_messages_allStickers = 0xedfd405f

type TL_messages_allStickers struct {
	Hash int32 // Hash:int
	Sets []TL  // sets:Vector<StickerSet>
}

// Encoding TL_messages_allStickers
func (e TL_messages_allStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_allStickers)
	x.Int(e.Hash)
	x.Vector(e.Sets)
	return x.buf
}

// disabledFeature#ae636f24 feature:string description:string = DisabledFeature;

const crc_disabledFeature = 0xae636f24

type TL_disabledFeature struct {
	Feature     string // feature:string
	Description string // description:string
}

// Encoding TL_disabledFeature
func (e TL_disabledFeature) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_disabledFeature)
	x.String(e.Feature)
	x.String(e.Description)
	return x.buf
}

// messages.affectedMessages#84d19185 pts:int pts_count:int = messages.AffectedMessages;

const crc_messages_affectedMessages = 0x84d19185

type TL_messages_affectedMessages struct {
	Pts       int32 // pts:int
	Pts_count int32 // pts_count:int
}

// Encoding TL_messages_affectedMessages
func (e TL_messages_affectedMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_affectedMessages)
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	return x.buf
}

// contactLinkUnknown#5f4f9247 = ContactLink;

const crc_contactLinkUnknown = 0x5f4f9247

type TL_contactLinkUnknown struct {
}

// Encoding TL_contactLinkUnknown
func (e TL_contactLinkUnknown) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactLinkUnknown)
	return x.buf
}

// contactLinkNone#feedd3ad = ContactLink;

const crc_contactLinkNone = 0xfeedd3ad

type TL_contactLinkNone struct {
}

// Encoding TL_contactLinkNone
func (e TL_contactLinkNone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactLinkNone)
	return x.buf
}

// contactLinkHasPhone#268f3f59 = ContactLink;

const crc_contactLinkHasPhone = 0x268f3f59

type TL_contactLinkHasPhone struct {
}

// Encoding TL_contactLinkHasPhone
func (e TL_contactLinkHasPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactLinkHasPhone)
	return x.buf
}

// contactLinkContact#d502c2d0 = ContactLink;

const crc_contactLinkContact = 0xd502c2d0

type TL_contactLinkContact struct {
}

// Encoding TL_contactLinkContact
func (e TL_contactLinkContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactLinkContact)
	return x.buf
}

// webPageEmpty#eb1477e8 Id:long = WebPage;

const crc_webPageEmpty = 0xeb1477e8

type TL_webPageEmpty struct {
	Id int64 // Id:long
}

// Encoding TL_webPageEmpty
func (e TL_webPageEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_webPageEmpty)
	x.Long(e.Id)
	return x.buf
}

// webPagePending#c586da1c Id:long date:int = WebPage;

const crc_webPagePending = 0xc586da1c

type TL_webPagePending struct {
	Id   int64 // Id:long
	Date int32 // date:int
}

// Encoding TL_webPagePending
func (e TL_webPagePending) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_webPagePending)
	x.Long(e.Id)
	x.Int(e.Date)
	return x.buf
}

// webPage#5f07b4bc flags:# Id:long url:string display_url:string Hash:int type:flags.0?string site_name:flags.1?string title:flags.2?string description:flags.3?string photo:flags.4?Photo embed_url:flags.5?string embed_type:flags.5?string embed_width:flags.6?int embed_height:flags.6?int duration:flags.7?int author:flags.8?string document:flags.9?Document cached_page:flags.10?Page = WebPage;

const crc_webPage = 0x5f07b4bc

type TL_webPage struct {
	Flags        int32
	Id           int64  // Id:long
	Url          string // url:string
	Display_url  string // display_url:string
	Hash         int32  // Hash:int
	Code_type    string // type:flags.0?string
	Site_name    string // site_name:flags.1?string
	Title        string // title:flags.2?string
	Description  string // description:flags.3?string
	Photo        TL     // photo:flags.4?Photo
	Embed_url    string // embed_url:flags.5?string
	Embed_type   string // embed_type:flags.5?string
	Embed_width  int32  // embed_width:flags.6?int
	Embed_height int32  // embed_height:flags.6?int
	Duration     int32  // duration:flags.7?int
	Author       string // author:flags.8?string
	Document     TL     // document:flags.9?Document
	Cached_page  TL     // cached_page:flags.10?Page
}

// Encoding TL_webPage
func (e TL_webPage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_webPage)
	var flags int32
	if e.Code_type != "" {
		flags |= (1 << 0)
	}
	if e.Site_name != "" {
		flags |= (1 << 1)
	}
	if e.Title != "" {
		flags |= (1 << 2)
	}
	if e.Description != "" {
		flags |= (1 << 3)
	}
	if _, ok := (e.Photo).(TL_null); !ok {
		flags |= (1 << 4)
	}
	if e.Embed_url != "" {
		flags |= (1 << 5)
	}
	if e.Embed_type != "" {
		flags |= (1 << 5)
	}
	if e.Embed_width > 0 {
		flags |= (1 << 6)
	}
	if e.Embed_height > 0 {
		flags |= (1 << 6)
	}
	if e.Duration > 0 {
		flags |= (1 << 7)
	}
	if e.Author != "" {
		flags |= (1 << 8)
	}
	if _, ok := (e.Document).(TL_null); !ok {
		flags |= (1 << 9)
	}
	if _, ok := (e.Cached_page).(TL_null); !ok {
		flags |= (1 << 10)
	}
	x.Int(flags)
	x.Long(e.Id)
	x.String(e.Url)
	x.String(e.Display_url)
	x.Int(e.Hash)
	if flags&(1<<0) != 0 {
		x.String(e.Code_type)
	}
	if flags&(1<<1) != 0 {
		x.String(e.Site_name)
	}
	if flags&(1<<2) != 0 {
		x.String(e.Title)
	}
	if flags&(1<<3) != 0 {
		x.String(e.Description)
	}
	if flags&(1<<4) != 0 {
		x.Bytes(e.Photo.encode())
	}
	if flags&(1<<5) != 0 {
		x.String(e.Embed_url)
	}
	if flags&(1<<5) != 0 {
		x.String(e.Embed_type)
	}
	if flags&(1<<6) != 0 {
		x.Int(e.Embed_width)
	}
	if flags&(1<<6) != 0 {
		x.Int(e.Embed_height)
	}
	if flags&(1<<7) != 0 {
		x.Int(e.Duration)
	}
	if flags&(1<<8) != 0 {
		x.String(e.Author)
	}
	if flags&(1<<9) != 0 {
		x.Bytes(e.Document.encode())
	}
	if flags&(1<<10) != 0 {
		x.Bytes(e.Cached_page.encode())
	}
	return x.buf
}

// webPageNotModified#85849473 = WebPage;

const crc_webPageNotModified = 0x85849473

type TL_webPageNotModified struct {
}

// Encoding TL_webPageNotModified
func (e TL_webPageNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_webPageNotModified)
	return x.buf
}

// authorization#7bf2e6f6 Hash:long flags:int device_model:string platform:string system_version:string api_id:int app_name:string app_version:string date_created:int date_active:int ip:string country:string region:string = Authorization;

const crc_authorization = 0x7bf2e6f6

type TL_authorization struct {
	Hash           int64  // Hash:long
	Flags          int32  // flags:int
	Device_model   string // device_model:string
	Platform       string // platform:string
	System_version string // system_version:string
	Api_id         int32  // api_id:int
	App_name       string // app_name:string
	App_version    string // app_version:string
	Date_created   int32  // date_created:int
	Date_active    int32  // date_active:int
	Ip             string // ip:string
	Country        string // country:string
	Region         string // region:string
}

// Encoding TL_authorization
func (e TL_authorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_authorization)
	x.Long(e.Hash)
	x.Int(e.Flags)
	x.String(e.Device_model)
	x.String(e.Platform)
	x.String(e.System_version)
	x.Int(e.Api_id)
	x.String(e.App_name)
	x.String(e.App_version)
	x.Int(e.Date_created)
	x.Int(e.Date_active)
	x.String(e.Ip)
	x.String(e.Country)
	x.String(e.Region)
	return x.buf
}

// account.authorizations#1250abde authorizations:Vector<Authorization> = account.Authorizations;

const crc_account_authorizations = 0x1250abde

type TL_account_authorizations struct {
	Authorizations []TL // authorizations:Vector<Authorization>
}

// Encoding TL_account_authorizations
func (e TL_account_authorizations) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_authorizations)
	x.Vector(e.Authorizations)
	return x.buf
}

// account.noPassword#96dabc18 new_salt:bytes email_unconfirmed_pattern:string = account.Password;

const crc_account_noPassword = 0x96dabc18

type TL_account_noPassword struct {
	New_salt                  []byte // new_salt:bytes
	Email_unconfirmed_pattern string // email_unconfirmed_pattern:string
}

// Encoding TL_account_noPassword
func (e TL_account_noPassword) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_noPassword)
	x.Bytes(e.New_salt)
	x.String(e.Email_unconfirmed_pattern)
	return x.buf
}

// account.password#7c18141c current_salt:bytes new_salt:bytes hint:string has_recovery:Bool email_unconfirmed_pattern:string = account.Password;

const crc_account_password = 0x7c18141c

type TL_account_password struct {
	Current_salt              []byte // current_salt:bytes
	New_salt                  []byte // new_salt:bytes
	Hint                      string // hint:string
	Has_recovery              TL     // has_recovery:Bool
	Email_unconfirmed_pattern string // email_unconfirmed_pattern:string
}

// Encoding TL_account_password
func (e TL_account_password) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_password)
	x.Bytes(e.Current_salt)
	x.Bytes(e.New_salt)
	x.String(e.Hint)
	x.Bytes(e.Has_recovery.encode())
	x.String(e.Email_unconfirmed_pattern)
	return x.buf
}

// account.passwordSettings#b7b72ab3 email:string = account.PasswordSettings;

const crc_account_passwordSettings = 0xb7b72ab3

type TL_account_passwordSettings struct {
	Email string // email:string
}

// Encoding TL_account_passwordSettings
func (e TL_account_passwordSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_passwordSettings)
	x.String(e.Email)
	return x.buf
}

// account.passwordInputSettings#86916deb flags:# new_salt:flags.0?bytes new_password_hash:flags.0?bytes hint:flags.0?string email:flags.1?string = account.PasswordInputSettings;

const crc_account_passwordInputSettings = 0x86916deb

type TL_account_passwordInputSettings struct {
	Flags             int32
	New_salt          []byte // new_salt:flags.0?bytes
	New_password_hash []byte // new_password_hash:flags.0?bytes
	Hint              string // hint:flags.0?string
	Email             string // email:flags.1?string
}

// Encoding TL_account_passwordInputSettings
func (e TL_account_passwordInputSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_passwordInputSettings)
	var flags int32
	if len(e.New_salt) != 0 {
		flags |= (1 << 0)
	}
	if len(e.New_password_hash) != 0 {
		flags |= (1 << 0)
	}
	if e.Hint != "" {
		flags |= (1 << 0)
	}
	if e.Email != "" {
		flags |= (1 << 1)
	}
	x.Int(flags)
	if flags&(1<<0) != 0 {
		x.Bytes(e.New_salt)
	}
	if flags&(1<<0) != 0 {
		x.Bytes(e.New_password_hash)
	}
	if flags&(1<<0) != 0 {
		x.String(e.Hint)
	}
	if flags&(1<<1) != 0 {
		x.String(e.Email)
	}
	return x.buf
}

// auth.passwordRecovery#137948a5 email_pattern:string = auth.PasswordRecovery;

const crc_auth_passwordRecovery = 0x137948a5

type TL_auth_passwordRecovery struct {
	Email_pattern string // email_pattern:string
}

// Encoding TL_auth_passwordRecovery
func (e TL_auth_passwordRecovery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_passwordRecovery)
	x.String(e.Email_pattern)
	return x.buf
}

// receivedNotifyMessage#a384b779 Id:int flags:int = ReceivedNotifyMessage;

const crc_receivedNotifyMessage = 0xa384b779

type TL_receivedNotifyMessage struct {
	Id    int32 // Id:int
	Flags int32 // flags:int
}

// Encoding TL_receivedNotifyMessage
func (e TL_receivedNotifyMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_receivedNotifyMessage)
	x.Int(e.Id)
	x.Int(e.Flags)
	return x.buf
}

// chatInviteEmpty#69df3769 = ExportedChatInvite;

const crc_chatInviteEmpty = 0x69df3769

type TL_chatInviteEmpty struct {
}

// Encoding TL_chatInviteEmpty
func (e TL_chatInviteEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatInviteEmpty)
	return x.buf
}

// chatInviteExported#fc2e05bc link:string = ExportedChatInvite;

const crc_chatInviteExported = 0xfc2e05bc

type TL_chatInviteExported struct {
	Link string // link:string
}

// Encoding TL_chatInviteExported
func (e TL_chatInviteExported) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatInviteExported)
	x.String(e.Link)
	return x.buf
}

// chatInviteAlready#5a686d7c chat:Chat = ChatInvite;

const crc_chatInviteAlready = 0x5a686d7c

type TL_chatInviteAlready struct {
	Chat TL // chat:Chat
}

// Encoding TL_chatInviteAlready
func (e TL_chatInviteAlready) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatInviteAlready)
	x.Bytes(e.Chat.encode())
	return x.buf
}

// chatInvite#db74f558 flags:# channel:flags.0?true broadcast:flags.1?true public:flags.2?true megagroup:flags.3?true title:string photo:ChatPhoto participants_count:int participants:flags.4?Vector<User> = ChatInvite;

const crc_chatInvite = 0xdb74f558

type TL_chatInvite struct {
	Flags              int32
	Channel            bool   // channel:flags.0?true
	Broadcast          bool   // broadcast:flags.1?true
	Public             bool   // public:flags.2?true
	Megagroup          bool   // megagroup:flags.3?true
	Title              string // title:string
	Photo              TL     // photo:ChatPhoto
	Participants_count int32  // participants_count:int
	Participants       []TL   // participants:flags.4?Vector<User>
}

// Encoding TL_chatInvite
func (e TL_chatInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatInvite)
	var flags int32
	if e.Channel {
		flags |= (1 << 0)
	}
	if e.Broadcast {
		flags |= (1 << 1)
	}
	if e.Public {
		flags |= (1 << 2)
	}
	if e.Megagroup {
		flags |= (1 << 3)
	}
	if len(e.Participants) != 0 {
		flags |= (1 << 4)
	}
	x.Int(flags)
	x.String(e.Title)
	x.Bytes(e.Photo.encode())
	x.Int(e.Participants_count)
	if flags&(1<<4) != 0 {
		x.Vector(e.Participants)
	}
	return x.buf
}

// inputStickerSetEmpty#ffb62b95 = InputStickerSet;

const crc_inputStickerSetEmpty = 0xffb62b95

type TL_inputStickerSetEmpty struct {
}

// Encoding TL_inputStickerSetEmpty
func (e TL_inputStickerSetEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputStickerSetEmpty)
	return x.buf
}

// inputStickerSetID#9de7a269 Id:long access_hash:long = InputStickerSet;

const crc_inputStickerSetID = 0x9de7a269

type TL_inputStickerSetID struct {
	Id          int64 // Id:long
	Access_hash int64 // access_hash:long
}

// Encoding TL_inputStickerSetID
func (e TL_inputStickerSetID) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputStickerSetID)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	return x.buf
}

// inputStickerSetShortName#861cc8a0 short_name:string = InputStickerSet;

const crc_inputStickerSetShortName = 0x861cc8a0

type TL_inputStickerSetShortName struct {
	Short_name string // short_name:string
}

// Encoding TL_inputStickerSetShortName
func (e TL_inputStickerSetShortName) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputStickerSetShortName)
	x.String(e.Short_name)
	return x.buf
}

// stickerSet#cd303b41 flags:# installed:flags.0?true archived:flags.1?true official:flags.2?true masks:flags.3?true Id:long access_hash:long title:string short_name:string count:int Hash:int = StickerSet;

const crc_stickerSet = 0xcd303b41

type TL_stickerSet struct {
	Flags       int32
	Installed   bool   // installed:flags.0?true
	Archived    bool   // archived:flags.1?true
	Official    bool   // official:flags.2?true
	Masks       bool   // masks:flags.3?true
	Id          int64  // Id:long
	Access_hash int64  // access_hash:long
	Title       string // title:string
	Short_name  string // short_name:string
	Count       int32  // count:int
	Hash        int32  // Hash:int
}

// Encoding TL_stickerSet
func (e TL_stickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_stickerSet)
	var flags int32
	if e.Installed {
		flags |= (1 << 0)
	}
	if e.Archived {
		flags |= (1 << 1)
	}
	if e.Official {
		flags |= (1 << 2)
	}
	if e.Masks {
		flags |= (1 << 3)
	}
	x.Int(flags)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	x.String(e.Title)
	x.String(e.Short_name)
	x.Int(e.Count)
	x.Int(e.Hash)
	return x.buf
}

// messages.stickerSet#b60a24a6 set:StickerSet packs:Vector<StickerPack> documents:Vector<Document> = messages.StickerSet;

const crc_messages_stickerSet = 0xb60a24a6

type TL_messages_stickerSet struct {
	Set       TL   // set:StickerSet
	Packs     []TL // packs:Vector<StickerPack>
	Documents []TL // documents:Vector<Document>
}

// Encoding TL_messages_stickerSet
func (e TL_messages_stickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_stickerSet)
	x.Bytes(e.Set.encode())
	x.Vector(e.Packs)
	x.Vector(e.Documents)
	return x.buf
}

// botCommand#c27ac8c7 command:string description:string = BotCommand;

const crc_botCommand = 0xc27ac8c7

type TL_botCommand struct {
	Command     string // command:string
	Description string // description:string
}

// Encoding TL_botCommand
func (e TL_botCommand) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botCommand)
	x.String(e.Command)
	x.String(e.Description)
	return x.buf
}

// botInfo#98e81d3a user_id:int description:string commands:Vector<BotCommand> = BotInfo;

const crc_botInfo = 0x98e81d3a

type TL_botInfo struct {
	User_id     int32  // user_id:int
	Description string // description:string
	Commands    []TL   // commands:Vector<BotCommand>
}

// Encoding TL_botInfo
func (e TL_botInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInfo)
	x.Int(e.User_id)
	x.String(e.Description)
	x.Vector(e.Commands)
	return x.buf
}

// keyboardButton#a2fa4880 text:string = KeyboardButton;

const crc_keyboardButton = 0xa2fa4880

type TL_keyboardButton struct {
	Text string // text:string
}

// Encoding TL_keyboardButton
func (e TL_keyboardButton) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButton)
	x.String(e.Text)
	return x.buf
}

// keyboardButtonUrl#258aff05 text:string url:string = KeyboardButton;

const crc_keyboardButtonUrl = 0x258aff05

type TL_keyboardButtonUrl struct {
	Text string // text:string
	Url  string // url:string
}

// Encoding TL_keyboardButtonUrl
func (e TL_keyboardButtonUrl) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonUrl)
	x.String(e.Text)
	x.String(e.Url)
	return x.buf
}

// keyboardButtonCallback#683a5e46 text:string data:bytes = KeyboardButton;

const crc_keyboardButtonCallback = 0x683a5e46

type TL_keyboardButtonCallback struct {
	Text string // text:string
	Data []byte // data:bytes
}

// Encoding TL_keyboardButtonCallback
func (e TL_keyboardButtonCallback) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonCallback)
	x.String(e.Text)
	x.Bytes(e.Data)
	return x.buf
}

// keyboardButtonRequestPhone#b16a6c29 text:string = KeyboardButton;

const crc_keyboardButtonRequestPhone = 0xb16a6c29

type TL_keyboardButtonRequestPhone struct {
	Text string // text:string
}

// Encoding TL_keyboardButtonRequestPhone
func (e TL_keyboardButtonRequestPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonRequestPhone)
	x.String(e.Text)
	return x.buf
}

// keyboardButtonRequestGeoLocation#fc796b3f text:string = KeyboardButton;

const crc_keyboardButtonRequestGeoLocation = 0xfc796b3f

type TL_keyboardButtonRequestGeoLocation struct {
	Text string // text:string
}

// Encoding TL_keyboardButtonRequestGeoLocation
func (e TL_keyboardButtonRequestGeoLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonRequestGeoLocation)
	x.String(e.Text)
	return x.buf
}

// keyboardButtonSwitchInline#568a748 flags:# same_peer:flags.0?true text:string query:string = KeyboardButton;

const crc_keyboardButtonSwitchInline = 0x568a748

type TL_keyboardButtonSwitchInline struct {
	Flags     int32
	Same_peer bool   // same_peer:flags.0?true
	Text      string // text:string
	Query     string // query:string
}

// Encoding TL_keyboardButtonSwitchInline
func (e TL_keyboardButtonSwitchInline) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonSwitchInline)
	var flags int32
	if e.Same_peer {
		flags |= (1 << 0)
	}
	x.Int(flags)
	x.String(e.Text)
	x.String(e.Query)
	return x.buf
}

// keyboardButtonGame#50f41ccf text:string = KeyboardButton;

const crc_keyboardButtonGame = 0x50f41ccf

type TL_keyboardButtonGame struct {
	Text string // text:string
}

// Encoding TL_keyboardButtonGame
func (e TL_keyboardButtonGame) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonGame)
	x.String(e.Text)
	return x.buf
}

// keyboardButtonBuy#afd93fbb text:string = KeyboardButton;

const crc_keyboardButtonBuy = 0xafd93fbb

type TL_keyboardButtonBuy struct {
	Text string // text:string
}

// Encoding TL_keyboardButtonBuy
func (e TL_keyboardButtonBuy) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonBuy)
	x.String(e.Text)
	return x.buf
}

// keyboardButtonRow#77608b83 buttons:Vector<KeyboardButton> = KeyboardButtonRow;

const crc_keyboardButtonRow = 0x77608b83

type TL_keyboardButtonRow struct {
	Buttons []TL // buttons:Vector<KeyboardButton>
}

// Encoding TL_keyboardButtonRow
func (e TL_keyboardButtonRow) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonRow)
	x.Vector(e.Buttons)
	return x.buf
}

// replyKeyboardHide#a03e5b85 flags:# selective:flags.2?true = ReplyMarkup;

const crc_replyKeyboardHide = 0xa03e5b85

type TL_replyKeyboardHide struct {
	Flags     int32
	Selective bool // selective:flags.2?true
}

// Encoding TL_replyKeyboardHide
func (e TL_replyKeyboardHide) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_replyKeyboardHide)
	var flags int32
	if e.Selective {
		flags |= (1 << 2)
	}
	x.Int(flags)
	return x.buf
}

// replyKeyboardForceReply#f4108aa0 flags:# single_use:flags.1?true selective:flags.2?true = ReplyMarkup;

const crc_replyKeyboardForceReply = 0xf4108aa0

type TL_replyKeyboardForceReply struct {
	Flags      int32
	Single_use bool // single_use:flags.1?true
	Selective  bool // selective:flags.2?true
}

// Encoding TL_replyKeyboardForceReply
func (e TL_replyKeyboardForceReply) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_replyKeyboardForceReply)
	var flags int32
	if e.Single_use {
		flags |= (1 << 1)
	}
	if e.Selective {
		flags |= (1 << 2)
	}
	x.Int(flags)
	return x.buf
}

// replyKeyboardMarkup#3502758c flags:# resize:flags.0?true single_use:flags.1?true selective:flags.2?true rows:Vector<KeyboardButtonRow> = ReplyMarkup;

const crc_replyKeyboardMarkup = 0x3502758c

type TL_replyKeyboardMarkup struct {
	Flags      int32
	Resize     bool // resize:flags.0?true
	Single_use bool // single_use:flags.1?true
	Selective  bool // selective:flags.2?true
	Rows       []TL // rows:Vector<KeyboardButtonRow>
}

// Encoding TL_replyKeyboardMarkup
func (e TL_replyKeyboardMarkup) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_replyKeyboardMarkup)
	var flags int32
	if e.Resize {
		flags |= (1 << 0)
	}
	if e.Single_use {
		flags |= (1 << 1)
	}
	if e.Selective {
		flags |= (1 << 2)
	}
	x.Int(flags)
	x.Vector(e.Rows)
	return x.buf
}

// replyInlineMarkup#48a30254 rows:Vector<KeyboardButtonRow> = ReplyMarkup;

const crc_replyInlineMarkup = 0x48a30254

type TL_replyInlineMarkup struct {
	Rows []TL // rows:Vector<KeyboardButtonRow>
}

// Encoding TL_replyInlineMarkup
func (e TL_replyInlineMarkup) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_replyInlineMarkup)
	x.Vector(e.Rows)
	return x.buf
}

// messageEntityUnknown#bb92ba95 offset:int length:int = MessageEntity;

const crc_messageEntityUnknown = 0xbb92ba95

type TL_messageEntityUnknown struct {
	Offset int32 // offset:int
	Length int32 // length:int
}

// Encoding TL_messageEntityUnknown
func (e TL_messageEntityUnknown) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityUnknown)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

// messageEntityMention#fa04579d offset:int length:int = MessageEntity;

const crc_messageEntityMention = 0xfa04579d

type TL_messageEntityMention struct {
	Offset int32 // offset:int
	Length int32 // length:int
}

// Encoding TL_messageEntityMention
func (e TL_messageEntityMention) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityMention)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

// messageEntityHashtag#6f635b0d offset:int length:int = MessageEntity;

const crc_messageEntityHashtag = 0x6f635b0d

type TL_messageEntityHashtag struct {
	Offset int32 // offset:int
	Length int32 // length:int
}

// Encoding TL_messageEntityHashtag
func (e TL_messageEntityHashtag) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityHashtag)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

// messageEntityBotCommand#6cef8ac7 offset:int length:int = MessageEntity;

const crc_messageEntityBotCommand = 0x6cef8ac7

type TL_messageEntityBotCommand struct {
	Offset int32 // offset:int
	Length int32 // length:int
}

// Encoding TL_messageEntityBotCommand
func (e TL_messageEntityBotCommand) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityBotCommand)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

// messageEntityUrl#6ed02538 offset:int length:int = MessageEntity;

const crc_messageEntityUrl = 0x6ed02538

type TL_messageEntityUrl struct {
	Offset int32 // offset:int
	Length int32 // length:int
}

// Encoding TL_messageEntityUrl
func (e TL_messageEntityUrl) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityUrl)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

// messageEntityEmail#64e475c2 offset:int length:int = MessageEntity;

const crc_messageEntityEmail = 0x64e475c2

type TL_messageEntityEmail struct {
	Offset int32 // offset:int
	Length int32 // length:int
}

// Encoding TL_messageEntityEmail
func (e TL_messageEntityEmail) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityEmail)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

// messageEntityBold#bd610bc9 offset:int length:int = MessageEntity;

const crc_messageEntityBold = 0xbd610bc9

type TL_messageEntityBold struct {
	Offset int32 // offset:int
	Length int32 // length:int
}

// Encoding TL_messageEntityBold
func (e TL_messageEntityBold) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityBold)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

// messageEntityItalic#826f8b60 offset:int length:int = MessageEntity;

const crc_messageEntityItalic = 0x826f8b60

type TL_messageEntityItalic struct {
	Offset int32 // offset:int
	Length int32 // length:int
}

// Encoding TL_messageEntityItalic
func (e TL_messageEntityItalic) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityItalic)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

// messageEntityCode#28a20571 offset:int length:int = MessageEntity;

const crc_messageEntityCode = 0x28a20571

type TL_messageEntityCode struct {
	Offset int32 // offset:int
	Length int32 // length:int
}

// Encoding TL_messageEntityCode
func (e TL_messageEntityCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityCode)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

// messageEntityPre#73924be0 offset:int length:int Language:string = MessageEntity;

const crc_messageEntityPre = 0x73924be0

type TL_messageEntityPre struct {
	Offset   int32  // offset:int
	Length   int32  // length:int
	Language string // Language:string
}

// Encoding TL_messageEntityPre
func (e TL_messageEntityPre) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityPre)
	x.Int(e.Offset)
	x.Int(e.Length)
	x.String(e.Language)
	return x.buf
}

// messageEntityTextUrl#76a6d327 offset:int length:int url:string = MessageEntity;

const crc_messageEntityTextUrl = 0x76a6d327

type TL_messageEntityTextUrl struct {
	Offset int32  // offset:int
	Length int32  // length:int
	Url    string // url:string
}

// Encoding TL_messageEntityTextUrl
func (e TL_messageEntityTextUrl) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityTextUrl)
	x.Int(e.Offset)
	x.Int(e.Length)
	x.String(e.Url)
	return x.buf
}

// messageEntityMentionName#352dca58 offset:int length:int user_id:int = MessageEntity;

const crc_messageEntityMentionName = 0x352dca58

type TL_messageEntityMentionName struct {
	Offset  int32 // offset:int
	Length  int32 // length:int
	User_id int32 // user_id:int
}

// Encoding TL_messageEntityMentionName
func (e TL_messageEntityMentionName) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityMentionName)
	x.Int(e.Offset)
	x.Int(e.Length)
	x.Int(e.User_id)
	return x.buf
}

// inputMessageEntityMentionName#208e68c9 offset:int length:int user_id:InputUser = MessageEntity;

const crc_inputMessageEntityMentionName = 0x208e68c9

type TL_inputMessageEntityMentionName struct {
	Offset  int32 // offset:int
	Length  int32 // length:int
	User_id TL    // user_id:InputUser
}

// Encoding TL_inputMessageEntityMentionName
func (e TL_inputMessageEntityMentionName) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessageEntityMentionName)
	x.Int(e.Offset)
	x.Int(e.Length)
	x.Bytes(e.User_id.encode())
	return x.buf
}

// inputChannelEmpty#ee8c1e86 = InputChannel;

const crc_inputChannelEmpty = 0xee8c1e86

type TL_inputChannelEmpty struct {
}

// Encoding TL_inputChannelEmpty
func (e TL_inputChannelEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputChannelEmpty)
	return x.buf
}

// inputChannel#afeb712e channel_id:int access_hash:long = InputChannel;

const crc_inputChannel = 0xafeb712e

type TL_inputChannel struct {
	Channel_id  int32 // channel_id:int
	Access_hash int64 // access_hash:long
}

// Encoding TL_inputChannel
func (e TL_inputChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputChannel)
	x.Int(e.Channel_id)
	x.Long(e.Access_hash)
	return x.buf
}

// contacts.resolvedPeer#7f077ad9 peer:Peer chats:Vector<Chat> users:Vector<User> = contacts.ResolvedPeer;

const crc_contacts_resolvedPeer = 0x7f077ad9

type TL_contacts_resolvedPeer struct {
	Peer  TL   // peer:Peer
	Chats []TL // chats:Vector<Chat>
	Users []TL // users:Vector<User>
}

// Encoding TL_contacts_resolvedPeer
func (e TL_contacts_resolvedPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_resolvedPeer)
	x.Bytes(e.Peer.encode())
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

// messageRange#ae30253 min_id:int max_id:int = MessageRange;

const crc_messageRange = 0xae30253

type TL_messageRange struct {
	Min_id int32 // min_id:int
	Max_id int32 // max_id:int
}

// Encoding TL_messageRange
func (e TL_messageRange) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageRange)
	x.Int(e.Min_id)
	x.Int(e.Max_id)
	return x.buf
}

// updates.channelDifferenceEmpty#3e11affb flags:# final:flags.0?true pts:int timeout:flags.1?int = updates.ChannelDifference;

const crc_updates_channelDifferenceEmpty = 0x3e11affb

type TL_updates_channelDifferenceEmpty struct {
	Flags   int32
	Final   bool  // final:flags.0?true
	Pts     int32 // pts:int
	Timeout int32 // timeout:flags.1?int
}

// Encoding TL_updates_channelDifferenceEmpty
func (e TL_updates_channelDifferenceEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_channelDifferenceEmpty)
	var flags int32
	if e.Final {
		flags |= (1 << 0)
	}
	if e.Timeout > 0 {
		flags |= (1 << 1)
	}
	x.Int(flags)
	x.Int(e.Pts)
	if flags&(1<<1) != 0 {
		x.Int(e.Timeout)
	}
	return x.buf
}

// updates.channelDifferenceTooLong#410dee07 flags:# final:flags.0?true pts:int timeout:flags.1?int top_message:int read_inbox_max_id:int read_outbox_max_id:int unread_count:int messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = updates.ChannelDifference;

const crc_updates_channelDifferenceTooLong = 0x410dee07

type TL_updates_channelDifferenceTooLong struct {
	Flags              int32
	Final              bool  // final:flags.0?true
	Pts                int32 // pts:int
	Timeout            int32 // timeout:flags.1?int
	Top_message        int32 // top_message:int
	Read_inbox_max_id  int32 // read_inbox_max_id:int
	Read_outbox_max_id int32 // read_outbox_max_id:int
	Unread_count       int32 // unread_count:int
	Messages           []TL  // messages:Vector<Message>
	Chats              []TL  // chats:Vector<Chat>
	Users              []TL  // users:Vector<User>
}

// Encoding TL_updates_channelDifferenceTooLong
func (e TL_updates_channelDifferenceTooLong) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_channelDifferenceTooLong)
	var flags int32
	if e.Final {
		flags |= (1 << 0)
	}
	if e.Timeout > 0 {
		flags |= (1 << 1)
	}
	x.Int(flags)
	x.Int(e.Pts)
	if flags&(1<<1) != 0 {
		x.Int(e.Timeout)
	}
	x.Int(e.Top_message)
	x.Int(e.Read_inbox_max_id)
	x.Int(e.Read_outbox_max_id)
	x.Int(e.Unread_count)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

// updates.channelDifference#2064674e flags:# final:flags.0?true pts:int timeout:flags.1?int new_messages:Vector<Message> other_updates:Vector<Update> chats:Vector<Chat> users:Vector<User> = updates.ChannelDifference;

const crc_updates_channelDifference = 0x2064674e

type TL_updates_channelDifference struct {
	Flags         int32
	Final         bool  // final:flags.0?true
	Pts           int32 // pts:int
	Timeout       int32 // timeout:flags.1?int
	New_messages  []TL  // new_messages:Vector<Message>
	Other_updates []TL  // other_updates:Vector<Update>
	Chats         []TL  // chats:Vector<Chat>
	Users         []TL  // users:Vector<User>
}

// Encoding TL_updates_channelDifference
func (e TL_updates_channelDifference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_channelDifference)
	var flags int32
	if e.Final {
		flags |= (1 << 0)
	}
	if e.Timeout > 0 {
		flags |= (1 << 1)
	}
	x.Int(flags)
	x.Int(e.Pts)
	if flags&(1<<1) != 0 {
		x.Int(e.Timeout)
	}
	x.Vector(e.New_messages)
	x.Vector(e.Other_updates)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

// channelMessagesFilterEmpty#94d42ee7 = ChannelMessagesFilter;

const crc_channelMessagesFilterEmpty = 0x94d42ee7

type TL_channelMessagesFilterEmpty struct {
}

// Encoding TL_channelMessagesFilterEmpty
func (e TL_channelMessagesFilterEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelMessagesFilterEmpty)
	return x.buf
}

// channelMessagesFilter#cd77d957 flags:# exclude_new_messages:flags.1?true ranges:Vector<MessageRange> = ChannelMessagesFilter;

const crc_channelMessagesFilter = 0xcd77d957

type TL_channelMessagesFilter struct {
	Flags                int32
	Exclude_new_messages bool // exclude_new_messages:flags.1?true
	Ranges               []TL // ranges:Vector<MessageRange>
}

// Encoding TL_channelMessagesFilter
func (e TL_channelMessagesFilter) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelMessagesFilter)
	var flags int32
	if e.Exclude_new_messages {
		flags |= (1 << 1)
	}
	x.Int(flags)
	x.Vector(e.Ranges)
	return x.buf
}

// channelParticipant#15ebac1d user_id:int date:int = ChannelParticipant;

const crc_channelParticipant = 0x15ebac1d

type TL_channelParticipant struct {
	User_id int32 // user_id:int
	Date    int32 // date:int
}

// Encoding TL_channelParticipant
func (e TL_channelParticipant) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipant)
	x.Int(e.User_id)
	x.Int(e.Date)
	return x.buf
}

// channelParticipantSelf#a3289a6d user_id:int inviter_id:int date:int = ChannelParticipant;

const crc_channelParticipantSelf = 0xa3289a6d

type TL_channelParticipantSelf struct {
	User_id    int32 // user_id:int
	Inviter_id int32 // inviter_id:int
	Date       int32 // date:int
}

// Encoding TL_channelParticipantSelf
func (e TL_channelParticipantSelf) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantSelf)
	x.Int(e.User_id)
	x.Int(e.Inviter_id)
	x.Int(e.Date)
	return x.buf
}

// channelParticipantModerator#91057fef user_id:int inviter_id:int date:int = ChannelParticipant;

const crc_channelParticipantModerator = 0x91057fef

type TL_channelParticipantModerator struct {
	User_id    int32 // user_id:int
	Inviter_id int32 // inviter_id:int
	Date       int32 // date:int
}

// Encoding TL_channelParticipantModerator
func (e TL_channelParticipantModerator) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantModerator)
	x.Int(e.User_id)
	x.Int(e.Inviter_id)
	x.Int(e.Date)
	return x.buf
}

// channelParticipantEditor#98192d61 user_id:int inviter_id:int date:int = ChannelParticipant;

const crc_channelParticipantEditor = 0x98192d61

type TL_channelParticipantEditor struct {
	User_id    int32 // user_id:int
	Inviter_id int32 // inviter_id:int
	Date       int32 // date:int
}

// Encoding TL_channelParticipantEditor
func (e TL_channelParticipantEditor) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantEditor)
	x.Int(e.User_id)
	x.Int(e.Inviter_id)
	x.Int(e.Date)
	return x.buf
}

// channelParticipantKicked#8cc5e69a user_id:int kicked_by:int date:int = ChannelParticipant;

const crc_channelParticipantKicked = 0x8cc5e69a

type TL_channelParticipantKicked struct {
	User_id   int32 // user_id:int
	Kicked_by int32 // kicked_by:int
	Date      int32 // date:int
}

// Encoding TL_channelParticipantKicked
func (e TL_channelParticipantKicked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantKicked)
	x.Int(e.User_id)
	x.Int(e.Kicked_by)
	x.Int(e.Date)
	return x.buf
}

// channelParticipantCreator#e3e2e1f9 user_id:int = ChannelParticipant;

const crc_channelParticipantCreator = 0xe3e2e1f9

type TL_channelParticipantCreator struct {
	User_id int32 // user_id:int
}

// Encoding TL_channelParticipantCreator
func (e TL_channelParticipantCreator) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantCreator)
	x.Int(e.User_id)
	return x.buf
}

// channelParticipantsRecent#de3f3c79 = ChannelParticipantsFilter;

const crc_channelParticipantsRecent = 0xde3f3c79

type TL_channelParticipantsRecent struct {
}

// Encoding TL_channelParticipantsRecent
func (e TL_channelParticipantsRecent) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantsRecent)
	return x.buf
}

// channelParticipantsAdmins#b4608969 = ChannelParticipantsFilter;

const crc_channelParticipantsAdmins = 0xb4608969

type TL_channelParticipantsAdmins struct {
}

// Encoding TL_channelParticipantsAdmins
func (e TL_channelParticipantsAdmins) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantsAdmins)
	return x.buf
}

// channelParticipantsKicked#3c37bb7a = ChannelParticipantsFilter;

const crc_channelParticipantsKicked = 0x3c37bb7a

type TL_channelParticipantsKicked struct {
}

// Encoding TL_channelParticipantsKicked
func (e TL_channelParticipantsKicked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantsKicked)
	return x.buf
}

// channelParticipantsBots#b0d1865b = ChannelParticipantsFilter;

const crc_channelParticipantsBots = 0xb0d1865b

type TL_channelParticipantsBots struct {
}

// Encoding TL_channelParticipantsBots
func (e TL_channelParticipantsBots) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantsBots)
	return x.buf
}

// channelRoleEmpty#b285a0c6 = ChannelParticipantRole;

const crc_channelRoleEmpty = 0xb285a0c6

type TL_channelRoleEmpty struct {
}

// Encoding TL_channelRoleEmpty
func (e TL_channelRoleEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelRoleEmpty)
	return x.buf
}

// channelRoleModerator#9618d975 = ChannelParticipantRole;

const crc_channelRoleModerator = 0x9618d975

type TL_channelRoleModerator struct {
}

// Encoding TL_channelRoleModerator
func (e TL_channelRoleModerator) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelRoleModerator)
	return x.buf
}

// channelRoleEditor#820bfe8c = ChannelParticipantRole;

const crc_channelRoleEditor = 0x820bfe8c

type TL_channelRoleEditor struct {
}

// Encoding TL_channelRoleEditor
func (e TL_channelRoleEditor) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelRoleEditor)
	return x.buf
}

// channels.channelParticipants#f56ee2a8 count:int participants:Vector<ChannelParticipant> users:Vector<User> = channels.ChannelParticipants;

const crc_channels_channelParticipants = 0xf56ee2a8

type TL_channels_channelParticipants struct {
	Count        int32 // count:int
	Participants []TL  // participants:Vector<ChannelParticipant>
	Users        []TL  // users:Vector<User>
}

// Encoding TL_channels_channelParticipants
func (e TL_channels_channelParticipants) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_channelParticipants)
	x.Int(e.Count)
	x.Vector(e.Participants)
	x.Vector(e.Users)
	return x.buf
}

// channels.channelParticipant#d0d9b163 participant:ChannelParticipant users:Vector<User> = channels.ChannelParticipant;

const crc_channels_channelParticipant = 0xd0d9b163

type TL_channels_channelParticipant struct {
	Participant TL   // participant:ChannelParticipant
	Users       []TL // users:Vector<User>
}

// Encoding TL_channels_channelParticipant
func (e TL_channels_channelParticipant) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_channelParticipant)
	x.Bytes(e.Participant.encode())
	x.Vector(e.Users)
	return x.buf
}

// help.termsOfService#f1ee3e90 text:string = help.TermsOfService;

const crc_help_termsOfService = 0xf1ee3e90

type TL_help_termsOfService struct {
	Text string // text:string
}

// Encoding TL_help_termsOfService
func (e TL_help_termsOfService) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_termsOfService)
	x.String(e.Text)
	return x.buf
}

// foundGif#162ecc1f url:string thumb_url:string content_url:string content_type:string w:int h:int = FoundGif;

const crc_foundGif = 0x162ecc1f

type TL_foundGif struct {
	Url          string // url:string
	Thumb_url    string // thumb_url:string
	Content_url  string // content_url:string
	Content_type string // content_type:string
	W            int32  // w:int
	H            int32  // h:int
}

// Encoding TL_foundGif
func (e TL_foundGif) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_foundGif)
	x.String(e.Url)
	x.String(e.Thumb_url)
	x.String(e.Content_url)
	x.String(e.Content_type)
	x.Int(e.W)
	x.Int(e.H)
	return x.buf
}

// foundGifCached#9c750409 url:string photo:Photo document:Document = FoundGif;

const crc_foundGifCached = 0x9c750409

type TL_foundGifCached struct {
	Url      string // url:string
	Photo    TL     // photo:Photo
	Document TL     // document:Document
}

// Encoding TL_foundGifCached
func (e TL_foundGifCached) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_foundGifCached)
	x.String(e.Url)
	x.Bytes(e.Photo.encode())
	x.Bytes(e.Document.encode())
	return x.buf
}

// messages.foundGifs#450a1c0a next_offset:int results:Vector<FoundGif> = messages.FoundGifs;

const crc_messages_foundGifs = 0x450a1c0a

type TL_messages_foundGifs struct {
	Next_offset int32 // next_offset:int
	Results     []TL  // results:Vector<FoundGif>
}

// Encoding TL_messages_foundGifs
func (e TL_messages_foundGifs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_foundGifs)
	x.Int(e.Next_offset)
	x.Vector(e.Results)
	return x.buf
}

// messages.savedGifsNotModified#e8025ca2 = messages.SavedGifs;

const crc_messages_savedGifsNotModified = 0xe8025ca2

type TL_messages_savedGifsNotModified struct {
}

// Encoding TL_messages_savedGifsNotModified
func (e TL_messages_savedGifsNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_savedGifsNotModified)
	return x.buf
}

// messages.savedGifs#2e0709a5 Hash:int gifs:Vector<Document> = messages.SavedGifs;

const crc_messages_savedGifs = 0x2e0709a5

type TL_messages_savedGifs struct {
	Hash int32 // Hash:int
	Gifs []TL  // gifs:Vector<Document>
}

// Encoding TL_messages_savedGifs
func (e TL_messages_savedGifs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_savedGifs)
	x.Int(e.Hash)
	x.Vector(e.Gifs)
	return x.buf
}

// inputBotInlineMessageMediaAuto#292fed13 flags:# caption:string reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;

const crc_inputBotInlineMessageMediaAuto = 0x292fed13

type TL_inputBotInlineMessageMediaAuto struct {
	Flags        int32
	Caption      string // caption:string
	Reply_markup TL     // reply_markup:flags.2?ReplyMarkup
}

// Encoding TL_inputBotInlineMessageMediaAuto
func (e TL_inputBotInlineMessageMediaAuto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineMessageMediaAuto)
	var flags int32
	if _, ok := (e.Reply_markup).(TL_null); !ok {
		flags |= (1 << 2)
	}
	x.Int(flags)
	x.String(e.Caption)
	if flags&(1<<2) != 0 {
		x.Bytes(e.Reply_markup.encode())
	}
	return x.buf
}

// inputBotInlineMessageText#3dcd7a87 flags:# no_webpage:flags.0?true message:string entities:flags.1?Vector<MessageEntity> reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;

const crc_inputBotInlineMessageText = 0x3dcd7a87

type TL_inputBotInlineMessageText struct {
	Flags        int32
	No_webpage   bool   // no_webpage:flags.0?true
	Message      string // message:string
	Entities     []TL   // entities:flags.1?Vector<MessageEntity>
	Reply_markup TL     // reply_markup:flags.2?ReplyMarkup
}

// Encoding TL_inputBotInlineMessageText
func (e TL_inputBotInlineMessageText) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineMessageText)
	var flags int32
	if e.No_webpage {
		flags |= (1 << 0)
	}
	if len(e.Entities) != 0 {
		flags |= (1 << 1)
	}
	if _, ok := (e.Reply_markup).(TL_null); !ok {
		flags |= (1 << 2)
	}
	x.Int(flags)
	x.String(e.Message)
	if flags&(1<<1) != 0 {
		x.Vector(e.Entities)
	}
	if flags&(1<<2) != 0 {
		x.Bytes(e.Reply_markup.encode())
	}
	return x.buf
}

// inputBotInlineMessageMediaGeo#f4a59de1 flags:# geo_point:InputGeoPoint reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;

const crc_inputBotInlineMessageMediaGeo = 0xf4a59de1

type TL_inputBotInlineMessageMediaGeo struct {
	Flags        int32
	Geo_point    TL // geo_point:InputGeoPoint
	Reply_markup TL // reply_markup:flags.2?ReplyMarkup
}

// Encoding TL_inputBotInlineMessageMediaGeo
func (e TL_inputBotInlineMessageMediaGeo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineMessageMediaGeo)
	var flags int32
	if _, ok := (e.Reply_markup).(TL_null); !ok {
		flags |= (1 << 2)
	}
	x.Int(flags)
	x.Bytes(e.Geo_point.encode())
	if flags&(1<<2) != 0 {
		x.Bytes(e.Reply_markup.encode())
	}
	return x.buf
}

// inputBotInlineMessageMediaVenue#aaafadc8 flags:# geo_point:InputGeoPoint title:string address:string provider:string venue_id:string reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;

const crc_inputBotInlineMessageMediaVenue = 0xaaafadc8

type TL_inputBotInlineMessageMediaVenue struct {
	Flags        int32
	Geo_point    TL     // geo_point:InputGeoPoint
	Title        string // title:string
	Address      string // address:string
	Provider     string // provider:string
	Venue_id     string // venue_id:string
	Reply_markup TL     // reply_markup:flags.2?ReplyMarkup
}

// Encoding TL_inputBotInlineMessageMediaVenue
func (e TL_inputBotInlineMessageMediaVenue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineMessageMediaVenue)
	var flags int32
	if _, ok := (e.Reply_markup).(TL_null); !ok {
		flags |= (1 << 2)
	}
	x.Int(flags)
	x.Bytes(e.Geo_point.encode())
	x.String(e.Title)
	x.String(e.Address)
	x.String(e.Provider)
	x.String(e.Venue_id)
	if flags&(1<<2) != 0 {
		x.Bytes(e.Reply_markup.encode())
	}
	return x.buf
}

// inputBotInlineMessageMediaContact#2daf01a7 flags:# phone_number:string first_name:string last_name:string reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;

const crc_inputBotInlineMessageMediaContact = 0x2daf01a7

type TL_inputBotInlineMessageMediaContact struct {
	Flags        int32
	Phone_number string // phone_number:string
	First_name   string // first_name:string
	Last_name    string // last_name:string
	Reply_markup TL     // reply_markup:flags.2?ReplyMarkup
}

// Encoding TL_inputBotInlineMessageMediaContact
func (e TL_inputBotInlineMessageMediaContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineMessageMediaContact)
	var flags int32
	if _, ok := (e.Reply_markup).(TL_null); !ok {
		flags |= (1 << 2)
	}
	x.Int(flags)
	x.String(e.Phone_number)
	x.String(e.First_name)
	x.String(e.Last_name)
	if flags&(1<<2) != 0 {
		x.Bytes(e.Reply_markup.encode())
	}
	return x.buf
}

// inputBotInlineMessageGame#4b425864 flags:# reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;

const crc_inputBotInlineMessageGame = 0x4b425864

type TL_inputBotInlineMessageGame struct {
	Flags        int32
	Reply_markup TL // reply_markup:flags.2?ReplyMarkup
}

// Encoding TL_inputBotInlineMessageGame
func (e TL_inputBotInlineMessageGame) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineMessageGame)
	var flags int32
	if _, ok := (e.Reply_markup).(TL_null); !ok {
		flags |= (1 << 2)
	}
	x.Int(flags)
	if flags&(1<<2) != 0 {
		x.Bytes(e.Reply_markup.encode())
	}
	return x.buf
}

// inputBotInlineResult#2cbbe15a flags:# Id:string type:string title:flags.1?string description:flags.2?string url:flags.3?string thumb_url:flags.4?string content_url:flags.5?string content_type:flags.5?string w:flags.6?int h:flags.6?int duration:flags.7?int send_message:InputBotInlineMessage = InputBotInlineResult;

const crc_inputBotInlineResult = 0x2cbbe15a

type TL_inputBotInlineResult struct {
	Flags        int32
	Id           string // Id:string
	Code_type    string // type:string
	Title        string // title:flags.1?string
	Description  string // description:flags.2?string
	Url          string // url:flags.3?string
	Thumb_url    string // thumb_url:flags.4?string
	Content_url  string // content_url:flags.5?string
	Content_type string // content_type:flags.5?string
	W            int32  // w:flags.6?int
	H            int32  // h:flags.6?int
	Duration     int32  // duration:flags.7?int
	Send_message TL     // send_message:InputBotInlineMessage
}

// Encoding TL_inputBotInlineResult
func (e TL_inputBotInlineResult) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineResult)
	var flags int32
	if e.Title != "" {
		flags |= (1 << 1)
	}
	if e.Description != "" {
		flags |= (1 << 2)
	}
	if e.Url != "" {
		flags |= (1 << 3)
	}
	if e.Thumb_url != "" {
		flags |= (1 << 4)
	}
	if e.Content_url != "" {
		flags |= (1 << 5)
	}
	if e.Content_type != "" {
		flags |= (1 << 5)
	}
	if e.W > 0 {
		flags |= (1 << 6)
	}
	if e.H > 0 {
		flags |= (1 << 6)
	}
	if e.Duration > 0 {
		flags |= (1 << 7)
	}
	x.Int(flags)
	x.String(e.Id)
	x.String(e.Code_type)
	if flags&(1<<1) != 0 {
		x.String(e.Title)
	}
	if flags&(1<<2) != 0 {
		x.String(e.Description)
	}
	if flags&(1<<3) != 0 {
		x.String(e.Url)
	}
	if flags&(1<<4) != 0 {
		x.String(e.Thumb_url)
	}
	if flags&(1<<5) != 0 {
		x.String(e.Content_url)
	}
	if flags&(1<<5) != 0 {
		x.String(e.Content_type)
	}
	if flags&(1<<6) != 0 {
		x.Int(e.W)
	}
	if flags&(1<<6) != 0 {
		x.Int(e.H)
	}
	if flags&(1<<7) != 0 {
		x.Int(e.Duration)
	}
	x.Bytes(e.Send_message.encode())
	return x.buf
}

// inputBotInlineResultPhoto#a8d864a7 Id:string type:string photo:InputPhoto send_message:InputBotInlineMessage = InputBotInlineResult;

const crc_inputBotInlineResultPhoto = 0xa8d864a7

type TL_inputBotInlineResultPhoto struct {
	Id           string // Id:string
	Code_type    string // type:string
	Photo        TL     // photo:InputPhoto
	Send_message TL     // send_message:InputBotInlineMessage
}

// Encoding TL_inputBotInlineResultPhoto
func (e TL_inputBotInlineResultPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineResultPhoto)
	x.String(e.Id)
	x.String(e.Code_type)
	x.Bytes(e.Photo.encode())
	x.Bytes(e.Send_message.encode())
	return x.buf
}

// inputBotInlineResultDocument#fff8fdc4 flags:# Id:string type:string title:flags.1?string description:flags.2?string document:InputDocument send_message:InputBotInlineMessage = InputBotInlineResult;

const crc_inputBotInlineResultDocument = 0xfff8fdc4

type TL_inputBotInlineResultDocument struct {
	Flags        int32
	Id           string // Id:string
	Code_type    string // type:string
	Title        string // title:flags.1?string
	Description  string // description:flags.2?string
	Document     TL     // document:InputDocument
	Send_message TL     // send_message:InputBotInlineMessage
}

// Encoding TL_inputBotInlineResultDocument
func (e TL_inputBotInlineResultDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineResultDocument)
	var flags int32
	if e.Title != "" {
		flags |= (1 << 1)
	}
	if e.Description != "" {
		flags |= (1 << 2)
	}
	x.Int(flags)
	x.String(e.Id)
	x.String(e.Code_type)
	if flags&(1<<1) != 0 {
		x.String(e.Title)
	}
	if flags&(1<<2) != 0 {
		x.String(e.Description)
	}
	x.Bytes(e.Document.encode())
	x.Bytes(e.Send_message.encode())
	return x.buf
}

// inputBotInlineResultGame#4fa417f2 Id:string short_name:string send_message:InputBotInlineMessage = InputBotInlineResult;

const crc_inputBotInlineResultGame = 0x4fa417f2

type TL_inputBotInlineResultGame struct {
	Id           string // Id:string
	Short_name   string // short_name:string
	Send_message TL     // send_message:InputBotInlineMessage
}

// Encoding TL_inputBotInlineResultGame
func (e TL_inputBotInlineResultGame) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineResultGame)
	x.String(e.Id)
	x.String(e.Short_name)
	x.Bytes(e.Send_message.encode())
	return x.buf
}

// botInlineMessageMediaAuto#a74b15b flags:# caption:string reply_markup:flags.2?ReplyMarkup = BotInlineMessage;

const crc_botInlineMessageMediaAuto = 0xa74b15b

type TL_botInlineMessageMediaAuto struct {
	Flags        int32
	Caption      string // caption:string
	Reply_markup TL     // reply_markup:flags.2?ReplyMarkup
}

// Encoding TL_botInlineMessageMediaAuto
func (e TL_botInlineMessageMediaAuto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInlineMessageMediaAuto)
	var flags int32
	if _, ok := (e.Reply_markup).(TL_null); !ok {
		flags |= (1 << 2)
	}
	x.Int(flags)
	x.String(e.Caption)
	if flags&(1<<2) != 0 {
		x.Bytes(e.Reply_markup.encode())
	}
	return x.buf
}

// botInlineMessageText#8c7f65e2 flags:# no_webpage:flags.0?true message:string entities:flags.1?Vector<MessageEntity> reply_markup:flags.2?ReplyMarkup = BotInlineMessage;

const crc_botInlineMessageText = 0x8c7f65e2

type TL_botInlineMessageText struct {
	Flags        int32
	No_webpage   bool   // no_webpage:flags.0?true
	Message      string // message:string
	Entities     []TL   // entities:flags.1?Vector<MessageEntity>
	Reply_markup TL     // reply_markup:flags.2?ReplyMarkup
}

// Encoding TL_botInlineMessageText
func (e TL_botInlineMessageText) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInlineMessageText)
	var flags int32
	if e.No_webpage {
		flags |= (1 << 0)
	}
	if len(e.Entities) != 0 {
		flags |= (1 << 1)
	}
	if _, ok := (e.Reply_markup).(TL_null); !ok {
		flags |= (1 << 2)
	}
	x.Int(flags)
	x.String(e.Message)
	if flags&(1<<1) != 0 {
		x.Vector(e.Entities)
	}
	if flags&(1<<2) != 0 {
		x.Bytes(e.Reply_markup.encode())
	}
	return x.buf
}

// botInlineMessageMediaGeo#3a8fd8b8 flags:# geo:GeoPoint reply_markup:flags.2?ReplyMarkup = BotInlineMessage;

const crc_botInlineMessageMediaGeo = 0x3a8fd8b8

type TL_botInlineMessageMediaGeo struct {
	Flags        int32
	Geo          TL // geo:GeoPoint
	Reply_markup TL // reply_markup:flags.2?ReplyMarkup
}

// Encoding TL_botInlineMessageMediaGeo
func (e TL_botInlineMessageMediaGeo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInlineMessageMediaGeo)
	var flags int32
	if _, ok := (e.Reply_markup).(TL_null); !ok {
		flags |= (1 << 2)
	}
	x.Int(flags)
	x.Bytes(e.Geo.encode())
	if flags&(1<<2) != 0 {
		x.Bytes(e.Reply_markup.encode())
	}
	return x.buf
}

// botInlineMessageMediaVenue#4366232e flags:# geo:GeoPoint title:string address:string provider:string venue_id:string reply_markup:flags.2?ReplyMarkup = BotInlineMessage;

const crc_botInlineMessageMediaVenue = 0x4366232e

type TL_botInlineMessageMediaVenue struct {
	Flags        int32
	Geo          TL     // geo:GeoPoint
	Title        string // title:string
	Address      string // address:string
	Provider     string // provider:string
	Venue_id     string // venue_id:string
	Reply_markup TL     // reply_markup:flags.2?ReplyMarkup
}

// Encoding TL_botInlineMessageMediaVenue
func (e TL_botInlineMessageMediaVenue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInlineMessageMediaVenue)
	var flags int32
	if _, ok := (e.Reply_markup).(TL_null); !ok {
		flags |= (1 << 2)
	}
	x.Int(flags)
	x.Bytes(e.Geo.encode())
	x.String(e.Title)
	x.String(e.Address)
	x.String(e.Provider)
	x.String(e.Venue_id)
	if flags&(1<<2) != 0 {
		x.Bytes(e.Reply_markup.encode())
	}
	return x.buf
}

// botInlineMessageMediaContact#35edb4d4 flags:# phone_number:string first_name:string last_name:string reply_markup:flags.2?ReplyMarkup = BotInlineMessage;

const crc_botInlineMessageMediaContact = 0x35edb4d4

type TL_botInlineMessageMediaContact struct {
	Flags        int32
	Phone_number string // phone_number:string
	First_name   string // first_name:string
	Last_name    string // last_name:string
	Reply_markup TL     // reply_markup:flags.2?ReplyMarkup
}

// Encoding TL_botInlineMessageMediaContact
func (e TL_botInlineMessageMediaContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInlineMessageMediaContact)
	var flags int32
	if _, ok := (e.Reply_markup).(TL_null); !ok {
		flags |= (1 << 2)
	}
	x.Int(flags)
	x.String(e.Phone_number)
	x.String(e.First_name)
	x.String(e.Last_name)
	if flags&(1<<2) != 0 {
		x.Bytes(e.Reply_markup.encode())
	}
	return x.buf
}

// botInlineResult#9bebaeb9 flags:# Id:string type:string title:flags.1?string description:flags.2?string url:flags.3?string thumb_url:flags.4?string content_url:flags.5?string content_type:flags.5?string w:flags.6?int h:flags.6?int duration:flags.7?int send_message:BotInlineMessage = BotInlineResult;

const crc_botInlineResult = 0x9bebaeb9

type TL_botInlineResult struct {
	Flags        int32
	Id           string // Id:string
	Code_type    string // type:string
	Title        string // title:flags.1?string
	Description  string // description:flags.2?string
	Url          string // url:flags.3?string
	Thumb_url    string // thumb_url:flags.4?string
	Content_url  string // content_url:flags.5?string
	Content_type string // content_type:flags.5?string
	W            int32  // w:flags.6?int
	H            int32  // h:flags.6?int
	Duration     int32  // duration:flags.7?int
	Send_message TL     // send_message:BotInlineMessage
}

// Encoding TL_botInlineResult
func (e TL_botInlineResult) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInlineResult)
	var flags int32
	if e.Title != "" {
		flags |= (1 << 1)
	}
	if e.Description != "" {
		flags |= (1 << 2)
	}
	if e.Url != "" {
		flags |= (1 << 3)
	}
	if e.Thumb_url != "" {
		flags |= (1 << 4)
	}
	if e.Content_url != "" {
		flags |= (1 << 5)
	}
	if e.Content_type != "" {
		flags |= (1 << 5)
	}
	if e.W > 0 {
		flags |= (1 << 6)
	}
	if e.H > 0 {
		flags |= (1 << 6)
	}
	if e.Duration > 0 {
		flags |= (1 << 7)
	}
	x.Int(flags)
	x.String(e.Id)
	x.String(e.Code_type)
	if flags&(1<<1) != 0 {
		x.String(e.Title)
	}
	if flags&(1<<2) != 0 {
		x.String(e.Description)
	}
	if flags&(1<<3) != 0 {
		x.String(e.Url)
	}
	if flags&(1<<4) != 0 {
		x.String(e.Thumb_url)
	}
	if flags&(1<<5) != 0 {
		x.String(e.Content_url)
	}
	if flags&(1<<5) != 0 {
		x.String(e.Content_type)
	}
	if flags&(1<<6) != 0 {
		x.Int(e.W)
	}
	if flags&(1<<6) != 0 {
		x.Int(e.H)
	}
	if flags&(1<<7) != 0 {
		x.Int(e.Duration)
	}
	x.Bytes(e.Send_message.encode())
	return x.buf
}

// botInlineMediaResult#17db940b flags:# Id:string type:string photo:flags.0?Photo document:flags.1?Document title:flags.2?string description:flags.3?string send_message:BotInlineMessage = BotInlineResult;

const crc_botInlineMediaResult = 0x17db940b

type TL_botInlineMediaResult struct {
	Flags        int32
	Id           string // Id:string
	Code_type    string // type:string
	Photo        TL     // photo:flags.0?Photo
	Document     TL     // document:flags.1?Document
	Title        string // title:flags.2?string
	Description  string // description:flags.3?string
	Send_message TL     // send_message:BotInlineMessage
}

// Encoding TL_botInlineMediaResult
func (e TL_botInlineMediaResult) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInlineMediaResult)
	var flags int32
	if _, ok := (e.Photo).(TL_null); !ok {
		flags |= (1 << 0)
	}
	if _, ok := (e.Document).(TL_null); !ok {
		flags |= (1 << 1)
	}
	if e.Title != "" {
		flags |= (1 << 2)
	}
	if e.Description != "" {
		flags |= (1 << 3)
	}
	x.Int(flags)
	x.String(e.Id)
	x.String(e.Code_type)
	if flags&(1<<0) != 0 {
		x.Bytes(e.Photo.encode())
	}
	if flags&(1<<1) != 0 {
		x.Bytes(e.Document.encode())
	}
	if flags&(1<<2) != 0 {
		x.String(e.Title)
	}
	if flags&(1<<3) != 0 {
		x.String(e.Description)
	}
	x.Bytes(e.Send_message.encode())
	return x.buf
}

// messages.botResults#ccd3563d flags:# gallery:flags.0?true query_id:long next_offset:flags.1?string switch_pm:flags.2?InlineBotSwitchPM results:Vector<BotInlineResult> cache_time:int = messages.BotResults;

const crc_messages_botResults = 0xccd3563d

type TL_messages_botResults struct {
	Flags       int32
	Gallery     bool   // gallery:flags.0?true
	Query_id    int64  // query_id:long
	Next_offset string // next_offset:flags.1?string
	Switch_pm   TL     // switch_pm:flags.2?InlineBotSwitchPM
	Results     []TL   // results:Vector<BotInlineResult>
	Cache_time  int32  // cache_time:int
}

// Encoding TL_messages_botResults
func (e TL_messages_botResults) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_botResults)
	var flags int32
	if e.Gallery {
		flags |= (1 << 0)
	}
	if e.Next_offset != "" {
		flags |= (1 << 1)
	}
	if _, ok := (e.Switch_pm).(TL_null); !ok {
		flags |= (1 << 2)
	}
	x.Int(flags)
	x.Long(e.Query_id)
	if flags&(1<<1) != 0 {
		x.String(e.Next_offset)
	}
	if flags&(1<<2) != 0 {
		x.Bytes(e.Switch_pm.encode())
	}
	x.Vector(e.Results)
	x.Int(e.Cache_time)
	return x.buf
}

// exportedMessageLink#1f486803 link:string = ExportedMessageLink;

const crc_exportedMessageLink = 0x1f486803

type TL_exportedMessageLink struct {
	Link string // link:string
}

// Encoding TL_exportedMessageLink
func (e TL_exportedMessageLink) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_exportedMessageLink)
	x.String(e.Link)
	return x.buf
}

// messageFwdHeader#c786ddcb flags:# from_id:flags.0?int date:int channel_id:flags.1?int channel_post:flags.2?int = MessageFwdHeader;

const crc_messageFwdHeader = 0xc786ddcb

type TL_messageFwdHeader struct {
	Flags        int32
	From_id      int32 // from_id:flags.0?int
	Date         int32 // date:int
	Channel_id   int32 // channel_id:flags.1?int
	Channel_post int32 // channel_post:flags.2?int
}

// Encoding TL_messageFwdHeader
func (e TL_messageFwdHeader) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageFwdHeader)
	var flags int32
	if e.From_id > 0 {
		flags |= (1 << 0)
	}
	if e.Channel_id > 0 {
		flags |= (1 << 1)
	}
	if e.Channel_post > 0 {
		flags |= (1 << 2)
	}
	x.Int(flags)
	if flags&(1<<0) != 0 {
		x.Int(e.From_id)
	}
	x.Int(e.Date)
	if flags&(1<<1) != 0 {
		x.Int(e.Channel_id)
	}
	if flags&(1<<2) != 0 {
		x.Int(e.Channel_post)
	}
	return x.buf
}

// auth.codeTypeSms#72a3158c = auth.CodeType;

const crc_auth_codeTypeSms = 0x72a3158c

type TL_auth_codeTypeSms struct {
}

// Encoding TL_auth_codeTypeSms
func (e TL_auth_codeTypeSms) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_codeTypeSms)
	return x.buf
}

// auth.codeTypeCall#741cd3e3 = auth.CodeType;

const crc_auth_codeTypeCall = 0x741cd3e3

type TL_auth_codeTypeCall struct {
}

// Encoding TL_auth_codeTypeCall
func (e TL_auth_codeTypeCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_codeTypeCall)
	return x.buf
}

// auth.codeTypeFlashCall#226ccefb = auth.CodeType;

const crc_auth_codeTypeFlashCall = 0x226ccefb

type TL_auth_codeTypeFlashCall struct {
}

// Encoding TL_auth_codeTypeFlashCall
func (e TL_auth_codeTypeFlashCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_codeTypeFlashCall)
	return x.buf
}

// auth.sentCodeTypeApp#3dbb5986 length:int = auth.SentCodeType;

const crc_auth_sentCodeTypeApp = 0x3dbb5986

type TL_auth_sentCodeTypeApp struct {
	Length int32 // length:int
}

// Encoding TL_auth_sentCodeTypeApp
func (e TL_auth_sentCodeTypeApp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_sentCodeTypeApp)
	x.Int(e.Length)
	return x.buf
}

// auth.sentCodeTypeSms#c000bba2 length:int = auth.SentCodeType;

const crc_auth_sentCodeTypeSms = 0xc000bba2

type TL_auth_sentCodeTypeSms struct {
	Length int32 // length:int
}

// Encoding TL_auth_sentCodeTypeSms
func (e TL_auth_sentCodeTypeSms) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_sentCodeTypeSms)
	x.Int(e.Length)
	return x.buf
}

// auth.sentCodeTypeCall#5353e5a7 length:int = auth.SentCodeType;

const crc_auth_sentCodeTypeCall = 0x5353e5a7

type TL_auth_sentCodeTypeCall struct {
	Length int32 // length:int
}

// Encoding TL_auth_sentCodeTypeCall
func (e TL_auth_sentCodeTypeCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_sentCodeTypeCall)
	x.Int(e.Length)
	return x.buf
}

// auth.sentCodeTypeFlashCall#ab03c6d9 pattern:string = auth.SentCodeType;

const crc_auth_sentCodeTypeFlashCall = 0xab03c6d9

type TL_auth_sentCodeTypeFlashCall struct {
	Pattern string // pattern:string
}

// Encoding TL_auth_sentCodeTypeFlashCall
func (e TL_auth_sentCodeTypeFlashCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_sentCodeTypeFlashCall)
	x.String(e.Pattern)
	return x.buf
}

// messages.botCallbackAnswer#36585ea4 flags:# alert:flags.1?true has_url:flags.3?true message:flags.0?string url:flags.2?string cache_time:int = messages.BotCallbackAnswer;

const crc_messages_botCallbackAnswer = 0x36585ea4

type TL_messages_botCallbackAnswer struct {
	Flags      int32
	Alert      bool   // alert:flags.1?true
	Has_url    bool   // has_url:flags.3?true
	Message    string // message:flags.0?string
	Url        string // url:flags.2?string
	Cache_time int32  // cache_time:int
}

// Encoding TL_messages_botCallbackAnswer
func (e TL_messages_botCallbackAnswer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_botCallbackAnswer)
	var flags int32
	if e.Alert {
		flags |= (1 << 1)
	}
	if e.Has_url {
		flags |= (1 << 3)
	}
	if e.Message != "" {
		flags |= (1 << 0)
	}
	if e.Url != "" {
		flags |= (1 << 2)
	}
	x.Int(flags)
	if flags&(1<<0) != 0 {
		x.String(e.Message)
	}
	if flags&(1<<2) != 0 {
		x.String(e.Url)
	}
	x.Int(e.Cache_time)
	return x.buf
}

// messages.messageEditData#26b5dde6 flags:# caption:flags.0?true = messages.MessageEditData;

const crc_messages_messageEditData = 0x26b5dde6

type TL_messages_messageEditData struct {
	Flags   int32
	Caption bool // caption:flags.0?true
}

// Encoding TL_messages_messageEditData
func (e TL_messages_messageEditData) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_messageEditData)
	var flags int32
	if e.Caption {
		flags |= (1 << 0)
	}
	x.Int(flags)
	return x.buf
}

// inputBotInlineMessageID#890c3d89 dc_id:int Id:long access_hash:long = InputBotInlineMessageID;

const crc_inputBotInlineMessageID = 0x890c3d89

type TL_inputBotInlineMessageID struct {
	Dc_id       int32 // dc_id:int
	Id          int64 // Id:long
	Access_hash int64 // access_hash:long
}

// Encoding TL_inputBotInlineMessageID
func (e TL_inputBotInlineMessageID) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineMessageID)
	x.Int(e.Dc_id)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	return x.buf
}

// inlineBotSwitchPM#3c20629f text:string start_param:string = InlineBotSwitchPM;

const crc_inlineBotSwitchPM = 0x3c20629f

type TL_inlineBotSwitchPM struct {
	Text        string // text:string
	Start_param string // start_param:string
}

// Encoding TL_inlineBotSwitchPM
func (e TL_inlineBotSwitchPM) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inlineBotSwitchPM)
	x.String(e.Text)
	x.String(e.Start_param)
	return x.buf
}

// messages.peerDialogs#3371c354 dialogs:Vector<Dialog> messages:Vector<Message> chats:Vector<Chat> users:Vector<User> state:updates.State = messages.PeerDialogs;

const crc_messages_peerDialogs = 0x3371c354

type TL_messages_peerDialogs struct {
	Dialogs  []TL // dialogs:Vector<Dialog>
	Messages []TL // messages:Vector<Message>
	Chats    []TL // chats:Vector<Chat>
	Users    []TL // users:Vector<User>
	State    TL   // state:updates.State
}

// Encoding TL_messages_peerDialogs
func (e TL_messages_peerDialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_peerDialogs)
	x.Vector(e.Dialogs)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	x.Bytes(e.State.encode())
	return x.buf
}

// topPeer#edcdc05b peer:Peer rating:double = TopPeer;

const crc_topPeer = 0xedcdc05b

type TL_topPeer struct {
	Peer   TL      // peer:Peer
	Rating float64 // rating:double
}

// Encoding TL_topPeer
func (e TL_topPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeer)
	x.Bytes(e.Peer.encode())
	x.Double(e.Rating)
	return x.buf
}

// topPeerCategoryBotsPM#ab661b5b = TopPeerCategory;

const crc_topPeerCategoryBotsPM = 0xab661b5b

type TL_topPeerCategoryBotsPM struct {
}

// Encoding TL_topPeerCategoryBotsPM
func (e TL_topPeerCategoryBotsPM) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeerCategoryBotsPM)
	return x.buf
}

// topPeerCategoryBotsInline#148677e2 = TopPeerCategory;

const crc_topPeerCategoryBotsInline = 0x148677e2

type TL_topPeerCategoryBotsInline struct {
}

// Encoding TL_topPeerCategoryBotsInline
func (e TL_topPeerCategoryBotsInline) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeerCategoryBotsInline)
	return x.buf
}

// topPeerCategoryCorrespondents#637b7ed = TopPeerCategory;

const crc_topPeerCategoryCorrespondents = 0x637b7ed

type TL_topPeerCategoryCorrespondents struct {
}

// Encoding TL_topPeerCategoryCorrespondents
func (e TL_topPeerCategoryCorrespondents) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeerCategoryCorrespondents)
	return x.buf
}

// topPeerCategoryGroups#bd17a14a = TopPeerCategory;

const crc_topPeerCategoryGroups = 0xbd17a14a

type TL_topPeerCategoryGroups struct {
}

// Encoding TL_topPeerCategoryGroups
func (e TL_topPeerCategoryGroups) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeerCategoryGroups)
	return x.buf
}

// topPeerCategoryChannels#161d9628 = TopPeerCategory;

const crc_topPeerCategoryChannels = 0x161d9628

type TL_topPeerCategoryChannels struct {
}

// Encoding TL_topPeerCategoryChannels
func (e TL_topPeerCategoryChannels) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeerCategoryChannels)
	return x.buf
}

// topPeerCategoryPeers#fb834291 category:TopPeerCategory count:int peers:Vector<TopPeer> = TopPeerCategoryPeers;

const crc_topPeerCategoryPeers = 0xfb834291

type TL_topPeerCategoryPeers struct {
	Category TL    // category:TopPeerCategory
	Count    int32 // count:int
	Peers    []TL  // peers:Vector<TopPeer>
}

// Encoding TL_topPeerCategoryPeers
func (e TL_topPeerCategoryPeers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeerCategoryPeers)
	x.Bytes(e.Category.encode())
	x.Int(e.Count)
	x.Vector(e.Peers)
	return x.buf
}

// contacts.topPeersNotModified#de266ef5 = contacts.TopPeers;

const crc_contacts_topPeersNotModified = 0xde266ef5

type TL_contacts_topPeersNotModified struct {
}

// Encoding TL_contacts_topPeersNotModified
func (e TL_contacts_topPeersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_topPeersNotModified)
	return x.buf
}

// contacts.topPeers#70b772a8 categories:Vector<TopPeerCategoryPeers> chats:Vector<Chat> users:Vector<User> = contacts.TopPeers;

const crc_contacts_topPeers = 0x70b772a8

type TL_contacts_topPeers struct {
	Categories []TL // categories:Vector<TopPeerCategoryPeers>
	Chats      []TL // chats:Vector<Chat>
	Users      []TL // users:Vector<User>
}

// Encoding TL_contacts_topPeers
func (e TL_contacts_topPeers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_topPeers)
	x.Vector(e.Categories)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

// draftMessageEmpty#ba4baec5 = DraftMessage;

const crc_draftMessageEmpty = 0xba4baec5

type TL_draftMessageEmpty struct {
}

// Encoding TL_draftMessageEmpty
func (e TL_draftMessageEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_draftMessageEmpty)
	return x.buf
}

// draftMessage#fd8e711f flags:# no_webpage:flags.1?true reply_to_msg_id:flags.0?int message:string entities:flags.3?Vector<MessageEntity> date:int = DraftMessage;

const crc_draftMessage = 0xfd8e711f

type TL_draftMessage struct {
	Flags           int32
	No_webpage      bool   // no_webpage:flags.1?true
	Reply_to_msg_id int32  // reply_to_msg_id:flags.0?int
	Message         string // message:string
	Entities        []TL   // entities:flags.3?Vector<MessageEntity>
	Date            int32  // date:int
}

// Encoding TL_draftMessage
func (e TL_draftMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_draftMessage)
	var flags int32
	if e.No_webpage {
		flags |= (1 << 1)
	}
	if e.Reply_to_msg_id > 0 {
		flags |= (1 << 0)
	}
	if len(e.Entities) != 0 {
		flags |= (1 << 3)
	}
	x.Int(flags)
	if flags&(1<<0) != 0 {
		x.Int(e.Reply_to_msg_id)
	}
	x.String(e.Message)
	if flags&(1<<3) != 0 {
		x.Vector(e.Entities)
	}
	x.Int(e.Date)
	return x.buf
}

// messages.featuredStickersNotModified#4ede3cf = messages.FeaturedStickers;

const crc_messages_featuredStickersNotModified = 0x4ede3cf

type TL_messages_featuredStickersNotModified struct {
}

// Encoding TL_messages_featuredStickersNotModified
func (e TL_messages_featuredStickersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_featuredStickersNotModified)
	return x.buf
}

// messages.featuredStickers#f89d88e5 Hash:int sets:Vector<StickerSetCovered> unread:Vector<long> = messages.FeaturedStickers;

const crc_messages_featuredStickers = 0xf89d88e5

type TL_messages_featuredStickers struct {
	Hash   int32   // Hash:int
	Sets   []TL    // sets:Vector<StickerSetCovered>
	Unread []int64 // unread:Vector<long>
}

// Encoding TL_messages_featuredStickers
func (e TL_messages_featuredStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_featuredStickers)
	x.Int(e.Hash)
	x.Vector(e.Sets)
	x.VectorLong(e.Unread)
	return x.buf
}

// messages.recentStickersNotModified#b17f890 = messages.RecentStickers;

const crc_messages_recentStickersNotModified = 0xb17f890

type TL_messages_recentStickersNotModified struct {
}

// Encoding TL_messages_recentStickersNotModified
func (e TL_messages_recentStickersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_recentStickersNotModified)
	return x.buf
}

// messages.recentStickers#5ce20970 Hash:int stickers:Vector<Document> = messages.RecentStickers;

const crc_messages_recentStickers = 0x5ce20970

type TL_messages_recentStickers struct {
	Hash     int32 // Hash:int
	Stickers []TL  // stickers:Vector<Document>
}

// Encoding TL_messages_recentStickers
func (e TL_messages_recentStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_recentStickers)
	x.Int(e.Hash)
	x.Vector(e.Stickers)
	return x.buf
}

// messages.archivedStickers#4fcba9c8 count:int sets:Vector<StickerSetCovered> = messages.ArchivedStickers;

const crc_messages_archivedStickers = 0x4fcba9c8

type TL_messages_archivedStickers struct {
	Count int32 // count:int
	Sets  []TL  // sets:Vector<StickerSetCovered>
}

// Encoding TL_messages_archivedStickers
func (e TL_messages_archivedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_archivedStickers)
	x.Int(e.Count)
	x.Vector(e.Sets)
	return x.buf
}

// messages.stickerSetInstallResultSuccess#38641628 = messages.StickerSetInstallResult;

const crc_messages_stickerSetInstallResultSuccess = 0x38641628

type TL_messages_stickerSetInstallResultSuccess struct {
}

// Encoding TL_messages_stickerSetInstallResultSuccess
func (e TL_messages_stickerSetInstallResultSuccess) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_stickerSetInstallResultSuccess)
	return x.buf
}

// messages.stickerSetInstallResultArchive#35e410a8 sets:Vector<StickerSetCovered> = messages.StickerSetInstallResult;

const crc_messages_stickerSetInstallResultArchive = 0x35e410a8

type TL_messages_stickerSetInstallResultArchive struct {
	Sets []TL // sets:Vector<StickerSetCovered>
}

// Encoding TL_messages_stickerSetInstallResultArchive
func (e TL_messages_stickerSetInstallResultArchive) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_stickerSetInstallResultArchive)
	x.Vector(e.Sets)
	return x.buf
}

// stickerSetCovered#6410a5d2 set:StickerSet cover:Document = StickerSetCovered;

const crc_stickerSetCovered = 0x6410a5d2

type TL_stickerSetCovered struct {
	Set   TL // set:StickerSet
	Cover TL // cover:Document
}

// Encoding TL_stickerSetCovered
func (e TL_stickerSetCovered) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_stickerSetCovered)
	x.Bytes(e.Set.encode())
	x.Bytes(e.Cover.encode())
	return x.buf
}

// stickerSetMultiCovered#3407e51b set:StickerSet covers:Vector<Document> = StickerSetCovered;

const crc_stickerSetMultiCovered = 0x3407e51b

type TL_stickerSetMultiCovered struct {
	Set    TL   // set:StickerSet
	Covers []TL // covers:Vector<Document>
}

// Encoding TL_stickerSetMultiCovered
func (e TL_stickerSetMultiCovered) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_stickerSetMultiCovered)
	x.Bytes(e.Set.encode())
	x.Vector(e.Covers)
	return x.buf
}

// maskCoords#aed6dbb2 n:int x:double y:double zoom:double = MaskCoords;

const crc_maskCoords = 0xaed6dbb2

type TL_maskCoords struct {
	N    int32   // n:int
	X    float64 // x:double
	Y    float64 // y:double
	Zoom float64 // zoom:double
}

// Encoding TL_maskCoords
func (e TL_maskCoords) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_maskCoords)
	x.Int(e.N)
	x.Double(e.X)
	x.Double(e.Y)
	x.Double(e.Zoom)
	return x.buf
}

// inputStickeredMediaPhoto#4a992157 Id:InputPhoto = InputStickeredMedia;

const crc_inputStickeredMediaPhoto = 0x4a992157

type TL_inputStickeredMediaPhoto struct {
	Id TL // Id:InputPhoto
}

// Encoding TL_inputStickeredMediaPhoto
func (e TL_inputStickeredMediaPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputStickeredMediaPhoto)
	x.Bytes(e.Id.encode())
	return x.buf
}

// inputStickeredMediaDocument#438865b Id:InputDocument = InputStickeredMedia;

const crc_inputStickeredMediaDocument = 0x438865b

type TL_inputStickeredMediaDocument struct {
	Id TL // Id:InputDocument
}

// Encoding TL_inputStickeredMediaDocument
func (e TL_inputStickeredMediaDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputStickeredMediaDocument)
	x.Bytes(e.Id.encode())
	return x.buf
}

// game#bdf9653b flags:# Id:long access_hash:long short_name:string title:string description:string photo:Photo document:flags.0?Document = Game;

const crc_game = 0xbdf9653b

type TL_game struct {
	Flags       int32
	Id          int64  // Id:long
	Access_hash int64  // access_hash:long
	Short_name  string // short_name:string
	Title       string // title:string
	Description string // description:string
	Photo       TL     // photo:Photo
	Document    TL     // document:flags.0?Document
}

// Encoding TL_game
func (e TL_game) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_game)
	var flags int32
	if _, ok := (e.Document).(TL_null); !ok {
		flags |= (1 << 0)
	}
	x.Int(flags)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	x.String(e.Short_name)
	x.String(e.Title)
	x.String(e.Description)
	x.Bytes(e.Photo.encode())
	if flags&(1<<0) != 0 {
		x.Bytes(e.Document.encode())
	}
	return x.buf
}

// inputGameID#32c3e77 Id:long access_hash:long = InputGame;

const crc_inputGameID = 0x32c3e77

type TL_inputGameID struct {
	Id          int64 // Id:long
	Access_hash int64 // access_hash:long
}

// Encoding TL_inputGameID
func (e TL_inputGameID) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputGameID)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	return x.buf
}

// inputGameShortName#c331e80a bot_id:InputUser short_name:string = InputGame;

const crc_inputGameShortName = 0xc331e80a

type TL_inputGameShortName struct {
	Bot_id     TL     // bot_id:InputUser
	Short_name string // short_name:string
}

// Encoding TL_inputGameShortName
func (e TL_inputGameShortName) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputGameShortName)
	x.Bytes(e.Bot_id.encode())
	x.String(e.Short_name)
	return x.buf
}

// highScore#58fffcd0 pos:int user_id:int score:int = HighScore;

const crc_highScore = 0x58fffcd0

type TL_highScore struct {
	Pos     int32 // pos:int
	User_id int32 // user_id:int
	Score   int32 // score:int
}

// Encoding TL_highScore
func (e TL_highScore) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_highScore)
	x.Int(e.Pos)
	x.Int(e.User_id)
	x.Int(e.Score)
	return x.buf
}

// messages.highScores#9a3bfd99 scores:Vector<HighScore> users:Vector<User> = messages.HighScores;

const crc_messages_highScores = 0x9a3bfd99

type TL_messages_highScores struct {
	Scores []TL // scores:Vector<HighScore>
	Users  []TL // users:Vector<User>
}

// Encoding TL_messages_highScores
func (e TL_messages_highScores) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_highScores)
	x.Vector(e.Scores)
	x.Vector(e.Users)
	return x.buf
}

// textEmpty#dc3d824f = RichText;

const crc_textEmpty = 0xdc3d824f

type TL_textEmpty struct {
}

// Encoding TL_textEmpty
func (e TL_textEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textEmpty)
	return x.buf
}

// textPlain#744694e0 text:string = RichText;

const crc_textPlain = 0x744694e0

type TL_textPlain struct {
	Text string // text:string
}

// Encoding TL_textPlain
func (e TL_textPlain) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textPlain)
	x.String(e.Text)
	return x.buf
}

// textBold#6724abc4 text:RichText = RichText;

const crc_textBold = 0x6724abc4

type TL_textBold struct {
	Text TL // text:RichText
}

// Encoding TL_textBold
func (e TL_textBold) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textBold)
	x.Bytes(e.Text.encode())
	return x.buf
}

// textItalic#d912a59c text:RichText = RichText;

const crc_textItalic = 0xd912a59c

type TL_textItalic struct {
	Text TL // text:RichText
}

// Encoding TL_textItalic
func (e TL_textItalic) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textItalic)
	x.Bytes(e.Text.encode())
	return x.buf
}

// textUnderline#c12622c4 text:RichText = RichText;

const crc_textUnderline = 0xc12622c4

type TL_textUnderline struct {
	Text TL // text:RichText
}

// Encoding TL_textUnderline
func (e TL_textUnderline) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textUnderline)
	x.Bytes(e.Text.encode())
	return x.buf
}

// textStrike#9bf8bb95 text:RichText = RichText;

const crc_textStrike = 0x9bf8bb95

type TL_textStrike struct {
	Text TL // text:RichText
}

// Encoding TL_textStrike
func (e TL_textStrike) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textStrike)
	x.Bytes(e.Text.encode())
	return x.buf
}

// textFixed#6c3f19b9 text:RichText = RichText;

const crc_textFixed = 0x6c3f19b9

type TL_textFixed struct {
	Text TL // text:RichText
}

// Encoding TL_textFixed
func (e TL_textFixed) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textFixed)
	x.Bytes(e.Text.encode())
	return x.buf
}

// textUrl#3c2884c1 text:RichText url:string webpage_id:long = RichText;

const crc_textUrl = 0x3c2884c1

type TL_textUrl struct {
	Text       TL     // text:RichText
	Url        string // url:string
	Webpage_id int64  // webpage_id:long
}

// Encoding TL_textUrl
func (e TL_textUrl) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textUrl)
	x.Bytes(e.Text.encode())
	x.String(e.Url)
	x.Long(e.Webpage_id)
	return x.buf
}

// textEmail#de5a0dd6 text:RichText email:string = RichText;

const crc_textEmail = 0xde5a0dd6

type TL_textEmail struct {
	Text  TL     // text:RichText
	Email string // email:string
}

// Encoding TL_textEmail
func (e TL_textEmail) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textEmail)
	x.Bytes(e.Text.encode())
	x.String(e.Email)
	return x.buf
}

// textConcat#7e6260d7 texts:Vector<RichText> = RichText;

const crc_textConcat = 0x7e6260d7

type TL_textConcat struct {
	Texts []TL // texts:Vector<RichText>
}

// Encoding TL_textConcat
func (e TL_textConcat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textConcat)
	x.Vector(e.Texts)
	return x.buf
}

// pageBlockUnsupported#13567e8a = PageBlock;

const crc_pageBlockUnsupported = 0x13567e8a

type TL_pageBlockUnsupported struct {
}

// Encoding TL_pageBlockUnsupported
func (e TL_pageBlockUnsupported) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockUnsupported)
	return x.buf
}

// pageBlockTitle#70abc3fd text:RichText = PageBlock;

const crc_pageBlockTitle = 0x70abc3fd

type TL_pageBlockTitle struct {
	Text TL // text:RichText
}

// Encoding TL_pageBlockTitle
func (e TL_pageBlockTitle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockTitle)
	x.Bytes(e.Text.encode())
	return x.buf
}

// pageBlockSubtitle#8ffa9a1f text:RichText = PageBlock;

const crc_pageBlockSubtitle = 0x8ffa9a1f

type TL_pageBlockSubtitle struct {
	Text TL // text:RichText
}

// Encoding TL_pageBlockSubtitle
func (e TL_pageBlockSubtitle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockSubtitle)
	x.Bytes(e.Text.encode())
	return x.buf
}

// pageBlockAuthorDate#baafe5e0 author:RichText published_date:int = PageBlock;

const crc_pageBlockAuthorDate = 0xbaafe5e0

type TL_pageBlockAuthorDate struct {
	Author         TL    // author:RichText
	Published_date int32 // published_date:int
}

// Encoding TL_pageBlockAuthorDate
func (e TL_pageBlockAuthorDate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockAuthorDate)
	x.Bytes(e.Author.encode())
	x.Int(e.Published_date)
	return x.buf
}

// pageBlockHeader#bfd064ec text:RichText = PageBlock;

const crc_pageBlockHeader = 0xbfd064ec

type TL_pageBlockHeader struct {
	Text TL // text:RichText
}

// Encoding TL_pageBlockHeader
func (e TL_pageBlockHeader) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockHeader)
	x.Bytes(e.Text.encode())
	return x.buf
}

// pageBlockSubheader#f12bb6e1 text:RichText = PageBlock;

const crc_pageBlockSubheader = 0xf12bb6e1

type TL_pageBlockSubheader struct {
	Text TL // text:RichText
}

// Encoding TL_pageBlockSubheader
func (e TL_pageBlockSubheader) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockSubheader)
	x.Bytes(e.Text.encode())
	return x.buf
}

// pageBlockParagraph#467a0766 text:RichText = PageBlock;

const crc_pageBlockParagraph = 0x467a0766

type TL_pageBlockParagraph struct {
	Text TL // text:RichText
}

// Encoding TL_pageBlockParagraph
func (e TL_pageBlockParagraph) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockParagraph)
	x.Bytes(e.Text.encode())
	return x.buf
}

// pageBlockPreformatted#c070d93e text:RichText Language:string = PageBlock;

const crc_pageBlockPreformatted = 0xc070d93e

type TL_pageBlockPreformatted struct {
	Text     TL     // text:RichText
	Language string // Language:string
}

// Encoding TL_pageBlockPreformatted
func (e TL_pageBlockPreformatted) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockPreformatted)
	x.Bytes(e.Text.encode())
	x.String(e.Language)
	return x.buf
}

// pageBlockFooter#48870999 text:RichText = PageBlock;

const crc_pageBlockFooter = 0x48870999

type TL_pageBlockFooter struct {
	Text TL // text:RichText
}

// Encoding TL_pageBlockFooter
func (e TL_pageBlockFooter) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockFooter)
	x.Bytes(e.Text.encode())
	return x.buf
}

// pageBlockDivider#db20b188 = PageBlock;

const crc_pageBlockDivider = 0xdb20b188

type TL_pageBlockDivider struct {
}

// Encoding TL_pageBlockDivider
func (e TL_pageBlockDivider) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockDivider)
	return x.buf
}

// pageBlockAnchor#ce0d37b0 name:string = PageBlock;

const crc_pageBlockAnchor = 0xce0d37b0

type TL_pageBlockAnchor struct {
	Name string // name:string
}

// Encoding TL_pageBlockAnchor
func (e TL_pageBlockAnchor) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockAnchor)
	x.String(e.Name)
	return x.buf
}

// pageBlockList#3a58c7f4 ordered:Bool items:Vector<RichText> = PageBlock;

const crc_pageBlockList = 0x3a58c7f4

type TL_pageBlockList struct {
	Ordered TL   // ordered:Bool
	Items   []TL // items:Vector<RichText>
}

// Encoding TL_pageBlockList
func (e TL_pageBlockList) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockList)
	x.Bytes(e.Ordered.encode())
	x.Vector(e.Items)
	return x.buf
}

// pageBlockBlockquote#263d7c26 text:RichText caption:RichText = PageBlock;

const crc_pageBlockBlockquote = 0x263d7c26

type TL_pageBlockBlockquote struct {
	Text    TL // text:RichText
	Caption TL // caption:RichText
}

// Encoding TL_pageBlockBlockquote
func (e TL_pageBlockBlockquote) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockBlockquote)
	x.Bytes(e.Text.encode())
	x.Bytes(e.Caption.encode())
	return x.buf
}

// pageBlockPullquote#4f4456d3 text:RichText caption:RichText = PageBlock;

const crc_pageBlockPullquote = 0x4f4456d3

type TL_pageBlockPullquote struct {
	Text    TL // text:RichText
	Caption TL // caption:RichText
}

// Encoding TL_pageBlockPullquote
func (e TL_pageBlockPullquote) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockPullquote)
	x.Bytes(e.Text.encode())
	x.Bytes(e.Caption.encode())
	return x.buf
}

// pageBlockPhoto#e9c69982 photo_id:long caption:RichText = PageBlock;

const crc_pageBlockPhoto = 0xe9c69982

type TL_pageBlockPhoto struct {
	Photo_id int64 // photo_id:long
	Caption  TL    // caption:RichText
}

// Encoding TL_pageBlockPhoto
func (e TL_pageBlockPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockPhoto)
	x.Long(e.Photo_id)
	x.Bytes(e.Caption.encode())
	return x.buf
}

// pageBlockVideo#d9d71866 flags:# autoplay:flags.0?true loop:flags.1?true video_id:long caption:RichText = PageBlock;

const crc_pageBlockVideo = 0xd9d71866

type TL_pageBlockVideo struct {
	Flags    int32
	Autoplay bool  // autoplay:flags.0?true
	Loop     bool  // loop:flags.1?true
	Video_id int64 // video_id:long
	Caption  TL    // caption:RichText
}

// Encoding TL_pageBlockVideo
func (e TL_pageBlockVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockVideo)
	var flags int32
	if e.Autoplay {
		flags |= (1 << 0)
	}
	if e.Loop {
		flags |= (1 << 1)
	}
	x.Int(flags)
	x.Long(e.Video_id)
	x.Bytes(e.Caption.encode())
	return x.buf
}

// pageBlockCover#39f23300 cover:PageBlock = PageBlock;

const crc_pageBlockCover = 0x39f23300

type TL_pageBlockCover struct {
	Cover TL // cover:PageBlock
}

// Encoding TL_pageBlockCover
func (e TL_pageBlockCover) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockCover)
	x.Bytes(e.Cover.encode())
	return x.buf
}

// pageBlockEmbed#cde200d1 flags:# full_width:flags.0?true allow_scrolling:flags.3?true url:flags.1?string html:flags.2?string poster_photo_id:flags.4?long w:int h:int caption:RichText = PageBlock;

const crc_pageBlockEmbed = 0xcde200d1

type TL_pageBlockEmbed struct {
	Flags           int32
	Full_width      bool   // full_width:flags.0?true
	Allow_scrolling bool   // allow_scrolling:flags.3?true
	Url             string // url:flags.1?string
	Html            string // html:flags.2?string
	Poster_photo_id int64  // poster_photo_id:flags.4?long
	W               int32  // w:int
	H               int32  // h:int
	Caption         TL     // caption:RichText
}

// Encoding TL_pageBlockEmbed
func (e TL_pageBlockEmbed) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockEmbed)
	var flags int32
	if e.Full_width {
		flags |= (1 << 0)
	}
	if e.Allow_scrolling {
		flags |= (1 << 3)
	}
	if e.Url != "" {
		flags |= (1 << 1)
	}
	if e.Html != "" {
		flags |= (1 << 2)
	}
	if e.Poster_photo_id > 0 {
		flags |= (1 << 4)
	}
	x.Int(flags)
	if flags&(1<<1) != 0 {
		x.String(e.Url)
	}
	if flags&(1<<2) != 0 {
		x.String(e.Html)
	}
	if flags&(1<<4) != 0 {
		x.Long(e.Poster_photo_id)
	}
	x.Int(e.W)
	x.Int(e.H)
	x.Bytes(e.Caption.encode())
	return x.buf
}

// pageBlockEmbedPost#292c7be9 url:string webpage_id:long author_photo_id:long author:string date:int blocks:Vector<PageBlock> caption:RichText = PageBlock;

const crc_pageBlockEmbedPost = 0x292c7be9

type TL_pageBlockEmbedPost struct {
	Url             string // url:string
	Webpage_id      int64  // webpage_id:long
	Author_photo_id int64  // author_photo_id:long
	Author          string // author:string
	Date            int32  // date:int
	Blocks          []TL   // blocks:Vector<PageBlock>
	Caption         TL     // caption:RichText
}

// Encoding TL_pageBlockEmbedPost
func (e TL_pageBlockEmbedPost) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockEmbedPost)
	x.String(e.Url)
	x.Long(e.Webpage_id)
	x.Long(e.Author_photo_id)
	x.String(e.Author)
	x.Int(e.Date)
	x.Vector(e.Blocks)
	x.Bytes(e.Caption.encode())
	return x.buf
}

// pageBlockCollage#8b31c4f items:Vector<PageBlock> caption:RichText = PageBlock;

const crc_pageBlockCollage = 0x8b31c4f

type TL_pageBlockCollage struct {
	Items   []TL // items:Vector<PageBlock>
	Caption TL   // caption:RichText
}

// Encoding TL_pageBlockCollage
func (e TL_pageBlockCollage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockCollage)
	x.Vector(e.Items)
	x.Bytes(e.Caption.encode())
	return x.buf
}

// pageBlockSlideshow#130c8963 items:Vector<PageBlock> caption:RichText = PageBlock;

const crc_pageBlockSlideshow = 0x130c8963

type TL_pageBlockSlideshow struct {
	Items   []TL // items:Vector<PageBlock>
	Caption TL   // caption:RichText
}

// Encoding TL_pageBlockSlideshow
func (e TL_pageBlockSlideshow) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockSlideshow)
	x.Vector(e.Items)
	x.Bytes(e.Caption.encode())
	return x.buf
}

// pagePart#8dee6c44 blocks:Vector<PageBlock> photos:Vector<Photo> videos:Vector<Document> = Page;

const crc_pagePart = 0x8dee6c44

type TL_pagePart struct {
	Blocks []TL // blocks:Vector<PageBlock>
	Photos []TL // photos:Vector<Photo>
	Videos []TL // videos:Vector<Document>
}

// Encoding TL_pagePart
func (e TL_pagePart) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pagePart)
	x.Vector(e.Blocks)
	x.Vector(e.Photos)
	x.Vector(e.Videos)
	return x.buf
}

// pageFull#d7a19d69 blocks:Vector<PageBlock> photos:Vector<Photo> videos:Vector<Document> = Page;

const crc_pageFull = 0xd7a19d69

type TL_pageFull struct {
	Blocks []TL // blocks:Vector<PageBlock>
	Photos []TL // photos:Vector<Photo>
	Videos []TL // videos:Vector<Document>
}

// Encoding TL_pageFull
func (e TL_pageFull) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageFull)
	x.Vector(e.Blocks)
	x.Vector(e.Photos)
	x.Vector(e.Videos)
	return x.buf
}

// phoneCallDiscardReasonMissed#85e42301 = PhoneCallDiscardReason;

const crc_phoneCallDiscardReasonMissed = 0x85e42301

type TL_phoneCallDiscardReasonMissed struct {
}

// Encoding TL_phoneCallDiscardReasonMissed
func (e TL_phoneCallDiscardReasonMissed) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallDiscardReasonMissed)
	return x.buf
}

// phoneCallDiscardReasonDisconnect#e095c1a0 = PhoneCallDiscardReason;

const crc_phoneCallDiscardReasonDisconnect = 0xe095c1a0

type TL_phoneCallDiscardReasonDisconnect struct {
}

// Encoding TL_phoneCallDiscardReasonDisconnect
func (e TL_phoneCallDiscardReasonDisconnect) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallDiscardReasonDisconnect)
	return x.buf
}

// phoneCallDiscardReasonHangup#57adc690 = PhoneCallDiscardReason;

const crc_phoneCallDiscardReasonHangup = 0x57adc690

type TL_phoneCallDiscardReasonHangup struct {
}

// Encoding TL_phoneCallDiscardReasonHangup
func (e TL_phoneCallDiscardReasonHangup) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallDiscardReasonHangup)
	return x.buf
}

// phoneCallDiscardReasonBusy#faf7e8c9 = PhoneCallDiscardReason;

const crc_phoneCallDiscardReasonBusy = 0xfaf7e8c9

type TL_phoneCallDiscardReasonBusy struct {
}

// Encoding TL_phoneCallDiscardReasonBusy
func (e TL_phoneCallDiscardReasonBusy) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallDiscardReasonBusy)
	return x.buf
}

// dataJSON#7d748d04 data:string = DataJSON;

const crc_dataJSON = 0x7d748d04

type TL_dataJSON struct {
	Data string // data:string
}

// Encoding TL_dataJSON
func (e TL_dataJSON) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_dataJSON)
	x.String(e.Data)
	return x.buf
}

// labeledPrice#cb296bf8 label:string amount:long = LabeledPrice;

const crc_labeledPrice = 0xcb296bf8

type TL_labeledPrice struct {
	Label  string // label:string
	Amount int64  // amount:long
}

// Encoding TL_labeledPrice
func (e TL_labeledPrice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_labeledPrice)
	x.String(e.Label)
	x.Long(e.Amount)
	return x.buf
}

// invoice#c30aa358 flags:# test:flags.0?true name_requested:flags.1?true phone_requested:flags.2?true email_requested:flags.3?true shipping_address_requested:flags.4?true flexible:flags.5?true currency:string prices:Vector<LabeledPrice> = Invoice;

const crc_invoice = 0xc30aa358

type TL_invoice struct {
	Flags                      int32
	Test                       bool   // test:flags.0?true
	Name_requested             bool   // name_requested:flags.1?true
	Phone_requested            bool   // phone_requested:flags.2?true
	Email_requested            bool   // email_requested:flags.3?true
	Shipping_address_requested bool   // shipping_address_requested:flags.4?true
	Flexible                   bool   // flexible:flags.5?true
	Currency                   string // currency:string
	Prices                     []TL   // prices:Vector<LabeledPrice>
}

// Encoding TL_invoice
func (e TL_invoice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_invoice)
	var flags int32
	if e.Test {
		flags |= (1 << 0)
	}
	if e.Name_requested {
		flags |= (1 << 1)
	}
	if e.Phone_requested {
		flags |= (1 << 2)
	}
	if e.Email_requested {
		flags |= (1 << 3)
	}
	if e.Shipping_address_requested {
		flags |= (1 << 4)
	}
	if e.Flexible {
		flags |= (1 << 5)
	}
	x.Int(flags)
	x.String(e.Currency)
	x.Vector(e.Prices)
	return x.buf
}

// paymentCharge#ea02c27e Id:string provider_charge_id:string = PaymentCharge;

const crc_paymentCharge = 0xea02c27e

type TL_paymentCharge struct {
	Id                 string // Id:string
	Provider_charge_id string // provider_charge_id:string
}

// Encoding TL_paymentCharge
func (e TL_paymentCharge) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_paymentCharge)
	x.String(e.Id)
	x.String(e.Provider_charge_id)
	return x.buf
}

// postAddress#1e8caaeb street_line1:string street_line2:string city:string state:string country_iso2:string post_code:string = PostAddress;

const crc_postAddress = 0x1e8caaeb

type TL_postAddress struct {
	Street_line1 string // street_line1:string
	Street_line2 string // street_line2:string
	City         string // city:string
	State        string // state:string
	Country_iso2 string // country_iso2:string
	Post_code    string // post_code:string
}

// Encoding TL_postAddress
func (e TL_postAddress) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_postAddress)
	x.String(e.Street_line1)
	x.String(e.Street_line2)
	x.String(e.City)
	x.String(e.State)
	x.String(e.Country_iso2)
	x.String(e.Post_code)
	return x.buf
}

// paymentRequestedInfo#909c3f94 flags:# name:flags.0?string phone:flags.1?string email:flags.2?string shipping_address:flags.3?PostAddress = PaymentRequestedInfo;

const crc_paymentRequestedInfo = 0x909c3f94

type TL_paymentRequestedInfo struct {
	Flags            int32
	Name             string // name:flags.0?string
	Phone            string // phone:flags.1?string
	Email            string // email:flags.2?string
	Shipping_address TL     // shipping_address:flags.3?PostAddress
}

// Encoding TL_paymentRequestedInfo
func (e TL_paymentRequestedInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_paymentRequestedInfo)
	var flags int32
	if e.Name != "" {
		flags |= (1 << 0)
	}
	if e.Phone != "" {
		flags |= (1 << 1)
	}
	if e.Email != "" {
		flags |= (1 << 2)
	}
	if _, ok := (e.Shipping_address).(TL_null); !ok {
		flags |= (1 << 3)
	}
	x.Int(flags)
	if flags&(1<<0) != 0 {
		x.String(e.Name)
	}
	if flags&(1<<1) != 0 {
		x.String(e.Phone)
	}
	if flags&(1<<2) != 0 {
		x.String(e.Email)
	}
	if flags&(1<<3) != 0 {
		x.Bytes(e.Shipping_address.encode())
	}
	return x.buf
}

// paymentSavedCredentialsCard#cdc27a1f Id:string title:string = PaymentSavedCredentials;

const crc_paymentSavedCredentialsCard = 0xcdc27a1f

type TL_paymentSavedCredentialsCard struct {
	Id    string // Id:string
	Title string // title:string
}

// Encoding TL_paymentSavedCredentialsCard
func (e TL_paymentSavedCredentialsCard) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_paymentSavedCredentialsCard)
	x.String(e.Id)
	x.String(e.Title)
	return x.buf
}

// webDocument#c61acbd8 url:string access_hash:long size:int mime_type:string attributes:Vector<DocumentAttribute> dc_id:int = WebDocument;

const crc_webDocument = 0xc61acbd8

type TL_webDocument struct {
	Url         string // url:string
	Access_hash int64  // access_hash:long
	Size        int32  // size:int
	Mime_type   string // mime_type:string
	Attributes  []TL   // attributes:Vector<DocumentAttribute>
	Dc_id       int32  // dc_id:int
}

// Encoding TL_webDocument
func (e TL_webDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_webDocument)
	x.String(e.Url)
	x.Long(e.Access_hash)
	x.Int(e.Size)
	x.String(e.Mime_type)
	x.Vector(e.Attributes)
	x.Int(e.Dc_id)
	return x.buf
}

// inputWebDocument#9bed434d url:string size:int mime_type:string attributes:Vector<DocumentAttribute> = InputWebDocument;

const crc_inputWebDocument = 0x9bed434d

type TL_inputWebDocument struct {
	Url        string // url:string
	Size       int32  // size:int
	Mime_type  string // mime_type:string
	Attributes []TL   // attributes:Vector<DocumentAttribute>
}

// Encoding TL_inputWebDocument
func (e TL_inputWebDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputWebDocument)
	x.String(e.Url)
	x.Int(e.Size)
	x.String(e.Mime_type)
	x.Vector(e.Attributes)
	return x.buf
}

// inputWebFileLocation#c239d686 url:string access_hash:long = InputWebFileLocation;

const crc_inputWebFileLocation = 0xc239d686

type TL_inputWebFileLocation struct {
	Url         string // url:string
	Access_hash int64  // access_hash:long
}

// Encoding TL_inputWebFileLocation
func (e TL_inputWebFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputWebFileLocation)
	x.String(e.Url)
	x.Long(e.Access_hash)
	return x.buf
}

// upload.webFile#21e753bc size:int mime_type:string file_type:storage.FileType mtime:int bytes:bytes = upload.WebFile;

const crc_upload_webFile = 0x21e753bc

type TL_upload_webFile struct {
	Size      int32  // size:int
	Mime_type string // mime_type:string
	File_type TL     // file_type:storage.FileType
	Mtime     int32  // mtime:int
	Bytes     []byte // bytes:bytes
}

// Encoding TL_upload_webFile
func (e TL_upload_webFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_webFile)
	x.Int(e.Size)
	x.String(e.Mime_type)
	x.Bytes(e.File_type.encode())
	x.Int(e.Mtime)
	x.Bytes(e.Bytes)
	return x.buf
}

// payments.paymentForm#3f56aea3 flags:# can_save_credentials:flags.2?true password_missing:flags.3?true bot_id:int invoice:Invoice provider_id:int url:string native_provider:flags.4?string native_params:flags.4?DataJSON saved_info:flags.0?PaymentRequestedInfo saved_credentials:flags.1?PaymentSavedCredentials users:Vector<User> = payments.PaymentForm;

const crc_payments_paymentForm = 0x3f56aea3

type TL_payments_paymentForm struct {
	Flags                int32
	Can_save_credentials bool   // can_save_credentials:flags.2?true
	Password_missing     bool   // password_missing:flags.3?true
	Bot_id               int32  // bot_id:int
	Invoice              TL     // invoice:Invoice
	Provider_id          int32  // provider_id:int
	Url                  string // url:string
	Native_provider      string // native_provider:flags.4?string
	Native_params        TL     // native_params:flags.4?DataJSON
	Saved_info           TL     // saved_info:flags.0?PaymentRequestedInfo
	Saved_credentials    TL     // saved_credentials:flags.1?PaymentSavedCredentials
	Users                []TL   // users:Vector<User>
}

// Encoding TL_payments_paymentForm
func (e TL_payments_paymentForm) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_payments_paymentForm)
	var flags int32
	if e.Can_save_credentials {
		flags |= (1 << 2)
	}
	if e.Password_missing {
		flags |= (1 << 3)
	}
	if e.Native_provider != "" {
		flags |= (1 << 4)
	}
	if _, ok := (e.Native_params).(TL_null); !ok {
		flags |= (1 << 4)
	}
	if _, ok := (e.Saved_info).(TL_null); !ok {
		flags |= (1 << 0)
	}
	if _, ok := (e.Saved_credentials).(TL_null); !ok {
		flags |= (1 << 1)
	}
	x.Int(flags)
	x.Int(e.Bot_id)
	x.Bytes(e.Invoice.encode())
	x.Int(e.Provider_id)
	x.String(e.Url)
	if flags&(1<<4) != 0 {
		x.String(e.Native_provider)
	}
	if flags&(1<<4) != 0 {
		x.Bytes(e.Native_params.encode())
	}
	if flags&(1<<0) != 0 {
		x.Bytes(e.Saved_info.encode())
	}
	if flags&(1<<1) != 0 {
		x.Bytes(e.Saved_credentials.encode())
	}
	x.Vector(e.Users)
	return x.buf
}

// payments.validatedRequestedInfo#d1451883 flags:# Id:flags.0?string shipping_options:flags.1?Vector<ShippingOption> = payments.ValidatedRequestedInfo;

const crc_payments_validatedRequestedInfo = 0xd1451883

type TL_payments_validatedRequestedInfo struct {
	Flags            int32
	Id               string // Id:flags.0?string
	Shipping_options []TL   // shipping_options:flags.1?Vector<ShippingOption>
}

// Encoding TL_payments_validatedRequestedInfo
func (e TL_payments_validatedRequestedInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_payments_validatedRequestedInfo)
	var flags int32
	if e.Id != "" {
		flags |= (1 << 0)
	}
	if len(e.Shipping_options) != 0 {
		flags |= (1 << 1)
	}
	x.Int(flags)
	if flags&(1<<0) != 0 {
		x.String(e.Id)
	}
	if flags&(1<<1) != 0 {
		x.Vector(e.Shipping_options)
	}
	return x.buf
}

// payments.paymentResult#4e5f810d updates:Updates = payments.PaymentResult;

const crc_payments_paymentResult = 0x4e5f810d

type TL_payments_paymentResult struct {
	Updates TL // updates:Updates
}

// Encoding TL_payments_paymentResult
func (e TL_payments_paymentResult) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_payments_paymentResult)
	x.Bytes(e.Updates.encode())
	return x.buf
}

// payments.paymentVerficationNeeded#6b56b921 url:string = payments.PaymentResult;

const crc_payments_paymentVerficationNeeded = 0x6b56b921

type TL_payments_paymentVerficationNeeded struct {
	Url string // url:string
}

// Encoding TL_payments_paymentVerficationNeeded
func (e TL_payments_paymentVerficationNeeded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_payments_paymentVerficationNeeded)
	x.String(e.Url)
	return x.buf
}

// payments.paymentReceipt#500911e1 flags:# date:int bot_id:int invoice:Invoice provider_id:int info:flags.0?PaymentRequestedInfo shipping:flags.1?ShippingOption currency:string total_amount:long credentials_title:string users:Vector<User> = payments.PaymentReceipt;

const crc_payments_paymentReceipt = 0x500911e1

type TL_payments_paymentReceipt struct {
	Flags             int32
	Date              int32  // date:int
	Bot_id            int32  // bot_id:int
	Invoice           TL     // invoice:Invoice
	Provider_id       int32  // provider_id:int
	Info              TL     // info:flags.0?PaymentRequestedInfo
	Shipping          TL     // shipping:flags.1?ShippingOption
	Currency          string // currency:string
	Total_amount      int64  // total_amount:long
	Credentials_title string // credentials_title:string
	Users             []TL   // users:Vector<User>
}

// Encoding TL_payments_paymentReceipt
func (e TL_payments_paymentReceipt) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_payments_paymentReceipt)
	var flags int32
	if _, ok := (e.Info).(TL_null); !ok {
		flags |= (1 << 0)
	}
	if _, ok := (e.Shipping).(TL_null); !ok {
		flags |= (1 << 1)
	}
	x.Int(flags)
	x.Int(e.Date)
	x.Int(e.Bot_id)
	x.Bytes(e.Invoice.encode())
	x.Int(e.Provider_id)
	if flags&(1<<0) != 0 {
		x.Bytes(e.Info.encode())
	}
	if flags&(1<<1) != 0 {
		x.Bytes(e.Shipping.encode())
	}
	x.String(e.Currency)
	x.Long(e.Total_amount)
	x.String(e.Credentials_title)
	x.Vector(e.Users)
	return x.buf
}

// payments.savedInfo#fb8fe43c flags:# has_saved_credentials:flags.1?true saved_info:flags.0?PaymentRequestedInfo = payments.SavedInfo;

const crc_payments_savedInfo = 0xfb8fe43c

type TL_payments_savedInfo struct {
	Flags                 int32
	Has_saved_credentials bool // has_saved_credentials:flags.1?true
	Saved_info            TL   // saved_info:flags.0?PaymentRequestedInfo
}

// Encoding TL_payments_savedInfo
func (e TL_payments_savedInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_payments_savedInfo)
	var flags int32
	if e.Has_saved_credentials {
		flags |= (1 << 1)
	}
	if _, ok := (e.Saved_info).(TL_null); !ok {
		flags |= (1 << 0)
	}
	x.Int(flags)
	if flags&(1<<0) != 0 {
		x.Bytes(e.Saved_info.encode())
	}
	return x.buf
}

// inputPaymentCredentialsSaved#c10eb2cf Id:string tmp_password:bytes = InputPaymentCredentials;

const crc_inputPaymentCredentialsSaved = 0xc10eb2cf

type TL_inputPaymentCredentialsSaved struct {
	Id           string // Id:string
	Tmp_password []byte // tmp_password:bytes
}

// Encoding TL_inputPaymentCredentialsSaved
func (e TL_inputPaymentCredentialsSaved) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPaymentCredentialsSaved)
	x.String(e.Id)
	x.Bytes(e.Tmp_password)
	return x.buf
}

// inputPaymentCredentials#3417d728 flags:# save:flags.0?true data:DataJSON = InputPaymentCredentials;

const crc_inputPaymentCredentials = 0x3417d728

type TL_inputPaymentCredentials struct {
	Flags int32
	Save  bool // save:flags.0?true
	Data  TL   // data:DataJSON
}

// Encoding TL_inputPaymentCredentials
func (e TL_inputPaymentCredentials) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPaymentCredentials)
	var flags int32
	if e.Save {
		flags |= (1 << 0)
	}
	x.Int(flags)
	x.Bytes(e.Data.encode())
	return x.buf
}

// account.tmpPassword#db64fd34 tmp_password:bytes valid_until:int = account.TmpPassword;

const crc_account_tmpPassword = 0xdb64fd34

type TL_account_tmpPassword struct {
	Tmp_password []byte // tmp_password:bytes
	Valid_until  int32  // valid_until:int
}

// Encoding TL_account_tmpPassword
func (e TL_account_tmpPassword) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_tmpPassword)
	x.Bytes(e.Tmp_password)
	x.Int(e.Valid_until)
	return x.buf
}

// shippingOption#b6213cdf Id:string title:string prices:Vector<LabeledPrice> = ShippingOption;

const crc_shippingOption = 0xb6213cdf

type TL_shippingOption struct {
	Id     string // Id:string
	Title  string // title:string
	Prices []TL   // prices:Vector<LabeledPrice>
}

// Encoding TL_shippingOption
func (e TL_shippingOption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_shippingOption)
	x.String(e.Id)
	x.String(e.Title)
	x.Vector(e.Prices)
	return x.buf
}

// inputPhoneCall#1e36fded Id:long access_hash:long = InputPhoneCall;

const crc_inputPhoneCall = 0x1e36fded

type TL_inputPhoneCall struct {
	Id          int64 // Id:long
	Access_hash int64 // access_hash:long
}

// Encoding TL_inputPhoneCall
func (e TL_inputPhoneCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPhoneCall)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	return x.buf
}

// phoneCallEmpty#5366c915 Id:long = PhoneCall;

const crc_phoneCallEmpty = 0x5366c915

type TL_phoneCallEmpty struct {
	Id int64 // Id:long
}

// Encoding TL_phoneCallEmpty
func (e TL_phoneCallEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallEmpty)
	x.Long(e.Id)
	return x.buf
}

// phoneCallWaiting#1b8f4ad1 flags:# Id:long access_hash:long date:int admin_id:int participant_id:int protocol:PhoneCallProtocol receive_date:flags.0?int = PhoneCall;

const crc_phoneCallWaiting = 0x1b8f4ad1

type TL_phoneCallWaiting struct {
	Flags          int32
	Id             int64 // Id:long
	Access_hash    int64 // access_hash:long
	Date           int32 // date:int
	Admin_id       int32 // admin_id:int
	Participant_id int32 // participant_id:int
	Protocol       TL    // protocol:PhoneCallProtocol
	Receive_date   int32 // receive_date:flags.0?int
}

// Encoding TL_phoneCallWaiting
func (e TL_phoneCallWaiting) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallWaiting)
	var flags int32
	if e.Receive_date > 0 {
		flags |= (1 << 0)
	}
	x.Int(flags)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	x.Int(e.Date)
	x.Int(e.Admin_id)
	x.Int(e.Participant_id)
	x.Bytes(e.Protocol.encode())
	if flags&(1<<0) != 0 {
		x.Int(e.Receive_date)
	}
	return x.buf
}

// phoneCallRequested#83761ce4 Id:long access_hash:long date:int admin_id:int participant_id:int g_a_hash:bytes protocol:PhoneCallProtocol = PhoneCall;

const crc_phoneCallRequested = 0x83761ce4

type TL_phoneCallRequested struct {
	Id             int64  // Id:long
	Access_hash    int64  // access_hash:long
	Date           int32  // date:int
	Admin_id       int32  // admin_id:int
	Participant_id int32  // participant_id:int
	G_a_hash       []byte // g_a_hash:bytes
	Protocol       TL     // protocol:PhoneCallProtocol
}

// Encoding TL_phoneCallRequested
func (e TL_phoneCallRequested) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallRequested)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	x.Int(e.Date)
	x.Int(e.Admin_id)
	x.Int(e.Participant_id)
	x.Bytes(e.G_a_hash)
	x.Bytes(e.Protocol.encode())
	return x.buf
}

// phoneCallAccepted#6d003d3f Id:long access_hash:long date:int admin_id:int participant_id:int g_b:bytes protocol:PhoneCallProtocol = PhoneCall;

const crc_phoneCallAccepted = 0x6d003d3f

type TL_phoneCallAccepted struct {
	Id             int64  // Id:long
	Access_hash    int64  // access_hash:long
	Date           int32  // date:int
	Admin_id       int32  // admin_id:int
	Participant_id int32  // participant_id:int
	G_b            []byte // g_b:bytes
	Protocol       TL     // protocol:PhoneCallProtocol
}

// Encoding TL_phoneCallAccepted
func (e TL_phoneCallAccepted) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallAccepted)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	x.Int(e.Date)
	x.Int(e.Admin_id)
	x.Int(e.Participant_id)
	x.Bytes(e.G_b)
	x.Bytes(e.Protocol.encode())
	return x.buf
}

// phoneCall#ffe6ab67 Id:long access_hash:long date:int admin_id:int participant_id:int g_a_or_b:bytes key_fingerprint:long protocol:PhoneCallProtocol connection:PhoneConnection alternative_connections:Vector<PhoneConnection> start_date:int = PhoneCall;

const crc_phoneCall = 0xffe6ab67

type TL_phoneCall struct {
	Id                      int64  // Id:long
	Access_hash             int64  // access_hash:long
	Date                    int32  // date:int
	Admin_id                int32  // admin_id:int
	Participant_id          int32  // participant_id:int
	G_a_or_b                []byte // g_a_or_b:bytes
	Key_fingerprint         int64  // key_fingerprint:long
	Protocol                TL     // protocol:PhoneCallProtocol
	Connection              TL     // connection:PhoneConnection
	Alternative_connections []TL   // alternative_connections:Vector<PhoneConnection>
	Start_date              int32  // start_date:int
}

// Encoding TL_phoneCall
func (e TL_phoneCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCall)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	x.Int(e.Date)
	x.Int(e.Admin_id)
	x.Int(e.Participant_id)
	x.Bytes(e.G_a_or_b)
	x.Long(e.Key_fingerprint)
	x.Bytes(e.Protocol.encode())
	x.Bytes(e.Connection.encode())
	x.Vector(e.Alternative_connections)
	x.Int(e.Start_date)
	return x.buf
}

// phoneCallDiscarded#50ca4de1 flags:# need_rating:flags.2?true need_debug:flags.3?true Id:long reason:flags.0?PhoneCallDiscardReason duration:flags.1?int = PhoneCall;

const crc_phoneCallDiscarded = 0x50ca4de1

type TL_phoneCallDiscarded struct {
	Flags       int32
	Need_rating bool  // need_rating:flags.2?true
	Need_debug  bool  // need_debug:flags.3?true
	Id          int64 // Id:long
	Reason      TL    // reason:flags.0?PhoneCallDiscardReason
	Duration    int32 // duration:flags.1?int
}

// Encoding TL_phoneCallDiscarded
func (e TL_phoneCallDiscarded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallDiscarded)
	var flags int32
	if e.Need_rating {
		flags |= (1 << 2)
	}
	if e.Need_debug {
		flags |= (1 << 3)
	}
	if _, ok := (e.Reason).(TL_null); !ok {
		flags |= (1 << 0)
	}
	if e.Duration > 0 {
		flags |= (1 << 1)
	}
	x.Int(flags)
	x.Long(e.Id)
	if flags&(1<<0) != 0 {
		x.Bytes(e.Reason.encode())
	}
	if flags&(1<<1) != 0 {
		x.Int(e.Duration)
	}
	return x.buf
}

// phoneConnection#9d4c17c0 Id:long ip:string ipv6:string port:int peer_tag:bytes = PhoneConnection;

const crc_phoneConnection = 0x9d4c17c0

type TL_phoneConnection struct {
	Id       int64  // Id:long
	Ip       string // ip:string
	Ipv6     string // ipv6:string
	Port     int32  // port:int
	Peer_tag []byte // peer_tag:bytes
}

// Encoding TL_phoneConnection
func (e TL_phoneConnection) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneConnection)
	x.Long(e.Id)
	x.String(e.Ip)
	x.String(e.Ipv6)
	x.Int(e.Port)
	x.Bytes(e.Peer_tag)
	return x.buf
}

// phoneCallProtocol#a2bb35cb flags:# udp_p2p:flags.0?true udp_reflector:flags.1?true min_layer:int max_layer:int = PhoneCallProtocol;

const crc_phoneCallProtocol = 0xa2bb35cb

type TL_phoneCallProtocol struct {
	Flags         int32
	Udp_p2p       bool  // udp_p2p:flags.0?true
	Udp_reflector bool  // udp_reflector:flags.1?true
	Min_layer     int32 // min_layer:int
	Max_layer     int32 // max_layer:int
}

// Encoding TL_phoneCallProtocol
func (e TL_phoneCallProtocol) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallProtocol)
	var flags int32
	if e.Udp_p2p {
		flags |= (1 << 0)
	}
	if e.Udp_reflector {
		flags |= (1 << 1)
	}
	x.Int(flags)
	x.Int(e.Min_layer)
	x.Int(e.Max_layer)
	return x.buf
}

// phone.phoneCall#ec82e140 phone_call:PhoneCall users:Vector<User> = phone.PhoneCall;

const crc_phone_phoneCall = 0xec82e140

type TL_phone_phoneCall struct {
	Phone_call TL   // phone_call:PhoneCall
	Users      []TL // users:Vector<User>
}

// Encoding TL_phone_phoneCall
func (e TL_phone_phoneCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phone_phoneCall)
	x.Bytes(e.Phone_call.encode())
	x.Vector(e.Users)
	return x.buf
}

// invokeAfterMsg#cb9f372d {X:Type} msg_id:long query:!X = X;

const crc_invokeAfterMsg = 0xcb9f372d

type TL_invokeAfterMsg struct {
	Msg_id int64 // msg_id:long
	Query  TL    // query:!X
}

// Encoding TL_invokeAfterMsg
func (e TL_invokeAfterMsg) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_invokeAfterMsg)
	x.Long(e.Msg_id)
	x.Bytes(e.Query.encode())
	return x.buf
}

// invokeAfterMsgs#3dc4b4f0 {X:Type} msg_ids:Vector<long> query:!X = X;

const crc_invokeAfterMsgs = 0x3dc4b4f0

type TL_invokeAfterMsgs struct {
	Msg_ids []int64 // msg_ids:Vector<long>
	Query   TL      // query:!X
}

// Encoding TL_invokeAfterMsgs
func (e TL_invokeAfterMsgs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_invokeAfterMsgs)
	x.VectorLong(e.Msg_ids)
	x.Bytes(e.Query.encode())
	return x.buf
}

// initConnection#69796de9 {X:Type} api_id:int device_model:string system_version:string app_version:string lang_code:string query:!X = X;

const crc_initConnection = 0x69796de9

type TL_initConnection struct {
	Api_id         int32  // api_id:int
	Device_model   string // device_model:string
	System_version string // system_version:string
	App_version    string // app_version:string
	Lang_code      string // lang_code:string
	Query          TL     // query:!X
}

// Encoding TL_initConnection
func (e TL_initConnection) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_initConnection)
	x.Int(e.Api_id)
	x.String(e.Device_model)
	x.String(e.System_version)
	x.String(e.App_version)
	x.String(e.Lang_code)
	x.Bytes(e.Query.encode())
	return x.buf
}

// invokeWithLayer#da9b0d0d {X:Type} layer:int query:!X = X;

const crc_invokeWithLayer = 0xda9b0d0d

type TL_invokeWithLayer struct {
	Layer int32 // layer:int
	Query TL    // query:!X
}

// Encoding TL_invokeWithLayer
func (e TL_invokeWithLayer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_invokeWithLayer)
	x.Int(e.Layer)
	x.Bytes(e.Query.encode())
	return x.buf
}

// invokeWithoutUpdates#bf9459b7 {X:Type} query:!X = X;

const crc_invokeWithoutUpdates = 0xbf9459b7

type TL_invokeWithoutUpdates struct {
	Query TL // query:!X
}

// Encoding TL_invokeWithoutUpdates
func (e TL_invokeWithoutUpdates) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_invokeWithoutUpdates)
	x.Bytes(e.Query.encode())
	return x.buf
}

// auth.checkPhone#6fe51dfb phone_number:string = auth.CheckedPhone;

const crc_auth_checkPhone = 0x6fe51dfb

type TL_auth_checkPhone struct {
	Phone_number string // phone_number:string
}

// Encoding TL_auth_checkPhone
func (e TL_auth_checkPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_checkPhone)
	x.String(e.Phone_number)
	return x.buf
}

// auth.sendCode#86aef0ec flags:# allow_flashcall:flags.0?true phone_number:string current_number:flags.0?Bool api_id:int api_hash:string = auth.SentCode;

const crc_auth_sendCode = 0x86aef0ec

type TL_auth_sendCode struct {
	Flags           int32
	Allow_flashcall bool   // allow_flashcall:flags.0?true
	Phone_number    string // phone_number:string
	Current_number  TL     // current_number:flags.0?Bool
	Api_id          int32  // api_id:int
	Api_hash        string // api_hash:string
}

// Encoding TL_auth_sendCode
func (e TL_auth_sendCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_sendCode)
	var flags int32
	if e.Allow_flashcall {
		flags |= (1 << 0)
	}
	if _, ok := (e.Current_number).(TL_null); !ok {
		flags |= (1 << 0)
	}
	x.Int(flags)
	x.String(e.Phone_number)
	if flags&(1<<0) != 0 {
		x.Bytes(e.Current_number.encode())
	}
	x.Int(e.Api_id)
	x.String(e.Api_hash)
	return x.buf
}

// auth.signUp#1b067634 phone_number:string phone_code_hash:string phone_code:string first_name:string last_name:string = auth.Authorization;

const crc_auth_signUp = 0x1b067634

type TL_auth_signUp struct {
	Phone_number    string // phone_number:string
	Phone_code_hash string // phone_code_hash:string
	Phone_code      string // phone_code:string
	First_name      string // first_name:string
	Last_name       string // last_name:string
}

// Encoding TL_auth_signUp
func (e TL_auth_signUp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_signUp)
	x.String(e.Phone_number)
	x.String(e.Phone_code_hash)
	x.String(e.Phone_code)
	x.String(e.First_name)
	x.String(e.Last_name)
	return x.buf
}

// auth.signIn#bcd51581 phone_number:string phone_code_hash:string phone_code:string = auth.Authorization;

const crc_auth_signIn = 0xbcd51581

type TL_auth_signIn struct {
	Phone_number    string // phone_number:string
	Phone_code_hash string // phone_code_hash:string
	Phone_code      string // phone_code:string
}

// Encoding TL_auth_signIn
func (e TL_auth_signIn) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_signIn)
	x.String(e.Phone_number)
	x.String(e.Phone_code_hash)
	x.String(e.Phone_code)
	return x.buf
}

// auth.logOut#5717da40 = Bool;

const crc_auth_logOut = 0x5717da40

type TL_auth_logOut struct {
}

// Encoding TL_auth_logOut
func (e TL_auth_logOut) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_logOut)
	return x.buf
}

// auth.resetAuthorizations#9fab0d1a = Bool;

const crc_auth_resetAuthorizations = 0x9fab0d1a

type TL_auth_resetAuthorizations struct {
}

// Encoding TL_auth_resetAuthorizations
func (e TL_auth_resetAuthorizations) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_resetAuthorizations)
	return x.buf
}

// auth.sendInvites#771c1d97 phone_numbers:Vector<string> message:string = Bool;

const crc_auth_sendInvites = 0x771c1d97

type TL_auth_sendInvites struct {
	Phone_numbers []string // phone_numbers:Vector<string>
	Message       string   // message:string
}

// Encoding TL_auth_sendInvites
func (e TL_auth_sendInvites) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_sendInvites)
	x.VectorString(e.Phone_numbers)
	x.String(e.Message)
	return x.buf
}

// auth.exportAuthorization#e5bfffcd dc_id:int = auth.ExportedAuthorization;

const crc_auth_exportAuthorization = 0xe5bfffcd

type TL_auth_exportAuthorization struct {
	Dc_id int32 // dc_id:int
}

// Encoding TL_auth_exportAuthorization
func (e TL_auth_exportAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_exportAuthorization)
	x.Int(e.Dc_id)
	return x.buf
}

// auth.importAuthorization#e3ef9613 Id:int bytes:bytes = auth.Authorization;

const crc_auth_importAuthorization = 0xe3ef9613

type TL_auth_importAuthorization struct {
	Id    int32  // Id:int
	Bytes []byte // bytes:bytes
}

// Encoding TL_auth_importAuthorization
func (e TL_auth_importAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_importAuthorization)
	x.Int(e.Id)
	x.Bytes(e.Bytes)
	return x.buf
}

// auth.bindTempAuthKey#cdd42a05 perm_auth_key_id:long nonce:long expires_at:int encrypted_message:bytes = Bool;

const crc_auth_bindTempAuthKey = 0xcdd42a05

type TL_auth_bindTempAuthKey struct {
	Perm_auth_key_id  int64  // perm_auth_key_id:long
	Nonce             int64  // nonce:long
	Expires_at        int32  // expires_at:int
	Encrypted_message []byte // encrypted_message:bytes
}

// Encoding TL_auth_bindTempAuthKey
func (e TL_auth_bindTempAuthKey) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_bindTempAuthKey)
	x.Long(e.Perm_auth_key_id)
	x.Long(e.Nonce)
	x.Int(e.Expires_at)
	x.Bytes(e.Encrypted_message)
	return x.buf
}

// auth.importBotAuthorization#67a3ff2c flags:int api_id:int api_hash:string bot_auth_token:string = auth.Authorization;

const crc_auth_importBotAuthorization = 0x67a3ff2c

type TL_auth_importBotAuthorization struct {
	Flags          int32  // flags:int
	Api_id         int32  // api_id:int
	Api_hash       string // api_hash:string
	Bot_auth_token string // bot_auth_token:string
}

// Encoding TL_auth_importBotAuthorization
func (e TL_auth_importBotAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_importBotAuthorization)
	x.Int(e.Flags)
	x.Int(e.Api_id)
	x.String(e.Api_hash)
	x.String(e.Bot_auth_token)
	return x.buf
}

// auth.checkPassword#a63011e password_hash:bytes = auth.Authorization;

const crc_auth_checkPassword = 0xa63011e

type TL_auth_checkPassword struct {
	Password_hash []byte // password_hash:bytes
}

// Encoding TL_auth_checkPassword
func (e TL_auth_checkPassword) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_checkPassword)
	x.Bytes(e.Password_hash)
	return x.buf
}

// auth.requestPasswordRecovery#d897bc66 = auth.PasswordRecovery;

const crc_auth_requestPasswordRecovery = 0xd897bc66

type TL_auth_requestPasswordRecovery struct {
}

// Encoding TL_auth_requestPasswordRecovery
func (e TL_auth_requestPasswordRecovery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_requestPasswordRecovery)
	return x.buf
}

// auth.recoverPassword#4ea56e92 code:string = auth.Authorization;

const crc_auth_recoverPassword = 0x4ea56e92

type TL_auth_recoverPassword struct {
	Code string // code:string
}

// Encoding TL_auth_recoverPassword
func (e TL_auth_recoverPassword) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_recoverPassword)
	x.String(e.Code)
	return x.buf
}

// auth.resendCode#3ef1a9bf phone_number:string phone_code_hash:string = auth.SentCode;

const crc_auth_resendCode = 0x3ef1a9bf

type TL_auth_resendCode struct {
	Phone_number    string // phone_number:string
	Phone_code_hash string // phone_code_hash:string
}

// Encoding TL_auth_resendCode
func (e TL_auth_resendCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_resendCode)
	x.String(e.Phone_number)
	x.String(e.Phone_code_hash)
	return x.buf
}

// auth.cancelCode#1f040578 phone_number:string phone_code_hash:string = Bool;

const crc_auth_cancelCode = 0x1f040578

type TL_auth_cancelCode struct {
	Phone_number    string // phone_number:string
	Phone_code_hash string // phone_code_hash:string
}

// Encoding TL_auth_cancelCode
func (e TL_auth_cancelCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_cancelCode)
	x.String(e.Phone_number)
	x.String(e.Phone_code_hash)
	return x.buf
}

// auth.dropTempAuthKeys#8e48a188 except_auth_keys:Vector<long> = Bool;

const crc_auth_dropTempAuthKeys = 0x8e48a188

type TL_auth_dropTempAuthKeys struct {
	Except_auth_keys []int64 // except_auth_keys:Vector<long>
}

// Encoding TL_auth_dropTempAuthKeys
func (e TL_auth_dropTempAuthKeys) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_dropTempAuthKeys)
	x.VectorLong(e.Except_auth_keys)
	return x.buf
}

// account.registerDevice#637ea878 token_type:int token:string = Bool;

const crc_account_registerDevice = 0x637ea878

type TL_account_registerDevice struct {
	Token_type int32  // token_type:int
	Token      string // token:string
}

// Encoding TL_account_registerDevice
func (e TL_account_registerDevice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_registerDevice)
	x.Int(e.Token_type)
	x.String(e.Token)
	return x.buf
}

// account.unregisterDevice#65c55b40 token_type:int token:string = Bool;

const crc_account_unregisterDevice = 0x65c55b40

type TL_account_unregisterDevice struct {
	Token_type int32  // token_type:int
	Token      string // token:string
}

// Encoding TL_account_unregisterDevice
func (e TL_account_unregisterDevice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_unregisterDevice)
	x.Int(e.Token_type)
	x.String(e.Token)
	return x.buf
}

// account.updateNotifySettings#84be5b93 peer:InputNotifyPeer settings:InputPeerNotifySettings = Bool;

const crc_account_updateNotifySettings = 0x84be5b93

type TL_account_updateNotifySettings struct {
	Peer     TL // peer:InputNotifyPeer
	Settings TL // settings:InputPeerNotifySettings
}

// Encoding TL_account_updateNotifySettings
func (e TL_account_updateNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_updateNotifySettings)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Settings.encode())
	return x.buf
}

// account.getNotifySettings#12b3ad31 peer:InputNotifyPeer = PeerNotifySettings;

const crc_account_getNotifySettings = 0x12b3ad31

type TL_account_getNotifySettings struct {
	Peer TL // peer:InputNotifyPeer
}

// Encoding TL_account_getNotifySettings
func (e TL_account_getNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getNotifySettings)
	x.Bytes(e.Peer.encode())
	return x.buf
}

// account.resetNotifySettings#db7e1747 = Bool;

const crc_account_resetNotifySettings = 0xdb7e1747

type TL_account_resetNotifySettings struct {
}

// Encoding TL_account_resetNotifySettings
func (e TL_account_resetNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_resetNotifySettings)
	return x.buf
}

// account.updateProfile#78515775 flags:# first_name:flags.0?string last_name:flags.1?string about:flags.2?string = User;

const crc_account_updateProfile = 0x78515775

type TL_account_updateProfile struct {
	Flags      int32
	First_name string // first_name:flags.0?string
	Last_name  string // last_name:flags.1?string
	About      string // about:flags.2?string
}

// Encoding TL_account_updateProfile
func (e TL_account_updateProfile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_updateProfile)
	var flags int32
	if e.First_name != "" {
		flags |= (1 << 0)
	}
	if e.Last_name != "" {
		flags |= (1 << 1)
	}
	if e.About != "" {
		flags |= (1 << 2)
	}
	x.Int(flags)
	if flags&(1<<0) != 0 {
		x.String(e.First_name)
	}
	if flags&(1<<1) != 0 {
		x.String(e.Last_name)
	}
	if flags&(1<<2) != 0 {
		x.String(e.About)
	}
	return x.buf
}

// account.updateStatus#6628562c offline:Bool = Bool;

const crc_account_updateStatus = 0x6628562c

type TL_account_updateStatus struct {
	Offline TL // offline:Bool
}

// Encoding TL_account_updateStatus
func (e TL_account_updateStatus) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_updateStatus)
	x.Bytes(e.Offline.encode())
	return x.buf
}

// account.getWallPapers#c04cfac2 = Vector<WallPaper>;

const crc_account_getWallPapers = 0xc04cfac2

type TL_account_getWallPapers struct {
}

// Encoding TL_account_getWallPapers
func (e TL_account_getWallPapers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getWallPapers)
	return x.buf
}

// account.reportPeer#ae189d5f peer:InputPeer reason:ReportReason = Bool;

const crc_account_reportPeer = 0xae189d5f

type TL_account_reportPeer struct {
	Peer   TL // peer:InputPeer
	Reason TL // reason:ReportReason
}

// Encoding TL_account_reportPeer
func (e TL_account_reportPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_reportPeer)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Reason.encode())
	return x.buf
}

// account.checkUsername#2714d86c username:string = Bool;

const crc_account_checkUsername = 0x2714d86c

type TL_account_checkUsername struct {
	Username string // username:string
}

// Encoding TL_account_checkUsername
func (e TL_account_checkUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_checkUsername)
	x.String(e.Username)
	return x.buf
}

// account.updateUsername#3e0bdd7c username:string = User;

const crc_account_updateUsername = 0x3e0bdd7c

type TL_account_updateUsername struct {
	Username string // username:string
}

// Encoding TL_account_updateUsername
func (e TL_account_updateUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_updateUsername)
	x.String(e.Username)
	return x.buf
}

// account.getPrivacy#dadbc950 key:InputPrivacyKey = account.PrivacyRules;

const crc_account_getPrivacy = 0xdadbc950

type TL_account_getPrivacy struct {
	Key TL // key:InputPrivacyKey
}

// Encoding TL_account_getPrivacy
func (e TL_account_getPrivacy) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getPrivacy)
	x.Bytes(e.Key.encode())
	return x.buf
}

// account.setPrivacy#c9f81ce8 key:InputPrivacyKey rules:Vector<InputPrivacyRule> = account.PrivacyRules;

const crc_account_setPrivacy = 0xc9f81ce8

type TL_account_setPrivacy struct {
	Key   TL   // key:InputPrivacyKey
	Rules []TL // rules:Vector<InputPrivacyRule>
}

// Encoding TL_account_setPrivacy
func (e TL_account_setPrivacy) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_setPrivacy)
	x.Bytes(e.Key.encode())
	x.Vector(e.Rules)
	return x.buf
}

// account.deleteAccount#418d4e0b reason:string = Bool;

const crc_account_deleteAccount = 0x418d4e0b

type TL_account_deleteAccount struct {
	Reason string // reason:string
}

// Encoding TL_account_deleteAccount
func (e TL_account_deleteAccount) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_deleteAccount)
	x.String(e.Reason)
	return x.buf
}

// account.getAccountTTL#8fc711d = AccountDaysTTL;

const crc_account_getAccountTTL = 0x8fc711d

type TL_account_getAccountTTL struct {
}

// Encoding TL_account_getAccountTTL
func (e TL_account_getAccountTTL) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getAccountTTL)
	return x.buf
}

// account.setAccountTTL#2442485e ttl:AccountDaysTTL = Bool;

const crc_account_setAccountTTL = 0x2442485e

type TL_account_setAccountTTL struct {
	Ttl TL // ttl:AccountDaysTTL
}

// Encoding TL_account_setAccountTTL
func (e TL_account_setAccountTTL) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_setAccountTTL)
	x.Bytes(e.Ttl.encode())
	return x.buf
}

// account.sendChangePhoneCode#8e57deb flags:# allow_flashcall:flags.0?true phone_number:string current_number:flags.0?Bool = auth.SentCode;

const crc_account_sendChangePhoneCode = 0x8e57deb

type TL_account_sendChangePhoneCode struct {
	Flags           int32
	Allow_flashcall bool   // allow_flashcall:flags.0?true
	Phone_number    string // phone_number:string
	Current_number  TL     // current_number:flags.0?Bool
}

// Encoding TL_account_sendChangePhoneCode
func (e TL_account_sendChangePhoneCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_sendChangePhoneCode)
	var flags int32
	if e.Allow_flashcall {
		flags |= (1 << 0)
	}
	if _, ok := (e.Current_number).(TL_null); !ok {
		flags |= (1 << 0)
	}
	x.Int(flags)
	x.String(e.Phone_number)
	if flags&(1<<0) != 0 {
		x.Bytes(e.Current_number.encode())
	}
	return x.buf
}

// account.changePhone#70c32edb phone_number:string phone_code_hash:string phone_code:string = User;

const crc_account_changePhone = 0x70c32edb

type TL_account_changePhone struct {
	Phone_number    string // phone_number:string
	Phone_code_hash string // phone_code_hash:string
	Phone_code      string // phone_code:string
}

// Encoding TL_account_changePhone
func (e TL_account_changePhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_changePhone)
	x.String(e.Phone_number)
	x.String(e.Phone_code_hash)
	x.String(e.Phone_code)
	return x.buf
}

// account.updateDeviceLocked#38df3532 period:int = Bool;

const crc_account_updateDeviceLocked = 0x38df3532

type TL_account_updateDeviceLocked struct {
	Period int32 // period:int
}

// Encoding TL_account_updateDeviceLocked
func (e TL_account_updateDeviceLocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_updateDeviceLocked)
	x.Int(e.Period)
	return x.buf
}

// account.getAuthorizations#e320c158 = account.Authorizations;

const crc_account_getAuthorizations = 0xe320c158

type TL_account_getAuthorizations struct {
}

// Encoding TL_account_getAuthorizations
func (e TL_account_getAuthorizations) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getAuthorizations)
	return x.buf
}

// account.resetAuthorization#df77f3bc Hash:long = Bool;

const crc_account_resetAuthorization = 0xdf77f3bc

type TL_account_resetAuthorization struct {
	Hash int64 // Hash:long
}

// Encoding TL_account_resetAuthorization
func (e TL_account_resetAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_resetAuthorization)
	x.Long(e.Hash)
	return x.buf
}

// account.getPassword#548a30f5 = account.Password;

const crc_account_getPassword = 0x548a30f5

type TL_account_getPassword struct {
}

// Encoding TL_account_getPassword
func (e TL_account_getPassword) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getPassword)
	return x.buf
}

// account.getPasswordSettings#bc8d11bb current_password_hash:bytes = account.PasswordSettings;

const crc_account_getPasswordSettings = 0xbc8d11bb

type TL_account_getPasswordSettings struct {
	Current_password_hash []byte // current_password_hash:bytes
}

// Encoding TL_account_getPasswordSettings
func (e TL_account_getPasswordSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getPasswordSettings)
	x.Bytes(e.Current_password_hash)
	return x.buf
}

// account.updatePasswordSettings#fa7c4b86 current_password_hash:bytes new_settings:account.PasswordInputSettings = Bool;

const crc_account_updatePasswordSettings = 0xfa7c4b86

type TL_account_updatePasswordSettings struct {
	Current_password_hash []byte // current_password_hash:bytes
	New_settings          TL     // new_settings:account.PasswordInputSettings
}

// Encoding TL_account_updatePasswordSettings
func (e TL_account_updatePasswordSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_updatePasswordSettings)
	x.Bytes(e.Current_password_hash)
	x.Bytes(e.New_settings.encode())
	return x.buf
}

// account.sendConfirmPhoneCode#1516d7bd flags:# allow_flashcall:flags.0?true Hash:string current_number:flags.0?Bool = auth.SentCode;

const crc_account_sendConfirmPhoneCode = 0x1516d7bd

type TL_account_sendConfirmPhoneCode struct {
	Flags           int32
	Allow_flashcall bool   // allow_flashcall:flags.0?true
	Hash            string // Hash:string
	Current_number  TL     // current_number:flags.0?Bool
}

// Encoding TL_account_sendConfirmPhoneCode
func (e TL_account_sendConfirmPhoneCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_sendConfirmPhoneCode)
	var flags int32
	if e.Allow_flashcall {
		flags |= (1 << 0)
	}
	if _, ok := (e.Current_number).(TL_null); !ok {
		flags |= (1 << 0)
	}
	x.Int(flags)
	x.String(e.Hash)
	if flags&(1<<0) != 0 {
		x.Bytes(e.Current_number.encode())
	}
	return x.buf
}

// account.confirmPhone#5f2178c3 phone_code_hash:string phone_code:string = Bool;

const crc_account_confirmPhone = 0x5f2178c3

type TL_account_confirmPhone struct {
	Phone_code_hash string // phone_code_hash:string
	Phone_code      string // phone_code:string
}

// Encoding TL_account_confirmPhone
func (e TL_account_confirmPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_confirmPhone)
	x.String(e.Phone_code_hash)
	x.String(e.Phone_code)
	return x.buf
}

// account.getTmpPassword#4a82327e password_hash:bytes period:int = account.TmpPassword;

const crc_account_getTmpPassword = 0x4a82327e

type TL_account_getTmpPassword struct {
	Password_hash []byte // password_hash:bytes
	Period        int32  // period:int
}

// Encoding TL_account_getTmpPassword
func (e TL_account_getTmpPassword) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getTmpPassword)
	x.Bytes(e.Password_hash)
	x.Int(e.Period)
	return x.buf
}

// users.getUsers#d91a548 Id:Vector<InputUser> = Vector<User>;

const crc_users_getUsers = 0xd91a548

type TL_users_getUsers struct {
	Id []TL // Id:Vector<InputUser>
}

// Encoding TL_users_getUsers
func (e TL_users_getUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_users_getUsers)
	x.Vector(e.Id)
	return x.buf
}

// users.getFullUser#ca30a5b1 Id:InputUser = UserFull;

const crc_users_getFullUser = 0xca30a5b1

type TL_users_getFullUser struct {
	Id TL // Id:InputUser
}

// Encoding TL_users_getFullUser
func (e TL_users_getFullUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_users_getFullUser)
	x.Bytes(e.Id.encode())
	return x.buf
}

// contacts.getStatuses#c4a353ee = Vector<ContactStatus>;

const crc_contacts_getStatuses = 0xc4a353ee

type TL_contacts_getStatuses struct {
}

// Encoding TL_contacts_getStatuses
func (e TL_contacts_getStatuses) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_getStatuses)
	return x.buf
}

// contacts.getContacts#22c6aa08 Hash:string = contacts.Contacts;

const crc_contacts_getContacts = 0x22c6aa08

type TL_contacts_getContacts struct {
	Hash string // Hash:string
}

// Encoding TL_contacts_getContacts
func (e TL_contacts_getContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_getContacts)
	x.String(e.Hash)
	return x.buf
}

// contacts.importContacts#da30b32d contacts:Vector<InputContact> replace:Bool = contacts.ImportedContacts;

const crc_contacts_importContacts = 0xda30b32d

type TL_contacts_importContacts struct {
	Contacts []TL // contacts:Vector<InputContact>
	Replace  TL   // replace:Bool
}

// Encoding TL_contacts_importContacts
func (e TL_contacts_importContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_importContacts)
	x.Vector(e.Contacts)
	x.Bytes(e.Replace.encode())
	return x.buf
}

// contacts.deleteContact#8e953744 Id:InputUser = contacts.Link;

const crc_contacts_deleteContact = 0x8e953744

type TL_contacts_deleteContact struct {
	Id TL // Id:InputUser
}

// Encoding TL_contacts_deleteContact
func (e TL_contacts_deleteContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_deleteContact)
	x.Bytes(e.Id.encode())
	return x.buf
}

// contacts.deleteContacts#59ab389e Id:Vector<InputUser> = Bool;

const crc_contacts_deleteContacts = 0x59ab389e

type TL_contacts_deleteContacts struct {
	Id []TL // Id:Vector<InputUser>
}

// Encoding TL_contacts_deleteContacts
func (e TL_contacts_deleteContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_deleteContacts)
	x.Vector(e.Id)
	return x.buf
}

// contacts.block#332b49fc Id:InputUser = Bool;

const crc_contacts_block = 0x332b49fc

type TL_contacts_block struct {
	Id TL // Id:InputUser
}

// Encoding TL_contacts_block
func (e TL_contacts_block) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_block)
	x.Bytes(e.Id.encode())
	return x.buf
}

// contacts.unblock#e54100bd Id:InputUser = Bool;

const crc_contacts_unblock = 0xe54100bd

type TL_contacts_unblock struct {
	Id TL // Id:InputUser
}

// Encoding TL_contacts_unblock
func (e TL_contacts_unblock) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_unblock)
	x.Bytes(e.Id.encode())
	return x.buf
}

// contacts.getBlocked#f57c350f offset:int limit:int = contacts.Blocked;

const crc_contacts_getBlocked = 0xf57c350f

type TL_contacts_getBlocked struct {
	Offset int32 // offset:int
	Limit  int32 // limit:int
}

// Encoding TL_contacts_getBlocked
func (e TL_contacts_getBlocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_getBlocked)
	x.Int(e.Offset)
	x.Int(e.Limit)
	return x.buf
}

// contacts.exportCard#84e53737 = Vector<int>;

const crc_contacts_exportCard = 0x84e53737

type TL_contacts_exportCard struct {
}

// Encoding TL_contacts_exportCard
func (e TL_contacts_exportCard) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_exportCard)
	return x.buf
}

// contacts.importCard#4fe196fe export_card:Vector<int> = User;

const crc_contacts_importCard = 0x4fe196fe

type TL_contacts_importCard struct {
	Export_card []int32 // export_card:Vector<int>
}

// Encoding TL_contacts_importCard
func (e TL_contacts_importCard) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_importCard)
	x.VectorInt(e.Export_card)
	return x.buf
}

// contacts.search#11f812d8 q:string limit:int = contacts.Found;

const crc_contacts_search = 0x11f812d8

type TL_contacts_search struct {
	Q     string // q:string
	Limit int32  // limit:int
}

// Encoding TL_contacts_search
func (e TL_contacts_search) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_search)
	x.String(e.Q)
	x.Int(e.Limit)
	return x.buf
}

// contacts.resolveUsername#f93ccba3 username:string = contacts.ResolvedPeer;

const crc_contacts_resolveUsername = 0xf93ccba3

type TL_contacts_resolveUsername struct {
	Username string // username:string
}

// Encoding TL_contacts_resolveUsername
func (e TL_contacts_resolveUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_resolveUsername)
	x.String(e.Username)
	return x.buf
}

// contacts.getTopPeers#d4982db5 flags:# correspondents:flags.0?true bots_pm:flags.1?true bots_inline:flags.2?true groups:flags.10?true channels:flags.15?true offset:int limit:int Hash:int = contacts.TopPeers;

const crc_contacts_getTopPeers = 0xd4982db5

type TL_contacts_getTopPeers struct {
	Flags          int32
	Correspondents bool  // correspondents:flags.0?true
	Bots_pm        bool  // bots_pm:flags.1?true
	Bots_inline    bool  // bots_inline:flags.2?true
	Groups         bool  // groups:flags.10?true
	Channels       bool  // channels:flags.15?true
	Offset         int32 // offset:int
	Limit          int32 // limit:int
	Hash           int32 // Hash:int
}

// Encoding TL_contacts_getTopPeers
func (e TL_contacts_getTopPeers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_getTopPeers)
	var flags int32
	if e.Correspondents {
		flags |= (1 << 0)
	}
	if e.Bots_pm {
		flags |= (1 << 1)
	}
	if e.Bots_inline {
		flags |= (1 << 2)
	}
	if e.Groups {
		flags |= (1 << 10)
	}
	if e.Channels {
		flags |= (1 << 15)
	}
	x.Int(flags)
	x.Int(e.Offset)
	x.Int(e.Limit)
	x.Int(e.Hash)
	return x.buf
}

// contacts.resetTopPeerRating#1ae373ac category:TopPeerCategory peer:InputPeer = Bool;

const crc_contacts_resetTopPeerRating = 0x1ae373ac

type TL_contacts_resetTopPeerRating struct {
	Category TL // category:TopPeerCategory
	Peer     TL // peer:InputPeer
}

// Encoding TL_contacts_resetTopPeerRating
func (e TL_contacts_resetTopPeerRating) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_resetTopPeerRating)
	x.Bytes(e.Category.encode())
	x.Bytes(e.Peer.encode())
	return x.buf
}

// messages.getMessages#4222fa74 Id:Vector<int> = messages.Messages;

const crc_messages_getMessages = 0x4222fa74

type TL_messages_getMessages struct {
	Id []int32 // Id:Vector<int>
}

// Encoding TL_messages_getMessages
func (e TL_messages_getMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getMessages)
	x.VectorInt(e.Id)
	return x.buf
}

// messages.getDialogs#191ba9c5 flags:# exclude_pinned:flags.0?true offset_date:int offset_id:int offset_peer:InputPeer limit:int = messages.Dialogs;

const crc_messages_getDialogs = 0x191ba9c5

type TL_messages_getDialogs struct {
	Flags          int32
	Exclude_pinned bool  // exclude_pinned:flags.0?true
	Offset_date    int32 // offset_date:int
	Offset_id      int32 // offset_id:int
	Offset_peer    TL    // offset_peer:InputPeer
	Limit          int32 // limit:int
}

// Encoding TL_messages_getDialogs
func (e TL_messages_getDialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getDialogs)
	var flags int32
	if e.Exclude_pinned {
		flags |= (1 << 0)
	}
	x.Int(flags)
	x.Int(e.Offset_date)
	x.Int(e.Offset_id)
	x.Bytes(e.Offset_peer.encode())
	x.Int(e.Limit)
	return x.buf
}

// messages.getHistory#afa92846 peer:InputPeer offset_id:int offset_date:int add_offset:int limit:int max_id:int min_id:int = messages.Messages;

const crc_messages_getHistory = 0xafa92846

type TL_messages_getHistory struct {
	Peer        TL    // peer:InputPeer
	Offset_id   int32 // offset_id:int
	Offset_date int32 // offset_date:int
	Add_offset  int32 // add_offset:int
	Limit       int32 // limit:int
	Max_id      int32 // max_id:int
	Min_id      int32 // min_id:int
}

// Encoding TL_messages_getHistory
func (e TL_messages_getHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getHistory)
	x.Bytes(e.Peer.encode())
	x.Int(e.Offset_id)
	x.Int(e.Offset_date)
	x.Int(e.Add_offset)
	x.Int(e.Limit)
	x.Int(e.Max_id)
	x.Int(e.Min_id)
	return x.buf
}

// messages.search#d4569248 flags:# peer:InputPeer q:string filter:MessagesFilter min_date:int max_date:int offset:int max_id:int limit:int = messages.Messages;

const crc_messages_search = 0xd4569248

type TL_messages_search struct {
	Flags    int32
	Peer     TL     // peer:InputPeer
	Q        string // q:string
	Filter   TL     // filter:MessagesFilter
	Min_date int32  // min_date:int
	Max_date int32  // max_date:int
	Offset   int32  // offset:int
	Max_id   int32  // max_id:int
	Limit    int32  // limit:int
}

// Encoding TL_messages_search
func (e TL_messages_search) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_search)
	var flags int32
	x.Int(flags)
	x.Bytes(e.Peer.encode())
	x.String(e.Q)
	x.Bytes(e.Filter.encode())
	x.Int(e.Min_date)
	x.Int(e.Max_date)
	x.Int(e.Offset)
	x.Int(e.Max_id)
	x.Int(e.Limit)
	return x.buf
}

// messages.readHistory#e306d3a peer:InputPeer max_id:int = messages.AffectedMessages;

const crc_messages_readHistory = 0xe306d3a

type TL_messages_readHistory struct {
	Peer   TL    // peer:InputPeer
	Max_id int32 // max_id:int
}

// Encoding TL_messages_readHistory
func (e TL_messages_readHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_readHistory)
	x.Bytes(e.Peer.encode())
	x.Int(e.Max_id)
	return x.buf
}

// messages.deleteHistory#1c015b09 flags:# just_clear:flags.0?true peer:InputPeer max_id:int = messages.AffectedHistory;

const crc_messages_deleteHistory = 0x1c015b09

type TL_messages_deleteHistory struct {
	Flags      int32
	Just_clear bool  // just_clear:flags.0?true
	Peer       TL    // peer:InputPeer
	Max_id     int32 // max_id:int
}

// Encoding TL_messages_deleteHistory
func (e TL_messages_deleteHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_deleteHistory)
	var flags int32
	if e.Just_clear {
		flags |= (1 << 0)
	}
	x.Int(flags)
	x.Bytes(e.Peer.encode())
	x.Int(e.Max_id)
	return x.buf
}

// messages.deleteMessages#e58e95d2 flags:# revoke:flags.0?true Id:Vector<int> = messages.AffectedMessages;

const crc_messages_deleteMessages = 0xe58e95d2

type TL_messages_deleteMessages struct {
	Flags  int32
	Revoke bool    // revoke:flags.0?true
	Id     []int32 // Id:Vector<int>
}

// Encoding TL_messages_deleteMessages
func (e TL_messages_deleteMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_deleteMessages)
	var flags int32
	if e.Revoke {
		flags |= (1 << 0)
	}
	x.Int(flags)
	x.VectorInt(e.Id)
	return x.buf
}

// messages.receivedMessages#5a954c0 max_id:int = Vector<ReceivedNotifyMessage>;

const crc_messages_receivedMessages = 0x5a954c0

type TL_messages_receivedMessages struct {
	Max_id int32 // max_id:int
}

// Encoding TL_messages_receivedMessages
func (e TL_messages_receivedMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_receivedMessages)
	x.Int(e.Max_id)
	return x.buf
}

// messages.setTyping#a3825e50 peer:InputPeer action:SendMessageAction = Bool;

const crc_messages_setTyping = 0xa3825e50

type TL_messages_setTyping struct {
	Peer   TL // peer:InputPeer
	Action TL // action:SendMessageAction
}

// Encoding TL_messages_setTyping
func (e TL_messages_setTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_setTyping)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Action.encode())
	return x.buf
}

// messages.sendMessage#fa88427a flags:# no_webpage:flags.1?true silent:flags.5?true background:flags.6?true clear_draft:flags.7?true peer:InputPeer reply_to_msg_id:flags.0?int message:string random_id:long reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> = Updates;

const crc_messages_sendMessage = 0xfa88427a

type TL_messages_sendMessage struct {
	Flags           int32
	No_webpage      bool   // no_webpage:flags.1?true
	Silent          bool   // silent:flags.5?true
	Background      bool   // background:flags.6?true
	Clear_draft     bool   // clear_draft:flags.7?true
	Peer            TL     // peer:InputPeer
	Reply_to_msg_id int32  // reply_to_msg_id:flags.0?int
	Message         string // message:string
	Random_id       int64  // random_id:long
	Reply_markup    TL     // reply_markup:flags.2?ReplyMarkup
	Entities        []TL   // entities:flags.3?Vector<MessageEntity>
}

// Encoding TL_messages_sendMessage
func (e TL_messages_sendMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sendMessage)
	var flags int32
	if e.No_webpage {
		flags |= (1 << 1)
	}
	if e.Silent {
		flags |= (1 << 5)
	}
	if e.Background {
		flags |= (1 << 6)
	}
	if e.Clear_draft {
		flags |= (1 << 7)
	}
	if e.Reply_to_msg_id > 0 {
		flags |= (1 << 0)
	}
	if _, ok := (e.Reply_markup).(TL_null); !ok {
		flags |= (1 << 2)
	}
	if len(e.Entities) != 0 {
		flags |= (1 << 3)
	}
	x.Int(flags)
	x.Bytes(e.Peer.encode())
	if flags&(1<<0) != 0 {
		x.Int(e.Reply_to_msg_id)
	}
	x.String(e.Message)
	x.Long(e.Random_id)
	if flags&(1<<2) != 0 {
		x.Bytes(e.Reply_markup.encode())
	}
	if flags&(1<<3) != 0 {
		x.Vector(e.Entities)
	}
	return x.buf
}

// messages.sendMedia#c8f16791 flags:# silent:flags.5?true background:flags.6?true clear_draft:flags.7?true peer:InputPeer reply_to_msg_id:flags.0?int media:InputMedia random_id:long reply_markup:flags.2?ReplyMarkup = Updates;

const crc_messages_sendMedia = 0xc8f16791

type TL_messages_sendMedia struct {
	Flags           int32
	Silent          bool  // silent:flags.5?true
	Background      bool  // background:flags.6?true
	Clear_draft     bool  // clear_draft:flags.7?true
	Peer            TL    // peer:InputPeer
	Reply_to_msg_id int32 // reply_to_msg_id:flags.0?int
	Media           TL    // media:InputMedia
	Random_id       int64 // random_id:long
	Reply_markup    TL    // reply_markup:flags.2?ReplyMarkup
}

// Encoding TL_messages_sendMedia
func (e TL_messages_sendMedia) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sendMedia)
	var flags int32
	if e.Silent {
		flags |= (1 << 5)
	}
	if e.Background {
		flags |= (1 << 6)
	}
	if e.Clear_draft {
		flags |= (1 << 7)
	}
	if e.Reply_to_msg_id > 0 {
		flags |= (1 << 0)
	}
	if _, ok := (e.Reply_markup).(TL_null); !ok {
		flags |= (1 << 2)
	}
	x.Int(flags)
	x.Bytes(e.Peer.encode())
	if flags&(1<<0) != 0 {
		x.Int(e.Reply_to_msg_id)
	}
	x.Bytes(e.Media.encode())
	x.Long(e.Random_id)
	if flags&(1<<2) != 0 {
		x.Bytes(e.Reply_markup.encode())
	}
	return x.buf
}

// messages.forwardMessages#708e0195 flags:# silent:flags.5?true background:flags.6?true with_my_score:flags.8?true from_peer:InputPeer Id:Vector<int> random_id:Vector<long> to_peer:InputPeer = Updates;

const crc_messages_forwardMessages = 0x708e0195

type TL_messages_forwardMessages struct {
	Flags         int32
	Silent        bool    // silent:flags.5?true
	Background    bool    // background:flags.6?true
	With_my_score bool    // with_my_score:flags.8?true
	From_peer     TL      // from_peer:InputPeer
	Id            []int32 // Id:Vector<int>
	Random_id     []int64 // random_id:Vector<long>
	To_peer       TL      // to_peer:InputPeer
}

// Encoding TL_messages_forwardMessages
func (e TL_messages_forwardMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_forwardMessages)
	var flags int32
	if e.Silent {
		flags |= (1 << 5)
	}
	if e.Background {
		flags |= (1 << 6)
	}
	if e.With_my_score {
		flags |= (1 << 8)
	}
	x.Int(flags)
	x.Bytes(e.From_peer.encode())
	x.VectorInt(e.Id)
	x.VectorLong(e.Random_id)
	x.Bytes(e.To_peer.encode())
	return x.buf
}

// messages.reportSpam#cf1592db peer:InputPeer = Bool;

const crc_messages_reportSpam = 0xcf1592db

type TL_messages_reportSpam struct {
	Peer TL // peer:InputPeer
}

// Encoding TL_messages_reportSpam
func (e TL_messages_reportSpam) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_reportSpam)
	x.Bytes(e.Peer.encode())
	return x.buf
}

// messages.hideReportSpam#a8f1709b peer:InputPeer = Bool;

const crc_messages_hideReportSpam = 0xa8f1709b

type TL_messages_hideReportSpam struct {
	Peer TL // peer:InputPeer
}

// Encoding TL_messages_hideReportSpam
func (e TL_messages_hideReportSpam) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_hideReportSpam)
	x.Bytes(e.Peer.encode())
	return x.buf
}

// messages.getPeerSettings#3672e09c peer:InputPeer = PeerSettings;

const crc_messages_getPeerSettings = 0x3672e09c

type TL_messages_getPeerSettings struct {
	Peer TL // peer:InputPeer
}

// Encoding TL_messages_getPeerSettings
func (e TL_messages_getPeerSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getPeerSettings)
	x.Bytes(e.Peer.encode())
	return x.buf
}

// messages.getChats#3c6aa187 Id:Vector<int> = messages.Chats;

const crc_messages_getChats = 0x3c6aa187

type TL_messages_getChats struct {
	Id []int32 // Id:Vector<int>
}

// Encoding TL_messages_getChats
func (e TL_messages_getChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getChats)
	x.VectorInt(e.Id)
	return x.buf
}

// messages.getFullChat#3b831c66 chat_id:int = messages.ChatFull;

const crc_messages_getFullChat = 0x3b831c66

type TL_messages_getFullChat struct {
	Chat_id int32 // chat_id:int
}

// Encoding TL_messages_getFullChat
func (e TL_messages_getFullChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getFullChat)
	x.Int(e.Chat_id)
	return x.buf
}

// messages.editChatTitle#dc452855 chat_id:int title:string = Updates;

const crc_messages_editChatTitle = 0xdc452855

type TL_messages_editChatTitle struct {
	Chat_id int32  // chat_id:int
	Title   string // title:string
}

// Encoding TL_messages_editChatTitle
func (e TL_messages_editChatTitle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_editChatTitle)
	x.Int(e.Chat_id)
	x.String(e.Title)
	return x.buf
}

// messages.editChatPhoto#ca4c79d8 chat_id:int photo:InputChatPhoto = Updates;

const crc_messages_editChatPhoto = 0xca4c79d8

type TL_messages_editChatPhoto struct {
	Chat_id int32 // chat_id:int
	Photo   TL    // photo:InputChatPhoto
}

// Encoding TL_messages_editChatPhoto
func (e TL_messages_editChatPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_editChatPhoto)
	x.Int(e.Chat_id)
	x.Bytes(e.Photo.encode())
	return x.buf
}

// messages.addChatUser#f9a0aa09 chat_id:int user_id:InputUser fwd_limit:int = Updates;

const crc_messages_addChatUser = 0xf9a0aa09

type TL_messages_addChatUser struct {
	Chat_id   int32 // chat_id:int
	User_id   TL    // user_id:InputUser
	Fwd_limit int32 // fwd_limit:int
}

// Encoding TL_messages_addChatUser
func (e TL_messages_addChatUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_addChatUser)
	x.Int(e.Chat_id)
	x.Bytes(e.User_id.encode())
	x.Int(e.Fwd_limit)
	return x.buf
}

// messages.deleteChatUser#e0611f16 chat_id:int user_id:InputUser = Updates;

const crc_messages_deleteChatUser = 0xe0611f16

type TL_messages_deleteChatUser struct {
	Chat_id int32 // chat_id:int
	User_id TL    // user_id:InputUser
}

// Encoding TL_messages_deleteChatUser
func (e TL_messages_deleteChatUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_deleteChatUser)
	x.Int(e.Chat_id)
	x.Bytes(e.User_id.encode())
	return x.buf
}

// messages.createChat#9cb126e users:Vector<InputUser> title:string = Updates;

const crc_messages_createChat = 0x9cb126e

type TL_messages_createChat struct {
	Users []TL   // users:Vector<InputUser>
	Title string // title:string
}

// Encoding TL_messages_createChat
func (e TL_messages_createChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_createChat)
	x.Vector(e.Users)
	x.String(e.Title)
	return x.buf
}

// messages.forwardMessage#33963bf9 peer:InputPeer Id:int random_id:long = Updates;

const crc_messages_forwardMessage = 0x33963bf9

type TL_messages_forwardMessage struct {
	Peer      TL    // peer:InputPeer
	Id        int32 // Id:int
	Random_id int64 // random_id:long
}

// Encoding TL_messages_forwardMessage
func (e TL_messages_forwardMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_forwardMessage)
	x.Bytes(e.Peer.encode())
	x.Int(e.Id)
	x.Long(e.Random_id)
	return x.buf
}

// messages.getDhConfig#26cf8950 Version:int random_length:int = messages.DhConfig;

const crc_messages_getDhConfig = 0x26cf8950

type TL_messages_getDhConfig struct {
	Version       int32 // Version:int
	Random_length int32 // random_length:int
}

// Encoding TL_messages_getDhConfig
func (e TL_messages_getDhConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getDhConfig)
	x.Int(e.Version)
	x.Int(e.Random_length)
	return x.buf
}

// messages.requestEncryption#f64daf43 user_id:InputUser random_id:int g_a:bytes = EncryptedChat;

const crc_messages_requestEncryption = 0xf64daf43

type TL_messages_requestEncryption struct {
	User_id   TL     // user_id:InputUser
	Random_id int32  // random_id:int
	G_a       []byte // g_a:bytes
}

// Encoding TL_messages_requestEncryption
func (e TL_messages_requestEncryption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_requestEncryption)
	x.Bytes(e.User_id.encode())
	x.Int(e.Random_id)
	x.Bytes(e.G_a)
	return x.buf
}

// messages.acceptEncryption#3dbc0415 peer:InputEncryptedChat g_b:bytes key_fingerprint:long = EncryptedChat;

const crc_messages_acceptEncryption = 0x3dbc0415

type TL_messages_acceptEncryption struct {
	Peer            TL     // peer:InputEncryptedChat
	G_b             []byte // g_b:bytes
	Key_fingerprint int64  // key_fingerprint:long
}

// Encoding TL_messages_acceptEncryption
func (e TL_messages_acceptEncryption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_acceptEncryption)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.G_b)
	x.Long(e.Key_fingerprint)
	return x.buf
}

// messages.discardEncryption#edd923c5 chat_id:int = Bool;

const crc_messages_discardEncryption = 0xedd923c5

type TL_messages_discardEncryption struct {
	Chat_id int32 // chat_id:int
}

// Encoding TL_messages_discardEncryption
func (e TL_messages_discardEncryption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_discardEncryption)
	x.Int(e.Chat_id)
	return x.buf
}

// messages.setEncryptedTyping#791451ed peer:InputEncryptedChat typing:Bool = Bool;

const crc_messages_setEncryptedTyping = 0x791451ed

type TL_messages_setEncryptedTyping struct {
	Peer   TL // peer:InputEncryptedChat
	Typing TL // typing:Bool
}

// Encoding TL_messages_setEncryptedTyping
func (e TL_messages_setEncryptedTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_setEncryptedTyping)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Typing.encode())
	return x.buf
}

// messages.readEncryptedHistory#7f4b690a peer:InputEncryptedChat max_date:int = Bool;

const crc_messages_readEncryptedHistory = 0x7f4b690a

type TL_messages_readEncryptedHistory struct {
	Peer     TL    // peer:InputEncryptedChat
	Max_date int32 // max_date:int
}

// Encoding TL_messages_readEncryptedHistory
func (e TL_messages_readEncryptedHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_readEncryptedHistory)
	x.Bytes(e.Peer.encode())
	x.Int(e.Max_date)
	return x.buf
}

// messages.sendEncrypted#a9776773 peer:InputEncryptedChat random_id:long data:bytes = messages.SentEncryptedMessage;

const crc_messages_sendEncrypted = 0xa9776773

type TL_messages_sendEncrypted struct {
	Peer      TL     // peer:InputEncryptedChat
	Random_id int64  // random_id:long
	Data      []byte // data:bytes
}

// Encoding TL_messages_sendEncrypted
func (e TL_messages_sendEncrypted) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sendEncrypted)
	x.Bytes(e.Peer.encode())
	x.Long(e.Random_id)
	x.Bytes(e.Data)
	return x.buf
}

// messages.sendEncryptedFile#9a901b66 peer:InputEncryptedChat random_id:long data:bytes file:InputEncryptedFile = messages.SentEncryptedMessage;

const crc_messages_sendEncryptedFile = 0x9a901b66

type TL_messages_sendEncryptedFile struct {
	Peer      TL     // peer:InputEncryptedChat
	Random_id int64  // random_id:long
	Data      []byte // data:bytes
	File      TL     // file:InputEncryptedFile
}

// Encoding TL_messages_sendEncryptedFile
func (e TL_messages_sendEncryptedFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sendEncryptedFile)
	x.Bytes(e.Peer.encode())
	x.Long(e.Random_id)
	x.Bytes(e.Data)
	x.Bytes(e.File.encode())
	return x.buf
}

// messages.sendEncryptedService#32d439a4 peer:InputEncryptedChat random_id:long data:bytes = messages.SentEncryptedMessage;

const crc_messages_sendEncryptedService = 0x32d439a4

type TL_messages_sendEncryptedService struct {
	Peer      TL     // peer:InputEncryptedChat
	Random_id int64  // random_id:long
	Data      []byte // data:bytes
}

// Encoding TL_messages_sendEncryptedService
func (e TL_messages_sendEncryptedService) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sendEncryptedService)
	x.Bytes(e.Peer.encode())
	x.Long(e.Random_id)
	x.Bytes(e.Data)
	return x.buf
}

// messages.receivedQueue#55a5bb66 max_qts:int = Vector<long>;

const crc_messages_receivedQueue = 0x55a5bb66

type TL_messages_receivedQueue struct {
	Max_qts int32 // max_qts:int
}

// Encoding TL_messages_receivedQueue
func (e TL_messages_receivedQueue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_receivedQueue)
	x.Int(e.Max_qts)
	return x.buf
}

// messages.reportEncryptedSpam#4b0c8c0f peer:InputEncryptedChat = Bool;

const crc_messages_reportEncryptedSpam = 0x4b0c8c0f

type TL_messages_reportEncryptedSpam struct {
	Peer TL // peer:InputEncryptedChat
}

// Encoding TL_messages_reportEncryptedSpam
func (e TL_messages_reportEncryptedSpam) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_reportEncryptedSpam)
	x.Bytes(e.Peer.encode())
	return x.buf
}

// messages.readMessageContents#36a73f77 Id:Vector<int> = messages.AffectedMessages;

const crc_messages_readMessageContents = 0x36a73f77

type TL_messages_readMessageContents struct {
	Id []int32 // Id:Vector<int>
}

// Encoding TL_messages_readMessageContents
func (e TL_messages_readMessageContents) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_readMessageContents)
	x.VectorInt(e.Id)
	return x.buf
}

// messages.getAllStickers#1c9618b1 Hash:int = messages.AllStickers;

const crc_messages_getAllStickers = 0x1c9618b1

type TL_messages_getAllStickers struct {
	Hash int32 // Hash:int
}

// Encoding TL_messages_getAllStickers
func (e TL_messages_getAllStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getAllStickers)
	x.Int(e.Hash)
	return x.buf
}

// messages.getWebPagePreview#25223e24 message:string = MessageMedia;

const crc_messages_getWebPagePreview = 0x25223e24

type TL_messages_getWebPagePreview struct {
	Message string // message:string
}

// Encoding TL_messages_getWebPagePreview
func (e TL_messages_getWebPagePreview) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getWebPagePreview)
	x.String(e.Message)
	return x.buf
}

// messages.exportChatInvite#7d885289 chat_id:int = ExportedChatInvite;

const crc_messages_exportChatInvite = 0x7d885289

type TL_messages_exportChatInvite struct {
	Chat_id int32 // chat_id:int
}

// Encoding TL_messages_exportChatInvite
func (e TL_messages_exportChatInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_exportChatInvite)
	x.Int(e.Chat_id)
	return x.buf
}

// messages.checkChatInvite#3eadb1bb Hash:string = ChatInvite;

const crc_messages_checkChatInvite = 0x3eadb1bb

type TL_messages_checkChatInvite struct {
	Hash string // Hash:string
}

// Encoding TL_messages_checkChatInvite
func (e TL_messages_checkChatInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_checkChatInvite)
	x.String(e.Hash)
	return x.buf
}

// messages.importChatInvite#6c50051c Hash:string = Updates;

const crc_messages_importChatInvite = 0x6c50051c

type TL_messages_importChatInvite struct {
	Hash string // Hash:string
}

// Encoding TL_messages_importChatInvite
func (e TL_messages_importChatInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_importChatInvite)
	x.String(e.Hash)
	return x.buf
}

// messages.getStickerSet#2619a90e stickerset:InputStickerSet = messages.StickerSet;

const crc_messages_getStickerSet = 0x2619a90e

type TL_messages_getStickerSet struct {
	Stickerset TL // stickerset:InputStickerSet
}

// Encoding TL_messages_getStickerSet
func (e TL_messages_getStickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getStickerSet)
	x.Bytes(e.Stickerset.encode())
	return x.buf
}

// messages.installStickerSet#c78fe460 stickerset:InputStickerSet archived:Bool = messages.StickerSetInstallResult;

const crc_messages_installStickerSet = 0xc78fe460

type TL_messages_installStickerSet struct {
	Stickerset TL // stickerset:InputStickerSet
	Archived   TL // archived:Bool
}

// Encoding TL_messages_installStickerSet
func (e TL_messages_installStickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_installStickerSet)
	x.Bytes(e.Stickerset.encode())
	x.Bytes(e.Archived.encode())
	return x.buf
}

// messages.uninstallStickerSet#f96e55de stickerset:InputStickerSet = Bool;

const crc_messages_uninstallStickerSet = 0xf96e55de

type TL_messages_uninstallStickerSet struct {
	Stickerset TL // stickerset:InputStickerSet
}

// Encoding TL_messages_uninstallStickerSet
func (e TL_messages_uninstallStickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_uninstallStickerSet)
	x.Bytes(e.Stickerset.encode())
	return x.buf
}

// messages.startBot#e6df7378 bot:InputUser peer:InputPeer random_id:long start_param:string = Updates;

const crc_messages_startBot = 0xe6df7378

type TL_messages_startBot struct {
	Bot         TL     // bot:InputUser
	Peer        TL     // peer:InputPeer
	Random_id   int64  // random_id:long
	Start_param string // start_param:string
}

// Encoding TL_messages_startBot
func (e TL_messages_startBot) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_startBot)
	x.Bytes(e.Bot.encode())
	x.Bytes(e.Peer.encode())
	x.Long(e.Random_id)
	x.String(e.Start_param)
	return x.buf
}

// messages.getMessagesViews#c4c8a55d peer:InputPeer Id:Vector<int> increment:Bool = Vector<int>;

const crc_messages_getMessagesViews = 0xc4c8a55d

type TL_messages_getMessagesViews struct {
	Peer      TL      // peer:InputPeer
	Id        []int32 // Id:Vector<int>
	Increment TL      // increment:Bool
}

// Encoding TL_messages_getMessagesViews
func (e TL_messages_getMessagesViews) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getMessagesViews)
	x.Bytes(e.Peer.encode())
	x.VectorInt(e.Id)
	x.Bytes(e.Increment.encode())
	return x.buf
}

// messages.toggleChatAdmins#ec8bd9e1 chat_id:int enabled:Bool = Updates;

const crc_messages_toggleChatAdmins = 0xec8bd9e1

type TL_messages_toggleChatAdmins struct {
	Chat_id int32 // chat_id:int
	Enabled TL    // enabled:Bool
}

// Encoding TL_messages_toggleChatAdmins
func (e TL_messages_toggleChatAdmins) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_toggleChatAdmins)
	x.Int(e.Chat_id)
	x.Bytes(e.Enabled.encode())
	return x.buf
}

// messages.editChatAdmin#a9e69f2e chat_id:int user_id:InputUser is_admin:Bool = Bool;

const crc_messages_editChatAdmin = 0xa9e69f2e

type TL_messages_editChatAdmin struct {
	Chat_id  int32 // chat_id:int
	User_id  TL    // user_id:InputUser
	Is_admin TL    // is_admin:Bool
}

// Encoding TL_messages_editChatAdmin
func (e TL_messages_editChatAdmin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_editChatAdmin)
	x.Int(e.Chat_id)
	x.Bytes(e.User_id.encode())
	x.Bytes(e.Is_admin.encode())
	return x.buf
}

// messages.migrateChat#15a3b8e3 chat_id:int = Updates;

const crc_messages_migrateChat = 0x15a3b8e3

type TL_messages_migrateChat struct {
	Chat_id int32 // chat_id:int
}

// Encoding TL_messages_migrateChat
func (e TL_messages_migrateChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_migrateChat)
	x.Int(e.Chat_id)
	return x.buf
}

// messages.searchGlobal#9e3cacb0 q:string offset_date:int offset_peer:InputPeer offset_id:int limit:int = messages.Messages;

const crc_messages_searchGlobal = 0x9e3cacb0

type TL_messages_searchGlobal struct {
	Q           string // q:string
	Offset_date int32  // offset_date:int
	Offset_peer TL     // offset_peer:InputPeer
	Offset_id   int32  // offset_id:int
	Limit       int32  // limit:int
}

// Encoding TL_messages_searchGlobal
func (e TL_messages_searchGlobal) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_searchGlobal)
	x.String(e.Q)
	x.Int(e.Offset_date)
	x.Bytes(e.Offset_peer.encode())
	x.Int(e.Offset_id)
	x.Int(e.Limit)
	return x.buf
}

// messages.reorderStickerSets#78337739 flags:# masks:flags.0?true order:Vector<long> = Bool;

const crc_messages_reorderStickerSets = 0x78337739

type TL_messages_reorderStickerSets struct {
	Flags int32
	Masks bool    // masks:flags.0?true
	Order []int64 // order:Vector<long>
}

// Encoding TL_messages_reorderStickerSets
func (e TL_messages_reorderStickerSets) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_reorderStickerSets)
	var flags int32
	if e.Masks {
		flags |= (1 << 0)
	}
	x.Int(flags)
	x.VectorLong(e.Order)
	return x.buf
}

// messages.getDocumentByHash#338e2464 sha256:bytes size:int mime_type:string = Document;

const crc_messages_getDocumentByHash = 0x338e2464

type TL_messages_getDocumentByHash struct {
	Sha256    []byte // sha256:bytes
	Size      int32  // size:int
	Mime_type string // mime_type:string
}

// Encoding TL_messages_getDocumentByHash
func (e TL_messages_getDocumentByHash) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getDocumentByHash)
	x.Bytes(e.Sha256)
	x.Int(e.Size)
	x.String(e.Mime_type)
	return x.buf
}

// messages.searchGifs#bf9a776b q:string offset:int = messages.FoundGifs;

const crc_messages_searchGifs = 0xbf9a776b

type TL_messages_searchGifs struct {
	Q      string // q:string
	Offset int32  // offset:int
}

// Encoding TL_messages_searchGifs
func (e TL_messages_searchGifs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_searchGifs)
	x.String(e.Q)
	x.Int(e.Offset)
	return x.buf
}

// messages.getSavedGifs#83bf3d52 Hash:int = messages.SavedGifs;

const crc_messages_getSavedGifs = 0x83bf3d52

type TL_messages_getSavedGifs struct {
	Hash int32 // Hash:int
}

// Encoding TL_messages_getSavedGifs
func (e TL_messages_getSavedGifs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getSavedGifs)
	x.Int(e.Hash)
	return x.buf
}

// messages.saveGif#327a30cb Id:InputDocument unsave:Bool = Bool;

const crc_messages_saveGif = 0x327a30cb

type TL_messages_saveGif struct {
	Id     TL // Id:InputDocument
	Unsave TL // unsave:Bool
}

// Encoding TL_messages_saveGif
func (e TL_messages_saveGif) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_saveGif)
	x.Bytes(e.Id.encode())
	x.Bytes(e.Unsave.encode())
	return x.buf
}

// messages.getInlineBotResults#514e999d flags:# bot:InputUser peer:InputPeer geo_point:flags.0?InputGeoPoint query:string offset:string = messages.BotResults;

const crc_messages_getInlineBotResults = 0x514e999d

type TL_messages_getInlineBotResults struct {
	Flags     int32
	Bot       TL     // bot:InputUser
	Peer      TL     // peer:InputPeer
	Geo_point TL     // geo_point:flags.0?InputGeoPoint
	Query     string // query:string
	Offset    string // offset:string
}

// Encoding TL_messages_getInlineBotResults
func (e TL_messages_getInlineBotResults) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getInlineBotResults)
	var flags int32
	if _, ok := (e.Geo_point).(TL_null); !ok {
		flags |= (1 << 0)
	}
	x.Int(flags)
	x.Bytes(e.Bot.encode())
	x.Bytes(e.Peer.encode())
	if flags&(1<<0) != 0 {
		x.Bytes(e.Geo_point.encode())
	}
	x.String(e.Query)
	x.String(e.Offset)
	return x.buf
}

// messages.setInlineBotResults#eb5ea206 flags:# gallery:flags.0?true private:flags.1?true query_id:long results:Vector<InputBotInlineResult> cache_time:int next_offset:flags.2?string switch_pm:flags.3?InlineBotSwitchPM = Bool;

const crc_messages_setInlineBotResults = 0xeb5ea206

type TL_messages_setInlineBotResults struct {
	Flags       int32
	Gallery     bool   // gallery:flags.0?true
	Private     bool   // private:flags.1?true
	Query_id    int64  // query_id:long
	Results     []TL   // results:Vector<InputBotInlineResult>
	Cache_time  int32  // cache_time:int
	Next_offset string // next_offset:flags.2?string
	Switch_pm   TL     // switch_pm:flags.3?InlineBotSwitchPM
}

// Encoding TL_messages_setInlineBotResults
func (e TL_messages_setInlineBotResults) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_setInlineBotResults)
	var flags int32
	if e.Gallery {
		flags |= (1 << 0)
	}
	if e.Private {
		flags |= (1 << 1)
	}
	if e.Next_offset != "" {
		flags |= (1 << 2)
	}
	if _, ok := (e.Switch_pm).(TL_null); !ok {
		flags |= (1 << 3)
	}
	x.Int(flags)
	x.Long(e.Query_id)
	x.Vector(e.Results)
	x.Int(e.Cache_time)
	if flags&(1<<2) != 0 {
		x.String(e.Next_offset)
	}
	if flags&(1<<3) != 0 {
		x.Bytes(e.Switch_pm.encode())
	}
	return x.buf
}

// messages.sendInlineBotResult#b16e06fe flags:# silent:flags.5?true background:flags.6?true clear_draft:flags.7?true peer:InputPeer reply_to_msg_id:flags.0?int random_id:long query_id:long Id:string = Updates;

const crc_messages_sendInlineBotResult = 0xb16e06fe

type TL_messages_sendInlineBotResult struct {
	Flags           int32
	Silent          bool   // silent:flags.5?true
	Background      bool   // background:flags.6?true
	Clear_draft     bool   // clear_draft:flags.7?true
	Peer            TL     // peer:InputPeer
	Reply_to_msg_id int32  // reply_to_msg_id:flags.0?int
	Random_id       int64  // random_id:long
	Query_id        int64  // query_id:long
	Id              string // Id:string
}

// Encoding TL_messages_sendInlineBotResult
func (e TL_messages_sendInlineBotResult) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sendInlineBotResult)
	var flags int32
	if e.Silent {
		flags |= (1 << 5)
	}
	if e.Background {
		flags |= (1 << 6)
	}
	if e.Clear_draft {
		flags |= (1 << 7)
	}
	if e.Reply_to_msg_id > 0 {
		flags |= (1 << 0)
	}
	x.Int(flags)
	x.Bytes(e.Peer.encode())
	if flags&(1<<0) != 0 {
		x.Int(e.Reply_to_msg_id)
	}
	x.Long(e.Random_id)
	x.Long(e.Query_id)
	x.String(e.Id)
	return x.buf
}

// messages.getMessageEditData#fda68d36 peer:InputPeer Id:int = messages.MessageEditData;

const crc_messages_getMessageEditData = 0xfda68d36

type TL_messages_getMessageEditData struct {
	Peer TL    // peer:InputPeer
	Id   int32 // Id:int
}

// Encoding TL_messages_getMessageEditData
func (e TL_messages_getMessageEditData) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getMessageEditData)
	x.Bytes(e.Peer.encode())
	x.Int(e.Id)
	return x.buf
}

// messages.editMessage#ce91e4ca flags:# no_webpage:flags.1?true peer:InputPeer Id:int message:flags.11?string reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> = Updates;

const crc_messages_editMessage = 0xce91e4ca

type TL_messages_editMessage struct {
	Flags        int32
	No_webpage   bool   // no_webpage:flags.1?true
	Peer         TL     // peer:InputPeer
	Id           int32  // Id:int
	Message      string // message:flags.11?string
	Reply_markup TL     // reply_markup:flags.2?ReplyMarkup
	Entities     []TL   // entities:flags.3?Vector<MessageEntity>
}

// Encoding TL_messages_editMessage
func (e TL_messages_editMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_editMessage)
	var flags int32
	if e.No_webpage {
		flags |= (1 << 1)
	}
	if e.Message != "" {
		flags |= (1 << 11)
	}
	if _, ok := (e.Reply_markup).(TL_null); !ok {
		flags |= (1 << 2)
	}
	if len(e.Entities) != 0 {
		flags |= (1 << 3)
	}
	x.Int(flags)
	x.Bytes(e.Peer.encode())
	x.Int(e.Id)
	if flags&(1<<11) != 0 {
		x.String(e.Message)
	}
	if flags&(1<<2) != 0 {
		x.Bytes(e.Reply_markup.encode())
	}
	if flags&(1<<3) != 0 {
		x.Vector(e.Entities)
	}
	return x.buf
}

// messages.editInlineBotMessage#130c2c85 flags:# no_webpage:flags.1?true Id:InputBotInlineMessageID message:flags.11?string reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> = Bool;

const crc_messages_editInlineBotMessage = 0x130c2c85

type TL_messages_editInlineBotMessage struct {
	Flags        int32
	No_webpage   bool   // no_webpage:flags.1?true
	Id           TL     // Id:InputBotInlineMessageID
	Message      string // message:flags.11?string
	Reply_markup TL     // reply_markup:flags.2?ReplyMarkup
	Entities     []TL   // entities:flags.3?Vector<MessageEntity>
}

// Encoding TL_messages_editInlineBotMessage
func (e TL_messages_editInlineBotMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_editInlineBotMessage)
	var flags int32
	if e.No_webpage {
		flags |= (1 << 1)
	}
	if e.Message != "" {
		flags |= (1 << 11)
	}
	if _, ok := (e.Reply_markup).(TL_null); !ok {
		flags |= (1 << 2)
	}
	if len(e.Entities) != 0 {
		flags |= (1 << 3)
	}
	x.Int(flags)
	x.Bytes(e.Id.encode())
	if flags&(1<<11) != 0 {
		x.String(e.Message)
	}
	if flags&(1<<2) != 0 {
		x.Bytes(e.Reply_markup.encode())
	}
	if flags&(1<<3) != 0 {
		x.Vector(e.Entities)
	}
	return x.buf
}

// messages.getBotCallbackAnswer#810a9fec flags:# game:flags.1?true peer:InputPeer msg_id:int data:flags.0?bytes = messages.BotCallbackAnswer;

const crc_messages_getBotCallbackAnswer = 0x810a9fec

type TL_messages_getBotCallbackAnswer struct {
	Flags  int32
	Game   bool   // game:flags.1?true
	Peer   TL     // peer:InputPeer
	Msg_id int32  // msg_id:int
	Data   []byte // data:flags.0?bytes
}

// Encoding TL_messages_getBotCallbackAnswer
func (e TL_messages_getBotCallbackAnswer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getBotCallbackAnswer)
	var flags int32
	if e.Game {
		flags |= (1 << 1)
	}
	if len(e.Data) != 0 {
		flags |= (1 << 0)
	}
	x.Int(flags)
	x.Bytes(e.Peer.encode())
	x.Int(e.Msg_id)
	if flags&(1<<0) != 0 {
		x.Bytes(e.Data)
	}
	return x.buf
}

// messages.setBotCallbackAnswer#d58f130a flags:# alert:flags.1?true query_id:long message:flags.0?string url:flags.2?string cache_time:int = Bool;

const crc_messages_setBotCallbackAnswer = 0xd58f130a

type TL_messages_setBotCallbackAnswer struct {
	Flags      int32
	Alert      bool   // alert:flags.1?true
	Query_id   int64  // query_id:long
	Message    string // message:flags.0?string
	Url        string // url:flags.2?string
	Cache_time int32  // cache_time:int
}

// Encoding TL_messages_setBotCallbackAnswer
func (e TL_messages_setBotCallbackAnswer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_setBotCallbackAnswer)
	var flags int32
	if e.Alert {
		flags |= (1 << 1)
	}
	if e.Message != "" {
		flags |= (1 << 0)
	}
	if e.Url != "" {
		flags |= (1 << 2)
	}
	x.Int(flags)
	x.Long(e.Query_id)
	if flags&(1<<0) != 0 {
		x.String(e.Message)
	}
	if flags&(1<<2) != 0 {
		x.String(e.Url)
	}
	x.Int(e.Cache_time)
	return x.buf
}

// messages.getPeerDialogs#2d9776b9 peers:Vector<InputPeer> = messages.PeerDialogs;

const crc_messages_getPeerDialogs = 0x2d9776b9

type TL_messages_getPeerDialogs struct {
	Peers []TL // peers:Vector<InputPeer>
}

// Encoding TL_messages_getPeerDialogs
func (e TL_messages_getPeerDialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getPeerDialogs)
	x.Vector(e.Peers)
	return x.buf
}

// messages.saveDraft#bc39e14b flags:# no_webpage:flags.1?true reply_to_msg_id:flags.0?int peer:InputPeer message:string entities:flags.3?Vector<MessageEntity> = Bool;

const crc_messages_saveDraft = 0xbc39e14b

type TL_messages_saveDraft struct {
	Flags           int32
	No_webpage      bool   // no_webpage:flags.1?true
	Reply_to_msg_id int32  // reply_to_msg_id:flags.0?int
	Peer            TL     // peer:InputPeer
	Message         string // message:string
	Entities        []TL   // entities:flags.3?Vector<MessageEntity>
}

// Encoding TL_messages_saveDraft
func (e TL_messages_saveDraft) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_saveDraft)
	var flags int32
	if e.No_webpage {
		flags |= (1 << 1)
	}
	if e.Reply_to_msg_id > 0 {
		flags |= (1 << 0)
	}
	if len(e.Entities) != 0 {
		flags |= (1 << 3)
	}
	x.Int(flags)
	if flags&(1<<0) != 0 {
		x.Int(e.Reply_to_msg_id)
	}
	x.Bytes(e.Peer.encode())
	x.String(e.Message)
	if flags&(1<<3) != 0 {
		x.Vector(e.Entities)
	}
	return x.buf
}

// messages.getAllDrafts#6a3f8d65 = Updates;

const crc_messages_getAllDrafts = 0x6a3f8d65

type TL_messages_getAllDrafts struct {
}

// Encoding TL_messages_getAllDrafts
func (e TL_messages_getAllDrafts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getAllDrafts)
	return x.buf
}

// messages.getFeaturedStickers#2dacca4f Hash:int = messages.FeaturedStickers;

const crc_messages_getFeaturedStickers = 0x2dacca4f

type TL_messages_getFeaturedStickers struct {
	Hash int32 // Hash:int
}

// Encoding TL_messages_getFeaturedStickers
func (e TL_messages_getFeaturedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getFeaturedStickers)
	x.Int(e.Hash)
	return x.buf
}

// messages.readFeaturedStickers#5b118126 Id:Vector<long> = Bool;

const crc_messages_readFeaturedStickers = 0x5b118126

type TL_messages_readFeaturedStickers struct {
	Id []int64 // Id:Vector<long>
}

// Encoding TL_messages_readFeaturedStickers
func (e TL_messages_readFeaturedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_readFeaturedStickers)
	x.VectorLong(e.Id)
	return x.buf
}

// messages.getRecentStickers#5ea192c9 flags:# attached:flags.0?true Hash:int = messages.RecentStickers;

const crc_messages_getRecentStickers = 0x5ea192c9

type TL_messages_getRecentStickers struct {
	Flags    int32
	Attached bool  // attached:flags.0?true
	Hash     int32 // Hash:int
}

// Encoding TL_messages_getRecentStickers
func (e TL_messages_getRecentStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getRecentStickers)
	var flags int32
	if e.Attached {
		flags |= (1 << 0)
	}
	x.Int(flags)
	x.Int(e.Hash)
	return x.buf
}

// messages.saveRecentSticker#392718f8 flags:# attached:flags.0?true Id:InputDocument unsave:Bool = Bool;

const crc_messages_saveRecentSticker = 0x392718f8

type TL_messages_saveRecentSticker struct {
	Flags    int32
	Attached bool // attached:flags.0?true
	Id       TL   // Id:InputDocument
	Unsave   TL   // unsave:Bool
}

// Encoding TL_messages_saveRecentSticker
func (e TL_messages_saveRecentSticker) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_saveRecentSticker)
	var flags int32
	if e.Attached {
		flags |= (1 << 0)
	}
	x.Int(flags)
	x.Bytes(e.Id.encode())
	x.Bytes(e.Unsave.encode())
	return x.buf
}

// messages.clearRecentStickers#8999602d flags:# attached:flags.0?true = Bool;

const crc_messages_clearRecentStickers = 0x8999602d

type TL_messages_clearRecentStickers struct {
	Flags    int32
	Attached bool // attached:flags.0?true
}

// Encoding TL_messages_clearRecentStickers
func (e TL_messages_clearRecentStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_clearRecentStickers)
	var flags int32
	if e.Attached {
		flags |= (1 << 0)
	}
	x.Int(flags)
	return x.buf
}

// messages.getArchivedStickers#57f17692 flags:# masks:flags.0?true offset_id:long limit:int = messages.ArchivedStickers;

const crc_messages_getArchivedStickers = 0x57f17692

type TL_messages_getArchivedStickers struct {
	Flags     int32
	Masks     bool  // masks:flags.0?true
	Offset_id int64 // offset_id:long
	Limit     int32 // limit:int
}

// Encoding TL_messages_getArchivedStickers
func (e TL_messages_getArchivedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getArchivedStickers)
	var flags int32
	if e.Masks {
		flags |= (1 << 0)
	}
	x.Int(flags)
	x.Long(e.Offset_id)
	x.Int(e.Limit)
	return x.buf
}

// messages.getMaskStickers#65b8c79f Hash:int = messages.AllStickers;

const crc_messages_getMaskStickers = 0x65b8c79f

type TL_messages_getMaskStickers struct {
	Hash int32 // Hash:int
}

// Encoding TL_messages_getMaskStickers
func (e TL_messages_getMaskStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getMaskStickers)
	x.Int(e.Hash)
	return x.buf
}

// messages.getAttachedStickers#cc5b67cc media:InputStickeredMedia = Vector<StickerSetCovered>;

const crc_messages_getAttachedStickers = 0xcc5b67cc

type TL_messages_getAttachedStickers struct {
	Media TL // media:InputStickeredMedia
}

// Encoding TL_messages_getAttachedStickers
func (e TL_messages_getAttachedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getAttachedStickers)
	x.Bytes(e.Media.encode())
	return x.buf
}

// messages.setGameScore#8ef8ecc0 flags:# edit_message:flags.0?true force:flags.1?true peer:InputPeer Id:int user_id:InputUser score:int = Updates;

const crc_messages_setGameScore = 0x8ef8ecc0

type TL_messages_setGameScore struct {
	Flags        int32
	Edit_message bool  // edit_message:flags.0?true
	Force        bool  // force:flags.1?true
	Peer         TL    // peer:InputPeer
	Id           int32 // Id:int
	User_id      TL    // user_id:InputUser
	Score        int32 // score:int
}

// Encoding TL_messages_setGameScore
func (e TL_messages_setGameScore) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_setGameScore)
	var flags int32
	if e.Edit_message {
		flags |= (1 << 0)
	}
	if e.Force {
		flags |= (1 << 1)
	}
	x.Int(flags)
	x.Bytes(e.Peer.encode())
	x.Int(e.Id)
	x.Bytes(e.User_id.encode())
	x.Int(e.Score)
	return x.buf
}

// messages.setInlineGameScore#15ad9f64 flags:# edit_message:flags.0?true force:flags.1?true Id:InputBotInlineMessageID user_id:InputUser score:int = Bool;

const crc_messages_setInlineGameScore = 0x15ad9f64

type TL_messages_setInlineGameScore struct {
	Flags        int32
	Edit_message bool  // edit_message:flags.0?true
	Force        bool  // force:flags.1?true
	Id           TL    // Id:InputBotInlineMessageID
	User_id      TL    // user_id:InputUser
	Score        int32 // score:int
}

// Encoding TL_messages_setInlineGameScore
func (e TL_messages_setInlineGameScore) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_setInlineGameScore)
	var flags int32
	if e.Edit_message {
		flags |= (1 << 0)
	}
	if e.Force {
		flags |= (1 << 1)
	}
	x.Int(flags)
	x.Bytes(e.Id.encode())
	x.Bytes(e.User_id.encode())
	x.Int(e.Score)
	return x.buf
}

// messages.getGameHighScores#e822649d peer:InputPeer Id:int user_id:InputUser = messages.HighScores;

const crc_messages_getGameHighScores = 0xe822649d

type TL_messages_getGameHighScores struct {
	Peer    TL    // peer:InputPeer
	Id      int32 // Id:int
	User_id TL    // user_id:InputUser
}

// Encoding TL_messages_getGameHighScores
func (e TL_messages_getGameHighScores) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getGameHighScores)
	x.Bytes(e.Peer.encode())
	x.Int(e.Id)
	x.Bytes(e.User_id.encode())
	return x.buf
}

// messages.getInlineGameHighScores#f635e1b Id:InputBotInlineMessageID user_id:InputUser = messages.HighScores;

const crc_messages_getInlineGameHighScores = 0xf635e1b

type TL_messages_getInlineGameHighScores struct {
	Id      TL // Id:InputBotInlineMessageID
	User_id TL // user_id:InputUser
}

// Encoding TL_messages_getInlineGameHighScores
func (e TL_messages_getInlineGameHighScores) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getInlineGameHighScores)
	x.Bytes(e.Id.encode())
	x.Bytes(e.User_id.encode())
	return x.buf
}

// messages.getCommonChats#d0a48c4 user_id:InputUser max_id:int limit:int = messages.Chats;

const crc_messages_getCommonChats = 0xd0a48c4

type TL_messages_getCommonChats struct {
	User_id TL    // user_id:InputUser
	Max_id  int32 // max_id:int
	Limit   int32 // limit:int
}

// Encoding TL_messages_getCommonChats
func (e TL_messages_getCommonChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getCommonChats)
	x.Bytes(e.User_id.encode())
	x.Int(e.Max_id)
	x.Int(e.Limit)
	return x.buf
}

// messages.getAllChats#eba80ff0 except_ids:Vector<int> = messages.Chats;

const crc_messages_getAllChats = 0xeba80ff0

type TL_messages_getAllChats struct {
	Except_ids []int32 // except_ids:Vector<int>
}

// Encoding TL_messages_getAllChats
func (e TL_messages_getAllChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getAllChats)
	x.VectorInt(e.Except_ids)
	return x.buf
}

// messages.getWebPage#32ca8f91 url:string Hash:int = WebPage;

const crc_messages_getWebPage = 0x32ca8f91

type TL_messages_getWebPage struct {
	Url  string // url:string
	Hash int32  // Hash:int
}

// Encoding TL_messages_getWebPage
func (e TL_messages_getWebPage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getWebPage)
	x.String(e.Url)
	x.Int(e.Hash)
	return x.buf
}

// messages.toggleDialogPin#3289be6a flags:# pinned:flags.0?true peer:InputPeer = Bool;

const crc_messages_toggleDialogPin = 0x3289be6a

type TL_messages_toggleDialogPin struct {
	Flags  int32
	Pinned bool // pinned:flags.0?true
	Peer   TL   // peer:InputPeer
}

// Encoding TL_messages_toggleDialogPin
func (e TL_messages_toggleDialogPin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_toggleDialogPin)
	var flags int32
	if e.Pinned {
		flags |= (1 << 0)
	}
	x.Int(flags)
	x.Bytes(e.Peer.encode())
	return x.buf
}

// messages.reorderPinnedDialogs#959ff644 flags:# force:flags.0?true order:Vector<InputPeer> = Bool;

const crc_messages_reorderPinnedDialogs = 0x959ff644

type TL_messages_reorderPinnedDialogs struct {
	Flags int32
	Force bool // force:flags.0?true
	Order []TL // order:Vector<InputPeer>
}

// Encoding TL_messages_reorderPinnedDialogs
func (e TL_messages_reorderPinnedDialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_reorderPinnedDialogs)
	var flags int32
	if e.Force {
		flags |= (1 << 0)
	}
	x.Int(flags)
	x.Vector(e.Order)
	return x.buf
}

// messages.getPinnedDialogs#e254d64e = messages.PeerDialogs;

const crc_messages_getPinnedDialogs = 0xe254d64e

type TL_messages_getPinnedDialogs struct {
}

// Encoding TL_messages_getPinnedDialogs
func (e TL_messages_getPinnedDialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getPinnedDialogs)
	return x.buf
}

// messages.setBotShippingResults#e5f672fa flags:# query_id:long error:flags.0?string shipping_options:flags.1?Vector<ShippingOption> = Bool;

const crc_messages_setBotShippingResults = 0xe5f672fa

type TL_messages_setBotShippingResults struct {
	Flags            int32
	Query_id         int64  // query_id:long
	Error            string // error:flags.0?string
	Shipping_options []TL   // shipping_options:flags.1?Vector<ShippingOption>
}

// Encoding TL_messages_setBotShippingResults
func (e TL_messages_setBotShippingResults) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_setBotShippingResults)
	var flags int32
	if e.Error != "" {
		flags |= (1 << 0)
	}
	if len(e.Shipping_options) != 0 {
		flags |= (1 << 1)
	}
	x.Int(flags)
	x.Long(e.Query_id)
	if flags&(1<<0) != 0 {
		x.String(e.Error)
	}
	if flags&(1<<1) != 0 {
		x.Vector(e.Shipping_options)
	}
	return x.buf
}

// messages.setBotPrecheckoutResults#9c2dd95 flags:# success:flags.1?true query_id:long error:flags.0?string = Bool;

const crc_messages_setBotPrecheckoutResults = 0x9c2dd95

type TL_messages_setBotPrecheckoutResults struct {
	Flags    int32
	Success  bool   // success:flags.1?true
	Query_id int64  // query_id:long
	Error    string // error:flags.0?string
}

// Encoding TL_messages_setBotPrecheckoutResults
func (e TL_messages_setBotPrecheckoutResults) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_setBotPrecheckoutResults)
	var flags int32
	if e.Success {
		flags |= (1 << 1)
	}
	if e.Error != "" {
		flags |= (1 << 0)
	}
	x.Int(flags)
	x.Long(e.Query_id)
	if flags&(1<<0) != 0 {
		x.String(e.Error)
	}
	return x.buf
}

// updates.getState#edd4882a = updates.State;

const crc_updates_getState = 0xedd4882a

type TL_updates_getState struct {
}

// Encoding TL_updates_getState
func (e TL_updates_getState) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_getState)
	return x.buf
}

// updates.getDifference#25939651 flags:# pts:int pts_total_limit:flags.0?int date:int qts:int = updates.Difference;

const crc_updates_getDifference = 0x25939651

type TL_updates_getDifference struct {
	Flags           int32
	Pts             int32 // pts:int
	Pts_total_limit int32 // pts_total_limit:flags.0?int
	Date            int32 // date:int
	Qts             int32 // qts:int
}

// Encoding TL_updates_getDifference
func (e TL_updates_getDifference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_getDifference)
	var flags int32
	if e.Pts_total_limit > 0 {
		flags |= (1 << 0)
	}
	x.Int(flags)
	x.Int(e.Pts)
	if flags&(1<<0) != 0 {
		x.Int(e.Pts_total_limit)
	}
	x.Int(e.Date)
	x.Int(e.Qts)
	return x.buf
}

// updates.getChannelDifference#3173d78 flags:# force:flags.0?true channel:InputChannel filter:ChannelMessagesFilter pts:int limit:int = updates.ChannelDifference;

const crc_updates_getChannelDifference = 0x3173d78

type TL_updates_getChannelDifference struct {
	Flags   int32
	Force   bool  // force:flags.0?true
	Channel TL    // channel:InputChannel
	Filter  TL    // filter:ChannelMessagesFilter
	Pts     int32 // pts:int
	Limit   int32 // limit:int
}

// Encoding TL_updates_getChannelDifference
func (e TL_updates_getChannelDifference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_getChannelDifference)
	var flags int32
	if e.Force {
		flags |= (1 << 0)
	}
	x.Int(flags)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.Filter.encode())
	x.Int(e.Pts)
	x.Int(e.Limit)
	return x.buf
}

// photos.updateProfilePhoto#f0bb5152 Id:InputPhoto = UserProfilePhoto;

const crc_photos_updateProfilePhoto = 0xf0bb5152

type TL_photos_updateProfilePhoto struct {
	Id TL // Id:InputPhoto
}

// Encoding TL_photos_updateProfilePhoto
func (e TL_photos_updateProfilePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photos_updateProfilePhoto)
	x.Bytes(e.Id.encode())
	return x.buf
}

// photos.uploadProfilePhoto#4f32c098 file:InputFile = photos.Photo;

const crc_photos_uploadProfilePhoto = 0x4f32c098

type TL_photos_uploadProfilePhoto struct {
	File TL // file:InputFile
}

// Encoding TL_photos_uploadProfilePhoto
func (e TL_photos_uploadProfilePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photos_uploadProfilePhoto)
	x.Bytes(e.File.encode())
	return x.buf
}

// photos.deletePhotos#87cf7f2f Id:Vector<InputPhoto> = Vector<long>;

const crc_photos_deletePhotos = 0x87cf7f2f

type TL_photos_deletePhotos struct {
	Id []TL // Id:Vector<InputPhoto>
}

// Encoding TL_photos_deletePhotos
func (e TL_photos_deletePhotos) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photos_deletePhotos)
	x.Vector(e.Id)
	return x.buf
}

// photos.getUserPhotos#91cd32a8 user_id:InputUser offset:int max_id:long limit:int = photos.Photos;

const crc_photos_getUserPhotos = 0x91cd32a8

type TL_photos_getUserPhotos struct {
	User_id TL    // user_id:InputUser
	Offset  int32 // offset:int
	Max_id  int64 // max_id:long
	Limit   int32 // limit:int
}

// Encoding TL_photos_getUserPhotos
func (e TL_photos_getUserPhotos) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photos_getUserPhotos)
	x.Bytes(e.User_id.encode())
	x.Int(e.Offset)
	x.Long(e.Max_id)
	x.Int(e.Limit)
	return x.buf
}

// upload.saveFilePart#b304a621 file_id:long file_part:int bytes:bytes = Bool;

const crc_upload_saveFilePart = 0xb304a621

type TL_upload_saveFilePart struct {
	File_id   int64  // file_id:long
	File_part int32  // file_part:int
	Bytes     []byte // bytes:bytes
}

// Encoding TL_upload_saveFilePart
func (e TL_upload_saveFilePart) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_saveFilePart)
	x.Long(e.File_id)
	x.Int(e.File_part)
	x.Bytes(e.Bytes)
	return x.buf
}

// upload.getFile#e3a6cfb5 location:InputFileLocation offset:int limit:int = upload.File;

const crc_upload_getFile = 0xe3a6cfb5

type TL_upload_getFile struct {
	Location TL    // location:InputFileLocation
	Offset   int32 // offset:int
	Limit    int32 // limit:int
}

// Encoding TL_upload_getFile
func (e TL_upload_getFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_getFile)
	x.Bytes(e.Location.encode())
	x.Int(e.Offset)
	x.Int(e.Limit)
	return x.buf
}

// upload.saveBigFilePart#de7b673d file_id:long file_part:int file_total_parts:int bytes:bytes = Bool;

const crc_upload_saveBigFilePart = 0xde7b673d

type TL_upload_saveBigFilePart struct {
	File_id          int64  // file_id:long
	File_part        int32  // file_part:int
	File_total_parts int32  // file_total_parts:int
	Bytes            []byte // bytes:bytes
}

// Encoding TL_upload_saveBigFilePart
func (e TL_upload_saveBigFilePart) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_saveBigFilePart)
	x.Long(e.File_id)
	x.Int(e.File_part)
	x.Int(e.File_total_parts)
	x.Bytes(e.Bytes)
	return x.buf
}

// upload.getWebFile#24e6818d location:InputWebFileLocation offset:int limit:int = upload.WebFile;

const crc_upload_getWebFile = 0x24e6818d

type TL_upload_getWebFile struct {
	Location TL    // location:InputWebFileLocation
	Offset   int32 // offset:int
	Limit    int32 // limit:int
}

// Encoding TL_upload_getWebFile
func (e TL_upload_getWebFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_getWebFile)
	x.Bytes(e.Location.encode())
	x.Int(e.Offset)
	x.Int(e.Limit)
	return x.buf
}

// help.getConfig#c4f9186b = Config;

const crc_help_getConfig = 0xc4f9186b

type TL_help_getConfig struct {
}

// Encoding TL_help_getConfig
func (e TL_help_getConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getConfig)
	return x.buf
}

// help.getNearestDc#1fb33026 = NearestDc;

const crc_help_getNearestDc = 0x1fb33026

type TL_help_getNearestDc struct {
}

// Encoding TL_help_getNearestDc
func (e TL_help_getNearestDc) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getNearestDc)
	return x.buf
}

// help.getAppUpdate#ae2de196 = help.AppUpdate;

const crc_help_getAppUpdate = 0xae2de196

type TL_help_getAppUpdate struct {
}

// Encoding TL_help_getAppUpdate
func (e TL_help_getAppUpdate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getAppUpdate)
	return x.buf
}

// help.saveAppLog#6f02f748 events:Vector<InputAppEvent> = Bool;

const crc_help_saveAppLog = 0x6f02f748

type TL_help_saveAppLog struct {
	Events []TL // events:Vector<InputAppEvent>
}

// Encoding TL_help_saveAppLog
func (e TL_help_saveAppLog) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_saveAppLog)
	x.Vector(e.Events)
	return x.buf
}

// help.getInviteText#4d392343 = help.InviteText;

const crc_help_getInviteText = 0x4d392343

type TL_help_getInviteText struct {
}

// Encoding TL_help_getInviteText
func (e TL_help_getInviteText) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getInviteText)
	return x.buf
}

// help.getSupport#9cdf08cd = help.Support;

const crc_help_getSupport = 0x9cdf08cd

type TL_help_getSupport struct {
}

// Encoding TL_help_getSupport
func (e TL_help_getSupport) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getSupport)
	return x.buf
}

// help.getAppChangelog#9010ef6f prev_app_version:string = Updates;

const crc_help_getAppChangelog = 0x9010ef6f

type TL_help_getAppChangelog struct {
	Prev_app_version string // prev_app_version:string
}

// Encoding TL_help_getAppChangelog
func (e TL_help_getAppChangelog) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getAppChangelog)
	x.String(e.Prev_app_version)
	return x.buf
}

// help.getTermsOfService#350170f3 = help.TermsOfService;

const crc_help_getTermsOfService = 0x350170f3

type TL_help_getTermsOfService struct {
}

// Encoding TL_help_getTermsOfService
func (e TL_help_getTermsOfService) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getTermsOfService)
	return x.buf
}

// help.setBotUpdatesStatus#ec22cfcd pending_updates_count:int message:string = Bool;

const crc_help_setBotUpdatesStatus = 0xec22cfcd

type TL_help_setBotUpdatesStatus struct {
	Pending_updates_count int32  // pending_updates_count:int
	Message               string // message:string
}

// Encoding TL_help_setBotUpdatesStatus
func (e TL_help_setBotUpdatesStatus) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_setBotUpdatesStatus)
	x.Int(e.Pending_updates_count)
	x.String(e.Message)
	return x.buf
}

// channels.readHistory#cc104937 channel:InputChannel max_id:int = Bool;

const crc_channels_readHistory = 0xcc104937

type TL_channels_readHistory struct {
	Channel TL    // channel:InputChannel
	Max_id  int32 // max_id:int
}

// Encoding TL_channels_readHistory
func (e TL_channels_readHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_readHistory)
	x.Bytes(e.Channel.encode())
	x.Int(e.Max_id)
	return x.buf
}

// channels.deleteMessages#84c1fd4e channel:InputChannel Id:Vector<int> = messages.AffectedMessages;

const crc_channels_deleteMessages = 0x84c1fd4e

type TL_channels_deleteMessages struct {
	Channel TL      // channel:InputChannel
	Id      []int32 // Id:Vector<int>
}

// Encoding TL_channels_deleteMessages
func (e TL_channels_deleteMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_deleteMessages)
	x.Bytes(e.Channel.encode())
	x.VectorInt(e.Id)
	return x.buf
}

// channels.deleteUserHistory#d10dd71b channel:InputChannel user_id:InputUser = messages.AffectedHistory;

const crc_channels_deleteUserHistory = 0xd10dd71b

type TL_channels_deleteUserHistory struct {
	Channel TL // channel:InputChannel
	User_id TL // user_id:InputUser
}

// Encoding TL_channels_deleteUserHistory
func (e TL_channels_deleteUserHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_deleteUserHistory)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.User_id.encode())
	return x.buf
}

// channels.reportSpam#fe087810 channel:InputChannel user_id:InputUser Id:Vector<int> = Bool;

const crc_channels_reportSpam = 0xfe087810

type TL_channels_reportSpam struct {
	Channel TL      // channel:InputChannel
	User_id TL      // user_id:InputUser
	Id      []int32 // Id:Vector<int>
}

// Encoding TL_channels_reportSpam
func (e TL_channels_reportSpam) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_reportSpam)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.User_id.encode())
	x.VectorInt(e.Id)
	return x.buf
}

// channels.getMessages#93d7b347 channel:InputChannel Id:Vector<int> = messages.Messages;

const crc_channels_getMessages = 0x93d7b347

type TL_channels_getMessages struct {
	Channel TL      // channel:InputChannel
	Id      []int32 // Id:Vector<int>
}

// Encoding TL_channels_getMessages
func (e TL_channels_getMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_getMessages)
	x.Bytes(e.Channel.encode())
	x.VectorInt(e.Id)
	return x.buf
}

// channels.getParticipants#24d98f92 channel:InputChannel filter:ChannelParticipantsFilter offset:int limit:int = channels.ChannelParticipants;

const crc_channels_getParticipants = 0x24d98f92

type TL_channels_getParticipants struct {
	Channel TL    // channel:InputChannel
	Filter  TL    // filter:ChannelParticipantsFilter
	Offset  int32 // offset:int
	Limit   int32 // limit:int
}

// Encoding TL_channels_getParticipants
func (e TL_channels_getParticipants) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_getParticipants)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.Filter.encode())
	x.Int(e.Offset)
	x.Int(e.Limit)
	return x.buf
}

// channels.getParticipant#546dd7a6 channel:InputChannel user_id:InputUser = channels.ChannelParticipant;

const crc_channels_getParticipant = 0x546dd7a6

type TL_channels_getParticipant struct {
	Channel TL // channel:InputChannel
	User_id TL // user_id:InputUser
}

// Encoding TL_channels_getParticipant
func (e TL_channels_getParticipant) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_getParticipant)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.User_id.encode())
	return x.buf
}

// channels.getChannels#a7f6bbb Id:Vector<InputChannel> = messages.Chats;

const crc_channels_getChannels = 0xa7f6bbb

type TL_channels_getChannels struct {
	Id []TL // Id:Vector<InputChannel>
}

// Encoding TL_channels_getChannels
func (e TL_channels_getChannels) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_getChannels)
	x.Vector(e.Id)
	return x.buf
}

// channels.getFullChannel#8736a09 channel:InputChannel = messages.ChatFull;

const crc_channels_getFullChannel = 0x8736a09

type TL_channels_getFullChannel struct {
	Channel TL // channel:InputChannel
}

// Encoding TL_channels_getFullChannel
func (e TL_channels_getFullChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_getFullChannel)
	x.Bytes(e.Channel.encode())
	return x.buf
}

// channels.createChannel#f4893d7f flags:# broadcast:flags.0?true megagroup:flags.1?true title:string about:string = Updates;

const crc_channels_createChannel = 0xf4893d7f

type TL_channels_createChannel struct {
	Flags     int32
	Broadcast bool   // broadcast:flags.0?true
	Megagroup bool   // megagroup:flags.1?true
	Title     string // title:string
	About     string // about:string
}

// Encoding TL_channels_createChannel
func (e TL_channels_createChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_createChannel)
	var flags int32
	if e.Broadcast {
		flags |= (1 << 0)
	}
	if e.Megagroup {
		flags |= (1 << 1)
	}
	x.Int(flags)
	x.String(e.Title)
	x.String(e.About)
	return x.buf
}

// channels.editAbout#13e27f1e channel:InputChannel about:string = Bool;

const crc_channels_editAbout = 0x13e27f1e

type TL_channels_editAbout struct {
	Channel TL     // channel:InputChannel
	About   string // about:string
}

// Encoding TL_channels_editAbout
func (e TL_channels_editAbout) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_editAbout)
	x.Bytes(e.Channel.encode())
	x.String(e.About)
	return x.buf
}

// channels.editAdmin#eb7611d0 channel:InputChannel user_id:InputUser role:ChannelParticipantRole = Updates;

const crc_channels_editAdmin = 0xeb7611d0

type TL_channels_editAdmin struct {
	Channel TL // channel:InputChannel
	User_id TL // user_id:InputUser
	Role    TL // role:ChannelParticipantRole
}

// Encoding TL_channels_editAdmin
func (e TL_channels_editAdmin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_editAdmin)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.User_id.encode())
	x.Bytes(e.Role.encode())
	return x.buf
}

// channels.editTitle#566decd0 channel:InputChannel title:string = Updates;

const crc_channels_editTitle = 0x566decd0

type TL_channels_editTitle struct {
	Channel TL     // channel:InputChannel
	Title   string // title:string
}

// Encoding TL_channels_editTitle
func (e TL_channels_editTitle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_editTitle)
	x.Bytes(e.Channel.encode())
	x.String(e.Title)
	return x.buf
}

// channels.editPhoto#f12e57c9 channel:InputChannel photo:InputChatPhoto = Updates;

const crc_channels_editPhoto = 0xf12e57c9

type TL_channels_editPhoto struct {
	Channel TL // channel:InputChannel
	Photo   TL // photo:InputChatPhoto
}

// Encoding TL_channels_editPhoto
func (e TL_channels_editPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_editPhoto)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.Photo.encode())
	return x.buf
}

// channels.checkUsername#10e6bd2c channel:InputChannel username:string = Bool;

const crc_channels_checkUsername = 0x10e6bd2c

type TL_channels_checkUsername struct {
	Channel  TL     // channel:InputChannel
	Username string // username:string
}

// Encoding TL_channels_checkUsername
func (e TL_channels_checkUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_checkUsername)
	x.Bytes(e.Channel.encode())
	x.String(e.Username)
	return x.buf
}

// channels.updateUsername#3514b3de channel:InputChannel username:string = Bool;

const crc_channels_updateUsername = 0x3514b3de

type TL_channels_updateUsername struct {
	Channel  TL     // channel:InputChannel
	Username string // username:string
}

// Encoding TL_channels_updateUsername
func (e TL_channels_updateUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_updateUsername)
	x.Bytes(e.Channel.encode())
	x.String(e.Username)
	return x.buf
}

// channels.joinChannel#24b524c5 channel:InputChannel = Updates;

const crc_channels_joinChannel = 0x24b524c5

type TL_channels_joinChannel struct {
	Channel TL // channel:InputChannel
}

// Encoding TL_channels_joinChannel
func (e TL_channels_joinChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_joinChannel)
	x.Bytes(e.Channel.encode())
	return x.buf
}

// channels.leaveChannel#f836aa95 channel:InputChannel = Updates;

const crc_channels_leaveChannel = 0xf836aa95

type TL_channels_leaveChannel struct {
	Channel TL // channel:InputChannel
}

// Encoding TL_channels_leaveChannel
func (e TL_channels_leaveChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_leaveChannel)
	x.Bytes(e.Channel.encode())
	return x.buf
}

// channels.inviteToChannel#199f3a6c channel:InputChannel users:Vector<InputUser> = Updates;

const crc_channels_inviteToChannel = 0x199f3a6c

type TL_channels_inviteToChannel struct {
	Channel TL   // channel:InputChannel
	Users   []TL // users:Vector<InputUser>
}

// Encoding TL_channels_inviteToChannel
func (e TL_channels_inviteToChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_inviteToChannel)
	x.Bytes(e.Channel.encode())
	x.Vector(e.Users)
	return x.buf
}

// channels.kickFromChannel#a672de14 channel:InputChannel user_id:InputUser kicked:Bool = Updates;

const crc_channels_kickFromChannel = 0xa672de14

type TL_channels_kickFromChannel struct {
	Channel TL // channel:InputChannel
	User_id TL // user_id:InputUser
	Kicked  TL // kicked:Bool
}

// Encoding TL_channels_kickFromChannel
func (e TL_channels_kickFromChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_kickFromChannel)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.User_id.encode())
	x.Bytes(e.Kicked.encode())
	return x.buf
}

// channels.exportInvite#c7560885 channel:InputChannel = ExportedChatInvite;

const crc_channels_exportInvite = 0xc7560885

type TL_channels_exportInvite struct {
	Channel TL // channel:InputChannel
}

// Encoding TL_channels_exportInvite
func (e TL_channels_exportInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_exportInvite)
	x.Bytes(e.Channel.encode())
	return x.buf
}

// channels.deleteChannel#c0111fe3 channel:InputChannel = Updates;

const crc_channels_deleteChannel = 0xc0111fe3

type TL_channels_deleteChannel struct {
	Channel TL // channel:InputChannel
}

// Encoding TL_channels_deleteChannel
func (e TL_channels_deleteChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_deleteChannel)
	x.Bytes(e.Channel.encode())
	return x.buf
}

// channels.toggleInvites#49609307 channel:InputChannel enabled:Bool = Updates;

const crc_channels_toggleInvites = 0x49609307

type TL_channels_toggleInvites struct {
	Channel TL // channel:InputChannel
	Enabled TL // enabled:Bool
}

// Encoding TL_channels_toggleInvites
func (e TL_channels_toggleInvites) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_toggleInvites)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.Enabled.encode())
	return x.buf
}

// channels.exportMessageLink#c846d22d channel:InputChannel Id:int = ExportedMessageLink;

const crc_channels_exportMessageLink = 0xc846d22d

type TL_channels_exportMessageLink struct {
	Channel TL    // channel:InputChannel
	Id      int32 // Id:int
}

// Encoding TL_channels_exportMessageLink
func (e TL_channels_exportMessageLink) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_exportMessageLink)
	x.Bytes(e.Channel.encode())
	x.Int(e.Id)
	return x.buf
}

// channels.toggleSignatures#1f69b606 channel:InputChannel enabled:Bool = Updates;

const crc_channels_toggleSignatures = 0x1f69b606

type TL_channels_toggleSignatures struct {
	Channel TL // channel:InputChannel
	Enabled TL // enabled:Bool
}

// Encoding TL_channels_toggleSignatures
func (e TL_channels_toggleSignatures) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_toggleSignatures)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.Enabled.encode())
	return x.buf
}

// channels.updatePinnedMessage#a72ded52 flags:# silent:flags.0?true channel:InputChannel Id:int = Updates;

const crc_channels_updatePinnedMessage = 0xa72ded52

type TL_channels_updatePinnedMessage struct {
	Flags   int32
	Silent  bool  // silent:flags.0?true
	Channel TL    // channel:InputChannel
	Id      int32 // Id:int
}

// Encoding TL_channels_updatePinnedMessage
func (e TL_channels_updatePinnedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_updatePinnedMessage)
	var flags int32
	if e.Silent {
		flags |= (1 << 0)
	}
	x.Int(flags)
	x.Bytes(e.Channel.encode())
	x.Int(e.Id)
	return x.buf
}

// channels.getAdminedPublicChannels#8d8d82d7 = messages.Chats;

const crc_channels_getAdminedPublicChannels = 0x8d8d82d7

type TL_channels_getAdminedPublicChannels struct {
}

// Encoding TL_channels_getAdminedPublicChannels
func (e TL_channels_getAdminedPublicChannels) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_getAdminedPublicChannels)
	return x.buf
}

// bots.sendCustomRequest#aa2769ed custom_method:string params:DataJSON = DataJSON;

const crc_bots_sendCustomRequest = 0xaa2769ed

type TL_bots_sendCustomRequest struct {
	Custom_method string // custom_method:string
	Params        TL     // params:DataJSON
}

// Encoding TL_bots_sendCustomRequest
func (e TL_bots_sendCustomRequest) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_bots_sendCustomRequest)
	x.String(e.Custom_method)
	x.Bytes(e.Params.encode())
	return x.buf
}

// bots.answerWebhookJSONQuery#e6213f4d query_id:long data:DataJSON = Bool;

const crc_bots_answerWebhookJSONQuery = 0xe6213f4d

type TL_bots_answerWebhookJSONQuery struct {
	Query_id int64 // query_id:long
	Data     TL    // data:DataJSON
}

// Encoding TL_bots_answerWebhookJSONQuery
func (e TL_bots_answerWebhookJSONQuery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_bots_answerWebhookJSONQuery)
	x.Long(e.Query_id)
	x.Bytes(e.Data.encode())
	return x.buf
}

// payments.getPaymentForm#99f09745 msg_id:int = payments.PaymentForm;

const crc_payments_getPaymentForm = 0x99f09745

type TL_payments_getPaymentForm struct {
	Msg_id int32 // msg_id:int
}

// Encoding TL_payments_getPaymentForm
func (e TL_payments_getPaymentForm) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_payments_getPaymentForm)
	x.Int(e.Msg_id)
	return x.buf
}

// payments.getPaymentReceipt#a092a980 msg_id:int = payments.PaymentReceipt;

const crc_payments_getPaymentReceipt = 0xa092a980

type TL_payments_getPaymentReceipt struct {
	Msg_id int32 // msg_id:int
}

// Encoding TL_payments_getPaymentReceipt
func (e TL_payments_getPaymentReceipt) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_payments_getPaymentReceipt)
	x.Int(e.Msg_id)
	return x.buf
}

// payments.validateRequestedInfo#770a8e74 flags:# save:flags.0?true msg_id:int info:PaymentRequestedInfo = payments.ValidatedRequestedInfo;

const crc_payments_validateRequestedInfo = 0x770a8e74

type TL_payments_validateRequestedInfo struct {
	Flags  int32
	Save   bool  // save:flags.0?true
	Msg_id int32 // msg_id:int
	Info   TL    // info:PaymentRequestedInfo
}

// Encoding TL_payments_validateRequestedInfo
func (e TL_payments_validateRequestedInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_payments_validateRequestedInfo)
	var flags int32
	if e.Save {
		flags |= (1 << 0)
	}
	x.Int(flags)
	x.Int(e.Msg_id)
	x.Bytes(e.Info.encode())
	return x.buf
}

// payments.sendPaymentForm#2b8879b3 flags:# msg_id:int requested_info_id:flags.0?string shipping_option_id:flags.1?string credentials:InputPaymentCredentials = payments.PaymentResult;

const crc_payments_sendPaymentForm = 0x2b8879b3

type TL_payments_sendPaymentForm struct {
	Flags              int32
	Msg_id             int32  // msg_id:int
	Requested_info_id  string // requested_info_id:flags.0?string
	Shipping_option_id string // shipping_option_id:flags.1?string
	Credentials        TL     // credentials:InputPaymentCredentials
}

// Encoding TL_payments_sendPaymentForm
func (e TL_payments_sendPaymentForm) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_payments_sendPaymentForm)
	var flags int32
	if e.Requested_info_id != "" {
		flags |= (1 << 0)
	}
	if e.Shipping_option_id != "" {
		flags |= (1 << 1)
	}
	x.Int(flags)
	x.Int(e.Msg_id)
	if flags&(1<<0) != 0 {
		x.String(e.Requested_info_id)
	}
	if flags&(1<<1) != 0 {
		x.String(e.Shipping_option_id)
	}
	x.Bytes(e.Credentials.encode())
	return x.buf
}

// payments.getSavedInfo#227d824b = payments.SavedInfo;

const crc_payments_getSavedInfo = 0x227d824b

type TL_payments_getSavedInfo struct {
}

// Encoding TL_payments_getSavedInfo
func (e TL_payments_getSavedInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_payments_getSavedInfo)
	return x.buf
}

// payments.clearSavedInfo#d83d70c1 flags:# credentials:flags.0?true info:flags.1?true = Bool;

const crc_payments_clearSavedInfo = 0xd83d70c1

type TL_payments_clearSavedInfo struct {
	Flags       int32
	Credentials bool // credentials:flags.0?true
	Info        bool // info:flags.1?true
}

// Encoding TL_payments_clearSavedInfo
func (e TL_payments_clearSavedInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_payments_clearSavedInfo)
	var flags int32
	if e.Credentials {
		flags |= (1 << 0)
	}
	if e.Info {
		flags |= (1 << 1)
	}
	x.Int(flags)
	return x.buf
}

// phone.getCallConfig#55451fa9 = DataJSON;

const crc_phone_getCallConfig = 0x55451fa9

type TL_phone_getCallConfig struct {
}

// Encoding TL_phone_getCallConfig
func (e TL_phone_getCallConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phone_getCallConfig)
	return x.buf
}

// phone.requestCall#5b95b3d4 user_id:InputUser random_id:int g_a_hash:bytes protocol:PhoneCallProtocol = phone.PhoneCall;

const crc_phone_requestCall = 0x5b95b3d4

type TL_phone_requestCall struct {
	User_id   TL     // user_id:InputUser
	Random_id int32  // random_id:int
	G_a_hash  []byte // g_a_hash:bytes
	Protocol  TL     // protocol:PhoneCallProtocol
}

// Encoding TL_phone_requestCall
func (e TL_phone_requestCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phone_requestCall)
	x.Bytes(e.User_id.encode())
	x.Int(e.Random_id)
	x.Bytes(e.G_a_hash)
	x.Bytes(e.Protocol.encode())
	return x.buf
}

// phone.acceptCall#3bd2b4a0 peer:InputPhoneCall g_b:bytes protocol:PhoneCallProtocol = phone.PhoneCall;

const crc_phone_acceptCall = 0x3bd2b4a0

type TL_phone_acceptCall struct {
	Peer     TL     // peer:InputPhoneCall
	G_b      []byte // g_b:bytes
	Protocol TL     // protocol:PhoneCallProtocol
}

// Encoding TL_phone_acceptCall
func (e TL_phone_acceptCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phone_acceptCall)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.G_b)
	x.Bytes(e.Protocol.encode())
	return x.buf
}

// phone.confirmCall#2efe1722 peer:InputPhoneCall g_a:bytes key_fingerprint:long protocol:PhoneCallProtocol = phone.PhoneCall;

const crc_phone_confirmCall = 0x2efe1722

type TL_phone_confirmCall struct {
	Peer            TL     // peer:InputPhoneCall
	G_a             []byte // g_a:bytes
	Key_fingerprint int64  // key_fingerprint:long
	Protocol        TL     // protocol:PhoneCallProtocol
}

// Encoding TL_phone_confirmCall
func (e TL_phone_confirmCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phone_confirmCall)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.G_a)
	x.Long(e.Key_fingerprint)
	x.Bytes(e.Protocol.encode())
	return x.buf
}

// phone.receivedCall#17d54f61 peer:InputPhoneCall = Bool;

const crc_phone_receivedCall = 0x17d54f61

type TL_phone_receivedCall struct {
	Peer TL // peer:InputPhoneCall
}

// Encoding TL_phone_receivedCall
func (e TL_phone_receivedCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phone_receivedCall)
	x.Bytes(e.Peer.encode())
	return x.buf
}

// phone.discardCall#78d413a6 peer:InputPhoneCall duration:int reason:PhoneCallDiscardReason connection_id:long = Updates;

const crc_phone_discardCall = 0x78d413a6

type TL_phone_discardCall struct {
	Peer          TL    // peer:InputPhoneCall
	Duration      int32 // duration:int
	Reason        TL    // reason:PhoneCallDiscardReason
	Connection_id int64 // connection_id:long
}

// Encoding TL_phone_discardCall
func (e TL_phone_discardCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phone_discardCall)
	x.Bytes(e.Peer.encode())
	x.Int(e.Duration)
	x.Bytes(e.Reason.encode())
	x.Long(e.Connection_id)
	return x.buf
}

// phone.setCallRating#1c536a34 peer:InputPhoneCall rating:int comment:string = Updates;

const crc_phone_setCallRating = 0x1c536a34

type TL_phone_setCallRating struct {
	Peer    TL     // peer:InputPhoneCall
	Rating  int32  // rating:int
	Comment string // comment:string
}

// Encoding TL_phone_setCallRating
func (e TL_phone_setCallRating) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phone_setCallRating)
	x.Bytes(e.Peer.encode())
	x.Int(e.Rating)
	x.String(e.Comment)
	return x.buf
}

// phone.saveCallDebug#277add7e peer:InputPhoneCall debug:DataJSON = Bool;

const crc_phone_saveCallDebug = 0x277add7e

type TL_phone_saveCallDebug struct {
	Peer  TL // peer:InputPhoneCall
	Debug TL // debug:DataJSON
}

// Encoding TL_phone_saveCallDebug
func (e TL_phone_saveCallDebug) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phone_saveCallDebug)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Debug.encode())
	return x.buf
}
func (m *DecodeBuf) ObjectGenerated(constructor uint32) (r TL) {
	switch constructor {
	case crc_boolFalse:
		r = TL_boolFalse{}
	case crc_boolTrue:
		r = TL_boolTrue{}
	case crc_true:
		r = TL_true{}
	case crc_error:
		r = TL_error{
			Code: m.Int(),
			Text: m.String(),
		}
	case crc_null:
		r = TL_null{}
	case crc_inputPeerEmpty:
		r = TL_inputPeerEmpty{}
	case crc_inputPeerSelf:
		r = TL_inputPeerSelf{}
	case crc_inputPeerChat:
		r = TL_inputPeerChat{
			Chat_id: m.Int(),
		}
	case crc_inputPeerUser:
		r = TL_inputPeerUser{
			User_id:     m.Int(),
			Access_hash: m.Long(),
		}
	case crc_inputPeerChannel:
		r = TL_inputPeerChannel{
			Channel_id:  m.Int(),
			Access_hash: m.Long(),
		}
	case crc_inputUserEmpty:
		r = TL_inputUserEmpty{}
	case crc_inputUserSelf:
		r = TL_inputUserSelf{}
	case crc_inputUser:
		r = TL_inputUser{
			User_id:     m.Int(),
			Access_hash: m.Long(),
		}
	case crc_inputPhoneContact:
		r = TL_inputPhoneContact{
			Client_id:  m.Long(),
			Phone:      m.String(),
			First_name: m.String(),
			Last_name:  m.String(),
		}
	case crc_inputFile:
		r = TL_inputFile{
			Id:           m.Long(),
			Parts:        m.Int(),
			Name:         m.String(),
			Md5_checksum: m.String(),
		}
	case crc_inputFileBig:
		r = TL_inputFileBig{
			Id:    m.Long(),
			Parts: m.Int(),
			Name:  m.String(),
		}
	case crc_inputMediaEmpty:
		r = TL_inputMediaEmpty{}
	case crc_inputMediaUploadedPhoto:
		flags := m.Int()
		file := m.Object()
		caption := m.String()
		var stickers []TL
		if flags&(1<<0) != 0 {
			stickers = m.Vector()
		}
		r = TL_inputMediaUploadedPhoto{
			Flags:    flags,
			File:     file,
			Caption:  caption,
			Stickers: stickers,
		}
	case crc_inputMediaPhoto:
		r = TL_inputMediaPhoto{
			Id:      m.Object(),
			Caption: m.String(),
		}
	case crc_inputMediaGeoPoint:
		r = TL_inputMediaGeoPoint{
			Geo_point: m.Object(),
		}
	case crc_inputMediaContact:
		r = TL_inputMediaContact{
			Phone_number: m.String(),
			First_name:   m.String(),
			Last_name:    m.String(),
		}
	case crc_inputMediaUploadedDocument:
		flags := m.Int()
		file := m.Object()
		mime_type := m.String()
		attributes := m.Vector()
		caption := m.String()
		var stickers []TL
		if flags&(1<<0) != 0 {
			stickers = m.Vector()
		}
		r = TL_inputMediaUploadedDocument{
			Flags:      flags,
			File:       file,
			Mime_type:  mime_type,
			Attributes: attributes,
			Caption:    caption,
			Stickers:   stickers,
		}
	case crc_inputMediaUploadedThumbDocument:
		flags := m.Int()
		file := m.Object()
		thumb := m.Object()
		mime_type := m.String()
		attributes := m.Vector()
		caption := m.String()
		var stickers []TL
		if flags&(1<<0) != 0 {
			stickers = m.Vector()
		}
		r = TL_inputMediaUploadedThumbDocument{
			Flags:      flags,
			File:       file,
			Thumb:      thumb,
			Mime_type:  mime_type,
			Attributes: attributes,
			Caption:    caption,
			Stickers:   stickers,
		}
	case crc_inputMediaDocument:
		r = TL_inputMediaDocument{
			Id:      m.Object(),
			Caption: m.String(),
		}
	case crc_inputMediaVenue:
		r = TL_inputMediaVenue{
			Geo_point: m.Object(),
			Title:     m.String(),
			Address:   m.String(),
			Provider:  m.String(),
			Venue_id:  m.String(),
		}
	case crc_inputMediaGifExternal:
		r = TL_inputMediaGifExternal{
			Url: m.String(),
			Q:   m.String(),
		}
	case crc_inputMediaPhotoExternal:
		r = TL_inputMediaPhotoExternal{
			Url:     m.String(),
			Caption: m.String(),
		}
	case crc_inputMediaDocumentExternal:
		r = TL_inputMediaDocumentExternal{
			Url:     m.String(),
			Caption: m.String(),
		}
	case crc_inputMediaGame:
		r = TL_inputMediaGame{
			Id: m.Object(),
		}
	case crc_inputMediaInvoice:
		flags := m.Int()
		title := m.String()
		description := m.String()
		var photo TL
		if flags&(1<<0) != 0 {
			photo = m.Object()
		}
		invoice := m.Object()
		payload := m.StringBytes()
		provider := m.String()
		start_param := m.String()
		r = TL_inputMediaInvoice{
			Flags:       flags,
			Title:       title,
			Description: description,
			Photo:       photo,
			Invoice:     invoice,
			Payload:     payload,
			Provider:    provider,
			Start_param: start_param,
		}
	case crc_inputChatPhotoEmpty:
		r = TL_inputChatPhotoEmpty{}
	case crc_inputChatUploadedPhoto:
		r = TL_inputChatUploadedPhoto{
			File: m.Object(),
		}
	case crc_inputChatPhoto:
		r = TL_inputChatPhoto{
			Id: m.Object(),
		}
	case crc_inputGeoPointEmpty:
		r = TL_inputGeoPointEmpty{}
	case crc_inputGeoPoint:
		r = TL_inputGeoPoint{
			Lat:  m.Double(),
			Long: m.Double(),
		}
	case crc_inputPhotoEmpty:
		r = TL_inputPhotoEmpty{}
	case crc_inputPhoto:
		r = TL_inputPhoto{
			Id:          m.Long(),
			Access_hash: m.Long(),
		}
	case crc_inputFileLocation:
		r = TL_inputFileLocation{
			Volume_id: m.Long(),
			Local_id:  m.Int(),
			Secret:    m.Long(),
		}
	case crc_inputEncryptedFileLocation:
		r = TL_inputEncryptedFileLocation{
			Id:          m.Long(),
			Access_hash: m.Long(),
		}
	case crc_inputDocumentFileLocation:
		r = TL_inputDocumentFileLocation{
			Id:          m.Long(),
			Access_hash: m.Long(),
			Version:     m.Int(),
		}
	case crc_inputAppEvent:
		r = TL_inputAppEvent{
			Time:      m.Double(),
			Code_type: m.String(),
			Peer:      m.Long(),
			Data:      m.String(),
		}
	case crc_peerUser:
		r = TL_peerUser{
			User_id: m.Int(),
		}
	case crc_peerChat:
		r = TL_peerChat{
			Chat_id: m.Int(),
		}
	case crc_peerChannel:
		r = TL_peerChannel{
			Channel_id: m.Int(),
		}
	case crc_storage_fileUnknown:
		r = TL_storage_fileUnknown{}
	case crc_storage_filePartial:
		r = TL_storage_filePartial{}
	case crc_storage_fileJpeg:
		r = TL_storage_fileJpeg{}
	case crc_storage_fileGif:
		r = TL_storage_fileGif{}
	case crc_storage_filePng:
		r = TL_storage_filePng{}
	case crc_storage_filePdf:
		r = TL_storage_filePdf{}
	case crc_storage_fileMp3:
		r = TL_storage_fileMp3{}
	case crc_storage_fileMov:
		r = TL_storage_fileMov{}
	case crc_storage_fileMp4:
		r = TL_storage_fileMp4{}
	case crc_storage_fileWebp:
		r = TL_storage_fileWebp{}
	case crc_fileLocationUnavailable:
		r = TL_fileLocationUnavailable{
			Volume_id: m.Long(),
			Local_id:  m.Int(),
			Secret:    m.Long(),
		}
	case crc_fileLocation:
		r = TL_fileLocation{
			Dc_id:     m.Int(),
			Volume_id: m.Long(),
			Local_id:  m.Int(),
			Secret:    m.Long(),
		}
	case crc_userEmpty:
		r = TL_userEmpty{
			Id: m.Int(),
		}
	case crc_user:
		flags := m.Int()
		self := flags&(1<<10) != 0
		contact := flags&(1<<11) != 0
		mutual_contact := flags&(1<<12) != 0
		deleted := flags&(1<<13) != 0
		bot := flags&(1<<14) != 0
		bot_chat_history := flags&(1<<15) != 0
		bot_nochats := flags&(1<<16) != 0
		verified := flags&(1<<17) != 0
		restricted := flags&(1<<18) != 0
		min := flags&(1<<20) != 0
		bot_inline_geo := flags&(1<<21) != 0
		id := m.Int()
		var access_hash int64
		if flags&(1<<0) != 0 {
			access_hash = m.Long()
		}
		var first_name string
		if flags&(1<<1) != 0 {
			first_name = m.String()
		}
		var last_name string
		if flags&(1<<2) != 0 {
			last_name = m.String()
		}
		var username string
		if flags&(1<<3) != 0 {
			username = m.String()
		}
		var phone string
		if flags&(1<<4) != 0 {
			phone = m.String()
		}
		var photo TL
		if flags&(1<<5) != 0 {
			photo = m.Object()
		}
		var status TL
		if flags&(1<<6) != 0 {
			status = m.Object()
		}
		var bot_info_version int32
		if flags&(1<<14) != 0 {
			bot_info_version = m.Int()
		}
		var restriction_reason string
		if flags&(1<<18) != 0 {
			restriction_reason = m.String()
		}
		var bot_inline_placeholder string
		if flags&(1<<19) != 0 {
			bot_inline_placeholder = m.String()
		}
		r = TL_user{
			Flags:                  flags,
			Self:                   self,
			Contact:                contact,
			Mutual_contact:         mutual_contact,
			Deleted:                deleted,
			Bot:                    bot,
			Bot_chat_history:       bot_chat_history,
			Bot_nochats:            bot_nochats,
			Verified:               verified,
			Restricted:             restricted,
			Min:                    min,
			Bot_inline_geo:         bot_inline_geo,
			Id:                     id,
			Access_hash:            access_hash,
			First_name:             first_name,
			Last_name:              last_name,
			Username:               username,
			Phone:                  phone,
			Photo:                  photo,
			Status:                 status,
			Bot_info_version:       bot_info_version,
			Restriction_reason:     restriction_reason,
			Bot_inline_placeholder: bot_inline_placeholder,
		}
	case crc_userProfilePhotoEmpty:
		r = TL_userProfilePhotoEmpty{}
	case crc_userProfilePhoto:
		r = TL_userProfilePhoto{
			Photo_id:    m.Long(),
			Photo_small: m.Object(),
			Photo_big:   m.Object(),
		}
	case crc_userStatusEmpty:
		r = TL_userStatusEmpty{}
	case crc_userStatusOnline:
		r = TL_userStatusOnline{
			Expires: m.Int(),
		}
	case crc_userStatusOffline:
		r = TL_userStatusOffline{
			Was_online: m.Int(),
		}
	case crc_userStatusRecently:
		r = TL_userStatusRecently{}
	case crc_userStatusLastWeek:
		r = TL_userStatusLastWeek{}
	case crc_userStatusLastMonth:
		r = TL_userStatusLastMonth{}
	case crc_chatEmpty:
		r = TL_chatEmpty{
			Id: m.Int(),
		}
	case crc_chat:
		flags := m.Int()
		creator := flags&(1<<0) != 0
		kicked := flags&(1<<1) != 0
		left := flags&(1<<2) != 0
		admins_enabled := flags&(1<<3) != 0
		admin := flags&(1<<4) != 0
		deactivated := flags&(1<<5) != 0
		id := m.Int()
		title := m.String()
		photo := m.Object()
		participants_count := m.Int()
		date := m.Int()
		version := m.Int()
		var migrated_to TL
		if flags&(1<<6) != 0 {
			migrated_to = m.Object()
		}
		r = TL_chat{
			Flags:              flags,
			Creator:            creator,
			Kicked:             kicked,
			Left:               left,
			Admins_enabled:     admins_enabled,
			Admin:              admin,
			Deactivated:        deactivated,
			Id:                 id,
			Title:              title,
			Photo:              photo,
			Participants_count: participants_count,
			Date:               date,
			Version:            version,
			Migrated_to:        migrated_to,
		}
	case crc_chatForbidden:
		r = TL_chatForbidden{
			Id:    m.Int(),
			Title: m.String(),
		}
	case crc_channel:
		flags := m.Int()
		creator := flags&(1<<0) != 0
		kicked := flags&(1<<1) != 0
		left := flags&(1<<2) != 0
		editor := flags&(1<<3) != 0
		moderator := flags&(1<<4) != 0
		broadcast := flags&(1<<5) != 0
		verified := flags&(1<<7) != 0
		megagroup := flags&(1<<8) != 0
		restricted := flags&(1<<9) != 0
		democracy := flags&(1<<10) != 0
		signatures := flags&(1<<11) != 0
		min := flags&(1<<12) != 0
		id := m.Int()
		var access_hash int64
		if flags&(1<<13) != 0 {
			access_hash = m.Long()
		}
		title := m.String()
		var username string
		if flags&(1<<6) != 0 {
			username = m.String()
		}
		photo := m.Object()
		date := m.Int()
		version := m.Int()
		var restriction_reason string
		if flags&(1<<9) != 0 {
			restriction_reason = m.String()
		}
		r = TL_channel{
			Flags:              flags,
			Creator:            creator,
			Kicked:             kicked,
			Left:               left,
			Editor:             editor,
			Moderator:          moderator,
			Broadcast:          broadcast,
			Verified:           verified,
			Megagroup:          megagroup,
			Restricted:         restricted,
			Democracy:          democracy,
			Signatures:         signatures,
			Min:                min,
			Id:                 id,
			Access_hash:        access_hash,
			Title:              title,
			Username:           username,
			Photo:              photo,
			Date:               date,
			Version:            version,
			Restriction_reason: restriction_reason,
		}
	case crc_channelForbidden:
		flags := m.Int()
		broadcast := flags&(1<<5) != 0
		megagroup := flags&(1<<8) != 0
		id := m.Int()
		access_hash := m.Long()
		title := m.String()
		r = TL_channelForbidden{
			Flags:       flags,
			Broadcast:   broadcast,
			Megagroup:   megagroup,
			Id:          id,
			Access_hash: access_hash,
			Title:       title,
		}
	case crc_chatFull:
		r = TL_chatFull{
			Id:              m.Int(),
			Participants:    m.Object(),
			Chat_photo:      m.Object(),
			Notify_settings: m.Object(),
			Exported_invite: m.Object(),
			Bot_info:        m.Vector(),
		}
	case crc_channelFull:
		flags := m.Int()
		can_view_participants := flags&(1<<3) != 0
		can_set_username := flags&(1<<6) != 0
		id := m.Int()
		about := m.String()
		var participants_count int32
		if flags&(1<<0) != 0 {
			participants_count = m.Int()
		}
		var admins_count int32
		if flags&(1<<1) != 0 {
			admins_count = m.Int()
		}
		var kicked_count int32
		if flags&(1<<2) != 0 {
			kicked_count = m.Int()
		}
		read_inbox_max_id := m.Int()
		read_outbox_max_id := m.Int()
		unread_count := m.Int()
		chat_photo := m.Object()
		notify_settings := m.Object()
		exported_invite := m.Object()
		bot_info := m.Vector()
		var migrated_from_chat_id int32
		if flags&(1<<4) != 0 {
			migrated_from_chat_id = m.Int()
		}
		var migrated_from_max_id int32
		if flags&(1<<4) != 0 {
			migrated_from_max_id = m.Int()
		}
		var pinned_msg_id int32
		if flags&(1<<5) != 0 {
			pinned_msg_id = m.Int()
		}
		r = TL_channelFull{
			Flags: flags,
			Can_view_participants: can_view_participants,
			Can_set_username:      can_set_username,
			Id:                    id,
			About:                 about,
			Participants_count:    participants_count,
			Admins_count:          admins_count,
			Kicked_count:          kicked_count,
			Read_inbox_max_id:     read_inbox_max_id,
			Read_outbox_max_id:    read_outbox_max_id,
			Unread_count:          unread_count,
			Chat_photo:            chat_photo,
			Notify_settings:       notify_settings,
			Exported_invite:       exported_invite,
			Bot_info:              bot_info,
			Migrated_from_chat_id: migrated_from_chat_id,
			Migrated_from_max_id:  migrated_from_max_id,
			Pinned_msg_id:         pinned_msg_id,
		}
	case crc_chatParticipant:
		r = TL_chatParticipant{
			User_id:    m.Int(),
			Inviter_id: m.Int(),
			Date:       m.Int(),
		}
	case crc_chatParticipantCreator:
		r = TL_chatParticipantCreator{
			User_id: m.Int(),
		}
	case crc_chatParticipantAdmin:
		r = TL_chatParticipantAdmin{
			User_id:    m.Int(),
			Inviter_id: m.Int(),
			Date:       m.Int(),
		}
	case crc_chatParticipantsForbidden:
		flags := m.Int()
		chat_id := m.Int()
		var self_participant TL
		if flags&(1<<0) != 0 {
			self_participant = m.Object()
		}
		r = TL_chatParticipantsForbidden{
			Flags:            flags,
			Chat_id:          chat_id,
			Self_participant: self_participant,
		}
	case crc_chatParticipants:
		r = TL_chatParticipants{
			Chat_id:      m.Int(),
			Participants: m.Vector(),
			Version:      m.Int(),
		}
	case crc_chatPhotoEmpty:
		r = TL_chatPhotoEmpty{}
	case crc_chatPhoto:
		r = TL_chatPhoto{
			Photo_small: m.Object(),
			Photo_big:   m.Object(),
		}
	case crc_messageEmpty:
		r = TL_messageEmpty{
			Id: m.Int(),
		}
	case crc_message:
		flags := m.Int()
		out := flags&(1<<1) != 0
		mentioned := flags&(1<<4) != 0
		media_unread := flags&(1<<5) != 0
		silent := flags&(1<<13) != 0
		post := flags&(1<<14) != 0
		id := m.Int()
		var from_id int32
		if flags&(1<<8) != 0 {
			from_id = m.Int()
		}
		to_id := m.Object()
		var fwd_from TL
		if flags&(1<<2) != 0 {
			fwd_from = m.Object()
		}
		var via_bot_id int32
		if flags&(1<<11) != 0 {
			via_bot_id = m.Int()
		}
		var reply_to_msg_id int32
		if flags&(1<<3) != 0 {
			reply_to_msg_id = m.Int()
		}
		date := m.Int()
		message := m.String()
		var media TL
		if flags&(1<<9) != 0 {
			media = m.Object()
		}
		var reply_markup TL
		if flags&(1<<6) != 0 {
			reply_markup = m.Object()
		}
		var entities []TL
		if flags&(1<<7) != 0 {
			entities = m.Vector()
		}
		var views int32
		if flags&(1<<10) != 0 {
			views = m.Int()
		}
		var edit_date int32
		if flags&(1<<15) != 0 {
			edit_date = m.Int()
		}
		r = TL_message{
			Flags:           flags,
			Out:             out,
			Mentioned:       mentioned,
			Media_unread:    media_unread,
			Silent:          silent,
			Post:            post,
			Id:              id,
			From_id:         from_id,
			To_id:           to_id,
			Fwd_from:        fwd_from,
			Via_bot_id:      via_bot_id,
			Reply_to_msg_id: reply_to_msg_id,
			Date:            date,
			Message:         message,
			Media:           media,
			Reply_markup:    reply_markup,
			Entities:        entities,
			Views:           views,
			Edit_date:       edit_date,
		}
	case crc_messageService:
		flags := m.Int()
		out := flags&(1<<1) != 0
		mentioned := flags&(1<<4) != 0
		media_unread := flags&(1<<5) != 0
		silent := flags&(1<<13) != 0
		post := flags&(1<<14) != 0
		id := m.Int()
		var from_id int32
		if flags&(1<<8) != 0 {
			from_id = m.Int()
		}
		to_id := m.Object()
		var reply_to_msg_id int32
		if flags&(1<<3) != 0 {
			reply_to_msg_id = m.Int()
		}
		date := m.Int()
		action := m.Object()
		r = TL_messageService{
			Flags:           flags,
			Out:             out,
			Mentioned:       mentioned,
			Media_unread:    media_unread,
			Silent:          silent,
			Post:            post,
			Id:              id,
			From_id:         from_id,
			To_id:           to_id,
			Reply_to_msg_id: reply_to_msg_id,
			Date:            date,
			Action:          action,
		}
	case crc_messageMediaEmpty:
		r = TL_messageMediaEmpty{}
	case crc_messageMediaPhoto:
		r = TL_messageMediaPhoto{
			Photo:   m.Object(),
			Caption: m.String(),
		}
	case crc_messageMediaGeo:
		r = TL_messageMediaGeo{
			Geo: m.Object(),
		}
	case crc_messageMediaContact:
		r = TL_messageMediaContact{
			Phone_number: m.String(),
			First_name:   m.String(),
			Last_name:    m.String(),
			User_id:      m.Int(),
		}
	case crc_messageMediaUnsupported:
		r = TL_messageMediaUnsupported{}
	case crc_messageMediaDocument:
		r = TL_messageMediaDocument{
			Document: m.Object(),
			Caption:  m.String(),
		}
	case crc_messageMediaWebPage:
		r = TL_messageMediaWebPage{
			Webpage: m.Object(),
		}
	case crc_messageMediaVenue:
		r = TL_messageMediaVenue{
			Geo:      m.Object(),
			Title:    m.String(),
			Address:  m.String(),
			Provider: m.String(),
			Venue_id: m.String(),
		}
	case crc_messageMediaGame:
		r = TL_messageMediaGame{
			Game: m.Object(),
		}
	case crc_messageMediaInvoice:
		flags := m.Int()
		shipping_address_requested := flags&(1<<1) != 0
		test := flags&(1<<3) != 0
		title := m.String()
		description := m.String()
		var photo TL
		if flags&(1<<0) != 0 {
			photo = m.Object()
		}
		var receipt_msg_id int32
		if flags&(1<<2) != 0 {
			receipt_msg_id = m.Int()
		}
		currency := m.String()
		total_amount := m.Long()
		start_param := m.String()
		r = TL_messageMediaInvoice{
			Flags: flags,
			Shipping_address_requested: shipping_address_requested,
			Test:           test,
			Title:          title,
			Description:    description,
			Photo:          photo,
			Receipt_msg_id: receipt_msg_id,
			Currency:       currency,
			Total_amount:   total_amount,
			Start_param:    start_param,
		}
	case crc_messageActionEmpty:
		r = TL_messageActionEmpty{}
	case crc_messageActionChatCreate:
		r = TL_messageActionChatCreate{
			Title: m.String(),
			Users: m.VectorInt(),
		}
	case crc_messageActionChatEditTitle:
		r = TL_messageActionChatEditTitle{
			Title: m.String(),
		}
	case crc_messageActionChatEditPhoto:
		r = TL_messageActionChatEditPhoto{
			Photo: m.Object(),
		}
	case crc_messageActionChatDeletePhoto:
		r = TL_messageActionChatDeletePhoto{}
	case crc_messageActionChatAddUser:
		r = TL_messageActionChatAddUser{
			Users: m.VectorInt(),
		}
	case crc_messageActionChatDeleteUser:
		r = TL_messageActionChatDeleteUser{
			User_id: m.Int(),
		}
	case crc_messageActionChatJoinedByLink:
		r = TL_messageActionChatJoinedByLink{
			Inviter_id: m.Int(),
		}
	case crc_messageActionChannelCreate:
		r = TL_messageActionChannelCreate{
			Title: m.String(),
		}
	case crc_messageActionChatMigrateTo:
		r = TL_messageActionChatMigrateTo{
			Channel_id: m.Int(),
		}
	case crc_messageActionChannelMigrateFrom:
		r = TL_messageActionChannelMigrateFrom{
			Title:   m.String(),
			Chat_id: m.Int(),
		}
	case crc_messageActionPinMessage:
		r = TL_messageActionPinMessage{}
	case crc_messageActionHistoryClear:
		r = TL_messageActionHistoryClear{}
	case crc_messageActionGameScore:
		r = TL_messageActionGameScore{
			Game_id: m.Long(),
			Score:   m.Int(),
		}
	case crc_messageActionPaymentSentMe:
		flags := m.Int()
		currency := m.String()
		total_amount := m.Long()
		payload := m.StringBytes()
		var info TL
		if flags&(1<<0) != 0 {
			info = m.Object()
		}
		var shipping_option_id string
		if flags&(1<<1) != 0 {
			shipping_option_id = m.String()
		}
		charge := m.Object()
		r = TL_messageActionPaymentSentMe{
			Flags:              flags,
			Currency:           currency,
			Total_amount:       total_amount,
			Payload:            payload,
			Info:               info,
			Shipping_option_id: shipping_option_id,
			Charge:             charge,
		}
	case crc_messageActionPaymentSent:
		r = TL_messageActionPaymentSent{
			Currency:     m.String(),
			Total_amount: m.Long(),
		}
	case crc_messageActionPhoneCall:
		flags := m.Int()
		call_id := m.Long()
		var reason TL
		if flags&(1<<0) != 0 {
			reason = m.Object()
		}
		var duration int32
		if flags&(1<<1) != 0 {
			duration = m.Int()
		}
		r = TL_messageActionPhoneCall{
			Flags:    flags,
			Call_id:  call_id,
			Reason:   reason,
			Duration: duration,
		}
	case crc_dialog:
		flags := m.Int()
		pinned := flags&(1<<2) != 0
		peer := m.Object()
		top_message := m.Int()
		read_inbox_max_id := m.Int()
		read_outbox_max_id := m.Int()
		unread_count := m.Int()
		notify_settings := m.Object()
		var pts int32
		if flags&(1<<0) != 0 {
			pts = m.Int()
		}
		var draft TL
		if flags&(1<<1) != 0 {
			draft = m.Object()
		}
		r = TL_dialog{
			Flags:              flags,
			Pinned:             pinned,
			Peer:               peer,
			Top_message:        top_message,
			Read_inbox_max_id:  read_inbox_max_id,
			Read_outbox_max_id: read_outbox_max_id,
			Unread_count:       unread_count,
			Notify_settings:    notify_settings,
			Pts:                pts,
			Draft:              draft,
		}
	case crc_photoEmpty:
		r = TL_photoEmpty{
			Id: m.Long(),
		}
	case crc_photo:
		flags := m.Int()
		has_stickers := flags&(1<<0) != 0
		id := m.Long()
		access_hash := m.Long()
		date := m.Int()
		sizes := m.Vector()
		r = TL_photo{
			Flags:        flags,
			Has_stickers: has_stickers,
			Id:           id,
			Access_hash:  access_hash,
			Date:         date,
			Sizes:        sizes,
		}
	case crc_photoSizeEmpty:
		r = TL_photoSizeEmpty{
			Code_type: m.String(),
		}
	case crc_photoSize:
		r = TL_photoSize{
			Code_type: m.String(),
			Location:  m.Object(),
			W:         m.Int(),
			H:         m.Int(),
			Size:      m.Int(),
		}
	case crc_photoCachedSize:
		r = TL_photoCachedSize{
			Code_type: m.String(),
			Location:  m.Object(),
			W:         m.Int(),
			H:         m.Int(),
			Bytes:     m.StringBytes(),
		}
	case crc_geoPointEmpty:
		r = TL_geoPointEmpty{}
	case crc_geoPoint:
		r = TL_geoPoint{
			Long: m.Double(),
			Lat:  m.Double(),
		}
	case crc_auth_checkedPhone:
		r = TL_auth_checkedPhone{
			Phone_registered: m.Object(),
		}
	case crc_auth_sentCode:
		flags := m.Int()
		phone_registered := flags&(1<<0) != 0
		code_type := m.Object()
		phone_code_hash := m.String()
		var next_type TL
		if flags&(1<<1) != 0 {
			next_type = m.Object()
		}
		var timeout int32
		if flags&(1<<2) != 0 {
			timeout = m.Int()
		}
		r = TL_auth_sentCode{
			Flags:            flags,
			Phone_registered: phone_registered,
			Code_type:        code_type,
			Phone_code_hash:  phone_code_hash,
			Next_type:        next_type,
			Timeout:          timeout,
		}
	case crc_auth_authorization:
		flags := m.Int()
		var tmp_sessions int32
		if flags&(1<<0) != 0 {
			tmp_sessions = m.Int()
		}
		user := m.Object()
		r = TL_auth_authorization{
			Flags:        flags,
			Tmp_sessions: tmp_sessions,
			User:         user,
		}
	case crc_auth_exportedAuthorization:
		r = TL_auth_exportedAuthorization{
			Id:    m.Int(),
			Bytes: m.StringBytes(),
		}
	case crc_inputNotifyPeer:
		r = TL_inputNotifyPeer{
			Peer: m.Object(),
		}
	case crc_inputNotifyUsers:
		r = TL_inputNotifyUsers{}
	case crc_inputNotifyChats:
		r = TL_inputNotifyChats{}
	case crc_inputNotifyAll:
		r = TL_inputNotifyAll{}
	case crc_inputPeerNotifyEventsEmpty:
		r = TL_inputPeerNotifyEventsEmpty{}
	case crc_inputPeerNotifyEventsAll:
		r = TL_inputPeerNotifyEventsAll{}
	case crc_inputPeerNotifySettings:
		flags := m.Int()
		show_previews := flags&(1<<0) != 0
		silent := flags&(1<<1) != 0
		mute_until := m.Int()
		sound := m.String()
		r = TL_inputPeerNotifySettings{
			Flags:         flags,
			Show_previews: show_previews,
			Silent:        silent,
			Mute_until:    mute_until,
			Sound:         sound,
		}
	case crc_peerNotifyEventsEmpty:
		r = TL_peerNotifyEventsEmpty{}
	case crc_peerNotifyEventsAll:
		r = TL_peerNotifyEventsAll{}
	case crc_peerNotifySettingsEmpty:
		r = TL_peerNotifySettingsEmpty{}
	case crc_peerNotifySettings:
		flags := m.Int()
		show_previews := flags&(1<<0) != 0
		silent := flags&(1<<1) != 0
		mute_until := m.Int()
		sound := m.String()
		r = TL_peerNotifySettings{
			Flags:         flags,
			Show_previews: show_previews,
			Silent:        silent,
			Mute_until:    mute_until,
			Sound:         sound,
		}
	case crc_peerSettings:
		flags := m.Int()
		report_spam := flags&(1<<0) != 0
		r = TL_peerSettings{
			Flags:       flags,
			Report_spam: report_spam,
		}
	case crc_wallPaper:
		r = TL_wallPaper{
			Id:    m.Int(),
			Title: m.String(),
			Sizes: m.Vector(),
			Color: m.Int(),
		}
	case crc_wallPaperSolid:
		r = TL_wallPaperSolid{
			Id:       m.Int(),
			Title:    m.String(),
			Bg_color: m.Int(),
			Color:    m.Int(),
		}
	case crc_inputReportReasonSpam:
		r = TL_inputReportReasonSpam{}
	case crc_inputReportReasonViolence:
		r = TL_inputReportReasonViolence{}
	case crc_inputReportReasonPornography:
		r = TL_inputReportReasonPornography{}
	case crc_inputReportReasonOther:
		r = TL_inputReportReasonOther{
			Text: m.String(),
		}
	case crc_userFull:
		flags := m.Int()
		blocked := flags&(1<<0) != 0
		phone_calls_available := flags&(1<<4) != 0
		phone_calls_private := flags&(1<<5) != 0
		user := m.Object()
		var about string
		if flags&(1<<1) != 0 {
			about = m.String()
		}
		link := m.Object()
		var profile_photo TL
		if flags&(1<<2) != 0 {
			profile_photo = m.Object()
		}
		notify_settings := m.Object()
		var bot_info TL
		if flags&(1<<3) != 0 {
			bot_info = m.Object()
		}
		common_chats_count := m.Int()
		r = TL_userFull{
			Flags:                 flags,
			Blocked:               blocked,
			Phone_calls_available: phone_calls_available,
			Phone_calls_private:   phone_calls_private,
			User:                  user,
			About:                 about,
			Link:                  link,
			Profile_photo:         profile_photo,
			Notify_settings:       notify_settings,
			Bot_info:              bot_info,
			Common_chats_count:    common_chats_count,
		}
	case crc_contact:
		r = TL_contact{
			User_id: m.Int(),
			Mutual:  m.Object(),
		}
	case crc_importedContact:
		r = TL_importedContact{
			User_id:   m.Int(),
			Client_id: m.Long(),
		}
	case crc_contactBlocked:
		r = TL_contactBlocked{
			User_id: m.Int(),
			Date:    m.Int(),
		}
	case crc_contactStatus:
		r = TL_contactStatus{
			User_id: m.Int(),
			Status:  m.Object(),
		}
	case crc_contacts_link:
		r = TL_contacts_link{
			My_link:      m.Object(),
			Foreign_link: m.Object(),
			User:         m.Object(),
		}
	case crc_contacts_contactsNotModified:
		r = TL_contacts_contactsNotModified{}
	case crc_contacts_contacts:
		r = TL_contacts_contacts{
			Contacts: m.Vector(),
			Users:    m.Vector(),
		}
	case crc_contacts_importedContacts:
		r = TL_contacts_importedContacts{
			Imported:       m.Vector(),
			Retry_contacts: m.VectorLong(),
			Users:          m.Vector(),
		}
	case crc_contacts_blocked:
		r = TL_contacts_blocked{
			Blocked: m.Vector(),
			Users:   m.Vector(),
		}
	case crc_contacts_blockedSlice:
		r = TL_contacts_blockedSlice{
			Count:   m.Int(),
			Blocked: m.Vector(),
			Users:   m.Vector(),
		}
	case crc_messages_dialogs:
		r = TL_messages_dialogs{
			Dialogs:  m.Vector(),
			Messages: m.Vector(),
			Chats:    m.Vector(),
			Users:    m.Vector(),
		}
	case crc_messages_dialogsSlice:
		r = TL_messages_dialogsSlice{
			Count:    m.Int(),
			Dialogs:  m.Vector(),
			Messages: m.Vector(),
			Chats:    m.Vector(),
			Users:    m.Vector(),
		}
	case crc_messages_messages:
		r = TL_messages_messages{
			Messages: m.Vector(),
			Chats:    m.Vector(),
			Users:    m.Vector(),
		}
	case crc_messages_messagesSlice:
		r = TL_messages_messagesSlice{
			Count:    m.Int(),
			Messages: m.Vector(),
			Chats:    m.Vector(),
			Users:    m.Vector(),
		}
	case crc_messages_channelMessages:
		flags := m.Int()
		pts := m.Int()
		count := m.Int()
		messages := m.Vector()
		chats := m.Vector()
		users := m.Vector()
		r = TL_messages_channelMessages{
			Flags:    flags,
			Pts:      pts,
			Count:    count,
			Messages: messages,
			Chats:    chats,
			Users:    users,
		}
	case crc_messages_chats:
		r = TL_messages_chats{
			Chats: m.Vector(),
		}
	case crc_messages_chatsSlice:
		r = TL_messages_chatsSlice{
			Count: m.Int(),
			Chats: m.Vector(),
		}
	case crc_messages_chatFull:
		r = TL_messages_chatFull{
			Full_chat: m.Object(),
			Chats:     m.Vector(),
			Users:     m.Vector(),
		}
	case crc_messages_affectedHistory:
		r = TL_messages_affectedHistory{
			Pts:       m.Int(),
			Pts_count: m.Int(),
			Offset:    m.Int(),
		}
	case crc_inputMessagesFilterEmpty:
		r = TL_inputMessagesFilterEmpty{}
	case crc_inputMessagesFilterPhotos:
		r = TL_inputMessagesFilterPhotos{}
	case crc_inputMessagesFilterVideo:
		r = TL_inputMessagesFilterVideo{}
	case crc_inputMessagesFilterPhotoVideo:
		r = TL_inputMessagesFilterPhotoVideo{}
	case crc_inputMessagesFilterPhotoVideoDocuments:
		r = TL_inputMessagesFilterPhotoVideoDocuments{}
	case crc_inputMessagesFilterDocument:
		r = TL_inputMessagesFilterDocument{}
	case crc_inputMessagesFilterUrl:
		r = TL_inputMessagesFilterUrl{}
	case crc_inputMessagesFilterGif:
		r = TL_inputMessagesFilterGif{}
	case crc_inputMessagesFilterVoice:
		r = TL_inputMessagesFilterVoice{}
	case crc_inputMessagesFilterMusic:
		r = TL_inputMessagesFilterMusic{}
	case crc_inputMessagesFilterChatPhotos:
		r = TL_inputMessagesFilterChatPhotos{}
	case crc_inputMessagesFilterPhoneCalls:
		flags := m.Int()
		missed := flags&(1<<0) != 0
		r = TL_inputMessagesFilterPhoneCalls{
			Flags:  flags,
			Missed: missed,
		}
	case crc_updateNewMessage:
		r = TL_updateNewMessage{
			Message:   m.Object(),
			Pts:       m.Int(),
			Pts_count: m.Int(),
		}
	case crc_updateMessageID:
		r = TL_updateMessageID{
			Id:        m.Int(),
			Random_id: m.Long(),
		}
	case crc_updateDeleteMessages:
		r = TL_updateDeleteMessages{
			Messages:  m.VectorInt(),
			Pts:       m.Int(),
			Pts_count: m.Int(),
		}
	case crc_updateUserTyping:
		r = TL_updateUserTyping{
			User_id: m.Int(),
			Action:  m.Object(),
		}
	case crc_updateChatUserTyping:
		r = TL_updateChatUserTyping{
			Chat_id: m.Int(),
			User_id: m.Int(),
			Action:  m.Object(),
		}
	case crc_updateChatParticipants:
		r = TL_updateChatParticipants{
			Participants: m.Object(),
		}
	case crc_updateUserStatus:
		r = TL_updateUserStatus{
			User_id: m.Int(),
			Status:  m.Object(),
		}
	case crc_updateUserName:
		r = TL_updateUserName{
			User_id:    m.Int(),
			First_name: m.String(),
			Last_name:  m.String(),
			Username:   m.String(),
		}
	case crc_updateUserPhoto:
		r = TL_updateUserPhoto{
			User_id:  m.Int(),
			Date:     m.Int(),
			Photo:    m.Object(),
			Previous: m.Object(),
		}
	case crc_updateContactRegistered:
		r = TL_updateContactRegistered{
			User_id: m.Int(),
			Date:    m.Int(),
		}
	case crc_updateContactLink:
		r = TL_updateContactLink{
			User_id:      m.Int(),
			My_link:      m.Object(),
			Foreign_link: m.Object(),
		}
	case crc_updateNewEncryptedMessage:
		r = TL_updateNewEncryptedMessage{
			Message: m.Object(),
			Qts:     m.Int(),
		}
	case crc_updateEncryptedChatTyping:
		r = TL_updateEncryptedChatTyping{
			Chat_id: m.Int(),
		}
	case crc_updateEncryption:
		r = TL_updateEncryption{
			Chat: m.Object(),
			Date: m.Int(),
		}
	case crc_updateEncryptedMessagesRead:
		r = TL_updateEncryptedMessagesRead{
			Chat_id:  m.Int(),
			Max_date: m.Int(),
			Date:     m.Int(),
		}
	case crc_updateChatParticipantAdd:
		r = TL_updateChatParticipantAdd{
			Chat_id:    m.Int(),
			User_id:    m.Int(),
			Inviter_id: m.Int(),
			Date:       m.Int(),
			Version:    m.Int(),
		}
	case crc_updateChatParticipantDelete:
		r = TL_updateChatParticipantDelete{
			Chat_id: m.Int(),
			User_id: m.Int(),
			Version: m.Int(),
		}
	case crc_updateDcOptions:
		r = TL_updateDcOptions{
			Dc_options: m.Vector(),
		}
	case crc_updateUserBlocked:
		r = TL_updateUserBlocked{
			User_id: m.Int(),
			Blocked: m.Object(),
		}
	case crc_updateNotifySettings:
		r = TL_updateNotifySettings{
			Peer:            m.Object(),
			Notify_settings: m.Object(),
		}
	case crc_updateServiceNotification:
		flags := m.Int()
		popup := flags&(1<<0) != 0
		var inbox_date int32
		if flags&(1<<1) != 0 {
			inbox_date = m.Int()
		}
		code_type := m.String()
		message := m.String()
		media := m.Object()
		entities := m.Vector()
		r = TL_updateServiceNotification{
			Flags:      flags,
			Popup:      popup,
			Inbox_date: inbox_date,
			Code_type:  code_type,
			Message:    message,
			Media:      media,
			Entities:   entities,
		}
	case crc_updatePrivacy:
		r = TL_updatePrivacy{
			Key:   m.Object(),
			Rules: m.Vector(),
		}
	case crc_updateUserPhone:
		r = TL_updateUserPhone{
			User_id: m.Int(),
			Phone:   m.String(),
		}
	case crc_updateReadHistoryInbox:
		r = TL_updateReadHistoryInbox{
			Peer:      m.Object(),
			Max_id:    m.Int(),
			Pts:       m.Int(),
			Pts_count: m.Int(),
		}
	case crc_updateReadHistoryOutbox:
		r = TL_updateReadHistoryOutbox{
			Peer:      m.Object(),
			Max_id:    m.Int(),
			Pts:       m.Int(),
			Pts_count: m.Int(),
		}
	case crc_updateWebPage:
		r = TL_updateWebPage{
			Webpage:   m.Object(),
			Pts:       m.Int(),
			Pts_count: m.Int(),
		}
	case crc_updateReadMessagesContents:
		r = TL_updateReadMessagesContents{
			Messages:  m.VectorInt(),
			Pts:       m.Int(),
			Pts_count: m.Int(),
		}
	case crc_updateChannelTooLong:
		flags := m.Int()
		channel_id := m.Int()
		var pts int32
		if flags&(1<<0) != 0 {
			pts = m.Int()
		}
		r = TL_updateChannelTooLong{
			Flags:      flags,
			Channel_id: channel_id,
			Pts:        pts,
		}
	case crc_updateChannel:
		r = TL_updateChannel{
			Channel_id: m.Int(),
		}
	case crc_updateNewChannelMessage:
		r = TL_updateNewChannelMessage{
			Message:   m.Object(),
			Pts:       m.Int(),
			Pts_count: m.Int(),
		}
	case crc_updateReadChannelInbox:
		r = TL_updateReadChannelInbox{
			Channel_id: m.Int(),
			Max_id:     m.Int(),
		}
	case crc_updateDeleteChannelMessages:
		r = TL_updateDeleteChannelMessages{
			Channel_id: m.Int(),
			Messages:   m.VectorInt(),
			Pts:        m.Int(),
			Pts_count:  m.Int(),
		}
	case crc_updateChannelMessageViews:
		r = TL_updateChannelMessageViews{
			Channel_id: m.Int(),
			Id:         m.Int(),
			Views:      m.Int(),
		}
	case crc_updateChatAdmins:
		r = TL_updateChatAdmins{
			Chat_id: m.Int(),
			Enabled: m.Object(),
			Version: m.Int(),
		}
	case crc_updateChatParticipantAdmin:
		r = TL_updateChatParticipantAdmin{
			Chat_id:  m.Int(),
			User_id:  m.Int(),
			Is_admin: m.Object(),
			Version:  m.Int(),
		}
	case crc_updateNewStickerSet:
		r = TL_updateNewStickerSet{
			Stickerset: m.Object(),
		}
	case crc_updateStickerSetsOrder:
		flags := m.Int()
		masks := flags&(1<<0) != 0
		order := m.VectorLong()
		r = TL_updateStickerSetsOrder{
			Flags: flags,
			Masks: masks,
			Order: order,
		}
	case crc_updateStickerSets:
		r = TL_updateStickerSets{}
	case crc_updateSavedGifs:
		r = TL_updateSavedGifs{}
	case crc_updateBotInlineQuery:
		flags := m.Int()
		query_id := m.Long()
		user_id := m.Int()
		query := m.String()
		var geo TL
		if flags&(1<<0) != 0 {
			geo = m.Object()
		}
		offset := m.String()
		r = TL_updateBotInlineQuery{
			Flags:    flags,
			Query_id: query_id,
			User_id:  user_id,
			Query:    query,
			Geo:      geo,
			Offset:   offset,
		}
	case crc_updateBotInlineSend:
		flags := m.Int()
		user_id := m.Int()
		query := m.String()
		var geo TL
		if flags&(1<<0) != 0 {
			geo = m.Object()
		}
		id := m.String()
		var msg_id TL
		if flags&(1<<1) != 0 {
			msg_id = m.Object()
		}
		r = TL_updateBotInlineSend{
			Flags:   flags,
			User_id: user_id,
			Query:   query,
			Geo:     geo,
			Id:      id,
			Msg_id:  msg_id,
		}
	case crc_updateEditChannelMessage:
		r = TL_updateEditChannelMessage{
			Message:   m.Object(),
			Pts:       m.Int(),
			Pts_count: m.Int(),
		}
	case crc_updateChannelPinnedMessage:
		r = TL_updateChannelPinnedMessage{
			Channel_id: m.Int(),
			Id:         m.Int(),
		}
	case crc_updateBotCallbackQuery:
		flags := m.Int()
		query_id := m.Long()
		user_id := m.Int()
		peer := m.Object()
		msg_id := m.Int()
		chat_instance := m.Long()
		var data []byte
		if flags&(1<<0) != 0 {
			data = m.StringBytes()
		}
		var game_short_name string
		if flags&(1<<1) != 0 {
			game_short_name = m.String()
		}
		r = TL_updateBotCallbackQuery{
			Flags:           flags,
			Query_id:        query_id,
			User_id:         user_id,
			Peer:            peer,
			Msg_id:          msg_id,
			Chat_instance:   chat_instance,
			Data:            data,
			Game_short_name: game_short_name,
		}
	case crc_updateEditMessage:
		r = TL_updateEditMessage{
			Message:   m.Object(),
			Pts:       m.Int(),
			Pts_count: m.Int(),
		}
	case crc_updateInlineBotCallbackQuery:
		flags := m.Int()
		query_id := m.Long()
		user_id := m.Int()
		msg_id := m.Object()
		chat_instance := m.Long()
		var data []byte
		if flags&(1<<0) != 0 {
			data = m.StringBytes()
		}
		var game_short_name string
		if flags&(1<<1) != 0 {
			game_short_name = m.String()
		}
		r = TL_updateInlineBotCallbackQuery{
			Flags:           flags,
			Query_id:        query_id,
			User_id:         user_id,
			Msg_id:          msg_id,
			Chat_instance:   chat_instance,
			Data:            data,
			Game_short_name: game_short_name,
		}
	case crc_updateReadChannelOutbox:
		r = TL_updateReadChannelOutbox{
			Channel_id: m.Int(),
			Max_id:     m.Int(),
		}
	case crc_updateDraftMessage:
		r = TL_updateDraftMessage{
			Peer:  m.Object(),
			Draft: m.Object(),
		}
	case crc_updateReadFeaturedStickers:
		r = TL_updateReadFeaturedStickers{}
	case crc_updateRecentStickers:
		r = TL_updateRecentStickers{}
	case crc_updateConfig:
		r = TL_updateConfig{}
	case crc_updatePtsChanged:
		r = TL_updatePtsChanged{}
	case crc_updateChannelWebPage:
		r = TL_updateChannelWebPage{
			Channel_id: m.Int(),
			Webpage:    m.Object(),
			Pts:        m.Int(),
			Pts_count:  m.Int(),
		}
	case crc_updateDialogPinned:
		flags := m.Int()
		pinned := flags&(1<<0) != 0
		peer := m.Object()
		r = TL_updateDialogPinned{
			Flags:  flags,
			Pinned: pinned,
			Peer:   peer,
		}
	case crc_updatePinnedDialogs:
		flags := m.Int()
		var order []TL
		if flags&(1<<0) != 0 {
			order = m.Vector()
		}
		r = TL_updatePinnedDialogs{
			Flags: flags,
			Order: order,
		}
	case crc_updateBotWebhookJSON:
		r = TL_updateBotWebhookJSON{
			Data: m.Object(),
		}
	case crc_updateBotWebhookJSONQuery:
		r = TL_updateBotWebhookJSONQuery{
			Query_id: m.Long(),
			Data:     m.Object(),
			Timeout:  m.Int(),
		}
	case crc_updateBotShippingQuery:
		r = TL_updateBotShippingQuery{
			Query_id:         m.Long(),
			User_id:          m.Int(),
			Payload:          m.StringBytes(),
			Shipping_address: m.Object(),
		}
	case crc_updateBotPrecheckoutQuery:
		flags := m.Int()
		query_id := m.Long()
		user_id := m.Int()
		payload := m.StringBytes()
		var info TL
		if flags&(1<<0) != 0 {
			info = m.Object()
		}
		var shipping_option_id string
		if flags&(1<<1) != 0 {
			shipping_option_id = m.String()
		}
		currency := m.String()
		total_amount := m.Long()
		r = TL_updateBotPrecheckoutQuery{
			Flags:              flags,
			Query_id:           query_id,
			User_id:            user_id,
			Payload:            payload,
			Info:               info,
			Shipping_option_id: shipping_option_id,
			Currency:           currency,
			Total_amount:       total_amount,
		}
	case crc_updatePhoneCall:
		r = TL_updatePhoneCall{
			Phone_call: m.Object(),
		}
	case crc_updates_state:
		r = TL_updates_state{
			Pts:          m.Int(),
			Qts:          m.Int(),
			Date:         m.Int(),
			Seq:          m.Int(),
			Unread_count: m.Int(),
		}
	case crc_updates_differenceEmpty:
		r = TL_updates_differenceEmpty{
			Date: m.Int(),
			Seq:  m.Int(),
		}
	case crc_updates_difference:
		r = TL_updates_difference{
			New_messages:           m.Vector(),
			New_encrypted_messages: m.Vector(),
			Other_updates:          m.Vector(),
			Chats:                  m.Vector(),
			Users:                  m.Vector(),
			State:                  m.Object(),
		}
	case crc_updates_differenceSlice:
		r = TL_updates_differenceSlice{
			New_messages:           m.Vector(),
			New_encrypted_messages: m.Vector(),
			Other_updates:          m.Vector(),
			Chats:                  m.Vector(),
			Users:                  m.Vector(),
			Intermediate_state:     m.Object(),
		}
	case crc_updates_differenceTooLong:
		r = TL_updates_differenceTooLong{
			Pts: m.Int(),
		}
	case crc_updatesTooLong:
		r = TL_updatesTooLong{}
	case crc_updateShortMessage:
		flags := m.Int()
		out := flags&(1<<1) != 0
		mentioned := flags&(1<<4) != 0
		media_unread := flags&(1<<5) != 0
		silent := flags&(1<<13) != 0
		id := m.Int()
		user_id := m.Int()
		message := m.String()
		pts := m.Int()
		pts_count := m.Int()
		date := m.Int()
		var fwd_from TL
		if flags&(1<<2) != 0 {
			fwd_from = m.Object()
		}
		var via_bot_id int32
		if flags&(1<<11) != 0 {
			via_bot_id = m.Int()
		}
		var reply_to_msg_id int32
		if flags&(1<<3) != 0 {
			reply_to_msg_id = m.Int()
		}
		var entities []TL
		if flags&(1<<7) != 0 {
			entities = m.Vector()
		}
		r = TL_updateShortMessage{
			Flags:           flags,
			Out:             out,
			Mentioned:       mentioned,
			Media_unread:    media_unread,
			Silent:          silent,
			Id:              id,
			User_id:         user_id,
			Message:         message,
			Pts:             pts,
			Pts_count:       pts_count,
			Date:            date,
			Fwd_from:        fwd_from,
			Via_bot_id:      via_bot_id,
			Reply_to_msg_id: reply_to_msg_id,
			Entities:        entities,
		}
	case crc_updateShortChatMessage:
		flags := m.Int()
		out := flags&(1<<1) != 0
		mentioned := flags&(1<<4) != 0
		media_unread := flags&(1<<5) != 0
		silent := flags&(1<<13) != 0
		id := m.Int()
		from_id := m.Int()
		chat_id := m.Int()
		message := m.String()
		pts := m.Int()
		pts_count := m.Int()
		date := m.Int()
		var fwd_from TL
		if flags&(1<<2) != 0 {
			fwd_from = m.Object()
		}
		var via_bot_id int32
		if flags&(1<<11) != 0 {
			via_bot_id = m.Int()
		}
		var reply_to_msg_id int32
		if flags&(1<<3) != 0 {
			reply_to_msg_id = m.Int()
		}
		var entities []TL
		if flags&(1<<7) != 0 {
			entities = m.Vector()
		}
		r = TL_updateShortChatMessage{
			Flags:           flags,
			Out:             out,
			Mentioned:       mentioned,
			Media_unread:    media_unread,
			Silent:          silent,
			Id:              id,
			From_id:         from_id,
			Chat_id:         chat_id,
			Message:         message,
			Pts:             pts,
			Pts_count:       pts_count,
			Date:            date,
			Fwd_from:        fwd_from,
			Via_bot_id:      via_bot_id,
			Reply_to_msg_id: reply_to_msg_id,
			Entities:        entities,
		}
	case crc_updateShort:
		r = TL_updateShort{
			Update: m.Object(),
			Date:   m.Int(),
		}
	case crc_updatesCombined:
		r = TL_updatesCombined{
			Updates:   m.Vector(),
			Users:     m.Vector(),
			Chats:     m.Vector(),
			Date:      m.Int(),
			Seq_start: m.Int(),
			Seq:       m.Int(),
		}
	case crc_updates:
		r = TL_updates{
			Updates: m.Vector(),
			Users:   m.Vector(),
			Chats:   m.Vector(),
			Date:    m.Int(),
			Seq:     m.Int(),
		}
	case crc_updateShortSentMessage:
		flags := m.Int()
		out := flags&(1<<1) != 0
		id := m.Int()
		pts := m.Int()
		pts_count := m.Int()
		date := m.Int()
		var media TL
		if flags&(1<<9) != 0 {
			media = m.Object()
		}
		var entities []TL
		if flags&(1<<7) != 0 {
			entities = m.Vector()
		}
		r = TL_updateShortSentMessage{
			Flags:     flags,
			Out:       out,
			Id:        id,
			Pts:       pts,
			Pts_count: pts_count,
			Date:      date,
			Media:     media,
			Entities:  entities,
		}
	case crc_photos_photos:
		r = TL_photos_photos{
			Photos: m.Vector(),
			Users:  m.Vector(),
		}
	case crc_photos_photosSlice:
		r = TL_photos_photosSlice{
			Count:  m.Int(),
			Photos: m.Vector(),
			Users:  m.Vector(),
		}
	case crc_photos_photo:
		r = TL_photos_photo{
			Photo: m.Object(),
			Users: m.Vector(),
		}
	case crc_upload_file:
		r = TL_upload_file{
			Code_type: m.Object(),
			Mtime:     m.Int(),
			Bytes:     m.StringBytes(),
		}
	case crc_dcOption:
		flags := m.Int()
		ipv6 := flags&(1<<0) != 0
		media_only := flags&(1<<1) != 0
		tcpo_only := flags&(1<<2) != 0
		id := m.Int()
		ip_address := m.String()
		port := m.Int()
		r = TL_dcOption{
			Flags:      flags,
			Ipv6:       ipv6,
			Media_only: media_only,
			Tcpo_only:  tcpo_only,
			Id:         id,
			Ip_address: ip_address,
			Port:       port,
		}
	case crc_config:
		flags := m.Int()
		phonecalls_enabled := flags&(1<<1) != 0
		date := m.Int()
		expires := m.Int()
		test_mode := m.Object()
		this_dc := m.Int()
		dc_options := m.Vector()
		chat_size_max := m.Int()
		megagroup_size_max := m.Int()
		forwarded_count_max := m.Int()
		online_update_period_ms := m.Int()
		offline_blur_timeout_ms := m.Int()
		offline_idle_timeout_ms := m.Int()
		online_cloud_timeout_ms := m.Int()
		notify_cloud_delay_ms := m.Int()
		notify_default_delay_ms := m.Int()
		chat_big_size := m.Int()
		push_chat_period_ms := m.Int()
		push_chat_limit := m.Int()
		saved_gifs_limit := m.Int()
		edit_time_limit := m.Int()
		rating_e_decay := m.Int()
		stickers_recent_limit := m.Int()
		var tmp_sessions int32
		if flags&(1<<0) != 0 {
			tmp_sessions = m.Int()
		}
		pinned_dialogs_count_max := m.Int()
		call_receive_timeout_ms := m.Int()
		call_ring_timeout_ms := m.Int()
		call_connect_timeout_ms := m.Int()
		call_packet_timeout_ms := m.Int()
		me_url_prefix := m.String()
		disabled_features := m.Vector()
		r = TL_config{
			Flags:                    flags,
			Phonecalls_enabled:       phonecalls_enabled,
			Date:                     date,
			Expires:                  expires,
			Test_mode:                test_mode,
			This_dc:                  this_dc,
			Dc_options:               dc_options,
			Chat_size_max:            chat_size_max,
			Megagroup_size_max:       megagroup_size_max,
			Forwarded_count_max:      forwarded_count_max,
			Online_update_period_ms:  online_update_period_ms,
			Offline_blur_timeout_ms:  offline_blur_timeout_ms,
			Offline_idle_timeout_ms:  offline_idle_timeout_ms,
			Online_cloud_timeout_ms:  online_cloud_timeout_ms,
			Notify_cloud_delay_ms:    notify_cloud_delay_ms,
			Notify_default_delay_ms:  notify_default_delay_ms,
			Chat_big_size:            chat_big_size,
			Push_chat_period_ms:      push_chat_period_ms,
			Push_chat_limit:          push_chat_limit,
			Saved_gifs_limit:         saved_gifs_limit,
			Edit_time_limit:          edit_time_limit,
			Rating_e_decay:           rating_e_decay,
			Stickers_recent_limit:    stickers_recent_limit,
			Tmp_sessions:             tmp_sessions,
			Pinned_dialogs_count_max: pinned_dialogs_count_max,
			Call_receive_timeout_ms:  call_receive_timeout_ms,
			Call_ring_timeout_ms:     call_ring_timeout_ms,
			Call_connect_timeout_ms:  call_connect_timeout_ms,
			Call_packet_timeout_ms:   call_packet_timeout_ms,
			Me_url_prefix:            me_url_prefix,
			Disabled_features:        disabled_features,
		}
	case crc_nearestDc:
		r = TL_nearestDc{
			Country:    m.String(),
			This_dc:    m.Int(),
			Nearest_dc: m.Int(),
		}
	case crc_help_appUpdate:
		r = TL_help_appUpdate{
			Id:       m.Int(),
			Critical: m.Object(),
			Url:      m.String(),
			Text:     m.String(),
		}
	case crc_help_noAppUpdate:
		r = TL_help_noAppUpdate{}
	case crc_help_inviteText:
		r = TL_help_inviteText{
			Message: m.String(),
		}
	case crc_encryptedChatEmpty:
		r = TL_encryptedChatEmpty{
			Id: m.Int(),
		}
	case crc_encryptedChatWaiting:
		r = TL_encryptedChatWaiting{
			Id:             m.Int(),
			Access_hash:    m.Long(),
			Date:           m.Int(),
			Admin_id:       m.Int(),
			Participant_id: m.Int(),
		}
	case crc_encryptedChatRequested:
		r = TL_encryptedChatRequested{
			Id:             m.Int(),
			Access_hash:    m.Long(),
			Date:           m.Int(),
			Admin_id:       m.Int(),
			Participant_id: m.Int(),
			G_a:            m.StringBytes(),
		}
	case crc_encryptedChat:
		r = TL_encryptedChat{
			Id:              m.Int(),
			Access_hash:     m.Long(),
			Date:            m.Int(),
			Admin_id:        m.Int(),
			Participant_id:  m.Int(),
			G_a_or_b:        m.StringBytes(),
			Key_fingerprint: m.Long(),
		}
	case crc_encryptedChatDiscarded:
		r = TL_encryptedChatDiscarded{
			Id: m.Int(),
		}
	case crc_inputEncryptedChat:
		r = TL_inputEncryptedChat{
			Chat_id:     m.Int(),
			Access_hash: m.Long(),
		}
	case crc_encryptedFileEmpty:
		r = TL_encryptedFileEmpty{}
	case crc_encryptedFile:
		r = TL_encryptedFile{
			Id:              m.Long(),
			Access_hash:     m.Long(),
			Size:            m.Int(),
			Dc_id:           m.Int(),
			Key_fingerprint: m.Int(),
		}
	case crc_inputEncryptedFileEmpty:
		r = TL_inputEncryptedFileEmpty{}
	case crc_inputEncryptedFileUploaded:
		r = TL_inputEncryptedFileUploaded{
			Id:              m.Long(),
			Parts:           m.Int(),
			Md5_checksum:    m.String(),
			Key_fingerprint: m.Int(),
		}
	case crc_inputEncryptedFile:
		r = TL_inputEncryptedFile{
			Id:          m.Long(),
			Access_hash: m.Long(),
		}
	case crc_inputEncryptedFileBigUploaded:
		r = TL_inputEncryptedFileBigUploaded{
			Id:              m.Long(),
			Parts:           m.Int(),
			Key_fingerprint: m.Int(),
		}
	case crc_encryptedMessage:
		r = TL_encryptedMessage{
			Random_id: m.Long(),
			Chat_id:   m.Int(),
			Date:      m.Int(),
			Bytes:     m.StringBytes(),
			File:      m.Object(),
		}
	case crc_encryptedMessageService:
		r = TL_encryptedMessageService{
			Random_id: m.Long(),
			Chat_id:   m.Int(),
			Date:      m.Int(),
			Bytes:     m.StringBytes(),
		}
	case crc_messages_dhConfigNotModified:
		r = TL_messages_dhConfigNotModified{
			Random: m.StringBytes(),
		}
	case crc_messages_dhConfig:
		r = TL_messages_dhConfig{
			G:       m.Int(),
			P:       m.StringBytes(),
			Version: m.Int(),
			Random:  m.StringBytes(),
		}
	case crc_messages_sentEncryptedMessage:
		r = TL_messages_sentEncryptedMessage{
			Date: m.Int(),
		}
	case crc_messages_sentEncryptedFile:
		r = TL_messages_sentEncryptedFile{
			Date: m.Int(),
			File: m.Object(),
		}
	case crc_inputDocumentEmpty:
		r = TL_inputDocumentEmpty{}
	case crc_inputDocument:
		r = TL_inputDocument{
			Id:          m.Long(),
			Access_hash: m.Long(),
		}
	case crc_documentEmpty:
		r = TL_documentEmpty{
			Id: m.Long(),
		}
	case crc_document:
		r = TL_document{
			Id:          m.Long(),
			Access_hash: m.Long(),
			Date:        m.Int(),
			Mime_type:   m.String(),
			Size:        m.Int(),
			Thumb:       m.Object(),
			Dc_id:       m.Int(),
			Version:     m.Int(),
			Attributes:  m.Vector(),
		}
	case crc_help_support:
		r = TL_help_support{
			Phone_number: m.String(),
			User:         m.Object(),
		}
	case crc_notifyPeer:
		r = TL_notifyPeer{
			Peer: m.Object(),
		}
	case crc_notifyUsers:
		r = TL_notifyUsers{}
	case crc_notifyChats:
		r = TL_notifyChats{}
	case crc_notifyAll:
		r = TL_notifyAll{}
	case crc_sendMessageTypingAction:
		r = TL_sendMessageTypingAction{}
	case crc_sendMessageCancelAction:
		r = TL_sendMessageCancelAction{}
	case crc_sendMessageRecordVideoAction:
		r = TL_sendMessageRecordVideoAction{}
	case crc_sendMessageUploadVideoAction:
		r = TL_sendMessageUploadVideoAction{
			Progress: m.Int(),
		}
	case crc_sendMessageRecordAudioAction:
		r = TL_sendMessageRecordAudioAction{}
	case crc_sendMessageUploadAudioAction:
		r = TL_sendMessageUploadAudioAction{
			Progress: m.Int(),
		}
	case crc_sendMessageUploadPhotoAction:
		r = TL_sendMessageUploadPhotoAction{
			Progress: m.Int(),
		}
	case crc_sendMessageUploadDocumentAction:
		r = TL_sendMessageUploadDocumentAction{
			Progress: m.Int(),
		}
	case crc_sendMessageGeoLocationAction:
		r = TL_sendMessageGeoLocationAction{}
	case crc_sendMessageChooseContactAction:
		r = TL_sendMessageChooseContactAction{}
	case crc_sendMessageGamePlayAction:
		r = TL_sendMessageGamePlayAction{}
	case crc_contacts_found:
		r = TL_contacts_found{
			Results: m.Vector(),
			Chats:   m.Vector(),
			Users:   m.Vector(),
		}
	case crc_inputPrivacyKeyStatusTimestamp:
		r = TL_inputPrivacyKeyStatusTimestamp{}
	case crc_inputPrivacyKeyChatInvite:
		r = TL_inputPrivacyKeyChatInvite{}
	case crc_inputPrivacyKeyPhoneCall:
		r = TL_inputPrivacyKeyPhoneCall{}
	case crc_privacyKeyStatusTimestamp:
		r = TL_privacyKeyStatusTimestamp{}
	case crc_privacyKeyChatInvite:
		r = TL_privacyKeyChatInvite{}
	case crc_privacyKeyPhoneCall:
		r = TL_privacyKeyPhoneCall{}
	case crc_inputPrivacyValueAllowContacts:
		r = TL_inputPrivacyValueAllowContacts{}
	case crc_inputPrivacyValueAllowAll:
		r = TL_inputPrivacyValueAllowAll{}
	case crc_inputPrivacyValueAllowUsers:
		r = TL_inputPrivacyValueAllowUsers{
			Users: m.Vector(),
		}
	case crc_inputPrivacyValueDisallowContacts:
		r = TL_inputPrivacyValueDisallowContacts{}
	case crc_inputPrivacyValueDisallowAll:
		r = TL_inputPrivacyValueDisallowAll{}
	case crc_inputPrivacyValueDisallowUsers:
		r = TL_inputPrivacyValueDisallowUsers{
			Users: m.Vector(),
		}
	case crc_privacyValueAllowContacts:
		r = TL_privacyValueAllowContacts{}
	case crc_privacyValueAllowAll:
		r = TL_privacyValueAllowAll{}
	case crc_privacyValueAllowUsers:
		r = TL_privacyValueAllowUsers{
			Users: m.VectorInt(),
		}
	case crc_privacyValueDisallowContacts:
		r = TL_privacyValueDisallowContacts{}
	case crc_privacyValueDisallowAll:
		r = TL_privacyValueDisallowAll{}
	case crc_privacyValueDisallowUsers:
		r = TL_privacyValueDisallowUsers{
			Users: m.VectorInt(),
		}
	case crc_account_privacyRules:
		r = TL_account_privacyRules{
			Rules: m.Vector(),
			Users: m.Vector(),
		}
	case crc_accountDaysTTL:
		r = TL_accountDaysTTL{
			Days: m.Int(),
		}
	case crc_documentAttributeImageSize:
		r = TL_documentAttributeImageSize{
			W: m.Int(),
			H: m.Int(),
		}
	case crc_documentAttributeAnimated:
		r = TL_documentAttributeAnimated{}
	case crc_documentAttributeSticker:
		flags := m.Int()
		mask := flags&(1<<1) != 0
		alt := m.String()
		stickerset := m.Object()
		var mask_coords TL
		if flags&(1<<0) != 0 {
			mask_coords = m.Object()
		}
		r = TL_documentAttributeSticker{
			Flags:       flags,
			Mask:        mask,
			Alt:         alt,
			Stickerset:  stickerset,
			Mask_coords: mask_coords,
		}
	case crc_documentAttributeVideo:
		r = TL_documentAttributeVideo{
			Duration: m.Int(),
			W:        m.Int(),
			H:        m.Int(),
		}
	case crc_documentAttributeAudio:
		flags := m.Int()
		voice := flags&(1<<10) != 0
		duration := m.Int()
		var title string
		if flags&(1<<0) != 0 {
			title = m.String()
		}
		var performer string
		if flags&(1<<1) != 0 {
			performer = m.String()
		}
		var waveform []byte
		if flags&(1<<2) != 0 {
			waveform = m.StringBytes()
		}
		r = TL_documentAttributeAudio{
			Flags:     flags,
			Voice:     voice,
			Duration:  duration,
			Title:     title,
			Performer: performer,
			Waveform:  waveform,
		}
	case crc_documentAttributeFilename:
		r = TL_documentAttributeFilename{
			File_name: m.String(),
		}
	case crc_documentAttributeHasStickers:
		r = TL_documentAttributeHasStickers{}
	case crc_messages_stickersNotModified:
		r = TL_messages_stickersNotModified{}
	case crc_messages_stickers:
		r = TL_messages_stickers{
			Hash:     m.String(),
			Stickers: m.Vector(),
		}
	case crc_stickerPack:
		r = TL_stickerPack{
			Emoticon:  m.String(),
			Documents: m.VectorLong(),
		}
	case crc_messages_allStickersNotModified:
		r = TL_messages_allStickersNotModified{}
	case crc_messages_allStickers:
		r = TL_messages_allStickers{
			Hash: m.Int(),
			Sets: m.Vector(),
		}
	case crc_disabledFeature:
		r = TL_disabledFeature{
			Feature:     m.String(),
			Description: m.String(),
		}
	case crc_messages_affectedMessages:
		r = TL_messages_affectedMessages{
			Pts:       m.Int(),
			Pts_count: m.Int(),
		}
	case crc_contactLinkUnknown:
		r = TL_contactLinkUnknown{}
	case crc_contactLinkNone:
		r = TL_contactLinkNone{}
	case crc_contactLinkHasPhone:
		r = TL_contactLinkHasPhone{}
	case crc_contactLinkContact:
		r = TL_contactLinkContact{}
	case crc_webPageEmpty:
		r = TL_webPageEmpty{
			Id: m.Long(),
		}
	case crc_webPagePending:
		r = TL_webPagePending{
			Id:   m.Long(),
			Date: m.Int(),
		}
	case crc_webPage:
		flags := m.Int()
		id := m.Long()
		url := m.String()
		display_url := m.String()
		hash := m.Int()
		var code_type string
		if flags&(1<<0) != 0 {
			code_type = m.String()
		}
		var site_name string
		if flags&(1<<1) != 0 {
			site_name = m.String()
		}
		var title string
		if flags&(1<<2) != 0 {
			title = m.String()
		}
		var description string
		if flags&(1<<3) != 0 {
			description = m.String()
		}
		var photo TL
		if flags&(1<<4) != 0 {
			photo = m.Object()
		}
		var embed_url string
		if flags&(1<<5) != 0 {
			embed_url = m.String()
		}
		var embed_type string
		if flags&(1<<5) != 0 {
			embed_type = m.String()
		}
		var embed_width int32
		if flags&(1<<6) != 0 {
			embed_width = m.Int()
		}
		var embed_height int32
		if flags&(1<<6) != 0 {
			embed_height = m.Int()
		}
		var duration int32
		if flags&(1<<7) != 0 {
			duration = m.Int()
		}
		var author string
		if flags&(1<<8) != 0 {
			author = m.String()
		}
		var document TL
		if flags&(1<<9) != 0 {
			document = m.Object()
		}
		var cached_page TL
		if flags&(1<<10) != 0 {
			cached_page = m.Object()
		}
		r = TL_webPage{
			Flags:        flags,
			Id:           id,
			Url:          url,
			Display_url:  display_url,
			Hash:         hash,
			Code_type:    code_type,
			Site_name:    site_name,
			Title:        title,
			Description:  description,
			Photo:        photo,
			Embed_url:    embed_url,
			Embed_type:   embed_type,
			Embed_width:  embed_width,
			Embed_height: embed_height,
			Duration:     duration,
			Author:       author,
			Document:     document,
			Cached_page:  cached_page,
		}
	case crc_webPageNotModified:
		r = TL_webPageNotModified{}
	case crc_authorization:
		r = TL_authorization{
			Hash:           m.Long(),
			Flags:          m.Int(),
			Device_model:   m.String(),
			Platform:       m.String(),
			System_version: m.String(),
			Api_id:         m.Int(),
			App_name:       m.String(),
			App_version:    m.String(),
			Date_created:   m.Int(),
			Date_active:    m.Int(),
			Ip:             m.String(),
			Country:        m.String(),
			Region:         m.String(),
		}
	case crc_account_authorizations:
		r = TL_account_authorizations{
			Authorizations: m.Vector(),
		}
	case crc_account_noPassword:
		r = TL_account_noPassword{
			New_salt:                  m.StringBytes(),
			Email_unconfirmed_pattern: m.String(),
		}
	case crc_account_password:
		r = TL_account_password{
			Current_salt:              m.StringBytes(),
			New_salt:                  m.StringBytes(),
			Hint:                      m.String(),
			Has_recovery:              m.Object(),
			Email_unconfirmed_pattern: m.String(),
		}
	case crc_account_passwordSettings:
		r = TL_account_passwordSettings{
			Email: m.String(),
		}
	case crc_account_passwordInputSettings:
		flags := m.Int()
		var new_salt []byte
		if flags&(1<<0) != 0 {
			new_salt = m.StringBytes()
		}
		var new_password_hash []byte
		if flags&(1<<0) != 0 {
			new_password_hash = m.StringBytes()
		}
		var hint string
		if flags&(1<<0) != 0 {
			hint = m.String()
		}
		var email string
		if flags&(1<<1) != 0 {
			email = m.String()
		}
		r = TL_account_passwordInputSettings{
			Flags:             flags,
			New_salt:          new_salt,
			New_password_hash: new_password_hash,
			Hint:              hint,
			Email:             email,
		}
	case crc_auth_passwordRecovery:
		r = TL_auth_passwordRecovery{
			Email_pattern: m.String(),
		}
	case crc_receivedNotifyMessage:
		r = TL_receivedNotifyMessage{
			Id:    m.Int(),
			Flags: m.Int(),
		}
	case crc_chatInviteEmpty:
		r = TL_chatInviteEmpty{}
	case crc_chatInviteExported:
		r = TL_chatInviteExported{
			Link: m.String(),
		}
	case crc_chatInviteAlready:
		r = TL_chatInviteAlready{
			Chat: m.Object(),
		}
	case crc_chatInvite:
		flags := m.Int()
		channel := flags&(1<<0) != 0
		broadcast := flags&(1<<1) != 0
		public := flags&(1<<2) != 0
		megagroup := flags&(1<<3) != 0
		title := m.String()
		photo := m.Object()
		participants_count := m.Int()
		var participants []TL
		if flags&(1<<4) != 0 {
			participants = m.Vector()
		}
		r = TL_chatInvite{
			Flags:              flags,
			Channel:            channel,
			Broadcast:          broadcast,
			Public:             public,
			Megagroup:          megagroup,
			Title:              title,
			Photo:              photo,
			Participants_count: participants_count,
			Participants:       participants,
		}
	case crc_inputStickerSetEmpty:
		r = TL_inputStickerSetEmpty{}
	case crc_inputStickerSetID:
		r = TL_inputStickerSetID{
			Id:          m.Long(),
			Access_hash: m.Long(),
		}
	case crc_inputStickerSetShortName:
		r = TL_inputStickerSetShortName{
			Short_name: m.String(),
		}
	case crc_stickerSet:
		flags := m.Int()
		installed := flags&(1<<0) != 0
		archived := flags&(1<<1) != 0
		official := flags&(1<<2) != 0
		masks := flags&(1<<3) != 0
		id := m.Long()
		access_hash := m.Long()
		title := m.String()
		short_name := m.String()
		count := m.Int()
		hash := m.Int()
		r = TL_stickerSet{
			Flags:       flags,
			Installed:   installed,
			Archived:    archived,
			Official:    official,
			Masks:       masks,
			Id:          id,
			Access_hash: access_hash,
			Title:       title,
			Short_name:  short_name,
			Count:       count,
			Hash:        hash,
		}
	case crc_messages_stickerSet:
		r = TL_messages_stickerSet{
			Set:       m.Object(),
			Packs:     m.Vector(),
			Documents: m.Vector(),
		}
	case crc_botCommand:
		r = TL_botCommand{
			Command:     m.String(),
			Description: m.String(),
		}
	case crc_botInfo:
		r = TL_botInfo{
			User_id:     m.Int(),
			Description: m.String(),
			Commands:    m.Vector(),
		}
	case crc_keyboardButton:
		r = TL_keyboardButton{
			Text: m.String(),
		}
	case crc_keyboardButtonUrl:
		r = TL_keyboardButtonUrl{
			Text: m.String(),
			Url:  m.String(),
		}
	case crc_keyboardButtonCallback:
		r = TL_keyboardButtonCallback{
			Text: m.String(),
			Data: m.StringBytes(),
		}
	case crc_keyboardButtonRequestPhone:
		r = TL_keyboardButtonRequestPhone{
			Text: m.String(),
		}
	case crc_keyboardButtonRequestGeoLocation:
		r = TL_keyboardButtonRequestGeoLocation{
			Text: m.String(),
		}
	case crc_keyboardButtonSwitchInline:
		flags := m.Int()
		same_peer := flags&(1<<0) != 0
		text := m.String()
		query := m.String()
		r = TL_keyboardButtonSwitchInline{
			Flags:     flags,
			Same_peer: same_peer,
			Text:      text,
			Query:     query,
		}
	case crc_keyboardButtonGame:
		r = TL_keyboardButtonGame{
			Text: m.String(),
		}
	case crc_keyboardButtonBuy:
		r = TL_keyboardButtonBuy{
			Text: m.String(),
		}
	case crc_keyboardButtonRow:
		r = TL_keyboardButtonRow{
			Buttons: m.Vector(),
		}
	case crc_replyKeyboardHide:
		flags := m.Int()
		selective := flags&(1<<2) != 0
		r = TL_replyKeyboardHide{
			Flags:     flags,
			Selective: selective,
		}
	case crc_replyKeyboardForceReply:
		flags := m.Int()
		single_use := flags&(1<<1) != 0
		selective := flags&(1<<2) != 0
		r = TL_replyKeyboardForceReply{
			Flags:      flags,
			Single_use: single_use,
			Selective:  selective,
		}
	case crc_replyKeyboardMarkup:
		flags := m.Int()
		resize := flags&(1<<0) != 0
		single_use := flags&(1<<1) != 0
		selective := flags&(1<<2) != 0
		rows := m.Vector()
		r = TL_replyKeyboardMarkup{
			Flags:      flags,
			Resize:     resize,
			Single_use: single_use,
			Selective:  selective,
			Rows:       rows,
		}
	case crc_replyInlineMarkup:
		r = TL_replyInlineMarkup{
			Rows: m.Vector(),
		}
	case crc_messageEntityUnknown:
		r = TL_messageEntityUnknown{
			Offset: m.Int(),
			Length: m.Int(),
		}
	case crc_messageEntityMention:
		r = TL_messageEntityMention{
			Offset: m.Int(),
			Length: m.Int(),
		}
	case crc_messageEntityHashtag:
		r = TL_messageEntityHashtag{
			Offset: m.Int(),
			Length: m.Int(),
		}
	case crc_messageEntityBotCommand:
		r = TL_messageEntityBotCommand{
			Offset: m.Int(),
			Length: m.Int(),
		}
	case crc_messageEntityUrl:
		r = TL_messageEntityUrl{
			Offset: m.Int(),
			Length: m.Int(),
		}
	case crc_messageEntityEmail:
		r = TL_messageEntityEmail{
			Offset: m.Int(),
			Length: m.Int(),
		}
	case crc_messageEntityBold:
		r = TL_messageEntityBold{
			Offset: m.Int(),
			Length: m.Int(),
		}
	case crc_messageEntityItalic:
		r = TL_messageEntityItalic{
			Offset: m.Int(),
			Length: m.Int(),
		}
	case crc_messageEntityCode:
		r = TL_messageEntityCode{
			Offset: m.Int(),
			Length: m.Int(),
		}
	case crc_messageEntityPre:
		r = TL_messageEntityPre{
			Offset:   m.Int(),
			Length:   m.Int(),
			Language: m.String(),
		}
	case crc_messageEntityTextUrl:
		r = TL_messageEntityTextUrl{
			Offset: m.Int(),
			Length: m.Int(),
			Url:    m.String(),
		}
	case crc_messageEntityMentionName:
		r = TL_messageEntityMentionName{
			Offset:  m.Int(),
			Length:  m.Int(),
			User_id: m.Int(),
		}
	case crc_inputMessageEntityMentionName:
		r = TL_inputMessageEntityMentionName{
			Offset:  m.Int(),
			Length:  m.Int(),
			User_id: m.Object(),
		}
	case crc_inputChannelEmpty:
		r = TL_inputChannelEmpty{}
	case crc_inputChannel:
		r = TL_inputChannel{
			Channel_id:  m.Int(),
			Access_hash: m.Long(),
		}
	case crc_contacts_resolvedPeer:
		r = TL_contacts_resolvedPeer{
			Peer:  m.Object(),
			Chats: m.Vector(),
			Users: m.Vector(),
		}
	case crc_messageRange:
		r = TL_messageRange{
			Min_id: m.Int(),
			Max_id: m.Int(),
		}
	case crc_updates_channelDifferenceEmpty:
		flags := m.Int()
		final := flags&(1<<0) != 0
		pts := m.Int()
		var timeout int32
		if flags&(1<<1) != 0 {
			timeout = m.Int()
		}
		r = TL_updates_channelDifferenceEmpty{
			Flags:   flags,
			Final:   final,
			Pts:     pts,
			Timeout: timeout,
		}
	case crc_updates_channelDifferenceTooLong:
		flags := m.Int()
		final := flags&(1<<0) != 0
		pts := m.Int()
		var timeout int32
		if flags&(1<<1) != 0 {
			timeout = m.Int()
		}
		top_message := m.Int()
		read_inbox_max_id := m.Int()
		read_outbox_max_id := m.Int()
		unread_count := m.Int()
		messages := m.Vector()
		chats := m.Vector()
		users := m.Vector()
		r = TL_updates_channelDifferenceTooLong{
			Flags:              flags,
			Final:              final,
			Pts:                pts,
			Timeout:            timeout,
			Top_message:        top_message,
			Read_inbox_max_id:  read_inbox_max_id,
			Read_outbox_max_id: read_outbox_max_id,
			Unread_count:       unread_count,
			Messages:           messages,
			Chats:              chats,
			Users:              users,
		}
	case crc_updates_channelDifference:
		flags := m.Int()
		final := flags&(1<<0) != 0
		pts := m.Int()
		var timeout int32
		if flags&(1<<1) != 0 {
			timeout = m.Int()
		}
		new_messages := m.Vector()
		other_updates := m.Vector()
		chats := m.Vector()
		users := m.Vector()
		r = TL_updates_channelDifference{
			Flags:         flags,
			Final:         final,
			Pts:           pts,
			Timeout:       timeout,
			New_messages:  new_messages,
			Other_updates: other_updates,
			Chats:         chats,
			Users:         users,
		}
	case crc_channelMessagesFilterEmpty:
		r = TL_channelMessagesFilterEmpty{}
	case crc_channelMessagesFilter:
		flags := m.Int()
		exclude_new_messages := flags&(1<<1) != 0
		ranges := m.Vector()
		r = TL_channelMessagesFilter{
			Flags:                flags,
			Exclude_new_messages: exclude_new_messages,
			Ranges:               ranges,
		}
	case crc_channelParticipant:
		r = TL_channelParticipant{
			User_id: m.Int(),
			Date:    m.Int(),
		}
	case crc_channelParticipantSelf:
		r = TL_channelParticipantSelf{
			User_id:    m.Int(),
			Inviter_id: m.Int(),
			Date:       m.Int(),
		}
	case crc_channelParticipantModerator:
		r = TL_channelParticipantModerator{
			User_id:    m.Int(),
			Inviter_id: m.Int(),
			Date:       m.Int(),
		}
	case crc_channelParticipantEditor:
		r = TL_channelParticipantEditor{
			User_id:    m.Int(),
			Inviter_id: m.Int(),
			Date:       m.Int(),
		}
	case crc_channelParticipantKicked:
		r = TL_channelParticipantKicked{
			User_id:   m.Int(),
			Kicked_by: m.Int(),
			Date:      m.Int(),
		}
	case crc_channelParticipantCreator:
		r = TL_channelParticipantCreator{
			User_id: m.Int(),
		}
	case crc_channelParticipantsRecent:
		r = TL_channelParticipantsRecent{}
	case crc_channelParticipantsAdmins:
		r = TL_channelParticipantsAdmins{}
	case crc_channelParticipantsKicked:
		r = TL_channelParticipantsKicked{}
	case crc_channelParticipantsBots:
		r = TL_channelParticipantsBots{}
	case crc_channelRoleEmpty:
		r = TL_channelRoleEmpty{}
	case crc_channelRoleModerator:
		r = TL_channelRoleModerator{}
	case crc_channelRoleEditor:
		r = TL_channelRoleEditor{}
	case crc_channels_channelParticipants:
		r = TL_channels_channelParticipants{
			Count:        m.Int(),
			Participants: m.Vector(),
			Users:        m.Vector(),
		}
	case crc_channels_channelParticipant:
		r = TL_channels_channelParticipant{
			Participant: m.Object(),
			Users:       m.Vector(),
		}
	case crc_help_termsOfService:
		r = TL_help_termsOfService{
			Text: m.String(),
		}
	case crc_foundGif:
		r = TL_foundGif{
			Url:          m.String(),
			Thumb_url:    m.String(),
			Content_url:  m.String(),
			Content_type: m.String(),
			W:            m.Int(),
			H:            m.Int(),
		}
	case crc_foundGifCached:
		r = TL_foundGifCached{
			Url:      m.String(),
			Photo:    m.Object(),
			Document: m.Object(),
		}
	case crc_messages_foundGifs:
		r = TL_messages_foundGifs{
			Next_offset: m.Int(),
			Results:     m.Vector(),
		}
	case crc_messages_savedGifsNotModified:
		r = TL_messages_savedGifsNotModified{}
	case crc_messages_savedGifs:
		r = TL_messages_savedGifs{
			Hash: m.Int(),
			Gifs: m.Vector(),
		}
	case crc_inputBotInlineMessageMediaAuto:
		flags := m.Int()
		caption := m.String()
		var reply_markup TL
		if flags&(1<<2) != 0 {
			reply_markup = m.Object()
		}
		r = TL_inputBotInlineMessageMediaAuto{
			Flags:        flags,
			Caption:      caption,
			Reply_markup: reply_markup,
		}
	case crc_inputBotInlineMessageText:
		flags := m.Int()
		no_webpage := flags&(1<<0) != 0
		message := m.String()
		var entities []TL
		if flags&(1<<1) != 0 {
			entities = m.Vector()
		}
		var reply_markup TL
		if flags&(1<<2) != 0 {
			reply_markup = m.Object()
		}
		r = TL_inputBotInlineMessageText{
			Flags:        flags,
			No_webpage:   no_webpage,
			Message:      message,
			Entities:     entities,
			Reply_markup: reply_markup,
		}
	case crc_inputBotInlineMessageMediaGeo:
		flags := m.Int()
		geo_point := m.Object()
		var reply_markup TL
		if flags&(1<<2) != 0 {
			reply_markup = m.Object()
		}
		r = TL_inputBotInlineMessageMediaGeo{
			Flags:        flags,
			Geo_point:    geo_point,
			Reply_markup: reply_markup,
		}
	case crc_inputBotInlineMessageMediaVenue:
		flags := m.Int()
		geo_point := m.Object()
		title := m.String()
		address := m.String()
		provider := m.String()
		venue_id := m.String()
		var reply_markup TL
		if flags&(1<<2) != 0 {
			reply_markup = m.Object()
		}
		r = TL_inputBotInlineMessageMediaVenue{
			Flags:        flags,
			Geo_point:    geo_point,
			Title:        title,
			Address:      address,
			Provider:     provider,
			Venue_id:     venue_id,
			Reply_markup: reply_markup,
		}
	case crc_inputBotInlineMessageMediaContact:
		flags := m.Int()
		phone_number := m.String()
		first_name := m.String()
		last_name := m.String()
		var reply_markup TL
		if flags&(1<<2) != 0 {
			reply_markup = m.Object()
		}
		r = TL_inputBotInlineMessageMediaContact{
			Flags:        flags,
			Phone_number: phone_number,
			First_name:   first_name,
			Last_name:    last_name,
			Reply_markup: reply_markup,
		}
	case crc_inputBotInlineMessageGame:
		flags := m.Int()
		var reply_markup TL
		if flags&(1<<2) != 0 {
			reply_markup = m.Object()
		}
		r = TL_inputBotInlineMessageGame{
			Flags:        flags,
			Reply_markup: reply_markup,
		}
	case crc_inputBotInlineResult:
		flags := m.Int()
		id := m.String()
		code_type := m.String()
		var title string
		if flags&(1<<1) != 0 {
			title = m.String()
		}
		var description string
		if flags&(1<<2) != 0 {
			description = m.String()
		}
		var url string
		if flags&(1<<3) != 0 {
			url = m.String()
		}
		var thumb_url string
		if flags&(1<<4) != 0 {
			thumb_url = m.String()
		}
		var content_url string
		if flags&(1<<5) != 0 {
			content_url = m.String()
		}
		var content_type string
		if flags&(1<<5) != 0 {
			content_type = m.String()
		}
		var w int32
		if flags&(1<<6) != 0 {
			w = m.Int()
		}
		var h int32
		if flags&(1<<6) != 0 {
			h = m.Int()
		}
		var duration int32
		if flags&(1<<7) != 0 {
			duration = m.Int()
		}
		send_message := m.Object()
		r = TL_inputBotInlineResult{
			Flags:        flags,
			Id:           id,
			Code_type:    code_type,
			Title:        title,
			Description:  description,
			Url:          url,
			Thumb_url:    thumb_url,
			Content_url:  content_url,
			Content_type: content_type,
			W:            w,
			H:            h,
			Duration:     duration,
			Send_message: send_message,
		}
	case crc_inputBotInlineResultPhoto:
		r = TL_inputBotInlineResultPhoto{
			Id:           m.String(),
			Code_type:    m.String(),
			Photo:        m.Object(),
			Send_message: m.Object(),
		}
	case crc_inputBotInlineResultDocument:
		flags := m.Int()
		id := m.String()
		code_type := m.String()
		var title string
		if flags&(1<<1) != 0 {
			title = m.String()
		}
		var description string
		if flags&(1<<2) != 0 {
			description = m.String()
		}
		document := m.Object()
		send_message := m.Object()
		r = TL_inputBotInlineResultDocument{
			Flags:        flags,
			Id:           id,
			Code_type:    code_type,
			Title:        title,
			Description:  description,
			Document:     document,
			Send_message: send_message,
		}
	case crc_inputBotInlineResultGame:
		r = TL_inputBotInlineResultGame{
			Id:           m.String(),
			Short_name:   m.String(),
			Send_message: m.Object(),
		}
	case crc_botInlineMessageMediaAuto:
		flags := m.Int()
		caption := m.String()
		var reply_markup TL
		if flags&(1<<2) != 0 {
			reply_markup = m.Object()
		}
		r = TL_botInlineMessageMediaAuto{
			Flags:        flags,
			Caption:      caption,
			Reply_markup: reply_markup,
		}
	case crc_botInlineMessageText:
		flags := m.Int()
		no_webpage := flags&(1<<0) != 0
		message := m.String()
		var entities []TL
		if flags&(1<<1) != 0 {
			entities = m.Vector()
		}
		var reply_markup TL
		if flags&(1<<2) != 0 {
			reply_markup = m.Object()
		}
		r = TL_botInlineMessageText{
			Flags:        flags,
			No_webpage:   no_webpage,
			Message:      message,
			Entities:     entities,
			Reply_markup: reply_markup,
		}
	case crc_botInlineMessageMediaGeo:
		flags := m.Int()
		geo := m.Object()
		var reply_markup TL
		if flags&(1<<2) != 0 {
			reply_markup = m.Object()
		}
		r = TL_botInlineMessageMediaGeo{
			Flags:        flags,
			Geo:          geo,
			Reply_markup: reply_markup,
		}
	case crc_botInlineMessageMediaVenue:
		flags := m.Int()
		geo := m.Object()
		title := m.String()
		address := m.String()
		provider := m.String()
		venue_id := m.String()
		var reply_markup TL
		if flags&(1<<2) != 0 {
			reply_markup = m.Object()
		}
		r = TL_botInlineMessageMediaVenue{
			Flags:        flags,
			Geo:          geo,
			Title:        title,
			Address:      address,
			Provider:     provider,
			Venue_id:     venue_id,
			Reply_markup: reply_markup,
		}
	case crc_botInlineMessageMediaContact:
		flags := m.Int()
		phone_number := m.String()
		first_name := m.String()
		last_name := m.String()
		var reply_markup TL
		if flags&(1<<2) != 0 {
			reply_markup = m.Object()
		}
		r = TL_botInlineMessageMediaContact{
			Flags:        flags,
			Phone_number: phone_number,
			First_name:   first_name,
			Last_name:    last_name,
			Reply_markup: reply_markup,
		}
	case crc_botInlineResult:
		flags := m.Int()
		id := m.String()
		code_type := m.String()
		var title string
		if flags&(1<<1) != 0 {
			title = m.String()
		}
		var description string
		if flags&(1<<2) != 0 {
			description = m.String()
		}
		var url string
		if flags&(1<<3) != 0 {
			url = m.String()
		}
		var thumb_url string
		if flags&(1<<4) != 0 {
			thumb_url = m.String()
		}
		var content_url string
		if flags&(1<<5) != 0 {
			content_url = m.String()
		}
		var content_type string
		if flags&(1<<5) != 0 {
			content_type = m.String()
		}
		var w int32
		if flags&(1<<6) != 0 {
			w = m.Int()
		}
		var h int32
		if flags&(1<<6) != 0 {
			h = m.Int()
		}
		var duration int32
		if flags&(1<<7) != 0 {
			duration = m.Int()
		}
		send_message := m.Object()
		r = TL_botInlineResult{
			Flags:        flags,
			Id:           id,
			Code_type:    code_type,
			Title:        title,
			Description:  description,
			Url:          url,
			Thumb_url:    thumb_url,
			Content_url:  content_url,
			Content_type: content_type,
			W:            w,
			H:            h,
			Duration:     duration,
			Send_message: send_message,
		}
	case crc_botInlineMediaResult:
		flags := m.Int()
		id := m.String()
		code_type := m.String()
		var photo TL
		if flags&(1<<0) != 0 {
			photo = m.Object()
		}
		var document TL
		if flags&(1<<1) != 0 {
			document = m.Object()
		}
		var title string
		if flags&(1<<2) != 0 {
			title = m.String()
		}
		var description string
		if flags&(1<<3) != 0 {
			description = m.String()
		}
		send_message := m.Object()
		r = TL_botInlineMediaResult{
			Flags:        flags,
			Id:           id,
			Code_type:    code_type,
			Photo:        photo,
			Document:     document,
			Title:        title,
			Description:  description,
			Send_message: send_message,
		}
	case crc_messages_botResults:
		flags := m.Int()
		gallery := flags&(1<<0) != 0
		query_id := m.Long()
		var next_offset string
		if flags&(1<<1) != 0 {
			next_offset = m.String()
		}
		var switch_pm TL
		if flags&(1<<2) != 0 {
			switch_pm = m.Object()
		}
		results := m.Vector()
		cache_time := m.Int()
		r = TL_messages_botResults{
			Flags:       flags,
			Gallery:     gallery,
			Query_id:    query_id,
			Next_offset: next_offset,
			Switch_pm:   switch_pm,
			Results:     results,
			Cache_time:  cache_time,
		}
	case crc_exportedMessageLink:
		r = TL_exportedMessageLink{
			Link: m.String(),
		}
	case crc_messageFwdHeader:
		flags := m.Int()
		var from_id int32
		if flags&(1<<0) != 0 {
			from_id = m.Int()
		}
		date := m.Int()
		var channel_id int32
		if flags&(1<<1) != 0 {
			channel_id = m.Int()
		}
		var channel_post int32
		if flags&(1<<2) != 0 {
			channel_post = m.Int()
		}
		r = TL_messageFwdHeader{
			Flags:        flags,
			From_id:      from_id,
			Date:         date,
			Channel_id:   channel_id,
			Channel_post: channel_post,
		}
	case crc_auth_codeTypeSms:
		r = TL_auth_codeTypeSms{}
	case crc_auth_codeTypeCall:
		r = TL_auth_codeTypeCall{}
	case crc_auth_codeTypeFlashCall:
		r = TL_auth_codeTypeFlashCall{}
	case crc_auth_sentCodeTypeApp:
		r = TL_auth_sentCodeTypeApp{
			Length: m.Int(),
		}
	case crc_auth_sentCodeTypeSms:
		r = TL_auth_sentCodeTypeSms{
			Length: m.Int(),
		}
	case crc_auth_sentCodeTypeCall:
		r = TL_auth_sentCodeTypeCall{
			Length: m.Int(),
		}
	case crc_auth_sentCodeTypeFlashCall:
		r = TL_auth_sentCodeTypeFlashCall{
			Pattern: m.String(),
		}
	case crc_messages_botCallbackAnswer:
		flags := m.Int()
		alert := flags&(1<<1) != 0
		has_url := flags&(1<<3) != 0
		var message string
		if flags&(1<<0) != 0 {
			message = m.String()
		}
		var url string
		if flags&(1<<2) != 0 {
			url = m.String()
		}
		cache_time := m.Int()
		r = TL_messages_botCallbackAnswer{
			Flags:      flags,
			Alert:      alert,
			Has_url:    has_url,
			Message:    message,
			Url:        url,
			Cache_time: cache_time,
		}
	case crc_messages_messageEditData:
		flags := m.Int()
		caption := flags&(1<<0) != 0
		r = TL_messages_messageEditData{
			Flags:   flags,
			Caption: caption,
		}
	case crc_inputBotInlineMessageID:
		r = TL_inputBotInlineMessageID{
			Dc_id:       m.Int(),
			Id:          m.Long(),
			Access_hash: m.Long(),
		}
	case crc_inlineBotSwitchPM:
		r = TL_inlineBotSwitchPM{
			Text:        m.String(),
			Start_param: m.String(),
		}
	case crc_messages_peerDialogs:
		r = TL_messages_peerDialogs{
			Dialogs:  m.Vector(),
			Messages: m.Vector(),
			Chats:    m.Vector(),
			Users:    m.Vector(),
			State:    m.Object(),
		}
	case crc_topPeer:
		r = TL_topPeer{
			Peer:   m.Object(),
			Rating: m.Double(),
		}
	case crc_topPeerCategoryBotsPM:
		r = TL_topPeerCategoryBotsPM{}
	case crc_topPeerCategoryBotsInline:
		r = TL_topPeerCategoryBotsInline{}
	case crc_topPeerCategoryCorrespondents:
		r = TL_topPeerCategoryCorrespondents{}
	case crc_topPeerCategoryGroups:
		r = TL_topPeerCategoryGroups{}
	case crc_topPeerCategoryChannels:
		r = TL_topPeerCategoryChannels{}
	case crc_topPeerCategoryPeers:
		r = TL_topPeerCategoryPeers{
			Category: m.Object(),
			Count:    m.Int(),
			Peers:    m.Vector(),
		}
	case crc_contacts_topPeersNotModified:
		r = TL_contacts_topPeersNotModified{}
	case crc_contacts_topPeers:
		r = TL_contacts_topPeers{
			Categories: m.Vector(),
			Chats:      m.Vector(),
			Users:      m.Vector(),
		}
	case crc_draftMessageEmpty:
		r = TL_draftMessageEmpty{}
	case crc_draftMessage:
		flags := m.Int()
		no_webpage := flags&(1<<1) != 0
		var reply_to_msg_id int32
		if flags&(1<<0) != 0 {
			reply_to_msg_id = m.Int()
		}
		message := m.String()
		var entities []TL
		if flags&(1<<3) != 0 {
			entities = m.Vector()
		}
		date := m.Int()
		r = TL_draftMessage{
			Flags:           flags,
			No_webpage:      no_webpage,
			Reply_to_msg_id: reply_to_msg_id,
			Message:         message,
			Entities:        entities,
			Date:            date,
		}
	case crc_messages_featuredStickersNotModified:
		r = TL_messages_featuredStickersNotModified{}
	case crc_messages_featuredStickers:
		r = TL_messages_featuredStickers{
			Hash:   m.Int(),
			Sets:   m.Vector(),
			Unread: m.VectorLong(),
		}
	case crc_messages_recentStickersNotModified:
		r = TL_messages_recentStickersNotModified{}
	case crc_messages_recentStickers:
		r = TL_messages_recentStickers{
			Hash:     m.Int(),
			Stickers: m.Vector(),
		}
	case crc_messages_archivedStickers:
		r = TL_messages_archivedStickers{
			Count: m.Int(),
			Sets:  m.Vector(),
		}
	case crc_messages_stickerSetInstallResultSuccess:
		r = TL_messages_stickerSetInstallResultSuccess{}
	case crc_messages_stickerSetInstallResultArchive:
		r = TL_messages_stickerSetInstallResultArchive{
			Sets: m.Vector(),
		}
	case crc_stickerSetCovered:
		r = TL_stickerSetCovered{
			Set:   m.Object(),
			Cover: m.Object(),
		}
	case crc_stickerSetMultiCovered:
		r = TL_stickerSetMultiCovered{
			Set:    m.Object(),
			Covers: m.Vector(),
		}
	case crc_maskCoords:
		r = TL_maskCoords{
			N:    m.Int(),
			X:    m.Double(),
			Y:    m.Double(),
			Zoom: m.Double(),
		}
	case crc_inputStickeredMediaPhoto:
		r = TL_inputStickeredMediaPhoto{
			Id: m.Object(),
		}
	case crc_inputStickeredMediaDocument:
		r = TL_inputStickeredMediaDocument{
			Id: m.Object(),
		}
	case crc_game:
		flags := m.Int()
		id := m.Long()
		access_hash := m.Long()
		short_name := m.String()
		title := m.String()
		description := m.String()
		photo := m.Object()
		var document TL
		if flags&(1<<0) != 0 {
			document = m.Object()
		}
		r = TL_game{
			Flags:       flags,
			Id:          id,
			Access_hash: access_hash,
			Short_name:  short_name,
			Title:       title,
			Description: description,
			Photo:       photo,
			Document:    document,
		}
	case crc_inputGameID:
		r = TL_inputGameID{
			Id:          m.Long(),
			Access_hash: m.Long(),
		}
	case crc_inputGameShortName:
		r = TL_inputGameShortName{
			Bot_id:     m.Object(),
			Short_name: m.String(),
		}
	case crc_highScore:
		r = TL_highScore{
			Pos:     m.Int(),
			User_id: m.Int(),
			Score:   m.Int(),
		}
	case crc_messages_highScores:
		r = TL_messages_highScores{
			Scores: m.Vector(),
			Users:  m.Vector(),
		}
	case crc_textEmpty:
		r = TL_textEmpty{}
	case crc_textPlain:
		r = TL_textPlain{
			Text: m.String(),
		}
	case crc_textBold:
		r = TL_textBold{
			Text: m.Object(),
		}
	case crc_textItalic:
		r = TL_textItalic{
			Text: m.Object(),
		}
	case crc_textUnderline:
		r = TL_textUnderline{
			Text: m.Object(),
		}
	case crc_textStrike:
		r = TL_textStrike{
			Text: m.Object(),
		}
	case crc_textFixed:
		r = TL_textFixed{
			Text: m.Object(),
		}
	case crc_textUrl:
		r = TL_textUrl{
			Text:       m.Object(),
			Url:        m.String(),
			Webpage_id: m.Long(),
		}
	case crc_textEmail:
		r = TL_textEmail{
			Text:  m.Object(),
			Email: m.String(),
		}
	case crc_textConcat:
		r = TL_textConcat{
			Texts: m.Vector(),
		}
	case crc_pageBlockUnsupported:
		r = TL_pageBlockUnsupported{}
	case crc_pageBlockTitle:
		r = TL_pageBlockTitle{
			Text: m.Object(),
		}
	case crc_pageBlockSubtitle:
		r = TL_pageBlockSubtitle{
			Text: m.Object(),
		}
	case crc_pageBlockAuthorDate:
		r = TL_pageBlockAuthorDate{
			Author:         m.Object(),
			Published_date: m.Int(),
		}
	case crc_pageBlockHeader:
		r = TL_pageBlockHeader{
			Text: m.Object(),
		}
	case crc_pageBlockSubheader:
		r = TL_pageBlockSubheader{
			Text: m.Object(),
		}
	case crc_pageBlockParagraph:
		r = TL_pageBlockParagraph{
			Text: m.Object(),
		}
	case crc_pageBlockPreformatted:
		r = TL_pageBlockPreformatted{
			Text:     m.Object(),
			Language: m.String(),
		}
	case crc_pageBlockFooter:
		r = TL_pageBlockFooter{
			Text: m.Object(),
		}
	case crc_pageBlockDivider:
		r = TL_pageBlockDivider{}
	case crc_pageBlockAnchor:
		r = TL_pageBlockAnchor{
			Name: m.String(),
		}
	case crc_pageBlockList:
		r = TL_pageBlockList{
			Ordered: m.Object(),
			Items:   m.Vector(),
		}
	case crc_pageBlockBlockquote:
		r = TL_pageBlockBlockquote{
			Text:    m.Object(),
			Caption: m.Object(),
		}
	case crc_pageBlockPullquote:
		r = TL_pageBlockPullquote{
			Text:    m.Object(),
			Caption: m.Object(),
		}
	case crc_pageBlockPhoto:
		r = TL_pageBlockPhoto{
			Photo_id: m.Long(),
			Caption:  m.Object(),
		}
	case crc_pageBlockVideo:
		flags := m.Int()
		autoplay := flags&(1<<0) != 0
		loop := flags&(1<<1) != 0
		video_id := m.Long()
		caption := m.Object()
		r = TL_pageBlockVideo{
			Flags:    flags,
			Autoplay: autoplay,
			Loop:     loop,
			Video_id: video_id,
			Caption:  caption,
		}
	case crc_pageBlockCover:
		r = TL_pageBlockCover{
			Cover: m.Object(),
		}
	case crc_pageBlockEmbed:
		flags := m.Int()
		full_width := flags&(1<<0) != 0
		allow_scrolling := flags&(1<<3) != 0
		var url string
		if flags&(1<<1) != 0 {
			url = m.String()
		}
		var html string
		if flags&(1<<2) != 0 {
			html = m.String()
		}
		var poster_photo_id int64
		if flags&(1<<4) != 0 {
			poster_photo_id = m.Long()
		}
		w := m.Int()
		h := m.Int()
		caption := m.Object()
		r = TL_pageBlockEmbed{
			Flags:           flags,
			Full_width:      full_width,
			Allow_scrolling: allow_scrolling,
			Url:             url,
			Html:            html,
			Poster_photo_id: poster_photo_id,
			W:               w,
			H:               h,
			Caption:         caption,
		}
	case crc_pageBlockEmbedPost:
		r = TL_pageBlockEmbedPost{
			Url:             m.String(),
			Webpage_id:      m.Long(),
			Author_photo_id: m.Long(),
			Author:          m.String(),
			Date:            m.Int(),
			Blocks:          m.Vector(),
			Caption:         m.Object(),
		}
	case crc_pageBlockCollage:
		r = TL_pageBlockCollage{
			Items:   m.Vector(),
			Caption: m.Object(),
		}
	case crc_pageBlockSlideshow:
		r = TL_pageBlockSlideshow{
			Items:   m.Vector(),
			Caption: m.Object(),
		}
	case crc_pagePart:
		r = TL_pagePart{
			Blocks: m.Vector(),
			Photos: m.Vector(),
			Videos: m.Vector(),
		}
	case crc_pageFull:
		r = TL_pageFull{
			Blocks: m.Vector(),
			Photos: m.Vector(),
			Videos: m.Vector(),
		}
	case crc_phoneCallDiscardReasonMissed:
		r = TL_phoneCallDiscardReasonMissed{}
	case crc_phoneCallDiscardReasonDisconnect:
		r = TL_phoneCallDiscardReasonDisconnect{}
	case crc_phoneCallDiscardReasonHangup:
		r = TL_phoneCallDiscardReasonHangup{}
	case crc_phoneCallDiscardReasonBusy:
		r = TL_phoneCallDiscardReasonBusy{}
	case crc_dataJSON:
		r = TL_dataJSON{
			Data: m.String(),
		}
	case crc_labeledPrice:
		r = TL_labeledPrice{
			Label:  m.String(),
			Amount: m.Long(),
		}
	case crc_invoice:
		flags := m.Int()
		test := flags&(1<<0) != 0
		name_requested := flags&(1<<1) != 0
		phone_requested := flags&(1<<2) != 0
		email_requested := flags&(1<<3) != 0
		shipping_address_requested := flags&(1<<4) != 0
		flexible := flags&(1<<5) != 0
		currency := m.String()
		prices := m.Vector()
		r = TL_invoice{
			Flags:                      flags,
			Test:                       test,
			Name_requested:             name_requested,
			Phone_requested:            phone_requested,
			Email_requested:            email_requested,
			Shipping_address_requested: shipping_address_requested,
			Flexible:                   flexible,
			Currency:                   currency,
			Prices:                     prices,
		}
	case crc_paymentCharge:
		r = TL_paymentCharge{
			Id:                 m.String(),
			Provider_charge_id: m.String(),
		}
	case crc_postAddress:
		r = TL_postAddress{
			Street_line1: m.String(),
			Street_line2: m.String(),
			City:         m.String(),
			State:        m.String(),
			Country_iso2: m.String(),
			Post_code:    m.String(),
		}
	case crc_paymentRequestedInfo:
		flags := m.Int()
		var name string
		if flags&(1<<0) != 0 {
			name = m.String()
		}
		var phone string
		if flags&(1<<1) != 0 {
			phone = m.String()
		}
		var email string
		if flags&(1<<2) != 0 {
			email = m.String()
		}
		var shipping_address TL
		if flags&(1<<3) != 0 {
			shipping_address = m.Object()
		}
		r = TL_paymentRequestedInfo{
			Flags:            flags,
			Name:             name,
			Phone:            phone,
			Email:            email,
			Shipping_address: shipping_address,
		}
	case crc_paymentSavedCredentialsCard:
		r = TL_paymentSavedCredentialsCard{
			Id:    m.String(),
			Title: m.String(),
		}
	case crc_webDocument:
		r = TL_webDocument{
			Url:         m.String(),
			Access_hash: m.Long(),
			Size:        m.Int(),
			Mime_type:   m.String(),
			Attributes:  m.Vector(),
			Dc_id:       m.Int(),
		}
	case crc_inputWebDocument:
		r = TL_inputWebDocument{
			Url:        m.String(),
			Size:       m.Int(),
			Mime_type:  m.String(),
			Attributes: m.Vector(),
		}
	case crc_inputWebFileLocation:
		r = TL_inputWebFileLocation{
			Url:         m.String(),
			Access_hash: m.Long(),
		}
	case crc_upload_webFile:
		r = TL_upload_webFile{
			Size:      m.Int(),
			Mime_type: m.String(),
			File_type: m.Object(),
			Mtime:     m.Int(),
			Bytes:     m.StringBytes(),
		}
	case crc_payments_paymentForm:
		flags := m.Int()
		can_save_credentials := flags&(1<<2) != 0
		password_missing := flags&(1<<3) != 0
		bot_id := m.Int()
		invoice := m.Object()
		provider_id := m.Int()
		url := m.String()
		var native_provider string
		if flags&(1<<4) != 0 {
			native_provider = m.String()
		}
		var native_params TL
		if flags&(1<<4) != 0 {
			native_params = m.Object()
		}
		var saved_info TL
		if flags&(1<<0) != 0 {
			saved_info = m.Object()
		}
		var saved_credentials TL
		if flags&(1<<1) != 0 {
			saved_credentials = m.Object()
		}
		users := m.Vector()
		r = TL_payments_paymentForm{
			Flags:                flags,
			Can_save_credentials: can_save_credentials,
			Password_missing:     password_missing,
			Bot_id:               bot_id,
			Invoice:              invoice,
			Provider_id:          provider_id,
			Url:                  url,
			Native_provider:      native_provider,
			Native_params:        native_params,
			Saved_info:           saved_info,
			Saved_credentials:    saved_credentials,
			Users:                users,
		}
	case crc_payments_validatedRequestedInfo:
		flags := m.Int()
		var id string
		if flags&(1<<0) != 0 {
			id = m.String()
		}
		var shipping_options []TL
		if flags&(1<<1) != 0 {
			shipping_options = m.Vector()
		}
		r = TL_payments_validatedRequestedInfo{
			Flags:            flags,
			Id:               id,
			Shipping_options: shipping_options,
		}
	case crc_payments_paymentResult:
		r = TL_payments_paymentResult{
			Updates: m.Object(),
		}
	case crc_payments_paymentVerficationNeeded:
		r = TL_payments_paymentVerficationNeeded{
			Url: m.String(),
		}
	case crc_payments_paymentReceipt:
		flags := m.Int()
		date := m.Int()
		bot_id := m.Int()
		invoice := m.Object()
		provider_id := m.Int()
		var info TL
		if flags&(1<<0) != 0 {
			info = m.Object()
		}
		var shipping TL
		if flags&(1<<1) != 0 {
			shipping = m.Object()
		}
		currency := m.String()
		total_amount := m.Long()
		credentials_title := m.String()
		users := m.Vector()
		r = TL_payments_paymentReceipt{
			Flags:             flags,
			Date:              date,
			Bot_id:            bot_id,
			Invoice:           invoice,
			Provider_id:       provider_id,
			Info:              info,
			Shipping:          shipping,
			Currency:          currency,
			Total_amount:      total_amount,
			Credentials_title: credentials_title,
			Users:             users,
		}
	case crc_payments_savedInfo:
		flags := m.Int()
		has_saved_credentials := flags&(1<<1) != 0
		var saved_info TL
		if flags&(1<<0) != 0 {
			saved_info = m.Object()
		}
		r = TL_payments_savedInfo{
			Flags: flags,
			Has_saved_credentials: has_saved_credentials,
			Saved_info:            saved_info,
		}
	case crc_inputPaymentCredentialsSaved:
		r = TL_inputPaymentCredentialsSaved{
			Id:           m.String(),
			Tmp_password: m.StringBytes(),
		}
	case crc_inputPaymentCredentials:
		flags := m.Int()
		save := flags&(1<<0) != 0
		data := m.Object()
		r = TL_inputPaymentCredentials{
			Flags: flags,
			Save:  save,
			Data:  data,
		}
	case crc_account_tmpPassword:
		r = TL_account_tmpPassword{
			Tmp_password: m.StringBytes(),
			Valid_until:  m.Int(),
		}
	case crc_shippingOption:
		r = TL_shippingOption{
			Id:     m.String(),
			Title:  m.String(),
			Prices: m.Vector(),
		}
	case crc_inputPhoneCall:
		r = TL_inputPhoneCall{
			Id:          m.Long(),
			Access_hash: m.Long(),
		}
	case crc_phoneCallEmpty:
		r = TL_phoneCallEmpty{
			Id: m.Long(),
		}
	case crc_phoneCallWaiting:
		flags := m.Int()
		id := m.Long()
		access_hash := m.Long()
		date := m.Int()
		admin_id := m.Int()
		participant_id := m.Int()
		protocol := m.Object()
		var receive_date int32
		if flags&(1<<0) != 0 {
			receive_date = m.Int()
		}
		r = TL_phoneCallWaiting{
			Flags:          flags,
			Id:             id,
			Access_hash:    access_hash,
			Date:           date,
			Admin_id:       admin_id,
			Participant_id: participant_id,
			Protocol:       protocol,
			Receive_date:   receive_date,
		}
	case crc_phoneCallRequested:
		r = TL_phoneCallRequested{
			Id:             m.Long(),
			Access_hash:    m.Long(),
			Date:           m.Int(),
			Admin_id:       m.Int(),
			Participant_id: m.Int(),
			G_a_hash:       m.StringBytes(),
			Protocol:       m.Object(),
		}
	case crc_phoneCallAccepted:
		r = TL_phoneCallAccepted{
			Id:             m.Long(),
			Access_hash:    m.Long(),
			Date:           m.Int(),
			Admin_id:       m.Int(),
			Participant_id: m.Int(),
			G_b:            m.StringBytes(),
			Protocol:       m.Object(),
		}
	case crc_phoneCall:
		r = TL_phoneCall{
			Id:                      m.Long(),
			Access_hash:             m.Long(),
			Date:                    m.Int(),
			Admin_id:                m.Int(),
			Participant_id:          m.Int(),
			G_a_or_b:                m.StringBytes(),
			Key_fingerprint:         m.Long(),
			Protocol:                m.Object(),
			Connection:              m.Object(),
			Alternative_connections: m.Vector(),
			Start_date:              m.Int(),
		}
	case crc_phoneCallDiscarded:
		flags := m.Int()
		need_rating := flags&(1<<2) != 0
		need_debug := flags&(1<<3) != 0
		id := m.Long()
		var reason TL
		if flags&(1<<0) != 0 {
			reason = m.Object()
		}
		var duration int32
		if flags&(1<<1) != 0 {
			duration = m.Int()
		}
		r = TL_phoneCallDiscarded{
			Flags:       flags,
			Need_rating: need_rating,
			Need_debug:  need_debug,
			Id:          id,
			Reason:      reason,
			Duration:    duration,
		}
	case crc_phoneConnection:
		r = TL_phoneConnection{
			Id:       m.Long(),
			Ip:       m.String(),
			Ipv6:     m.String(),
			Port:     m.Int(),
			Peer_tag: m.StringBytes(),
		}
	case crc_phoneCallProtocol:
		flags := m.Int()
		udp_p2p := flags&(1<<0) != 0
		udp_reflector := flags&(1<<1) != 0
		min_layer := m.Int()
		max_layer := m.Int()
		r = TL_phoneCallProtocol{
			Flags:         flags,
			Udp_p2p:       udp_p2p,
			Udp_reflector: udp_reflector,
			Min_layer:     min_layer,
			Max_layer:     max_layer,
		}
	case crc_phone_phoneCall:
		r = TL_phone_phoneCall{
			Phone_call: m.Object(),
			Users:      m.Vector(),
		}
	case crc_invokeAfterMsg:
		r = TL_invokeAfterMsg{
			Msg_id: m.Long(),
			Query:  m.Object(),
		}
	case crc_invokeAfterMsgs:
		r = TL_invokeAfterMsgs{
			Msg_ids: m.VectorLong(),
			Query:   m.Object(),
		}
	case crc_initConnection:
		r = TL_initConnection{
			Api_id:         m.Int(),
			Device_model:   m.String(),
			System_version: m.String(),
			App_version:    m.String(),
			Lang_code:      m.String(),
			Query:          m.Object(),
		}
	case crc_invokeWithLayer:
		r = TL_invokeWithLayer{
			Layer: m.Int(),
			Query: m.Object(),
		}
	case crc_invokeWithoutUpdates:
		r = TL_invokeWithoutUpdates{
			Query: m.Object(),
		}
	case crc_auth_checkPhone:
		r = TL_auth_checkPhone{
			Phone_number: m.String(),
		}
	case crc_auth_sendCode:
		flags := m.Int()
		allow_flashcall := flags&(1<<0) != 0
		phone_number := m.String()
		var current_number TL
		if flags&(1<<0) != 0 {
			current_number = m.Object()
		}
		api_id := m.Int()
		api_hash := m.String()
		r = TL_auth_sendCode{
			Flags:           flags,
			Allow_flashcall: allow_flashcall,
			Phone_number:    phone_number,
			Current_number:  current_number,
			Api_id:          api_id,
			Api_hash:        api_hash,
		}
	case crc_auth_signUp:
		r = TL_auth_signUp{
			Phone_number:    m.String(),
			Phone_code_hash: m.String(),
			Phone_code:      m.String(),
			First_name:      m.String(),
			Last_name:       m.String(),
		}
	case crc_auth_signIn:
		r = TL_auth_signIn{
			Phone_number:    m.String(),
			Phone_code_hash: m.String(),
			Phone_code:      m.String(),
		}
	case crc_auth_logOut:
		r = TL_auth_logOut{}
	case crc_auth_resetAuthorizations:
		r = TL_auth_resetAuthorizations{}
	case crc_auth_sendInvites:
		r = TL_auth_sendInvites{
			Phone_numbers: m.VectorString(),
			Message:       m.String(),
		}
	case crc_auth_exportAuthorization:
		r = TL_auth_exportAuthorization{
			Dc_id: m.Int(),
		}
	case crc_auth_importAuthorization:
		r = TL_auth_importAuthorization{
			Id:    m.Int(),
			Bytes: m.StringBytes(),
		}
	case crc_auth_bindTempAuthKey:
		r = TL_auth_bindTempAuthKey{
			Perm_auth_key_id:  m.Long(),
			Nonce:             m.Long(),
			Expires_at:        m.Int(),
			Encrypted_message: m.StringBytes(),
		}
	case crc_auth_importBotAuthorization:
		r = TL_auth_importBotAuthorization{
			Flags:          m.Int(),
			Api_id:         m.Int(),
			Api_hash:       m.String(),
			Bot_auth_token: m.String(),
		}
	case crc_auth_checkPassword:
		r = TL_auth_checkPassword{
			Password_hash: m.StringBytes(),
		}
	case crc_auth_requestPasswordRecovery:
		r = TL_auth_requestPasswordRecovery{}
	case crc_auth_recoverPassword:
		r = TL_auth_recoverPassword{
			Code: m.String(),
		}
	case crc_auth_resendCode:
		r = TL_auth_resendCode{
			Phone_number:    m.String(),
			Phone_code_hash: m.String(),
		}
	case crc_auth_cancelCode:
		r = TL_auth_cancelCode{
			Phone_number:    m.String(),
			Phone_code_hash: m.String(),
		}
	case crc_auth_dropTempAuthKeys:
		r = TL_auth_dropTempAuthKeys{
			Except_auth_keys: m.VectorLong(),
		}
	case crc_account_registerDevice:
		r = TL_account_registerDevice{
			Token_type: m.Int(),
			Token:      m.String(),
		}
	case crc_account_unregisterDevice:
		r = TL_account_unregisterDevice{
			Token_type: m.Int(),
			Token:      m.String(),
		}
	case crc_account_updateNotifySettings:
		r = TL_account_updateNotifySettings{
			Peer:     m.Object(),
			Settings: m.Object(),
		}
	case crc_account_getNotifySettings:
		r = TL_account_getNotifySettings{
			Peer: m.Object(),
		}
	case crc_account_resetNotifySettings:
		r = TL_account_resetNotifySettings{}
	case crc_account_updateProfile:
		flags := m.Int()
		var first_name string
		if flags&(1<<0) != 0 {
			first_name = m.String()
		}
		var last_name string
		if flags&(1<<1) != 0 {
			last_name = m.String()
		}
		var about string
		if flags&(1<<2) != 0 {
			about = m.String()
		}
		r = TL_account_updateProfile{
			Flags:      flags,
			First_name: first_name,
			Last_name:  last_name,
			About:      about,
		}
	case crc_account_updateStatus:
		r = TL_account_updateStatus{
			Offline: m.Object(),
		}
	case crc_account_getWallPapers:
		r = TL_account_getWallPapers{}
	case crc_account_reportPeer:
		r = TL_account_reportPeer{
			Peer:   m.Object(),
			Reason: m.Object(),
		}
	case crc_account_checkUsername:
		r = TL_account_checkUsername{
			Username: m.String(),
		}
	case crc_account_updateUsername:
		r = TL_account_updateUsername{
			Username: m.String(),
		}
	case crc_account_getPrivacy:
		r = TL_account_getPrivacy{
			Key: m.Object(),
		}
	case crc_account_setPrivacy:
		r = TL_account_setPrivacy{
			Key:   m.Object(),
			Rules: m.Vector(),
		}
	case crc_account_deleteAccount:
		r = TL_account_deleteAccount{
			Reason: m.String(),
		}
	case crc_account_getAccountTTL:
		r = TL_account_getAccountTTL{}
	case crc_account_setAccountTTL:
		r = TL_account_setAccountTTL{
			Ttl: m.Object(),
		}
	case crc_account_sendChangePhoneCode:
		flags := m.Int()
		allow_flashcall := flags&(1<<0) != 0
		phone_number := m.String()
		var current_number TL
		if flags&(1<<0) != 0 {
			current_number = m.Object()
		}
		r = TL_account_sendChangePhoneCode{
			Flags:           flags,
			Allow_flashcall: allow_flashcall,
			Phone_number:    phone_number,
			Current_number:  current_number,
		}
	case crc_account_changePhone:
		r = TL_account_changePhone{
			Phone_number:    m.String(),
			Phone_code_hash: m.String(),
			Phone_code:      m.String(),
		}
	case crc_account_updateDeviceLocked:
		r = TL_account_updateDeviceLocked{
			Period: m.Int(),
		}
	case crc_account_getAuthorizations:
		r = TL_account_getAuthorizations{}
	case crc_account_resetAuthorization:
		r = TL_account_resetAuthorization{
			Hash: m.Long(),
		}
	case crc_account_getPassword:
		r = TL_account_getPassword{}
	case crc_account_getPasswordSettings:
		r = TL_account_getPasswordSettings{
			Current_password_hash: m.StringBytes(),
		}
	case crc_account_updatePasswordSettings:
		r = TL_account_updatePasswordSettings{
			Current_password_hash: m.StringBytes(),
			New_settings:          m.Object(),
		}
	case crc_account_sendConfirmPhoneCode:
		flags := m.Int()
		allow_flashcall := flags&(1<<0) != 0
		hash := m.String()
		var current_number TL
		if flags&(1<<0) != 0 {
			current_number = m.Object()
		}
		r = TL_account_sendConfirmPhoneCode{
			Flags:           flags,
			Allow_flashcall: allow_flashcall,
			Hash:            hash,
			Current_number:  current_number,
		}
	case crc_account_confirmPhone:
		r = TL_account_confirmPhone{
			Phone_code_hash: m.String(),
			Phone_code:      m.String(),
		}
	case crc_account_getTmpPassword:
		r = TL_account_getTmpPassword{
			Password_hash: m.StringBytes(),
			Period:        m.Int(),
		}
	case crc_users_getUsers:
		r = TL_users_getUsers{
			Id: m.Vector(),
		}
	case crc_users_getFullUser:
		r = TL_users_getFullUser{
			Id: m.Object(),
		}
	case crc_contacts_getStatuses:
		r = TL_contacts_getStatuses{}
	case crc_contacts_getContacts:
		r = TL_contacts_getContacts{
			Hash: m.String(),
		}
	case crc_contacts_importContacts:
		r = TL_contacts_importContacts{
			Contacts: m.Vector(),
			Replace:  m.Object(),
		}
	case crc_contacts_deleteContact:
		r = TL_contacts_deleteContact{
			Id: m.Object(),
		}
	case crc_contacts_deleteContacts:
		r = TL_contacts_deleteContacts{
			Id: m.Vector(),
		}
	case crc_contacts_block:
		r = TL_contacts_block{
			Id: m.Object(),
		}
	case crc_contacts_unblock:
		r = TL_contacts_unblock{
			Id: m.Object(),
		}
	case crc_contacts_getBlocked:
		r = TL_contacts_getBlocked{
			Offset: m.Int(),
			Limit:  m.Int(),
		}
	case crc_contacts_exportCard:
		r = TL_contacts_exportCard{}
	case crc_contacts_importCard:
		r = TL_contacts_importCard{
			Export_card: m.VectorInt(),
		}
	case crc_contacts_search:
		r = TL_contacts_search{
			Q:     m.String(),
			Limit: m.Int(),
		}
	case crc_contacts_resolveUsername:
		r = TL_contacts_resolveUsername{
			Username: m.String(),
		}
	case crc_contacts_getTopPeers:
		flags := m.Int()
		correspondents := flags&(1<<0) != 0
		bots_pm := flags&(1<<1) != 0
		bots_inline := flags&(1<<2) != 0
		groups := flags&(1<<10) != 0
		channels := flags&(1<<15) != 0
		offset := m.Int()
		limit := m.Int()
		hash := m.Int()
		r = TL_contacts_getTopPeers{
			Flags:          flags,
			Correspondents: correspondents,
			Bots_pm:        bots_pm,
			Bots_inline:    bots_inline,
			Groups:         groups,
			Channels:       channels,
			Offset:         offset,
			Limit:          limit,
			Hash:           hash,
		}
	case crc_contacts_resetTopPeerRating:
		r = TL_contacts_resetTopPeerRating{
			Category: m.Object(),
			Peer:     m.Object(),
		}
	case crc_messages_getMessages:
		r = TL_messages_getMessages{
			Id: m.VectorInt(),
		}
	case crc_messages_getDialogs:
		flags := m.Int()
		exclude_pinned := flags&(1<<0) != 0
		offset_date := m.Int()
		offset_id := m.Int()
		offset_peer := m.Object()
		limit := m.Int()
		r = TL_messages_getDialogs{
			Flags:          flags,
			Exclude_pinned: exclude_pinned,
			Offset_date:    offset_date,
			Offset_id:      offset_id,
			Offset_peer:    offset_peer,
			Limit:          limit,
		}
	case crc_messages_getHistory:
		r = TL_messages_getHistory{
			Peer:        m.Object(),
			Offset_id:   m.Int(),
			Offset_date: m.Int(),
			Add_offset:  m.Int(),
			Limit:       m.Int(),
			Max_id:      m.Int(),
			Min_id:      m.Int(),
		}
	case crc_messages_search:
		flags := m.Int()
		peer := m.Object()
		q := m.String()
		filter := m.Object()
		min_date := m.Int()
		max_date := m.Int()
		offset := m.Int()
		max_id := m.Int()
		limit := m.Int()
		r = TL_messages_search{
			Flags:    flags,
			Peer:     peer,
			Q:        q,
			Filter:   filter,
			Min_date: min_date,
			Max_date: max_date,
			Offset:   offset,
			Max_id:   max_id,
			Limit:    limit,
		}
	case crc_messages_readHistory:
		r = TL_messages_readHistory{
			Peer:   m.Object(),
			Max_id: m.Int(),
		}
	case crc_messages_deleteHistory:
		flags := m.Int()
		just_clear := flags&(1<<0) != 0
		peer := m.Object()
		max_id := m.Int()
		r = TL_messages_deleteHistory{
			Flags:      flags,
			Just_clear: just_clear,
			Peer:       peer,
			Max_id:     max_id,
		}
	case crc_messages_deleteMessages:
		flags := m.Int()
		revoke := flags&(1<<0) != 0
		id := m.VectorInt()
		r = TL_messages_deleteMessages{
			Flags:  flags,
			Revoke: revoke,
			Id:     id,
		}
	case crc_messages_receivedMessages:
		r = TL_messages_receivedMessages{
			Max_id: m.Int(),
		}
	case crc_messages_setTyping:
		r = TL_messages_setTyping{
			Peer:   m.Object(),
			Action: m.Object(),
		}
	case crc_messages_sendMessage:
		flags := m.Int()
		no_webpage := flags&(1<<1) != 0
		silent := flags&(1<<5) != 0
		background := flags&(1<<6) != 0
		clear_draft := flags&(1<<7) != 0
		peer := m.Object()
		var reply_to_msg_id int32
		if flags&(1<<0) != 0 {
			reply_to_msg_id = m.Int()
		}
		message := m.String()
		random_id := m.Long()
		var reply_markup TL
		if flags&(1<<2) != 0 {
			reply_markup = m.Object()
		}
		var entities []TL
		if flags&(1<<3) != 0 {
			entities = m.Vector()
		}
		r = TL_messages_sendMessage{
			Flags:           flags,
			No_webpage:      no_webpage,
			Silent:          silent,
			Background:      background,
			Clear_draft:     clear_draft,
			Peer:            peer,
			Reply_to_msg_id: reply_to_msg_id,
			Message:         message,
			Random_id:       random_id,
			Reply_markup:    reply_markup,
			Entities:        entities,
		}
	case crc_messages_sendMedia:
		flags := m.Int()
		silent := flags&(1<<5) != 0
		background := flags&(1<<6) != 0
		clear_draft := flags&(1<<7) != 0
		peer := m.Object()
		var reply_to_msg_id int32
		if flags&(1<<0) != 0 {
			reply_to_msg_id = m.Int()
		}
		media := m.Object()
		random_id := m.Long()
		var reply_markup TL
		if flags&(1<<2) != 0 {
			reply_markup = m.Object()
		}
		r = TL_messages_sendMedia{
			Flags:           flags,
			Silent:          silent,
			Background:      background,
			Clear_draft:     clear_draft,
			Peer:            peer,
			Reply_to_msg_id: reply_to_msg_id,
			Media:           media,
			Random_id:       random_id,
			Reply_markup:    reply_markup,
		}
	case crc_messages_forwardMessages:
		flags := m.Int()
		silent := flags&(1<<5) != 0
		background := flags&(1<<6) != 0
		with_my_score := flags&(1<<8) != 0
		from_peer := m.Object()
		id := m.VectorInt()
		random_id := m.VectorLong()
		to_peer := m.Object()
		r = TL_messages_forwardMessages{
			Flags:         flags,
			Silent:        silent,
			Background:    background,
			With_my_score: with_my_score,
			From_peer:     from_peer,
			Id:            id,
			Random_id:     random_id,
			To_peer:       to_peer,
		}
	case crc_messages_reportSpam:
		r = TL_messages_reportSpam{
			Peer: m.Object(),
		}
	case crc_messages_hideReportSpam:
		r = TL_messages_hideReportSpam{
			Peer: m.Object(),
		}
	case crc_messages_getPeerSettings:
		r = TL_messages_getPeerSettings{
			Peer: m.Object(),
		}
	case crc_messages_getChats:
		r = TL_messages_getChats{
			Id: m.VectorInt(),
		}
	case crc_messages_getFullChat:
		r = TL_messages_getFullChat{
			Chat_id: m.Int(),
		}
	case crc_messages_editChatTitle:
		r = TL_messages_editChatTitle{
			Chat_id: m.Int(),
			Title:   m.String(),
		}
	case crc_messages_editChatPhoto:
		r = TL_messages_editChatPhoto{
			Chat_id: m.Int(),
			Photo:   m.Object(),
		}
	case crc_messages_addChatUser:
		r = TL_messages_addChatUser{
			Chat_id:   m.Int(),
			User_id:   m.Object(),
			Fwd_limit: m.Int(),
		}
	case crc_messages_deleteChatUser:
		r = TL_messages_deleteChatUser{
			Chat_id: m.Int(),
			User_id: m.Object(),
		}
	case crc_messages_createChat:
		r = TL_messages_createChat{
			Users: m.Vector(),
			Title: m.String(),
		}
	case crc_messages_forwardMessage:
		r = TL_messages_forwardMessage{
			Peer:      m.Object(),
			Id:        m.Int(),
			Random_id: m.Long(),
		}
	case crc_messages_getDhConfig:
		r = TL_messages_getDhConfig{
			Version:       m.Int(),
			Random_length: m.Int(),
		}
	case crc_messages_requestEncryption:
		r = TL_messages_requestEncryption{
			User_id:   m.Object(),
			Random_id: m.Int(),
			G_a:       m.StringBytes(),
		}
	case crc_messages_acceptEncryption:
		r = TL_messages_acceptEncryption{
			Peer:            m.Object(),
			G_b:             m.StringBytes(),
			Key_fingerprint: m.Long(),
		}
	case crc_messages_discardEncryption:
		r = TL_messages_discardEncryption{
			Chat_id: m.Int(),
		}
	case crc_messages_setEncryptedTyping:
		r = TL_messages_setEncryptedTyping{
			Peer:   m.Object(),
			Typing: m.Object(),
		}
	case crc_messages_readEncryptedHistory:
		r = TL_messages_readEncryptedHistory{
			Peer:     m.Object(),
			Max_date: m.Int(),
		}
	case crc_messages_sendEncrypted:
		r = TL_messages_sendEncrypted{
			Peer:      m.Object(),
			Random_id: m.Long(),
			Data:      m.StringBytes(),
		}
	case crc_messages_sendEncryptedFile:
		r = TL_messages_sendEncryptedFile{
			Peer:      m.Object(),
			Random_id: m.Long(),
			Data:      m.StringBytes(),
			File:      m.Object(),
		}
	case crc_messages_sendEncryptedService:
		r = TL_messages_sendEncryptedService{
			Peer:      m.Object(),
			Random_id: m.Long(),
			Data:      m.StringBytes(),
		}
	case crc_messages_receivedQueue:
		r = TL_messages_receivedQueue{
			Max_qts: m.Int(),
		}
	case crc_messages_reportEncryptedSpam:
		r = TL_messages_reportEncryptedSpam{
			Peer: m.Object(),
		}
	case crc_messages_readMessageContents:
		r = TL_messages_readMessageContents{
			Id: m.VectorInt(),
		}
	case crc_messages_getAllStickers:
		r = TL_messages_getAllStickers{
			Hash: m.Int(),
		}
	case crc_messages_getWebPagePreview:
		r = TL_messages_getWebPagePreview{
			Message: m.String(),
		}
	case crc_messages_exportChatInvite:
		r = TL_messages_exportChatInvite{
			Chat_id: m.Int(),
		}
	case crc_messages_checkChatInvite:
		r = TL_messages_checkChatInvite{
			Hash: m.String(),
		}
	case crc_messages_importChatInvite:
		r = TL_messages_importChatInvite{
			Hash: m.String(),
		}
	case crc_messages_getStickerSet:
		r = TL_messages_getStickerSet{
			Stickerset: m.Object(),
		}
	case crc_messages_installStickerSet:
		r = TL_messages_installStickerSet{
			Stickerset: m.Object(),
			Archived:   m.Object(),
		}
	case crc_messages_uninstallStickerSet:
		r = TL_messages_uninstallStickerSet{
			Stickerset: m.Object(),
		}
	case crc_messages_startBot:
		r = TL_messages_startBot{
			Bot:         m.Object(),
			Peer:        m.Object(),
			Random_id:   m.Long(),
			Start_param: m.String(),
		}
	case crc_messages_getMessagesViews:
		r = TL_messages_getMessagesViews{
			Peer:      m.Object(),
			Id:        m.VectorInt(),
			Increment: m.Object(),
		}
	case crc_messages_toggleChatAdmins:
		r = TL_messages_toggleChatAdmins{
			Chat_id: m.Int(),
			Enabled: m.Object(),
		}
	case crc_messages_editChatAdmin:
		r = TL_messages_editChatAdmin{
			Chat_id:  m.Int(),
			User_id:  m.Object(),
			Is_admin: m.Object(),
		}
	case crc_messages_migrateChat:
		r = TL_messages_migrateChat{
			Chat_id: m.Int(),
		}
	case crc_messages_searchGlobal:
		r = TL_messages_searchGlobal{
			Q:           m.String(),
			Offset_date: m.Int(),
			Offset_peer: m.Object(),
			Offset_id:   m.Int(),
			Limit:       m.Int(),
		}
	case crc_messages_reorderStickerSets:
		flags := m.Int()
		masks := flags&(1<<0) != 0
		order := m.VectorLong()
		r = TL_messages_reorderStickerSets{
			Flags: flags,
			Masks: masks,
			Order: order,
		}
	case crc_messages_getDocumentByHash:
		r = TL_messages_getDocumentByHash{
			Sha256:    m.StringBytes(),
			Size:      m.Int(),
			Mime_type: m.String(),
		}
	case crc_messages_searchGifs:
		r = TL_messages_searchGifs{
			Q:      m.String(),
			Offset: m.Int(),
		}
	case crc_messages_getSavedGifs:
		r = TL_messages_getSavedGifs{
			Hash: m.Int(),
		}
	case crc_messages_saveGif:
		r = TL_messages_saveGif{
			Id:     m.Object(),
			Unsave: m.Object(),
		}
	case crc_messages_getInlineBotResults:
		flags := m.Int()
		bot := m.Object()
		peer := m.Object()
		var geo_point TL
		if flags&(1<<0) != 0 {
			geo_point = m.Object()
		}
		query := m.String()
		offset := m.String()
		r = TL_messages_getInlineBotResults{
			Flags:     flags,
			Bot:       bot,
			Peer:      peer,
			Geo_point: geo_point,
			Query:     query,
			Offset:    offset,
		}
	case crc_messages_setInlineBotResults:
		flags := m.Int()
		gallery := flags&(1<<0) != 0
		private := flags&(1<<1) != 0
		query_id := m.Long()
		results := m.Vector()
		cache_time := m.Int()
		var next_offset string
		if flags&(1<<2) != 0 {
			next_offset = m.String()
		}
		var switch_pm TL
		if flags&(1<<3) != 0 {
			switch_pm = m.Object()
		}
		r = TL_messages_setInlineBotResults{
			Flags:       flags,
			Gallery:     gallery,
			Private:     private,
			Query_id:    query_id,
			Results:     results,
			Cache_time:  cache_time,
			Next_offset: next_offset,
			Switch_pm:   switch_pm,
		}
	case crc_messages_sendInlineBotResult:
		flags := m.Int()
		silent := flags&(1<<5) != 0
		background := flags&(1<<6) != 0
		clear_draft := flags&(1<<7) != 0
		peer := m.Object()
		var reply_to_msg_id int32
		if flags&(1<<0) != 0 {
			reply_to_msg_id = m.Int()
		}
		random_id := m.Long()
		query_id := m.Long()
		id := m.String()
		r = TL_messages_sendInlineBotResult{
			Flags:           flags,
			Silent:          silent,
			Background:      background,
			Clear_draft:     clear_draft,
			Peer:            peer,
			Reply_to_msg_id: reply_to_msg_id,
			Random_id:       random_id,
			Query_id:        query_id,
			Id:              id,
		}
	case crc_messages_getMessageEditData:
		r = TL_messages_getMessageEditData{
			Peer: m.Object(),
			Id:   m.Int(),
		}
	case crc_messages_editMessage:
		flags := m.Int()
		no_webpage := flags&(1<<1) != 0
		peer := m.Object()
		id := m.Int()
		var message string
		if flags&(1<<11) != 0 {
			message = m.String()
		}
		var reply_markup TL
		if flags&(1<<2) != 0 {
			reply_markup = m.Object()
		}
		var entities []TL
		if flags&(1<<3) != 0 {
			entities = m.Vector()
		}
		r = TL_messages_editMessage{
			Flags:        flags,
			No_webpage:   no_webpage,
			Peer:         peer,
			Id:           id,
			Message:      message,
			Reply_markup: reply_markup,
			Entities:     entities,
		}
	case crc_messages_editInlineBotMessage:
		flags := m.Int()
		no_webpage := flags&(1<<1) != 0
		id := m.Object()
		var message string
		if flags&(1<<11) != 0 {
			message = m.String()
		}
		var reply_markup TL
		if flags&(1<<2) != 0 {
			reply_markup = m.Object()
		}
		var entities []TL
		if flags&(1<<3) != 0 {
			entities = m.Vector()
		}
		r = TL_messages_editInlineBotMessage{
			Flags:        flags,
			No_webpage:   no_webpage,
			Id:           id,
			Message:      message,
			Reply_markup: reply_markup,
			Entities:     entities,
		}
	case crc_messages_getBotCallbackAnswer:
		flags := m.Int()
		game := flags&(1<<1) != 0
		peer := m.Object()
		msg_id := m.Int()
		var data []byte
		if flags&(1<<0) != 0 {
			data = m.StringBytes()
		}
		r = TL_messages_getBotCallbackAnswer{
			Flags:  flags,
			Game:   game,
			Peer:   peer,
			Msg_id: msg_id,
			Data:   data,
		}
	case crc_messages_setBotCallbackAnswer:
		flags := m.Int()
		alert := flags&(1<<1) != 0
		query_id := m.Long()
		var message string
		if flags&(1<<0) != 0 {
			message = m.String()
		}
		var url string
		if flags&(1<<2) != 0 {
			url = m.String()
		}
		cache_time := m.Int()
		r = TL_messages_setBotCallbackAnswer{
			Flags:      flags,
			Alert:      alert,
			Query_id:   query_id,
			Message:    message,
			Url:        url,
			Cache_time: cache_time,
		}
	case crc_messages_getPeerDialogs:
		r = TL_messages_getPeerDialogs{
			Peers: m.Vector(),
		}
	case crc_messages_saveDraft:
		flags := m.Int()
		no_webpage := flags&(1<<1) != 0
		var reply_to_msg_id int32
		if flags&(1<<0) != 0 {
			reply_to_msg_id = m.Int()
		}
		peer := m.Object()
		message := m.String()
		var entities []TL
		if flags&(1<<3) != 0 {
			entities = m.Vector()
		}
		r = TL_messages_saveDraft{
			Flags:           flags,
			No_webpage:      no_webpage,
			Reply_to_msg_id: reply_to_msg_id,
			Peer:            peer,
			Message:         message,
			Entities:        entities,
		}
	case crc_messages_getAllDrafts:
		r = TL_messages_getAllDrafts{}
	case crc_messages_getFeaturedStickers:
		r = TL_messages_getFeaturedStickers{
			Hash: m.Int(),
		}
	case crc_messages_readFeaturedStickers:
		r = TL_messages_readFeaturedStickers{
			Id: m.VectorLong(),
		}
	case crc_messages_getRecentStickers:
		flags := m.Int()
		attached := flags&(1<<0) != 0
		hash := m.Int()
		r = TL_messages_getRecentStickers{
			Flags:    flags,
			Attached: attached,
			Hash:     hash,
		}
	case crc_messages_saveRecentSticker:
		flags := m.Int()
		attached := flags&(1<<0) != 0
		id := m.Object()
		unsave := m.Object()
		r = TL_messages_saveRecentSticker{
			Flags:    flags,
			Attached: attached,
			Id:       id,
			Unsave:   unsave,
		}
	case crc_messages_clearRecentStickers:
		flags := m.Int()
		attached := flags&(1<<0) != 0
		r = TL_messages_clearRecentStickers{
			Flags:    flags,
			Attached: attached,
		}
	case crc_messages_getArchivedStickers:
		flags := m.Int()
		masks := flags&(1<<0) != 0
		offset_id := m.Long()
		limit := m.Int()
		r = TL_messages_getArchivedStickers{
			Flags:     flags,
			Masks:     masks,
			Offset_id: offset_id,
			Limit:     limit,
		}
	case crc_messages_getMaskStickers:
		r = TL_messages_getMaskStickers{
			Hash: m.Int(),
		}
	case crc_messages_getAttachedStickers:
		r = TL_messages_getAttachedStickers{
			Media: m.Object(),
		}
	case crc_messages_setGameScore:
		flags := m.Int()
		edit_message := flags&(1<<0) != 0
		force := flags&(1<<1) != 0
		peer := m.Object()
		id := m.Int()
		user_id := m.Object()
		score := m.Int()
		r = TL_messages_setGameScore{
			Flags:        flags,
			Edit_message: edit_message,
			Force:        force,
			Peer:         peer,
			Id:           id,
			User_id:      user_id,
			Score:        score,
		}
	case crc_messages_setInlineGameScore:
		flags := m.Int()
		edit_message := flags&(1<<0) != 0
		force := flags&(1<<1) != 0
		id := m.Object()
		user_id := m.Object()
		score := m.Int()
		r = TL_messages_setInlineGameScore{
			Flags:        flags,
			Edit_message: edit_message,
			Force:        force,
			Id:           id,
			User_id:      user_id,
			Score:        score,
		}
	case crc_messages_getGameHighScores:
		r = TL_messages_getGameHighScores{
			Peer:    m.Object(),
			Id:      m.Int(),
			User_id: m.Object(),
		}
	case crc_messages_getInlineGameHighScores:
		r = TL_messages_getInlineGameHighScores{
			Id:      m.Object(),
			User_id: m.Object(),
		}
	case crc_messages_getCommonChats:
		r = TL_messages_getCommonChats{
			User_id: m.Object(),
			Max_id:  m.Int(),
			Limit:   m.Int(),
		}
	case crc_messages_getAllChats:
		r = TL_messages_getAllChats{
			Except_ids: m.VectorInt(),
		}
	case crc_messages_getWebPage:
		r = TL_messages_getWebPage{
			Url:  m.String(),
			Hash: m.Int(),
		}
	case crc_messages_toggleDialogPin:
		flags := m.Int()
		pinned := flags&(1<<0) != 0
		peer := m.Object()
		r = TL_messages_toggleDialogPin{
			Flags:  flags,
			Pinned: pinned,
			Peer:   peer,
		}
	case crc_messages_reorderPinnedDialogs:
		flags := m.Int()
		force := flags&(1<<0) != 0
		order := m.Vector()
		r = TL_messages_reorderPinnedDialogs{
			Flags: flags,
			Force: force,
			Order: order,
		}
	case crc_messages_getPinnedDialogs:
		r = TL_messages_getPinnedDialogs{}
	case crc_messages_setBotShippingResults:
		flags := m.Int()
		query_id := m.Long()
		var error string
		if flags&(1<<0) != 0 {
			error = m.String()
		}
		var shipping_options []TL
		if flags&(1<<1) != 0 {
			shipping_options = m.Vector()
		}
		r = TL_messages_setBotShippingResults{
			Flags:            flags,
			Query_id:         query_id,
			Error:            error,
			Shipping_options: shipping_options,
		}
	case crc_messages_setBotPrecheckoutResults:
		flags := m.Int()
		success := flags&(1<<1) != 0
		query_id := m.Long()
		var error string
		if flags&(1<<0) != 0 {
			error = m.String()
		}
		r = TL_messages_setBotPrecheckoutResults{
			Flags:    flags,
			Success:  success,
			Query_id: query_id,
			Error:    error,
		}
	case crc_updates_getState:
		r = TL_updates_getState{}
	case crc_updates_getDifference:
		flags := m.Int()
		pts := m.Int()
		var pts_total_limit int32
		if flags&(1<<0) != 0 {
			pts_total_limit = m.Int()
		}
		date := m.Int()
		qts := m.Int()
		r = TL_updates_getDifference{
			Flags:           flags,
			Pts:             pts,
			Pts_total_limit: pts_total_limit,
			Date:            date,
			Qts:             qts,
		}
	case crc_updates_getChannelDifference:
		flags := m.Int()
		force := flags&(1<<0) != 0
		channel := m.Object()
		filter := m.Object()
		pts := m.Int()
		limit := m.Int()
		r = TL_updates_getChannelDifference{
			Flags:   flags,
			Force:   force,
			Channel: channel,
			Filter:  filter,
			Pts:     pts,
			Limit:   limit,
		}
	case crc_photos_updateProfilePhoto:
		r = TL_photos_updateProfilePhoto{
			Id: m.Object(),
		}
	case crc_photos_uploadProfilePhoto:
		r = TL_photos_uploadProfilePhoto{
			File: m.Object(),
		}
	case crc_photos_deletePhotos:
		r = TL_photos_deletePhotos{
			Id: m.Vector(),
		}
	case crc_photos_getUserPhotos:
		r = TL_photos_getUserPhotos{
			User_id: m.Object(),
			Offset:  m.Int(),
			Max_id:  m.Long(),
			Limit:   m.Int(),
		}
	case crc_upload_saveFilePart:
		r = TL_upload_saveFilePart{
			File_id:   m.Long(),
			File_part: m.Int(),
			Bytes:     m.StringBytes(),
		}
	case crc_upload_getFile:
		r = TL_upload_getFile{
			Location: m.Object(),
			Offset:   m.Int(),
			Limit:    m.Int(),
		}
	case crc_upload_saveBigFilePart:
		r = TL_upload_saveBigFilePart{
			File_id:          m.Long(),
			File_part:        m.Int(),
			File_total_parts: m.Int(),
			Bytes:            m.StringBytes(),
		}
	case crc_upload_getWebFile:
		r = TL_upload_getWebFile{
			Location: m.Object(),
			Offset:   m.Int(),
			Limit:    m.Int(),
		}
	case crc_help_getConfig:
		r = TL_help_getConfig{}
	case crc_help_getNearestDc:
		r = TL_help_getNearestDc{}
	case crc_help_getAppUpdate:
		r = TL_help_getAppUpdate{}
	case crc_help_saveAppLog:
		r = TL_help_saveAppLog{
			Events: m.Vector(),
		}
	case crc_help_getInviteText:
		r = TL_help_getInviteText{}
	case crc_help_getSupport:
		r = TL_help_getSupport{}
	case crc_help_getAppChangelog:
		r = TL_help_getAppChangelog{
			Prev_app_version: m.String(),
		}
	case crc_help_getTermsOfService:
		r = TL_help_getTermsOfService{}
	case crc_help_setBotUpdatesStatus:
		r = TL_help_setBotUpdatesStatus{
			Pending_updates_count: m.Int(),
			Message:               m.String(),
		}
	case crc_channels_readHistory:
		r = TL_channels_readHistory{
			Channel: m.Object(),
			Max_id:  m.Int(),
		}
	case crc_channels_deleteMessages:
		r = TL_channels_deleteMessages{
			Channel: m.Object(),
			Id:      m.VectorInt(),
		}
	case crc_channels_deleteUserHistory:
		r = TL_channels_deleteUserHistory{
			Channel: m.Object(),
			User_id: m.Object(),
		}
	case crc_channels_reportSpam:
		r = TL_channels_reportSpam{
			Channel: m.Object(),
			User_id: m.Object(),
			Id:      m.VectorInt(),
		}
	case crc_channels_getMessages:
		r = TL_channels_getMessages{
			Channel: m.Object(),
			Id:      m.VectorInt(),
		}
	case crc_channels_getParticipants:
		r = TL_channels_getParticipants{
			Channel: m.Object(),
			Filter:  m.Object(),
			Offset:  m.Int(),
			Limit:   m.Int(),
		}
	case crc_channels_getParticipant:
		r = TL_channels_getParticipant{
			Channel: m.Object(),
			User_id: m.Object(),
		}
	case crc_channels_getChannels:
		r = TL_channels_getChannels{
			Id: m.Vector(),
		}
	case crc_channels_getFullChannel:
		r = TL_channels_getFullChannel{
			Channel: m.Object(),
		}
	case crc_channels_createChannel:
		flags := m.Int()
		broadcast := flags&(1<<0) != 0
		megagroup := flags&(1<<1) != 0
		title := m.String()
		about := m.String()
		r = TL_channels_createChannel{
			Flags:     flags,
			Broadcast: broadcast,
			Megagroup: megagroup,
			Title:     title,
			About:     about,
		}
	case crc_channels_editAbout:
		r = TL_channels_editAbout{
			Channel: m.Object(),
			About:   m.String(),
		}
	case crc_channels_editAdmin:
		r = TL_channels_editAdmin{
			Channel: m.Object(),
			User_id: m.Object(),
			Role:    m.Object(),
		}
	case crc_channels_editTitle:
		r = TL_channels_editTitle{
			Channel: m.Object(),
			Title:   m.String(),
		}
	case crc_channels_editPhoto:
		r = TL_channels_editPhoto{
			Channel: m.Object(),
			Photo:   m.Object(),
		}
	case crc_channels_checkUsername:
		r = TL_channels_checkUsername{
			Channel:  m.Object(),
			Username: m.String(),
		}
	case crc_channels_updateUsername:
		r = TL_channels_updateUsername{
			Channel:  m.Object(),
			Username: m.String(),
		}
	case crc_channels_joinChannel:
		r = TL_channels_joinChannel{
			Channel: m.Object(),
		}
	case crc_channels_leaveChannel:
		r = TL_channels_leaveChannel{
			Channel: m.Object(),
		}
	case crc_channels_inviteToChannel:
		r = TL_channels_inviteToChannel{
			Channel: m.Object(),
			Users:   m.Vector(),
		}
	case crc_channels_kickFromChannel:
		r = TL_channels_kickFromChannel{
			Channel: m.Object(),
			User_id: m.Object(),
			Kicked:  m.Object(),
		}
	case crc_channels_exportInvite:
		r = TL_channels_exportInvite{
			Channel: m.Object(),
		}
	case crc_channels_deleteChannel:
		r = TL_channels_deleteChannel{
			Channel: m.Object(),
		}
	case crc_channels_toggleInvites:
		r = TL_channels_toggleInvites{
			Channel: m.Object(),
			Enabled: m.Object(),
		}
	case crc_channels_exportMessageLink:
		r = TL_channels_exportMessageLink{
			Channel: m.Object(),
			Id:      m.Int(),
		}
	case crc_channels_toggleSignatures:
		r = TL_channels_toggleSignatures{
			Channel: m.Object(),
			Enabled: m.Object(),
		}
	case crc_channels_updatePinnedMessage:
		flags := m.Int()
		silent := flags&(1<<0) != 0
		channel := m.Object()
		id := m.Int()
		r = TL_channels_updatePinnedMessage{
			Flags:   flags,
			Silent:  silent,
			Channel: channel,
			Id:      id,
		}
	case crc_channels_getAdminedPublicChannels:
		r = TL_channels_getAdminedPublicChannels{}
	case crc_bots_sendCustomRequest:
		r = TL_bots_sendCustomRequest{
			Custom_method: m.String(),
			Params:        m.Object(),
		}
	case crc_bots_answerWebhookJSONQuery:
		r = TL_bots_answerWebhookJSONQuery{
			Query_id: m.Long(),
			Data:     m.Object(),
		}
	case crc_payments_getPaymentForm:
		r = TL_payments_getPaymentForm{
			Msg_id: m.Int(),
		}
	case crc_payments_getPaymentReceipt:
		r = TL_payments_getPaymentReceipt{
			Msg_id: m.Int(),
		}
	case crc_payments_validateRequestedInfo:
		flags := m.Int()
		save := flags&(1<<0) != 0
		msg_id := m.Int()
		info := m.Object()
		r = TL_payments_validateRequestedInfo{
			Flags:  flags,
			Save:   save,
			Msg_id: msg_id,
			Info:   info,
		}
	case crc_payments_sendPaymentForm:
		flags := m.Int()
		msg_id := m.Int()
		var requested_info_id string
		if flags&(1<<0) != 0 {
			requested_info_id = m.String()
		}
		var shipping_option_id string
		if flags&(1<<1) != 0 {
			shipping_option_id = m.String()
		}
		credentials := m.Object()
		r = TL_payments_sendPaymentForm{
			Flags:              flags,
			Msg_id:             msg_id,
			Requested_info_id:  requested_info_id,
			Shipping_option_id: shipping_option_id,
			Credentials:        credentials,
		}
	case crc_payments_getSavedInfo:
		r = TL_payments_getSavedInfo{}
	case crc_payments_clearSavedInfo:
		flags := m.Int()
		credentials := flags&(1<<0) != 0
		info := flags&(1<<1) != 0
		r = TL_payments_clearSavedInfo{
			Flags:       flags,
			Credentials: credentials,
			Info:        info,
		}
	case crc_phone_getCallConfig:
		r = TL_phone_getCallConfig{}
	case crc_phone_requestCall:
		r = TL_phone_requestCall{
			User_id:   m.Object(),
			Random_id: m.Int(),
			G_a_hash:  m.StringBytes(),
			Protocol:  m.Object(),
		}
	case crc_phone_acceptCall:
		r = TL_phone_acceptCall{
			Peer:     m.Object(),
			G_b:      m.StringBytes(),
			Protocol: m.Object(),
		}
	case crc_phone_confirmCall:
		r = TL_phone_confirmCall{
			Peer:            m.Object(),
			G_a:             m.StringBytes(),
			Key_fingerprint: m.Long(),
			Protocol:        m.Object(),
		}
	case crc_phone_receivedCall:
		r = TL_phone_receivedCall{
			Peer: m.Object(),
		}
	case crc_phone_discardCall:
		r = TL_phone_discardCall{
			Peer:          m.Object(),
			Duration:      m.Int(),
			Reason:        m.Object(),
			Connection_id: m.Long(),
		}
	case crc_phone_setCallRating:
		r = TL_phone_setCallRating{
			Peer:    m.Object(),
			Rating:  m.Int(),
			Comment: m.String(),
		}
	case crc_phone_saveCallDebug:
		r = TL_phone_saveCallDebug{
			Peer:  m.Object(),
			Debug: m.Object(),
		}
	default:
		m.err = fmt.Errorf("Unknown constructor: %x", constructor)
		return nil
	}
	return
}
