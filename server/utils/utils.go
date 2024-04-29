package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"userContact/common"
)

// 定义一个 Transfer
type Transfer struct {
	// 分析字段
	Conn net.Conn
	// 就是函数的参数
	Buf [8096]byte // 使用切片使用, 其实可以反复利用,就是把和一个层关联的方法中共同的属性中抽象出来
}

func (this *Transfer) ReadPkg() (mes common.Message, err error) {
	for {
		n, err := this.Conn.Read(this.Buf[:4]) // 注意读取的长度
		if err != nil {
			fmt.Println(err)
			err = errors.New("read pkg header error")
			return mes, err
		}
		// 开始服务端接收信息,判断用户消息的合法性并且判断信息的合理性
		// 根据读取到的长度尽心一个转换
		var pkgLen uint32
		pkgLen = binary.BigEndian.Uint32(this.Buf[:4]) // 就是把切片解析成一个数字,就是一个反向的转换
		// 根据一个 pkgLen 读取消息内容
		n, err = this.Conn.Read(this.Buf[:int(pkgLen)])
		if n != int(pkgLen) || err != nil {
			fmt.Println("read pkg body error")
			//err = errors.New()
			return mes, err
		}
		// 此时值已经被写入到 buf 中
		// 开始把 pkgLen  反序列化
		// 可以使用返回值中的一个参数
		err = json.Unmarshal(this.Buf[:int(pkgLen)], &mes) // 注意之后的一个 &msg , 地址传递才可以改变值
		if err != nil {
			fmt.Println("反序列化失败 , err =", err)
		}

		return mes, nil
	}
}

// 定义一个工具函数,封装写包的一个函数
func (this *Transfer) WritePkg(mesData []byte) (err error) {
	// 发送一个一个长度
	var pkgLen uint32
	pkgLen = uint32(len(mesData))

	binary.BigEndian.PutUint32(this.Buf[:4], pkgLen)
	n, err := this.Conn.Write(this.Buf[:4])
	fmt.Println(err)
	if n != 4 || err != nil {
		fmt.Println("发送长度失败")
		return
	}
	// 发送 data 本身
	n, err = this.Conn.Write(mesData)
	if n != int(pkgLen) || err != nil {
		fmt.Println("发送数据长度不正确")
		return
	}
	return
}
