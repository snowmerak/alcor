#!/bin/bash

protoc --go_out=./ proto/protobuf/*.proto
protoc --dart_out=./client/lib proto/protobuf/*.proto