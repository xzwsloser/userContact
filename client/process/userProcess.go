package userprocess

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"userContact/common"
	"userContact/server/utils"
)

type UserProcess struct {
	// 暂时不用字段
}

// 完成登录的函数
func (up *UserProcess) Login(userId int, userPwd string) (err error) {
	// 判断登录是否成功,最好返回一个 error
	// 开始定一个协议

	// 首先连接服务器端, 一般用于读取配置文件
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial error =", err)
		return
	}
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
	n, err := conn.Write(bytes[0:4])
	if n != 4 || err != nil {
		fmt.Println("net.Dial , err =", err)
		return
	}
	// 演示关闭
	defer conn.Close()
	fmt.Println("客户端发送消息长度成功")
	// 发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("发送消息失败, err =", err)
	}
	// 其实这里可以创建一个 全局变量 Tf 进行操作
	var tf utils.Transfer = utils.Transfer{
		Conn: conn,
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
	if loginResMes.Code == 200 {
		//fmt.Println("登陆成功")
		// 显示登录成功之后的一个菜单,但是需要利用循环显示
		for {
			ShowMenu() // 显示菜单
			// 这里还要开启一个协程
			// 这一个协程时刻监听服务器的响应,如果服务器有数据推送到客户端,并且显示在客户端
			// 这一个协程的作用就是不断读取信息
			go ServerProcess(conn)
		}
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	// 最后转换为一个对象
	return
}
