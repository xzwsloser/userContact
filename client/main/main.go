package main

import (
	"fmt"
	"net"
	userprocess "userContact/client/process"
)

func main() {
	// 首先书初始化 conn
	conn, err := net.Dial("tcp", "localhost:8889")
	defer conn.Close()
	if err != nil {
		fmt.Println("获取连接失败")
		return
	}
	// 接受用户的选择和菜单的显示
	var key int
	// 判断是否循环显示菜单
	//var loop = true
	// 这里其实就是一个二级菜单
	for {
		fmt.Println("============欢迎来到多人聊天系统================")
		fmt.Println("\t\t\t 1 登录聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择(1-3)")

		fmt.Scanf("%d\n", &key)
		// 根据选择进入功能
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			var userId int
			var userPwd string
			fmt.Println("请输入用户的 ID 号")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n", &userPwd)
			var up *userprocess.UserProcess
			up = &userprocess.UserProcess{
				Conn: conn,
			}
			err := up.Login(userId, userPwd)
			if err != nil {
				fmt.Println("登录的业务逻辑有误")
			}
			//loop = false
		case 2:
			fmt.Println("注册用户")
			var userId int
			var userPwd string
			var userName string
			fmt.Println("请输入用户的 ID 号")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户的昵称")
			fmt.Scanf("%s\n", &userName)
			// 调用一个 UserProcess 示例中的方法
			up := &userprocess.UserProcess{
				Conn: conn,
			}
			err := up.Register(userId, userPwd, userName)
			if err != nil {
				fmt.Println("注册失败请重试")
			}
			//loop = false
		case 3:
			fmt.Println("退出系统")
			//loop = false
		default:
			fmt.Println("你的输入有误,请重新选择")
		}
	}
	// 根据新的提示信息
	//if key == 1 {
	//	// 登录聊天室
	//
	//
	//
	//
	//
	//
	//	// 登录的业务 , 写到另外一个文件中
	//	err := process.Login(userId, userPwd)
	//	fmt.Println(err)
	//	if err != nil {
	//		// 失败
	//	} else {
	//		// 成功
	//	}
	//}

}
