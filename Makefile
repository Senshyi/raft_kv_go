PROJECT= raft_kv_go
GOPATH ?= $(shell go env GOPATH)

gen:
	protoc --proto_path=./proto ./proto/*.proto --plugin=$(GOPATH)/bin/protoc-gen-go-grpc --go-grpc_out=. --go_out=.

