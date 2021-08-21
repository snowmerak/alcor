#!/bin/bash

protoc --go_out=./ proto/protobuf/*.proto
protoc --dart_out=./client/lib proto/protobuf/*.proto

cp proto/protobuf/*.proto ./observer/src/model
pb-rs ./observer/src/model/*.proto
rm ./observer/src/model/*.proto