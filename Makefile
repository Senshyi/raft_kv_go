PROJECT= raft_kv_go
GOPATH ?= $(shell go env GOPATH)

gen:
	protoc --proto_path=./proto ./proto/*.proto --go-grpc_out=. --go_out=.

