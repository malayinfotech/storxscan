#!/usr/bin/env bash

cd "$(dirname "${BASH_SOURCE[0]}")/.."

docker run -v `pwd`:/opt/storxscan -w /opt/storxscan ghcr.io/elek/storx-build:20220901-2 go build -o cmd/storxscan/storxscan storx.io/storxscan/cmd/storxscan
