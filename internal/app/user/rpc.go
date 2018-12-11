package user

import (
	"log"

	proto "github.com/weisd/web-kit/api/protobuf/user"
)

// RPCServer RPCServer
type RPCServer interface {
	proto.RPCServiceServer
	Init(dsn string) error
}

var userRPCServerRegister map[string]RPCServer

// RPCRegister RPCRegister
func RPCRegister(name string, rpc RPCServer) {
	if _, has := userRPCServerRegister[name]; has {
		log.Fatalf("rpc server exists: %s", name)
	}

	userRPCServerRegister[name] = rpc
}
