#!/bin/sh
# where should put duckdb libs (generate duckdb.a first)
CGO_LDFLAGS="-L$(pwd)/.libs/win/" \
go build -tags=duckdb_use_lib,all -o engine.exe \
-ldflags="-X main.Version=$(Git describe --abbrev=0 --tags) -s -w" . && \
# upx to compress the executable.
upx --lzma engine.exe
