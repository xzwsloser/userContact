package userprocess

import (
	"encoding/json"
	"fmt"
	"userContact/common"
)

func outputGroupMes(mes *common.Message) {
	// 开始显示消息
	// 首先序列化
	var smsMes common.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("反序列化失败 ,err = ", err)
		return
	}
	// 开始显示内容
	content := smsMes.Context
	// 显示信息格式化
	info := fmt.Sprintf("用户id\t 对大家说:\t%s", smsMes.UserId, content)
	fmt.Println(info)
	fmt.Println()
}
