package model

import "errors"

// 根据业务逻辑中出现的错误,做成一个全局变量,应为返回变量的函数不可以用于定义常量
var (
	// 定义错误
	ERROR_USER_NOTEXISTS = errors.New("用户不存在")
	ERROR_USER_EXISTS    = errors.New("用户已经存在")
	ERROR_USER_PWD       = errors.New("用户密码不存在")
)
