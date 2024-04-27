package main

import (
	"fmt"
	"net"
	"userContact/common"
	"userContact/server/process"
	"userContact/server/utils"
)

type Processor struct {
	Conn net.Conn
}

// 根据客户端发送的请求的种类不同,调用不同函数, 相当于 Controller 层
// 这一个相当于一个主控
func (this *Processor) ServiceProcessMes(mes *common.Message) (err error) {
	switch mes.Type {
	case common.LoginMesType:
		// 登录业务逻辑
		var userProcess process.UserProcess = process.UserProcess{
			Conn: this.Conn,
		}
		err := userProcess.ServiceProcessLogin(mes)
		if err != nil {
			fmt.Println("总控处理注册逻辑错误")
		}
	case common.RegisterMesType:
		// 处理登录的逻辑
		var userProcess process.UserProcess = process.UserProcess{
			Conn: this.Conn,
		}
		err := userProcess.ServiceProcessRegister(mes)
		if err != nil {
			fmt.Println("总控处理注册逻辑错误")
		}
		// 创建一个方法处理注册请求
		// 类似上面的一个方法

	default:

		fmt.Println("消息类型不存在")
	}
	return
}

func (this *Processor) process2() (err error) {
	// 循环读取客户发送的消息
	// 创建 Transfer 对象
	var tf utils.Transfer = utils.Transfer{
		Conn: this.Conn,
	}
	mes, err := tf.ReadPkg()
	if err != nil {
		fmt.Println("获取对象失败, err =", err)
		return err // 条件判断中的 return 一定需要利用实例对象传出
	}
	fmt.Println("mes =", mes)
	// 创建对象
	// 这个其实就是处理对象
	err = this.ServiceProcessMes(&mes)
	if err != nil {
		fmt.Println("处理信息出错")
		return err
	}
	return
}
