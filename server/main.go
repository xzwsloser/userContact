package main

import (
	"encoding/json"
	"fmt"
	"net"
	"userContact/common"
	"userContact/utils"
)

// 根据客户端发送的请求的种类不同,调用不同函数, 相当于 Controller 层
func ServiceProcessMes(conn net.Conn, mes *common.Message) (err error) {
	switch mes.Type {
	case common.LoginMesType:
		// 登录业务逻辑
		ServiceProcessLogin(conn, mes)
	case common.RegisterMesType:
		// 处理登录的逻辑
	default:

		fmt.Println("消息类型不存在")
	}
	return
}

// 处理登录请求
func ServiceProcessLogin(conn net.Conn, mes *common.Message) (err error) {
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
	err = utils.WritePkg(conn, mesData)
	// 最后发送信息, 首先还是确定长度
	return
}

// 处理和客户端的通信
func Process(conn net.Conn) {
	// 用于读取客户端发送的信息
	// 演示关闭 conn
	defer conn.Close()
	// 循环读取客户发送的消息
	mes, err := utils.ReadPkg(conn)
	if err != nil {
		fmt.Println("获取对象失败, err =", err)
	}
	fmt.Println("mes =", mes)
	ServiceProcessMes(conn, &mes)
}

// 将读取数据包封装成一个函数

func main() {
	// 服务器端监听
	fmt.Println("服务器在8889端口进行监听")
	listen, err := net.Listen("tcp", "localhost:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("监听失败")
	}
	// 监听成功,等待客户端连接服务器
	for {
		fmt.Println("等待客户端连接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept Error = ", err)
		}
		// 连接成功就可以启动一个协程,则启动一个客户端保持通信
		go Process(conn)
	}
}
