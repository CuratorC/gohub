// Package auth 授权相关逻辑
package auth

import (
	"errors"
	"gohub/app/models/user"
)

// Attempt 尝试登录
func Attempt(email string, password string) (user.User, error) {
	_user := user.GetByMulti(email)
	if _user.ID == 0 {
		return user.User{}, errors.New("账号不存在")
	}

	if !_user.ComparePassword(password) {
		return user.User{}, errors.New("密码错误")
	}

	return _user, nil
}

// LoginByPhone 登录指定用户
func LoginByPhone(phone string) (user.User, error) {
	_user := user.GetByPhone(phone)
	if _user.ID == 0 {
		return user.User{}, errors.New("手机号未注册")
	}

	return _user, nil
}
