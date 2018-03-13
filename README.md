***There was a big change on Mar 14. If you are looking for the old one, please use v0.1.3.***

MTProto
===
Telegram MTProto and proxy (over gRPC) in Go (golang).
Telegram API layer: ***71***

## Quick start

## Usage
### Without proxy
You can find the real code at [simpleshell](https://github.com/cjongseok/mtproto/blob/master/examples/simpleshell/main.go).
#### Sign-in with key
```go
// Mew MTProto manager
config, _ := core.NewConfiguration(apiId, apiHash, appVersion, deviceModel, systemVersion, language, 0, 0, key)
manager, _ := core.NewManager(config)

// Sign-in by key
mconn, _ := manager.LoadAuthentication(phoneNumber, preferredAddr)
```
#### New sign-in
```go
// New MTProto manager
config, _ := core.NewConfiguration(apiId, apiHash, appVersion, deviceModel, systemVersion, language, 0, 0, "new-key.mtproto")
manager, _ := core.NewManager(config)

// Request to send an authentication code to mobile
mconn, sentCode, err := manager.NewAuthentication(phoneNumber, telegramAddress, false)

// Get the code from user input
fmt.Scanf("%s", &code)

// Sign-in and generate the new key
_, err = mconn.SignIn(phoneNumber, code, sentCode.GetValue().PhoneCodeHash)
```
#### Telegram RPC over mtproto
You can do RPCs by function calls as below examples.
##### Get dialogs
```go
// New RPC caller
caller := core.RPCaller{mconn}

// New input peer
emptyPeer := &core.TypeInputPeer{&core.TypeInputPeer_InputPeerEmpty{&core.PredInputPeerEmpty{}}

// Query (it is declared as a method, "messages.getDialogs", in TL-schema Layer 71)
dialogs, _ := caller.MessagesGetDialogs(context.Background(), &core.ReqMessagesGetDialogs{
    OffsetDate: 	0,
    OffsetId: 	0,
    OffsetPeer: 	emptyPeer,
    Limit: 		1,
})
```
##### Send a message to a channel
```go
// New RPC caller
caller := core.RPCaller{mconn}

// New input peer
channelPeer := &core.TypeInputPeer{&core.TypeInputPeer_InputPeerChannel{
    &core.PredInputPeerChannel{
        yourChannelId, yourChannelHash,
    }}}

// Query (it is declared as a method, "messages.sendMessage", in TL-schema Layer 71)
caller.MessagesSendMessage(context.Background(), &core.ReqMessagesSendMessage{
    Peer:      peer,
    Message:   "Hello MTProto",
    RandomId: rand.Int63(),
})
```

### With proxy
You can use the proxy in two purposes:
* MTProto session sharing: Many proxy clients can use the same MTProto session on the proxy server.
* MTProto in other languages: The proxy enables various languages on its client side, since it uses gRPC.

You can find the real code at [proxy_test.go](https://github.com/cjongseok/mtproto/blob/master/proxy/proxy_test.go)
#### Proxy server
For now you should write your own proxy in Go as below.<br>
A standalone proxy daemon would be ready in the near future.
```go
// New proxy server
config, _ := core.NewConfiguration(apiId, apiHash, appVersion, deviceModel, systemVersion, language, 0, 0, key)
server = proxy.NewServer(port)

// Start the server
server.Start(config, phone, telegramAddr)
```
#### Proxy client in Go
```go
// New proxy client
client, _ := proxy.NewClient(proxyAddr)

// Query dialogs. It is same with the 'Get dialogs' section but the RPC caller
emptyPeer := &core.TypeInputPeer{&core.TypeInputPeer_InputPeerEmpty{&core.PredInputPeerEmpty{}}
dialogs, err := client.MessagesGetDialogs(context.Background(), &core.ReqMessagesGetDialogs{
    OffsetDate: 0,
    OffsetId:   0,
    OffsetPeer: emptyPeer,
    Limit:      1,
})
```
#### Proxy client in other languages
By compiling core/types.tl.proto and proxy/tl_update.proto, 
you can create clients in another language.
For the compiliation, you need [Google Protobuf](https://developers.google.com/protocol-buffers/).<br>
If Protobuf ready, you can compile the files with a couple of options.
For example, you can generate Go codes as below.
```bash
# At the mtproto home
mkdir -p out/core out/proxy
protoc -I core -I <PROTOC_HOME>/include core/types.tl.proto --go_out=plugins=grpc:./out/core
protoc -I $GOPATH/src -I proxy tl_update.proto --go_out=plugins=grpc:./out/proxy
```
For other languages, you can alternate '--go_out' with '-xxx_out'. (I didn't try it, but I believe it would work as other gRPC projects. If you try it, please share the result.)<br>
It seems protoc-3.5.1 supports below language options.
* --cpp_out=OUT_DIR           Generate C++ header and source.
* --csharp_out=OUT_DIR        Generate C# source file.
* --java_out=OUT_DIR          Generate Java source file.
* --javanano_out=OUT_DIR      Generate Java Nano source file.
* --js_out=OUT_DIR            Generate JavaScript source.
* --objc_out=OUT_DIR          Generate Objective C header and source.
* --php_out=OUT_DIR           Generate PHP source file.
* --python_out=OUT_DIR        Generate Python source file.
* --ruby_out=OUT_DIR          Generate Ruby source file.

### Types and Predicates
### X and !X

## Tools
### Keygen
### Dumplayer

## Compiler


## Acknowledgement
* https://github.com/sdidyk/mtproto: It is the backend of most MTProto Go implementations.
I referred its MTProto schema compiler, (de)serializer, handshaking, and encryption.
* https://github.com/shelomentsevd/mtproto: I referred its layer 65 implementation and API wrappers.
* https://github.com/ronaksoft/mtproto: I referred its backend changes for layer 71.


## License
Apache 2.0
