MTProto
===
Telegram MTProto and proxy (over gRPC) in Go (golang).
Telegram API layer: ***71***

## Quick start
```sh
# It is vendored in 'dep'. Refer to https://github.com/golang/dep for 'dep' installation.
dep ensure

# Run simple shell with your Telegram API id, hash, and, server address with your phone number.
# If you don't have Telegram API stuffs, get them from 'https://my.telegram.org/apps'.
go run examples/simpleshell/main.go <APIID> <APIHASH> <PHONE> <IP> <PORT>

# Then you can see 'Enter code:' message
# Telegram sends you an authentication code. Check it on your mobile or desktop app and put it.
Enter code: <YOUR_CODE>

# Now signed-in. Let's get your recent dialogs. 
# You can see them in JSON.
$ dialogs
....

# Quit the shell.
$ exit

# You can find 'credentials.json' file which keeps your MTProto secerets.
ls -al credentials.json

# You can check if the scerets correct by sign-in with it.
go run examples/simpleshell/main.go  credentials.json
```

## Usage
You can find the real code at [simpleshell](https://github.com/cjongseok/mtproto/blob/master/examples/simpleshell/main.go).
### Sign-in with key
```go
// Mew MTProto manager
config, _ := mtproto.NewConfiguration(appVersion, deviceModel, systemVersion, language, 0, 0, "credentials.json")
manager, _ := mtproto.NewManager(config)

// Sign-in by key
mconn, _ := manager.LoadAuthentication()
```
### Sign-in without key
```go
// New MTProto manager
config, _ := mtproto.NewConfiguration(appVersion, deviceModel, systemVersion, language, 0, 0, "new-credentials.json")
manager, _ := mtproto.NewManager(config)

// Request to send an authentication code
// It needs your phone number and Telegram API stuffs you can check in https://my.telegram.org/apps 
mconn, sentCode, err := manager.NewAuthentication(phoneNumber, apiID, apiHash, ip, port)

// Get the code from user input
fmt.Scanf("%s", &code)

// Sign-in and generate the new key
_, err = mconn.SignIn(phoneNumber, code, sentCode.GetValue().PhoneCodeHash)
```
### Telegram RPC in Protocol Buffer
cjongseok/mtproto implements [TL-schema](https://core.telegram.org/schema) [functions](https://github.com/cjongseok/mtproto/blob/142b7f31f963a074dac9dd6759e0ae054e4a894c/compiler/scheme-71.tl#L951) in [Protocol Buffer](https://developers.google.com/protocol-buffers/). These are declared in [types.tl.proto](https://github.com/cjongseok/mtproto/blob/142b7f31f963a074dac9dd6759e0ae054e4a894c/types.tl.proto#L6385) as RPCs, and implemented in Go at [procs.tl.go](https://github.com/cjongseok/mtproto/blob/master/procs.tl.go). With the same interface, you can call functions not only in direct connection to the Telegram server, but also over a mtproto proxy.

Let's have two direct call examples, [messages.getDialogs](https://github.com/cjongseok/mtproto/blob/142b7f31f963a074dac9dd6759e0ae054e4a894c/compiler/scheme-71.tl#L1025) and [messages.sendMessage](https://github.com/cjongseok/mtproto/blob/142b7f31f963a074dac9dd6759e0ae054e4a894c/compiler/scheme-71.tl#L1033).
#### Get dialogs
```go
// New RPC caller with a connection to Telegram server. 
// By alternating mconn with a proxy connection, you can call same functions over proxy. It is covered later.
caller := mtproto.RPCaller{mconn}

// New input peer
// In Telegram DSL, Predicates inherit a Type.
// Here we create a Predicate, InputPeerEmpty whose parent is InputPeer.
// More details about these types are covered later.
emptyPeer := &mtproto.TypeInputPeer{&mtproto.TypeInputPeer_InputPeerEmpty{&mtproto.PredInputPeerEmpty{}}

// Query to Telegram
dialogs, _ := caller.MessagesGetDialogs(context.Background(), &mtproto.ReqMessagesGetDialogs{
    OffsetDate: 0,
    OffsetId: 	0,
    OffsetPeer: emptyPeer,
    Limit: 		1,
})
```
#### Send a message to a channel
```go
// New RPC caller with a connection to Telegram server. 
caller := mtproto.RPCaller{mconn}

// New input peer
// Create a Predicate, InputPeerChannel, wraped by its parent Type, InputPeer.
channelPeer := &mtproto.TypeInputPeer{&mtproto.TypeInputPeer_InputPeerChannel{
    &mtproto.PredInputPeerChannel{
        yourChannelId, yourChannelHash,
    }}}

// Send a request to Telegram
caller.MessagesSendMessage(context.Background(), &mtproto.ReqMessagesSendMessage{
    Peer:      peer,
    Message:   "Hello MTProto",
    RandomId:  rand.Int63(),
})
```

## How mtproto is impelemented in Protocol Buffer
### Types
Telegram's mtproto has three kinds of types, Type, Predicate, and Method. A Type is a kind of a data structure interface which has no fields, and a Predicate implements a Type. In the above case, mtproto.[PredInputPeerChannel](https://github.com/cjongseok/mtproto/blob/142b7f31f963a074dac9dd6759e0ae054e4a894c/types.tl.proto#L3046) is a Predicate of a Type mtproto.[TypeInputPeer](https://github.com/cjongseok/mtproto/blob/142b7f31f963a074dac9dd6759e0ae054e4a894c/types.tl.proto#L19). gRPC recommends to implement this kind of polymorphism with [Oneof](https://developers.google.com/protocol-buffers/docs/proto3#oneof), so [InputPeer](https://github.com/cjongseok/mtproto/blob/142b7f31f963a074dac9dd6759e0ae054e4a894c/compiler/scheme-71.tl#L38) is defined in Protocol Buffer as below:
```protobuf
// types.tl.proto:19
message TypeInputPeer {
	oneof Value {
		PredInputPeerEmpty InputPeerEmpty = 1;
		PredInputPeerSelf InputPeerSelf = 2;
		PredInputPeerChat InputPeerChat = 3;
		PredInputPeerUser InputPeerUser = 4;
		PredInputPeerChannel InputPeerChannel = 5;
    }
}
```
The use of gRPC Oneof in Go is complex, because Go does not allow hierarchical relations among types, e.g., inheritance. I believe, however, gRPC guys did their best and it would be the best implementation of such polymorphism in Go with RPC. For more details about the use of OneOf in Go, please refer to [this document](https://developers.google.com/protocol-buffers/docs/reference/go-generated#oneof).

### Methods
Mtproto methods have a namespace as you can see in [TL-schema](https://github.com/cjongseok/mtproto/blob/142b7f31f963a074dac9dd6759e0ae054e4a894c/compiler/scheme-71.tl#L951), e.g., auth, account, users, ... . Instead of managing these namespaces as separate [Protocol Buffer services](https://developers.google.com/protocol-buffers/docs/proto3#services), these are integrated into one Protocol Buffer Service, [Mtproto](https://github.com/cjongseok/mtproto/blob/142b7f31f963a074dac9dd6759e0ae054e4a894c/types.tl.proto#L6385), and namesapces are remained as method name prefixes. In the above example, the original name of [MessagesSendMessage](https://github.com/cjongseok/mtproto/blob/142b7f31f963a074dac9dd6759e0ae054e4a894c/types.tl.proto#L6425) is [messages.sendMessage](https://github.com/cjongseok/mtproto/blob/142b7f31f963a074dac9dd6759e0ae054e4a894c/compiler/scheme-71.tl#L1033).

In the original schema, a method can have multiple parameters. These paremeters are declared into a new data structure in Protocol Buffer whose name starts with 'Req'. For example, [messages.sendMessage](https://github.com/cjongseok/mtproto/blob/142b7f31f963a074dac9dd6759e0ae054e4a894c/compiler/scheme-71.tl#L1033) requires many parameters, and these are transformed into [ReqMessagesSendMessage](https://github.com/cjongseok/mtproto/blob/142b7f31f963a074dac9dd6759e0ae054e4a894c/types.tl.proto#L5165) in [MessagesSendMessage](https://github.com/cjongseok/mtproto/blob/142b7f31f963a074dac9dd6759e0ae054e4a894c/types.tl.proto#L6425).

# Proxy
You can use the proxy in two purposes:
* MTProto session sharing: Many proxy clients can use the same MTProto session on the proxy server.
* MTProto in other languages: The proxy enables various languages on its client side, since it uses gRPC.

## Server
### As a stand-alone daemon
[mtprotod](https://github.com/cjongseok/mtproto/tree/master/mtprotod) is a stand-alone proxy daemon containing Telegram MTProto implementation in Go.

#### Quick Start
```bash
# start mtprotod at port 11011
docker run \
-p 11011: 11011 \
-v /your/mtproto/secrets/directory:/opt \
cjongseok/mtprotod start  \
--port 11011 \
--addr <Your_Telegram_server_address> \
--apiid <Your_Telegram_api_id> \
--apihash <Your_Telegram_api_hash> \
--phone <Your_Telegram_phone_number> \
--secrets /opt/<Your_MTProto_secrets_file_name>

# At mtproto/proxy, let's get dialogs through over the proxy
dep ensure
go test proxy/proxy_test.go --run TestDialogs
```

#### Build & Run
```bash
# In mtproto directory
dep ensure
go run mtprotod/main.go start \
--addr <Your_Telegram_server_address> \
--apiid <Your_Telegram_api_id> \
--apihash <Your_Telegram_api_hash> \
--phone <Your_Telegram_phone_number> \
--port <Proxy_port> \
--secrets /opt/<Your_MTProto_secrets_file_name>
```

### As a part of Go application
Use mtproto/proxy package.
```go
// New proxy server
config, _ := mtproto.NewConfiguration(apiId, apiHash, appVersion, deviceModel, systemVersion, language, 0, 0, key)
server = proxy.NewServer(port)

// Start the server
server.Start(config, phone)
```
## Client in Go
```go
// New proxy client
client, _ := proxy.NewClient(proxyAddr)

// Telegram RPC over proxy. It is same with the previous 'Get dialogs' but the RPC caller
emptyPeer := &mtproto.TypeInputPeer{&mtproto.TypeInputPeer_InputPeerEmpty{&mtproto.PredInputPeerEmpty{}}
dialogs, err := client.MessagesGetDialogs(context.Background(), &mtproto.ReqMessagesGetDialogs{
    OffsetDate: 0,
    OffsetId:   0,
    OffsetPeer: emptyPeer,
    Limit:      1,
})
```
## Client in Python
See [py](https://github.com/cjongseok/mtproto/tree/master/py).
## Client in other languages
By compiling [types.tl.proto](https://github.com/cjongseok/mtproto/tree/master/types.tl.proto) and [proxy/tl_update.proto](https://github.com/cjongseok/mtproto/tree/master/proxy/tl_update.proto), 
you can create clients in your preferred language.<br>
For this, you need to install [Protocol Buffer](https://developers.google.com/protocol-buffers/) together with gRPC library of the target language.
Then you can compile Protocol Buffer files with this kind of command lines:
* [Go](https://github.com/cjongseok/mtproto/blob/master/compiler/build.sh)
* [Python](https://github.com/cjongseok/mtproto/blob/master/compiler/build_py.sh)

You can find these command lines for other languages in [gRPC tutorial](https://grpc.io/docs/).

<!--
### Types and Predicates
### X and !X

## Tools
### Keygen
### Dumplayer

## Compiler
-->


## Acknowledgement
* https://github.com/sdidyk/mtproto: It is the backend of most MTProto Go implementations.
I referred its MTProto schema compiler, (de)serializer, handshaking, and encryption.
* https://github.com/shelomentsevd/mtproto: I referred its layer 65 implementation and API wrappers.
* https://github.com/ronaksoft/mtproto: I referred its backend changes for layer 71.


## License
Apache 2.0
