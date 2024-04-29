package userprocess

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"userContact/common"
	"userContact/server/utils"
)

// 显示登录成功的见面
// 开启协程保持和服务器端的联系

// 显示登录成功之后的界面
func ShowMenu() {
	fmt.Println("-----------恭喜  登录成功--------")
	fmt.Println("-------------1. 显示用户列表------------")
	fmt.Println("-------------2. 发送消息------------")
	fmt.Println("-------------3. 信息列表----------------")
	fmt.Println("-------------4. 退出系统------------")
	fmt.Println("请选择(1-4)")
	var key int
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		// 显示在用户列表
		outputOnlineUser()
	case 2:
		// 发送消息
	case 3:
		// 查看信息列表
	case 4:
		// 退出系统
		os.Exit(0) // 直接退出
	default:
		fmt.Println("输出错误,请重新输入")
	}
}

func ServerProcess(conn net.Conn) {
	// 创建一个读取的结构体
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		// 保持通信, 注意这里会发生一个阻塞
		// 此时就会阻塞在这里
		// 注意里面的Read 函数们没有读取到数据不会报错,而是等待一个特定的时间之后就会报错
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err=", err)
			return
		}
		// 读取到消息,进行下一步逻辑处理
		// 开始处理消息
		switch mes.Type {
		case common.NotifyUserStatusMesType: // 上线消息
			// 取出 mes.Data
			// 把 当前用户的状态保存到用户的 map 中 map[int]User
			// 开始处理信息
			// 把状态保存到用户列表
			// 编写方法处理返回的信息
			var notifyUserStatusMes common.NotifyUserStatusMes
			err := json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			if err != nil {
				fmt.Println("返回信息序列化失败")
			}
			updateUserStatus(&notifyUserStatusMes)

		default:
			// 无法处理的消息
			fmt.Println("服务器发送协程无法处理的消息")
		}
	}
}
