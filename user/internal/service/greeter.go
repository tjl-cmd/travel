package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/afocus/captcha"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jakehl/goid"
	"github.com/tjl-cmd/travel/api/proto/user"
	v1 "github.com/tjl-cmd/travel/api/proto/user"
	"github.com/tjl-cmd/travel/user/internal/biz"
	"github.com/tjl-cmd/travel/user/internal/data"
	"image/png"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedUserServer
	uc  *biz.GreeterUsecase
	log *log.Helper
}

func (s *GreeterService) Login(ctx context.Context, msg *user.LoginInfoMsg) (*user.LoginInfo, error) {
	code, err := data.Redispool.Get().Do("get", msg.CodeId)
	if err != nil {
		return nil, err
	}
	if fmt.Sprintf("%s", code) != msg.Code {
		return nil, errors.New("验证码错误")
	}
	if msg.Username == "" || msg.Password == "" {
		return nil, errors.New("账户或密码不能为空")
	}
}

func (s *GreeterService) Register(ctx context.Context, registry *user.Registry) (*user.Response, error) {
	if registry.Username == "" || registry.Password == "" {
		return nil, errors.New("账户或密码不能为空")
	}
}

func (s *GreeterService) UpdateUser(ctx context.Context, info *user.UpdateUserInfo) (*user.Response, error) {
	panic("implement me")
}

func (s *GreeterService) Logout(ctx context.Context, info *user.LogoutInfo) (*user.Empty, error) {
	panic("implement me")
}

func (s *GreeterService) CheckToken(ctx context.Context, info *user.LogoutInfo) (*user.CheckTokenResp, error) {
	panic("implement me")
}

func (s *GreeterService) GetUserInfoById(ctx context.Context, id *user.GetUserByID) (*user.UserResp, error) {
	panic("implement me")
}

func (s *GreeterService) GetUserInfoByIds(ctx context.Context, ds *user.GetUserByIDs) (*user.UserRespS, error) {
	panic("implement me")
}

func (s *GreeterService) GreeterAuthCode(ctx context.Context, empty *user.Empty) (*user.CodeResp, error) {
	cap := captcha.New()
	if err := cap.SetFont("comic.ttf"); err != nil {
		return nil, err
	}
	imag, str := cap.Create(4, captcha.NUM)
	emptyBuff := bytes.NewBuffer(nil)
	err := png.Encode(emptyBuff, imag)
	if err != nil {
		return nil, err
	}
	base64Str := "data:image/png;base64," + base64.StdEncoding.EncodeToString(emptyBuff.Bytes())
	codeId := goid.NewV4UUID().String()
	_, err = data.Redispool.Get().Do("SET", codeId, str, "EX", "300")
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	res := &user.CodeResp{
		CodeId: codeId,
		Images: base64Str,
	}
	return res, nil
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase, logger log.Logger) *GreeterService {
	return &GreeterService{uc: uc, log: log.NewHelper(logger)}
}
