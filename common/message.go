package common

// 定义消息类型, 比如上限的消息
const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
	RegisterResType = "RegisterRes"
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
	Code  int    `json:"code"`  // 返回状态码
	Error string `json:"error"` // 返回错误信息
}

type RegisterMes struct {
}
