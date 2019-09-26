#!/bin/bash

protoc feedbackpb/survey.proto --go_out=plugins=grpc:.
