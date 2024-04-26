package main

import (
	"fmt"
	"net"
)

// 处理和客户端的通信,哪一个连接启动一个协程,可以
func Process(conn net.Conn) {
	// 用于读取客户端发送的信息
	// 演示关闭 conn
	defer conn.Close()
	// 调用主控
	processor := Processor{
		Conn: conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("客户端和服务端协程出错, err =", err)
		return
	}
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
