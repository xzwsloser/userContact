package userprocess

import (
	"encoding/json"
	"fmt"
	"userContact/common"
	"userContact/server/utils"
)

// 转发信息
type SmsProcess struct {
}

// 发送群聊的消息
func (this *SmsProcess) SendGroupMes(context string) (err error) {
	// 开始发送消息
	//1. 创建一个 mes
	var mes common.Message
	mes.Type = common.SmsMesType
	// 2. 创建一个 SmsMes 对象
	var smsMes common.SmsMes
	smsMes.Context = context
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus
	// 开始序列化sms
	smsData, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("序列化失败")
		return err
	}
	mes.Data = string(smsData)
	// 对于 mes 序列化
	mesData, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化失败")
		return err
	}
	// 把 mes 发送给服务器端
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(mesData)
	if err != nil {
		fmt.Println("信息发送失败 , err =", err)
		return err
	}
	return nil
}
