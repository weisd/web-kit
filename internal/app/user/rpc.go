package user

import (
	"context"

	google_protobuf2 "github.com/gogo/protobuf/types"
	proto "github.com/weisd/web-kit/api/protobuf/user"
)

var _ proto.RPCServiceServer = &RPCServer{}

// RPCServer RPCServer
type RPCServer struct{}

// NewRPCServer NewRPCServer
func NewRPCServer() *RPCServer {
	return &RPCServer{}
}

// Create 创建用户
func (p *RPCServer) Create(ctx context.Context, in *proto.User) (out *google_protobuf2.Empty, err error) {
	return nil, nil
}

// UpdatePassword 更新密码
func (p *RPCServer) UpdatePassword(ctx context.Context, in *proto.IDPassword) (out *google_protobuf2.Empty, err error) {
	return nil, nil
}

// UpdatePhone  更新手机号
func (p *RPCServer) UpdatePhone(ctx context.Context, in *proto.IDPhone) (out *google_protobuf2.Empty, err error) {
	return nil, nil
}

// UpdateEmail 更新email
func (p *RPCServer) UpdateEmail(ctx context.Context, in *proto.IDEmail) (out *google_protobuf2.Empty, err error) {
	return nil, nil
}

// UpdateNickname 更新昵称
func (p *RPCServer) UpdateNickname(ctx context.Context, in *proto.IDNickname) (out *google_protobuf2.Empty, err error) {
	return nil, nil
}

// UpdateAvatar 更新头像
func (p *RPCServer) UpdateAvatar(ctx context.Context, in *proto.IDAvatar) (out *google_protobuf2.Empty, err error) {
	return nil, nil
}

// UpdateStatus 更新状态
func (p *RPCServer) UpdateStatus(ctx context.Context, in *proto.IDStatus) (out *google_protobuf2.Empty, err error) {
	return nil, nil
}

// InfoByID 通过id查询
func (p *RPCServer) InfoByID(ctx context.Context, in *proto.ID) (out *proto.User, err error) {
	return nil, nil
}

// InfoByPhone 通过手机查询
func (p *RPCServer) InfoByPhone(ctx context.Context, in *proto.Phone) (out *proto.User, err error) {
	return nil, nil
}

// InfoByEmail 通过email查询
func (p *RPCServer) InfoByEmail(ctx context.Context, in *proto.Email) (out *proto.User, err error) {
	return nil, nil
}

// InfoByAccount 通过 手机、email、昵称查询
func (p *RPCServer) InfoByAccount(ctx context.Context, in *proto.Account) (out *proto.User, err error) {
	return nil, nil
}
