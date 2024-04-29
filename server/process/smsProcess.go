package process

import (
	"encoding/json"
	"fmt"
	"net"
	"userContact/common"
	"userContact/server/utils"
)

type SmsProcess struct {
	// 展示不用字段
}

// 1. 转发消息
func (this *SmsProcess) SendGroupMes(mes *common.Message) {
	// 开始遍历服务器端的一个 onlineUser , 将消息转发
	// 开始取出内容
	var smsmes common.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &mes)
	if err != nil {
		fmt.Println("反序列化失败,err = ", err)
	}
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化失败,err = ", err)
	}
	for id, up := range UserMgrObj.onlineUses {
		if id == smsmes.UserId {
			continue
		}
		this.SendMesToEachOnlineUser(data, up.Conn)
	}
}

// 发送信息的函数
func (this *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {
	// 开始发送消息给每一个客户端
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败,err = ", err)
	}
}
