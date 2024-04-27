package process

import (
	"encoding/json"
	"fmt"
	"net"
	"userContact/common"
	"userContact/server/model"
	"userContact/server/utils"
)

// 处理用户需要的字段
type UserProcess struct {
	Conn net.Conn
}

// 处理注册请求
func (this *UserProcess) ServiceProcessRegister(mes *common.Message) (err error) {
	// 取出一个Data , 进行相应的判断,发送相应的信息
	var registerMes common.RegisterMes
	fmt.Println(mes)
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal fall err = ", err)
		return err
	}
	// 开始返回一个响应
	var resMes common.Message
	resMes.Type = common.RegisterResMesType
	// 定义一个返回结构
	var registerResMeg common.RegisterResMes
	// 数据库中完成注册
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMeg.Code = 505 // 表示用户信息不存在
			registerResMeg.Error = err.Error()
		} else {
			registerResMeg.Code = 506 // 表示用户信息不存在
			registerResMeg.Error = "注册时发生位置错误"
		}
	} else {
		registerResMeg.Code = 200
	}
	// 进行反序列化操作
	// 写入一个连接
	// 完成序列化
	// 1. 首先序列化 loginResMsg
	data, err := json.Marshal(registerResMeg)
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
	// err = utils.WritePkg(conn, mesData)
	// 首先创建一个实例
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(mesData)
	// 最后发送信息, 首先还是确定长度
	return
}

// 处理登录请求

func (this *UserProcess) ServiceProcessLogin(mes *common.Message) (err error) {
	// 取出一个Data , 进行相应的判断,发送相应的信息
	var loginMes common.LoginMes
	fmt.Println(mes)
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
	// 这里得到了一个用户信息
	// 检验用户信息
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)

	fmt.Println("登陆成功的用户信息为", user)
	// 数据库中查找用户的逻辑
	// 校验用户信息,直接在 redis 数据库中寻找数据
	//client := redis.NewClient(&redis.Options{
	//	Addr:     "192.168.59.132:6379",
	//	Password: "808453",
	//	DB:       0,
	//})
	//var userDao = model.UserDao{Client: client}
	//user, err := userDao.GetUserById(loginMes.UserId)
	//fmt.Println(user)
	//// 用于用户登录的逻辑
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMeg.Code = 500
			loginResMeg.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMeg.Code = 500
			loginResMeg.Error = err.Error()
		}
	} else {
		// 开始封装正确信息
		loginResMeg.Code = 200
		fmt.Println(loginMes, "登陆成功")
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
	// err = utils.WritePkg(conn, mesData)
	// 首先创建一个实例
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(mesData)
	// 最后发送信息, 首先还是确定长度
	return
}
