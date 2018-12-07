package user

import (
	"context"

	proto "github.com/weisd/web-kit/api/protobuf/user"
)

var _ proto.APIServiceServer = &APIServer{}

// APIServer APIServer
type APIServer struct{}

// NewAPIServer NewAPIServer
func NewAPIServer() *APIServer {
	return &APIServer{}
}

// Login 创建用户
func (p *APIServer) Login(ctx context.Context, in *proto.LoginRequest) (*proto.User, error) {
	return nil, nil
}
