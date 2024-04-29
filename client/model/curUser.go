package clientmodel

import (
	"net"
	"userContact/server/model"
)

// 定义一个全局变量 curUser

type CurUser struct {
	Conn net.Conn
	model.User
}
