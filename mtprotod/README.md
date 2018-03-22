***It is broken now. It is under repair.***

mtprotod
===
MTProto proxy daemon.
It contains Telegram MTProto implementation in Go, its clients can share one Telegram MTProto session. Or you can also make your own session scheduler in Go.

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

Quick test
---
You can test if the proxy running using [proxy_test.go](https://github.com/cjongseok/mtproto/blob/master/proxy/proxy_test.go).
```bash
# At mtproto/proxy,
# let's get dialogs through over the proxy
dep ensure
go run main.go test --proxy <proxy_address>
```
