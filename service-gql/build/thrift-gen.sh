#!/bin/bash

find ./src/thrift | grep -v .thrift$ | xargs rm -rf
mkdir -p ./src/thrift/{user,post}
thrift -r -out ./src/thrift/user --gen js:node ./src/thrift/user.thrift
thrift -r -out ./src/thrift/post --gen js:node ./src/thrift/post.thrift
