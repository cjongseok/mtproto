#!/bin/bash
# build a linux executable
dep ensure
CGO_ENABLED=0 GOOS=linux go build -ldflags '-s' -a -installsuffix cgo -o main

# build & push a Docker image
proxyVersion=$(go run main.go --version | awk '{print $3}')
dockerImage=cjongseok/mtprotod:${proxyVersion}
latestImage=cjongseok/mtprotod:latest
docker build -t ${dockerImage} .
docker tag ${dockerImage} ${latestImage}
docker login
docker push ${dockerImage}
docker push ${latestImage}
