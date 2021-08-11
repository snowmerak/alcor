#!/bin/bash

protoc --go_out=../server protobuf/*.proto
protoc --dart_out=../client/lib protobuf/*.proto