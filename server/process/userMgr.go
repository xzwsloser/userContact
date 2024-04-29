package process

import "fmt"

// 定义成一个全局变量
// 因为很多地方需要使用,所以需要定义成一个全局变量
var (
	UserMgrObj *UserMgr
)

// 定义一个结构体
type UserMgr struct {
	onlineUses map[int]*UserProcess
}

// 初始化函数
func init() {
	// 一个包中的 init 函数
	UserMgrObj = &UserMgr{
		// map,chan 容器需要 make 才可以使用
		onlineUses: make(map[int]*UserProcess, 1024),
	}
}

// 添加用户
func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.onlineUses[up.UserId] = up
}

// 删除用户
func (this *UserMgr) DeleteOnlineUser(up *UserProcess) {
	delete(this.onlineUses, up.UserId)
}

// 显示当前所用用户对象
func (this *UserMgr) GetAllOnlineUser(up *UserProcess) map[int]*UserProcess {
	return this.onlineUses
}

// 根据 Id 地址返回在线用户
func (this *UserMgr) GetOnlineUserById(userid int) (up *UserProcess, err error) {
	up, ok := this.onlineUses[userid]
	if !ok {
		// 当前在线列表不在,所以直接返回一个 err
		err = fmt.Errorf("用户 %d 不在线", userid)
		return up, err
	}
	return
}

// 修改信息其实就是一个 Add 方法的运用,map中键值可以进行覆盖
