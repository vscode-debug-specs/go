#!/bin/bash
protoc \
    --go_out=. \
    --go_opt=paths=source_relative \
    "--go_opt=Mgateways/grpc/service.proto=github.com/vscode-debug-specs/go/gateways/grpc;grpc" \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    "--go-grpc_opt=Mgateways/grpc/service.proto=github.com/vscode-debug-specs/go/gateways/grpc;grpc" \
    gateways/grpc/service.proto
