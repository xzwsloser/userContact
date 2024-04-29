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
	Conn   net.Conn
	UserId int // 表示当前连接所属用户的 Id 号码
}

// 通知所有用户在线信息的方法
func (this *UserProcess) NotifyOthersOnline(userId int) {
	// 需要遍历一个上线的列表
	// 每一次客户端和服务端之间的协程把用户 Id 发送到服务器端
	// 服务器端发送这一个信息即可
	// 发送给之前定义的在线用户列表即可
	for id, up := range UserMgrObj.onlineUses {
		if id == userId {
			continue
		}
		// 开始拿到 up 中的连接,发送连接
		up.NotifyMeOnline(userId) // up 也是一个 UserProcess 对象
	}

}

func (this *UserProcess) NotifyMeOnline(userId int) {
	// 拿到连接发送状态
	// 开始组装消息
	// NotifyStatusMes
	var mes common.Message
	// 开始封装
	mes.Type = common.NotifyUserStatusMesType
	var notifyUserStatusMes common.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = common.UserOnline
	// 开始序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("序列化出错")
		return
	}
	// 开始赋值
	mes.Data = string(data)
	// 开始把 mes 序列化
	mesData, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化失败")
		return
	}
	// 开始发送 mesData
	// 创建发送实例对象
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	fmt.Println("显示用户", string(mesData))
	err = tf.WritePkg(mesData)
	if err != nil {
		fmt.Println("服务器端发送信息失败 ")
		return
	}
}

// 处理注册请求
func (this *UserProcess) ServiceProcessRegister(mes *common.Message) (err error) {
	fmt.Println("执行了 ServiceProcessRegister 方法")
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
		fmt.Println("err 不为空", err)
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
		// 登录成功,可以放入字段了
		// 登陆成功的 UserId 复制该用户信息
		this.UserId = loginMes.UserId
		// 相当于连接到信息
		UserMgrObj.AddOnlineUser(this)
		// 通知其他用户上线信息
		// 但是客户端如何发送信息呢
		// 把在线用户放在一个key中
		// 遍历里面的一个 map 容器
		for k, _ := range UserMgrObj.onlineUses {
			loginResMeg.UserIds = append(loginResMeg.UserIds, k)
		}
		this.NotifyOthersOnline(loginMes.UserId)
		fmt.Println(loginMes, "登录成功")
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
