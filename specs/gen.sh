#!/bin/sh

protoc -I . demo.proto --go_out=plugins=grpc:../src