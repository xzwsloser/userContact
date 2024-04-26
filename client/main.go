package main

import (
	"fmt"
	"userContact/client/User"
)

func main() {
	// 接受用户的选择和菜单的显示
	var key int
	// 判断是否循环显示菜单
	var loop = true

	for loop {
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
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("你的输入有误,请重新选择")
		}
	}
	// 根据新的提示信息
	if key == 1 {
		// 登录聊天室
		var userId int
		var userPwd string
		fmt.Println("请输入用户的 ID 号")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入用户的密码")
		fmt.Scanf("%s\n", &userPwd)
		// 登录的业务 , 写到另外一个文件中
		err := User.Login(userId, userPwd)
		fmt.Println(err)
		if err != nil {
			// 失败
		} else {
			// 成功
		}
	}
}
