package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"userContact/common"
)

func readPkg(conn net.Conn) (mes common.Message, err error) {
	for {
		buf := make([]byte, 1024*4)
		n, err := conn.Read(buf[:4]) // 注意读取的长度
		if n != 4 || err != nil {
			err = errors.New("read pkg header error")
			return common.Message{}, err
		}
		// 开始服务端接收信息,判断用户消息的合法性并且判断信息的合理性
		// 根据读取到的长度尽心一个转换
		var pkgLen uint32
		pkgLen = binary.BigEndian.Uint32(buf[:4]) // 就是把切片解析成一个数字,就是一个反向的转换
		// 根据一个 pkgLen 读取消息内容
		n, err = conn.Read(buf[:pkgLen])
		if n != int(pkgLen) || err != nil {
			err = errors.New("read pkg body error")
			return common.Message{}, err
		}
		// 此时值已经被写入到 buf 中
		// 开始把 pkgLen  反序列化
		// 可以使用返回值中的一个参数
		err = json.Unmarshal(buf[:pkgLen], &mes) // 注意之后的一个 &msg , 地址传递才可以改变值
		if err != nil {
			fmt.Println("反序列化失败 , err =", err)
		}
		return mes, err
	}
}

// 处理和客户端的通信
func Process(conn net.Conn) {
	// 用于读取客户端发送的信息
	// 演示关闭 conn
	defer conn.Close()
	// 循环读取客户发送的消息
	mes, err := readPkg(conn)
	if err != nil {
		fmt.Println("获取对象失败, err =", err)
	}
	fmt.Println("mes =", mes)
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
