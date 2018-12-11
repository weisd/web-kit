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

// UserRPCServerRegister UserRPCServerRegister
var UserRPCServerRegister = map[string]RPCServer{}

// RPCRegister RPCRegister
func RPCRegister(name string, rpc RPCServer) {
	if _, has := UserRPCServerRegister[name]; has {
		log.Fatalf("rpc server exists: %s", name)
	}

	UserRPCServerRegister[name] = rpc
}
