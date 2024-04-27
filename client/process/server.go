package userprocess

import (
	"fmt"
	"net"
	"os"
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
			fmt.Println("协程读取服务器端信息失败")
			return
		}
		// 读取到消息,进行下一步逻辑处理
		fmt.Println(mes)
	}
}
