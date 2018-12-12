package user

import (
	"context"
	"sync"
	"time"

	google_protobuf2 "github.com/gogo/protobuf/types"
	proto "github.com/weisd/web-kit/api/protobuf/user"
	"github.com/weisd/web-kit/internal/pkg/ierrors"
	"github.com/weisd/web-kit/internal/pkg/istring"
)

// var _ proto.RPCServiceServer = &MemoryServer{}
var _ RPCServer = &MemoryServer{}

func init() {
	RPCRegister("memory", &MemoryServer{})
}

// MemoryServer MemoryServer
type MemoryServer struct {
	data  map[int64]*proto.User
	lock  *sync.RWMutex
	idGen int64
}

// Init Init
func (p *MemoryServer) Init(ctx context.Context, dsn string) error {
	p.data = make(map[int64]*proto.User)
	p.lock = new(sync.RWMutex)
	return nil
}

// Create 创建用户
func (p *MemoryServer) Create(ctx context.Context, in *proto.User) (out *google_protobuf2.Empty, err error) {
	out = &google_protobuf2.Empty{}
	p.lock.Lock()
	p.idGen++
	in.ID = p.idGen
	in.Status = proto.UserStatus_Normal
	in.CreatedAt = time.Now()
	in.UpdatedAt = in.CreatedAt
	p.data[in.ID] = in
	p.lock.Unlock()
	return
}

// UpdatePassword 更新密码
func (p *MemoryServer) UpdatePassword(ctx context.Context, in *proto.IDPassword) (out *google_protobuf2.Empty, err error) {
	out = &google_protobuf2.Empty{}
	p.lock.Lock()
	info, has := p.data[in.ID]
	if !has {
		p.lock.Unlock()
		return
	}

	info.Password = in.Password
	info.Salt = in.Salt
	p.lock.Unlock()
	return
}

// UpdatePhone  更新手机号
func (p *MemoryServer) UpdatePhone(ctx context.Context, in *proto.IDPhone) (out *google_protobuf2.Empty, err error) {
	out = &google_protobuf2.Empty{}
	p.lock.Lock()
	info, has := p.data[in.ID]
	if !has {
		p.lock.Unlock()
		return
	}

	info.Phone = in.Phone
	p.lock.Unlock()
	return
}

// UpdateEmail 更新email
func (p *MemoryServer) UpdateEmail(ctx context.Context, in *proto.IDEmail) (out *google_protobuf2.Empty, err error) {
	out = &google_protobuf2.Empty{}
	p.lock.Lock()
	info, has := p.data[in.ID]
	if !has {
		p.lock.Unlock()
		return
	}

	info.Email = in.Email
	p.lock.Unlock()
	return
}

// UpdateNickname 更新昵称
func (p *MemoryServer) UpdateNickname(ctx context.Context, in *proto.IDNickname) (out *google_protobuf2.Empty, err error) {
	out = &google_protobuf2.Empty{}
	p.lock.Lock()
	info, has := p.data[in.ID]
	if !has {
		p.lock.Unlock()
		return
	}

	info.NickName = in.NickName
	p.lock.Unlock()
	return
}

// UpdateAvatar 更新头像
func (p *MemoryServer) UpdateAvatar(ctx context.Context, in *proto.IDAvatar) (out *google_protobuf2.Empty, err error) {
	out = &google_protobuf2.Empty{}
	p.lock.Lock()
	info, has := p.data[in.ID]
	if !has {
		p.lock.Unlock()
		return
	}

	info.Avatar = in.Avatar
	p.lock.Unlock()
	return
}

// UpdateStatus 更新状态
func (p *MemoryServer) UpdateStatus(ctx context.Context, in *proto.IDStatus) (out *google_protobuf2.Empty, err error) {
	out = &google_protobuf2.Empty{}
	p.lock.Lock()
	info, has := p.data[in.ID]
	if !has {
		p.lock.Unlock()
		return
	}

	info.Status = in.Status
	p.lock.Unlock()
	return
}

// InfoByID 通过id查询
func (p *MemoryServer) InfoByID(ctx context.Context, in *proto.ID) (out *proto.User, err error) {
	out = &proto.User{}
	p.lock.RLock()
	info, has := p.data[in.ID]
	if !has {
		p.lock.RUnlock()
		return
	}

	*out = *info
	p.lock.RUnlock()
	return
}

// InfoByPhone 通过手机查询
func (p *MemoryServer) InfoByPhone(ctx context.Context, in *proto.Phone) (out *proto.User, err error) {
	out = &proto.User{}
	p.lock.RLock()
	for _, v := range p.data {
		if v.Phone == in.Phone {
			*out = *v
			break
		}
	}
	p.lock.RUnlock()
	return
}

// InfoByEmail 通过email查询
func (p *MemoryServer) InfoByEmail(ctx context.Context, in *proto.Email) (out *proto.User, err error) {
	out = &proto.User{}
	p.lock.RLock()
	for _, v := range p.data {
		if v.Email == in.Email {
			*out = *v
			break
		}
	}
	p.lock.RUnlock()
	return
}

// InfoByNickname 通过Nickname查询
func (p *MemoryServer) InfoByNickname(ctx context.Context, in *proto.Nickname) (out *proto.User, err error) {
	out = &proto.User{}
	p.lock.RLock()
	for _, v := range p.data {
		if v.NickName == in.Nickname {
			*out = *v
			break
		}
	}
	p.lock.RUnlock()
	return
}

// InfoByAccount 通过 手机、email、昵称查询
func (p *MemoryServer) InfoByAccount(ctx context.Context, in *proto.Account) (out *proto.User, err error) {
	switch {
	case istring.IsPhone(in.Account):
		return p.InfoByPhone(ctx, &proto.Phone{Phone: in.Account})
	case istring.IsEmail(in.Account):
		return p.InfoByEmail(ctx, &proto.Email{Email: in.Account})
	case istring.IsValidNickname(in.Account):
		return p.InfoByNickname(ctx, &proto.Nickname{Nickname: in.Account})
	default:
		err = ierrors.ErrInValidNickname
		return
	}
	return
}
