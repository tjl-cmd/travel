package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tjl-cmd/travel/user/internal/data"
)

type UserRepo interface {
	Login(username string, password string) (int64, error)
	Register(username string, password string) (int64, error)
	GetUserFromId(int64) (*data.User, error)
	GetUserFormByIds([]int64) ([]*data.User, error)
	UpdateUserInfo(*data.User) (int64, error)
}
type UserCase struct {
	log  *log.Helper
	repo UserRepo
}

func NewUserRepo(repo UserRepo, logger log.Logger) *UserCase {
	return &UserCase{repo: repo, log: log.NewHelper(logger)}
}
