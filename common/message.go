package common

// 定义消息类型, 比如上限的消息
const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
	RegisterResType = "RegisterRes"
)

// 定义客户端发送的结构体
type Message struct {
	Type string `json:"type"` // 消息类型
	Data string `json:"data"` // 消息对象
}

type LoginMes struct {
	UserId   int    `json:"userid"`   // ID
	UserPwd  string `json:"userpwd"`  // 密码
	userName string `json:"username"` // 用户名
}

type LoginResMes struct {
	Code  int    `json:"code"`  // 返回状态码
	Error string `json:"error"` // 返回错误信息
}

type RegisterMes struct {
}
