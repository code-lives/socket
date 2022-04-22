package redis

var OnlineKey = GetOnlineKey()

type Info struct {
	NickName string
	Image    string
}

// AddUser 把用户添加到redis 中
func AddUser(uid string) {
	Rdb.SAdd(OnlineKey, uid)
}

// DelUser 把用户踢出redis
func DelUser(uid string) {
	Rdb.SRem(OnlineKey, uid)
}

// GetUser 获取用户信息
func GetUser(uid string) *Info {
	//比如从 redis 中获取用户信息
	return &Info{
		Image:    "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fb-ssl.duitang.com%2Fuploads%2Fitem%2F201412%2F28%2F20141228171331_L5T2n.jpeg&refer=http%3A%2F%2Fb-ssl.duitang.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1652863955&t=e81f65e13d7be8ae4e15ef86cd7cc43a",
		NickName: "张三",
	}
}
