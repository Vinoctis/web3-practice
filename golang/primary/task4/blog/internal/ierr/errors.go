package ierr 

import (
	"errors"
)

var (
	ErrUserNotFound = errors.New("用户不存在")
	ErrPassword     = errors.New("密码有误")
)