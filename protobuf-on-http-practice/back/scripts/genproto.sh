#!/bin/zsh

set -x

PROTO_DIR=./proto
NODE_MODULES_DIR=./node_modules

for f in ${PROTO_DIR}/**/*.proto
do
    echo ${f}
    protoc \
        --go_opt=paths=source_relative \
        --plugin=protoc-gen-grpc=${NODE_MODULES_DIR}/.bin/grpc_tools_node_protoc_plugin \
        --plugin=protoc-gen-ts=${NODE_MODULES_DIR}/.bin/protoc-gen-ts \
        --ts_out=service=grpc-node:. \
        --go_out=plugins=grpc:. \
        --js_out=import_style=commonjs,binary:. \
        ${f}
done
