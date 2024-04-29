package common

import "userContact/server/model"

// 定义用户在线状态的常量

// 定义消息类型, 比如上限的消息
const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

// 这里注意结构体之间的包含关系,首先第一个结构体 Message 结构体中的第二个字段其实就是相应的对象序列化之后的结构
// 第一个字段就是类型,传递消息时,首先序列化第二个字段,成为第一个结构体的一个元素,在发送 Message 即可
// 定义客户端发送的结构体
type Message struct {
	Type string `json:"type"` // 消息类型
	Data string `json:"data"` // 消息对象
}

type LoginMes struct {
	UserId   int    `json:"userid"`   // ID
	UserPwd  string `json:"userpwd"`  // 密码
	UserName string `json:"username"` // 用户名
}

type LoginResMes struct {
	Code    int    `json:"code"` // 返回状态码
	UserIds []int  // 返回用户ID 号地址							// 增加一个用户切片
	Error   string `json:"error"` // 返回错误信息
}

// 相当于一个继承关系
type RegisterMes struct {
	User model.User `json:"user"` // 注册信息类型
}

// 注册相应的状态码
type RegisterResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

// 定义一个传递状态的结构体
// 为了配合服务器端推送上线通知,创建新的类型
// 可以利用客户端和服务器端的交换协程进行作用
type NotifyUserStatusMes struct {
	UserId int `json:"userid"`
	Status int `json:"userstatus"`
}

// 新增一个消息结构体
type SmsMes struct {
	Context    string `json:"context"`
	model.User        // 匿名结构体,相当于一个内嵌结构体
}
