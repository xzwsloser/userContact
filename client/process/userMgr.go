package userprocess

import (
	"fmt"
	"userContact/common"
	"userContact/server/model"
)

// 定义一个全局变量
// 定义一个客户端需要维护的一个 map[int]User
var onlineUsers map[int]*model.User = make(map[int]*model.User, 10)

// 但是初始化工作在哪里
// 初始化工作应该在登陆成功的时候
func updateUserStatus(notifyUserStatusMes *common.NotifyUserStatusMes) {
	// 首先判断原来有没有信息
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	// 没有用户的话就可以更新状态,有的话就可以直接赋值
	if !ok {
		// 更新状态
		user = &model.User{
			UserId:     notifyUserStatusMes.UserId,
			UserStatus: notifyUserStatusMes.Status,
		}
	}
	onlineUsers[notifyUserStatusMes.UserId] = user
	outputOnlineUser()
}

// 在客户端显示当前在线用户
func outputOnlineUser() {
	fmt.Println("当前在线用户列表为: ")
	for id, _ := range onlineUsers {
		fmt.Println("用户 id:\t", id)
	}
}
