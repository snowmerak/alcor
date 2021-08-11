#!/bin/bash

protoc --go_out=./ protobuf/*.proto
protoc --dart_out=./proto_dart protobuf/*.proto