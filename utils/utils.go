package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"userContact/common"
)

func ReadPkg(conn net.Conn) (mes common.Message, err error) {
	for {
		buf := make([]byte, 1024*4)
		n, err := conn.Read(buf[:4]) // 注意读取的长度
		if err != nil {
			fmt.Println(err)
			err = errors.New("read pkg header error")
			return mes, err
		}
		// 开始服务端接收信息,判断用户消息的合法性并且判断信息的合理性
		// 根据读取到的长度尽心一个转换
		var pkgLen uint32
		pkgLen = binary.BigEndian.Uint32(buf[:4]) // 就是把切片解析成一个数字,就是一个反向的转换
		// 根据一个 pkgLen 读取消息内容
		n, err = conn.Read(buf[:int(pkgLen)])
		if n != int(pkgLen) || err != nil {
			fmt.Println("read pkg body error")
			//err = errors.New()
			return mes, err
		}
		// 此时值已经被写入到 buf 中
		// 开始把 pkgLen  反序列化
		// 可以使用返回值中的一个参数
		err = json.Unmarshal(buf[:int(pkgLen)], &mes) // 注意之后的一个 &msg , 地址传递才可以改变值
		if err != nil {
			fmt.Println("反序列化失败 , err =", err)
		}

		return mes, nil
	}
}

// 定义一个工具函数,封装写包的一个函数
func WritePkg(conn net.Conn, mesData []byte) (err error) {
	// 发送一个一个长度
	var pkgLen uint32
	pkgLen = uint32(len(mesData))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("发送长度失败")
		return
	}
	// 发送 data 本身
	n, err = conn.Write(mesData)
	if n != int(pkgLen) || err != nil {
		fmt.Println("发送数据长度不正确")
		return
	}

	return
}
