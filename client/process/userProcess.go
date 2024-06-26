package userprocess

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"userContact/common"
	"userContact/server/model"
	"userContact/server/utils"
)

type UserProcess struct {
	Conn net.Conn
	// 暂时不用字段
}

func (up *UserProcess) Register(userid int, userpwd string, username string) (err error) {
	// 1. 连接到服务器

	if err != nil {
		fmt.Println("net.Dial error = ", err)
		return err
	}
	var mes common.Message
	mes.Type = common.RegisterMesType
	// 实例化对象信息,用于装入到结构体中
	var regiterMes common.RegisterMes
	regiterMes.User.UserName = username
	regiterMes.User.UserId = userid
	regiterMes.User.UserPwd = userpwd
	// 开始进行信息的序列化
	registerData, err := json.Marshal(regiterMes)
	if err != nil {
		fmt.Println("注册序列化失败")
		return err
	}
	mes.Data = string(registerData)
	// 开始把 mes 序列化
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("发送信息序列化失败")
		return err
	}
	// 开始发送长度,其实就是 WritePkg 方法
	// 开始创建一个
	tf := &utils.Transfer{
		Conn: up.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("writePkg wrong,err = ", err)
		return err

	}
	// 开始读一个包
	Newmes, err := tf.ReadPkg()
	// 返回的就是一个信息
	var registerResMes common.RegisterResMes
	err = json.Unmarshal([]byte(Newmes.Data), &registerResMes)
	if err != nil {
		fmt.Println("反序列化失败")
		return
	}
	// 问什么会打印失败呢
	if registerResMes.Code == 200 {
		//fmt.Println("登陆成功")
		fmt.Println("注册成功,请重新登录")
		os.Exit(0)
	} else if registerResMes.Code == 500 {
		fmt.Println(registerResMes.Error)
		os.Exit(0)

	}
	// 最后转换为一个对象
	os.Exit(0)
	return

}

// 完成登录的函数
func (up *UserProcess) Login(userId int, userPwd string) (err error) {
	// 判断登录是否成功,最好返回一个 error
	// 开始定一个协议

	// 首先连接服务器端, 一般用于读取配置文件

	// 序列化信息发送消息给服务端
	var mes common.Message
	mes.Type = common.LoginMesType

	var loginMes common.LoginMes
	// 用户信息
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	// 开始序列化
	userStr, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("用户信息序列化失败")
		return
	}
	mes.Data = string(userStr)
	// 最后把 mes 序列化
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("发送信息序列化失败")
		return
	}
	// 1. 首先发送长度
	// 但是发送数据时可以直接发送一个 len(data)
	// 由于发送的一个 write 方法发送的就是一个 byte 切片, 首先需要把 len(data) 转换为一个 []byte
	var pkgLen uint32
	// 发送一个字节数组
	pkgLen = uint32(len(data))
	// 首先定义一个 byte 切片
	var bytes [4]byte
	// 一个字节就是 8 位
	// 这里需要利用 数组创建切片
	binary.BigEndian.PutUint32(bytes[0:4], pkgLen)
	// 发送一个长度
	n, err := up.Conn.Write(bytes[0:4])
	if n != 4 || err != nil {
		fmt.Println("net.Dial , err =", err)
		return
	}
	// 演示关闭

	fmt.Println("客户端发送消息长度成功")
	// 发送消息本身
	_, err = up.Conn.Write(data)
	if err != nil {
		fmt.Println("发送消息失败, err =", err)
	}
	// 其实这里可以创建一个 全局变量 Tf 进行操作
	var tf utils.Transfer = utils.Transfer{
		Conn: up.Conn,
	}
	mes, err = tf.ReadPkg()
	fmt.Println(mes, err)
	// 开始反序列化成一个 LoginResMes
	if err != nil {
		fmt.Println("读取服务器信息出错")
		return
	}
	var loginResMes common.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if err != nil {
		fmt.Println("反序列化失败")
		return
	}
	// 问什么会打印失败呢
	if loginResMes.Code == 200 {
		CurUser.Conn = up.Conn
		CurUser.User.UserId = userId
		CurUser.User.UserStatus = common.UserOnline
		//fmt.Println("登陆成功")
		// 显示登录成功之后的一个菜单,但是需要利用循环显示
		// 登录成功, 之后可以显示在线用户的列表
		// 遍历一个 Userid
		for _, v := range loginResMes.UserIds {
			if v == userId {
				continue
			}

			// 完成一个全局变量的初始化
			user := &model.User{
				UserId:     v,
				UserStatus: common.UserOnline,
			}
			// 不需要其他信息
			// 其实就是初始信息
			onlineUsers[v] = user
		}
		fmt.Println("\n\n")
		go ServerProcess(up.Conn)
		for {

			ShowMenu() // 显示菜单
			// 这一个协程的作用就是不断读取信息
			// 这一个协程时刻监听服务器的响应,如果服务器有数据推送到客户端,并且显示在客户端
			// 这里还要开启一个协程
			// 完成一个初始化

		}
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	// 最后转换为一个对象
	return
}
