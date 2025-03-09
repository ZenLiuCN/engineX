#!/bin/sh
# where should put duckdb libs (generate duckdb.a first)
CGO_LDFLAGS="-L$(pwd)/../../Loader/.libs/win/" \
go build -tags=duckdb_use_lib,all -o engine.exe \
-ldflags="-X main.Version=$(Git describe --abbrev=0 --tags) -s -w" . && \
# upx to compress the executable.
upx --lzma engine.exe
# HTTP_PROXY=127.0.0.1:1079 GOPROXY=https://proxy.golang.org GO111MODULE=on  go get github.com/ZenLiuCN/engine@v0.6.8
