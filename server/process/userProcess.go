package process

import (
	"encoding/json"
	"fmt"
	"net"
	"userContact/common"
	"userContact/server/utils"
)

// 处理用户需要的字段
type UserProcess struct {
	Conn net.Conn
}

// 处理登录请求
func (this *UserProcess) ServiceProcessLogin(mes *common.Message) (err error) {
	// 取出一个Data , 进行相应的判断,发送相应的信息
	var loginMes common.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fall err = ", err)
		return err
	}
	// 定义一个返回类型, 注意 Message 统一用于封装
	var resMes common.Message
	resMes.Type = common.LoginResMesType
	// 定义一个返回结构
	var loginResMeg common.LoginResMes

	// 数据库中查找用户的逻辑
	if loginMes.UserId == 100 && loginMes.UserPwd == "loser" {
		// 合法
		loginResMeg.Code = 200

	} else {
		// 不合法
		loginResMeg.Code = 500 // 用户不存在
		loginResMeg.Error = "用户不存在,请注册之后再使用..."
	}
	// 完成序列化
	// 1. 首先序列化 loginResMsg
	data, err := json.Marshal(loginResMeg)
	if err != nil {
		fmt.Println("响应结构体序列化失败")
		return
	}
	// 2. 开始序列化另外一个
	resMes.Data = string(data)
	mesData, err := json.Marshal(resMes)
	if err != nil {
		fmt.Println("响应信息序列化失败")
	}
	// 得到的就是响应信息
	// err = utils.WritePkg(conn, mesData)
	// 首先创建一个实例
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(mesData)
	// 最后发送信息, 首先还是确定长度
	return
}
