package model

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
)

// 定义这一个userDao , 需要完成对于结构体的操作
var (
	MyUserDao *UserDao
)

type UserDao struct {
	// redis 客户端对象
	Client *redis.Client
}

// 这里使用一个工厂模式创建一个 UserDao 对象
// redis 连接池放在哪里
func NewUserDao(client *redis.Client) *UserDao {
	userDao := &UserDao{Client: client}
	return userDao
}

// 查找用户
func (this *UserDao) GetUserById(id int) (user *User, err error) {
	// 创建上下文对象
	ctx := context.Background()
	// 利用 rdb 进行操作
	res, err := this.Client.HGet(ctx, "users", strconv.Itoa(id)).Result()
	if err != nil {
		// 没有找到用户, 其实就就是错误
		err = ERROR_USER_NOTEXISTS
		return user, err
	}
	// 利用指针的方法的话,其实返回的无论如何都是同一个值
	user = &User{}
	// 之后把这一字符串转换为一个对象
	// json 数据的反序列化
	// 注意 res 是一个字符串
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("反序列化出错")
		return user, err
	}
	return
}

// 完成对于用户的校验, 如何密码和用户栋正确,但是如果ID和密码都正确都会返回一个信息
func (this *UserDao) Login(userid int, userpwd string) (user *User, err error) {
	user, err = this.GetUserById(userid)
	// 检验密码和用户名是否正确
	if err != nil {
		return user, err
	}
	if user.UserPwd != userpwd {
		err = ERROR_USER_PWD
		return user, err
	}
	// 传递一个对象
	return user, nil
}
