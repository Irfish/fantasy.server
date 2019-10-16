#!/usr/bin/env bash
protoc --go_out=plugins=grpc:. rpc_*.proto


