#!/bin/bash

protoc --proto_path ../01-protobuffer survey.proto --go_out=plugins=grpc:./../01-protobuffer/
