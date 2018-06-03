mtprotod
===
MTProto proxy daemon.
It contains Telegram MTProto implementation in Go.
Its clients can share one Telegram MTProto session, or you can make your own session scheduler.

Quick Start
---
You can test if the proxy running using [proxy_test.go](https://github.com/cjongseok/mtproto/blob/master/proxy/proxy_test.go).
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
go test --run TestDialogs
```

Usage
---
### Run as a container
mtprotod Docker image is at Docker hub, [cjongseok/mtprotod](https://hub.docker.com/r/cjongseok/mtprotod/).
```bash
docker run \
-p <Proxy_port>:<Proxy_port> \
-v /your/mtproto/secrets/directory:/opt \
cjongseok/mtprotod start  \
--addr <Your_Telegram_server_address> \
--apiid <Your_Telegram_api_id> \
--apihash <Your_Telegram_api_hash> \
--phone <Your_Telegram_phone_number> \
--port <Proxy_port> \
--secrets /opt/<Your_MTProto_secrets_file_name>
```
### Build & Run
mtprotod is vendored in [***dep***](https://github.com/golang/dep).
```bash
# In mtprotod directory
dep ensure
go run main.go start \
--addr <Your_Telegram_server_address> \
--apiid <Your_Telegram_api_id> \
--apihash <Your_Telegram_api_hash> \
--phone <Your_Telegram_phone_number> \
--port <Proxy_port> \
--secrets /opt/<Your_MTProto_secrets_file_name>
```

Clients
---
* [Go client](https://github.com/cjongseok/mtproto#client-in-go)
* [Python client](https://github.com/cjongseok/mtproto/tree/master/py)
