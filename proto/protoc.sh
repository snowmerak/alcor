#!/bin/bash

protoc --go_out=./ proto/protobuf/*.proto
protoc --dart_out=./client/lib proto/protobuf/*.proto
protoc --dart_out=./admin/lib proto/protobuf/*.proto
protoc --python_out=./edgepy proto/protobuf/*.proto