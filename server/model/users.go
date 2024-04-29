package model

// 其实 model 层就是结构体的定义和结构体方法的定义

// 1. 定义用户结构体
type User struct {
	// 绑定对象时,根据标签值判断哪一个字段的值赋值给哪一个字段
	// 和数据库中的字段名称一致
	UserId     int    `json:"userid"`
	UserPwd    string `json:"userpwd"` // 序列化或者反序列化成功
	UserName   string `json:"username"`
	UserStatus int    `json:"userstatus"`
}
