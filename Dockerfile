#syntax=docker/dockerfile:1
FROM golang:1.16-alpine
WORKDIR /watcher-ethereum
COPY . .
RUN apk add build-base gcc
RUN go env -w GOPROXY=https://goproxy.io,direct
RUN go env -w GOPRIVATE=github.com/itsabgr/*
RUN go mod download
RUN go test ./...
ENV NET=https://mainnet.infura.io/v3/{Project ID}
CMD go run ./cmd/watch-ethereum