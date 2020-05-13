#!/bin/bash

rm -rf ./internal/thrift-rpc/*
mkdir -p ./internal/thrift-rpc/
thrift -r -out ./internal/thrift-rpc --gen go ./internal/thrift/post.thrift