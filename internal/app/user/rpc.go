package user

import (
	"context"

	google_protobuf2 "github.com/gogo/protobuf/types"
	proto "github.com/weisd/web-kit/api/protobuf/user"
)

var _ proto.UserRpcServiceServer = &RpcServer{}

// RpcServer RpcServer
type RpcServer struct{}

func NewRcpServer() *RpcServer {
	return &RpcServer{}
}

// Create 创建用户
func (p *RpcServer) Create(ctx context.Context, in *proto.User) (*google_protobuf2.Empty, error) {
	return nil, nil
}

// UpdatePassword 更新密码
func (p *RpcServer) UpdatePassword(ctx context.Context, in *proto.IDPassword) (*google_protobuf2.Empty, error) {
	return nil, nil
}

// UpdatePhone  更新手机号
func (p *RpcServer) UpdatePhone(ctx context.Context, in *proto.IDPhone) (*google_protobuf2.Empty, error) {
	return nil, nil
}

// UpdateEmail 更新email
func (p *RpcServer) UpdateEmail(ctx context.Context, in *proto.IDEmail) (*google_protobuf2.Empty, error) {
	return nil, nil
}

// UpdateNickname 更新昵称
func (p *RpcServer) UpdateNickname(ctx context.Context, in *proto.IDNickname) (*google_protobuf2.Empty, error) {
	return nil, nil
}

// UpdateAvatar 更新头像
func (p *RpcServer) UpdateAvatar(ctx context.Context, in *proto.IDAvatar) (*google_protobuf2.Empty, error) {
	return nil, nil
}

// UpdateStatus 更新状态
func (p *RpcServer) UpdateStatus(ctx context.Context, in *proto.IDStatus) (*google_protobuf2.Empty, error) {
	return nil, nil
}

// InfoByID 通过id查询
func (p *RpcServer) InfoByID(ctx context.Context, in *proto.ID) (*proto.User, error) {
	return nil, nil
}

// InfoByPhone 通过手机查询
func (p *RpcServer) InfoByPhone(ctx context.Context, in *proto.Phone) (*proto.User, error) {
	return nil, nil
}

// InfoByEmail 通过email查询
func (p *RpcServer) InfoByEmail(ctx context.Context, in *proto.Email) (*proto.User, error) {
	return nil, nil
}

// InfoByAccount 通过 手机、email、昵称查询
func (p *RpcServer) InfoByAccount(ctx context.Context, in *proto.Account) (*proto.User, error) {
	return nil, nil
}
