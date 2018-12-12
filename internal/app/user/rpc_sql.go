package user

import (
	"context"
	"strings"

	google_protobuf2 "github.com/gogo/protobuf/types"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	proto "github.com/weisd/web-kit/api/protobuf/user"

	_ "github.com/go-sql-driver/mysql"
	"github.com/weisd/web-kit/internal/pkg/ierrors"
	"github.com/weisd/web-kit/internal/pkg/imodel"
	"github.com/weisd/web-kit/internal/pkg/istring"
)

// var _ proto.RPCServiceServer = &SQLServer{}
var _ RPCServer = &SQLServer{}

func init() {
	RPCRegister("sql", &SQLServer{})
}

// SQLServer SQLServer
type SQLServer struct {
	db *gorm.DB
}

// Init Init
func (p *SQLServer) Init(ctx context.Context, dsn string) error {
	idx := strings.Index(dsn, ":")
	if idx == -1 {
		return errors.New("invalid dsn format")
	}

	cancel, db, err := imodel.NewGromDB(string(dsn[:idx]), string(dsn[idx+1:]))
	if err != nil {
		return errors.Wrap(err, "NewGromDB")
	}

	p.db = db

	// 创建表
	p.initTables()

	go func() {
		select {
		case <-ctx.Done():
			cancel()
		}
	}()

	return nil
}

func (p *SQLServer) initTables() {
	if !p.db.HasTable(&proto.User{}) {
		p.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARACTER SET=utf8").AutoMigrate(&proto.User{})
		p.db.Model(&proto.User{}).AddIndex("idx_user_phone", "phone")
	}
}

// Create CreateCreate
func (p *SQLServer) Create(ctx context.Context, in *proto.User) (out *google_protobuf2.Empty, err error) {

	if len(in.Password) > 0 {
		in.Salt = istring.RandString(4)
		in.Password = istring.HashPassword(in.Password, in.Salt)
	}

	db := p.db.Create(in)

	return &google_protobuf2.Empty{}, db.Error
}

// UpdatePassword 更新密码
func (p *SQLServer) UpdatePassword(ctx context.Context, in *proto.IDPassword) (out *google_protobuf2.Empty, err error) {

	if len(in.Password) > 0 {
		in.Salt = istring.RandString(4)
		in.Password = istring.HashPassword(in.Password, in.Salt)
	}

	db := p.db.Model(&proto.User{ID: in.ID}).Updates(proto.User{
		Password: in.Password,
		Salt:     in.Salt,
	})

	return &google_protobuf2.Empty{}, db.Error
}

// UpdatePhone  更新手机号
func (p *SQLServer) UpdatePhone(ctx context.Context, in *proto.IDPhone) (out *google_protobuf2.Empty, err error) {
	db := p.db.Model(&proto.User{ID: in.ID}).Updates(proto.User{
		Phone: in.Phone,
	})

	return &google_protobuf2.Empty{}, db.Error
}

// UpdateEmail 更新email
func (p *SQLServer) UpdateEmail(ctx context.Context, in *proto.IDEmail) (out *google_protobuf2.Empty, err error) {
	db := p.db.Model(&proto.User{ID: in.ID}).Updates(proto.User{
		Email: in.Email,
	})

	return &google_protobuf2.Empty{}, db.Error
}

// UpdateNickname 更新昵称
func (p *SQLServer) UpdateNickname(ctx context.Context, in *proto.IDNickname) (out *google_protobuf2.Empty, err error) {
	db := p.db.Model(&proto.User{ID: in.ID}).Updates(proto.User{
		NickName: in.NickName,
	})

	return &google_protobuf2.Empty{}, db.Error
}

// UpdateAvatar 更新头像
func (p *SQLServer) UpdateAvatar(ctx context.Context, in *proto.IDAvatar) (out *google_protobuf2.Empty, err error) {
	db := p.db.Model(&proto.User{ID: in.ID}).Updates(proto.User{
		Avatar: in.Avatar,
	})

	return &google_protobuf2.Empty{}, db.Error
}

// UpdateStatus 更新状态
func (p *SQLServer) UpdateStatus(ctx context.Context, in *proto.IDStatus) (out *google_protobuf2.Empty, err error) {
	db := p.db.Model(&proto.User{ID: in.ID}).Update("status", in.Status)

	return &google_protobuf2.Empty{}, db.Error
}

// InfoByID 通过id查询
func (p *SQLServer) InfoByID(ctx context.Context, in *proto.ID) (out *proto.User, err error) {
	out = &proto.User{}
	db := p.db.First(out, in.ID)
	return out, db.Error
}

// InfoByPhone 通过手机查询
func (p *SQLServer) InfoByPhone(ctx context.Context, in *proto.Phone) (out *proto.User, err error) {
	out = &proto.User{}
	db := p.db.Where(&proto.User{Phone: in.Phone}).First(out)
	return out, db.Error
}

// InfoByEmail 通过email查询
func (p *SQLServer) InfoByEmail(ctx context.Context, in *proto.Email) (out *proto.User, err error) {
	out = &proto.User{}
	db := p.db.Where(&proto.User{Email: in.Email}).First(out)
	return out, db.Error
}

// InfoByNickname 通过nickname查询
func (p *SQLServer) InfoByNickname(ctx context.Context, in *proto.Nickname) (out *proto.User, err error) {
	out = &proto.User{}
	db := p.db.Where(&proto.User{NickName: in.Nickname}).First(out)
	return out, db.Error
}

// InfoByAccount 通过 手机、email、昵称查询
func (p *SQLServer) InfoByAccount(ctx context.Context, in *proto.Account) (out *proto.User, err error) {
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
}
