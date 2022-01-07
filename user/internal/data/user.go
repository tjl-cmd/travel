package data

import (
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tjl-cmd/travel/common"
	"time"
)

type User struct {
	Id          int64     `json:"id" gorm:"primary_key;not_null;auto_increment"`
	Username    string    `json:"username" gorm:"unique;not_null"` // 用户登录账户
	Password    string    `json:"password" gorm:"not_null"`        // 用户密码
	PwdScrypt   string    `json:"pwd_scrypt"`                      // 加密秘钥
	Birthday    time.Time `json:"birthday"`                        // 用户生日
	Sex         int       `json:"sex"`                             // 用户性别
	Phone       string    `json:"phone"`                           // 用户手机号
	Avatar      string    `json:"avatar"`                          // 用户头像
	Nickname    string    `json:"nickname"`                        // 用户昵称
	Card        string    `json:"card"`                            // 用户身份证号
	CardEncrypt string    `json:"card_encrypt"`                    // 用户身份证加密
	Email       string    `json:"email"`                           // 用户邮箱
	Name        string    `json:"name"`                            // 用户真实姓名
	Status      int       `json:"status"`                          // 用户状态

}

func (u User) TableName() string {
	return "user"
}

type userRepo struct {
	data *Data
	log  *log.Helper
}

func (u *userRepo) Login(username string, password string) (int64, error) {
	var user User
	u.data.db.Where("username = ?", username).Find(&user)
	if user.Id == 0 {
		return 0, errors.New("当前用户不存在")
	}
	if user.Password != common.AesDecrypt(password, user.PwdScrypt) {
		return 0, errors.New("账户或密码错误")
	}
	return user.Id, nil
}
func (u *userRepo) Register(username string, password string) (int64, error) {
	var user User
	u.data.db.Where("username = ?", username).Find(&user)
	if user.Id > 0 {
		return 0, errors.New("账户已存在")
	}
	user.Username = username
	user.PwdScrypt = common.RandString(10)
	user.Password = common.AesEncrypt(password, user.PwdScrypt)
	err := u.data.db.Create(&user).Error
	if err != nil {
		u.log.Debugf("Create user err=%s", err.Error())
		return 0, err
	}
	return user.Id, nil
}
func (u *userRepo) GetUserFromId(id int64) (user *User, err error) {
	return user, u.data.db.Where("id = ?", id).Find(&user).Error
}
func (u *userRepo) GetUserFormByIds(ids []int64) (users []*User, err error) {
	return users, u.data.db.Where("id in (?)", ids).Find(&users).Error
}
func (u *userRepo) UpdateUserInfo(*User) (int64, error) {
	return 0, nil
}
