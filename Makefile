.PHONY: proto
proto:
	protoc --proto_path=$(GOPATH)/src --proto_path=$(GOPATH)/src/github.com/google/protobuf/src --proto_path=proto --go_out=plugins=grpc:proto --govalidators_out=proto proto/*.proto

