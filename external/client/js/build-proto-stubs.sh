#!/usr/bin/env bash

yarn run grpc_tools_node_protoc \
--js_out=import_style=commonjs,binary:./src/stubs \
--grpc_out=grpc_js:./src/stubs \
-I ../../../internal/infrastructure/proto\
 ../../../internal/infrastructure/proto/*.proto

yarn run grpc_tools_node_protoc \
--plugin=protoc-gen-ts=./node_modules/grpc_tools_node_protoc_ts/bin/protoc-gen-ts \
--ts_out=grpc_js:./src/stubs \
-I ../../../internal/infrastructure/proto/ \
 ../../../internal/infrastructure/proto/*.proto