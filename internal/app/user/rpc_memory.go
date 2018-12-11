package user

import (
	"context"

	google_protobuf2 "github.com/gogo/protobuf/types"
	proto "github.com/weisd/web-kit/api/protobuf/user"
)

// var _ proto.RPCServiceServer = &MemoryServer{}
var _ RPCServer = &MemoryServer{}

func init() {
	RPCRegister("memory", &MemoryServer{})
}

// MemoryServer MemoryServer
type MemoryServer struct {
	data map[int64]*proto.User
}

// Init Init
func (p *MemoryServer) Init(dsn string) error {
	return nil
}

// Create 创建用户
func (p *MemoryServer) Create(ctx context.Context, in *proto.User) (out *google_protobuf2.Empty, err error) {
	return
}

// UpdatePassword 更新密码
func (p *MemoryServer) UpdatePassword(ctx context.Context, in *proto.IDPassword) (out *google_protobuf2.Empty, err error) {
	return
}

// UpdatePhone  更新手机号
func (p *MemoryServer) UpdatePhone(ctx context.Context, in *proto.IDPhone) (out *google_protobuf2.Empty, err error) {
	return
}

// UpdateEmail 更新email
func (p *MemoryServer) UpdateEmail(ctx context.Context, in *proto.IDEmail) (out *google_protobuf2.Empty, err error) {
	return
}

// UpdateNickname 更新昵称
func (p *MemoryServer) UpdateNickname(ctx context.Context, in *proto.IDNickname) (out *google_protobuf2.Empty, err error) {
	return
}

// UpdateAvatar 更新头像
func (p *MemoryServer) UpdateAvatar(ctx context.Context, in *proto.IDAvatar) (out *google_protobuf2.Empty, err error) {
	return
}

// UpdateStatus 更新状态
func (p *MemoryServer) UpdateStatus(ctx context.Context, in *proto.IDStatus) (out *google_protobuf2.Empty, err error) {
	return
}

// InfoByID 通过id查询
func (p *MemoryServer) InfoByID(ctx context.Context, in *proto.ID) (out *proto.User, err error) {
	return
}

// InfoByPhone 通过手机查询
func (p *MemoryServer) InfoByPhone(ctx context.Context, in *proto.Phone) (out *proto.User, err error) {
	return
}

// InfoByEmail 通过email查询
func (p *MemoryServer) InfoByEmail(ctx context.Context, in *proto.Email) (out *proto.User, err error) {
	return
}

// InfoByNickname 通过Nickname查询
func (p *MemoryServer) InfoByNickname(ctx context.Context, in *proto.Nickname) (out *proto.User, err error) {
	return
}

// InfoByAccount 通过 手机、email、昵称查询
func (p *MemoryServer) InfoByAccount(ctx context.Context, in *proto.Account) (out *proto.User, err error) {
	return
}
