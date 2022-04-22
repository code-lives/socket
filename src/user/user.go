package user

import (
	"chat/src/msg"
	"chat/src/rabbitmq"
	"chat/src/redis"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

type User struct {
	Info *redis.Info     `json:"Info"`
	Uid  string          `json:"Uid"`
	Addr string          `json:"Addr"`
	Conn *websocket.Conn `json:"Conn"`
	C    chan string
}

// NewUser 存储用户信息
func NewUser(uid string, OnlineMap map[string]*User, conn *websocket.Conn) *User {
	userAdder := conn.RemoteAddr().String()
	user := &User{
		Uid:  uid,
		Info: redis.GetUser(uid), //获取用户信息
		Addr: userAdder,
		C:    make(chan string),
		Conn: conn,
	}
	Online(uid, OnlineMap, user)
	return user
}

// Online 用户上线
func Online(uid string, OnlineMap map[string]*User, user *User) {
	OnlineMap[uid] = user
	redis.AddUser(uid)
	fmt.Printf("用户【%v】上线\n", uid)
}

// Offline 用户离开
func Offline(uid string, OnlineMap map[string]*User) {
	delete(OnlineMap, uid)
	redis.DelUser(uid)
	fmt.Printf("用户【%v】溜了\n", uid)
}

// SendMsg 发送消息
func SendMsg(m msg.Msg, OnlineMap map[string]*User) {

	if _, ok := OnlineMap[m.RecipientAccount]; ok {
		ToMsg := AssembleParam(m.RecipientAccount, m.SendAccount, m.MessageType, m.Content, OnlineMap[m.RecipientAccount].Info)
		OnlineMap[m.RecipientAccount].Conn.WriteMessage(1, []byte(ToMsg))
	} else {
		ToMsg := AssembleParam(m.SendAccount, m.RecipientAccount, msg.LEAVE, msg.GetMsgContent(msg.LEAVE), OnlineMap[m.SendAccount].Info)
		OnlineMap[m.SendAccount].Conn.WriteMessage(1, []byte(ToMsg))
	}
}

// AssembleParam 组装发送数据
func AssembleParam(SendAccount string, RecipientAccount string, MessageType int, Content string, u *redis.Info) string {
	var msg msg.Msg
	now := time.Now()
	msg.Content = Content
	msg.SendAccount = SendAccount
	msg.RecipientAccount = RecipientAccount
	msg.MessageType = MessageType
	msg.Time = time.Now().Unix()
	msg.TimeStamp = fmt.Sprintf("%d-%d-%d %d:%d:%d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	msg.Info = u
	data, err := json.Marshal(msg)
	if err != nil {
		panic("msg 结构体转 json 失败" + err.Error())
	}
	s := string(data)
	//存入rabbitmq 写入数据库
	rabbitmq.Queue("chat", s)
	return s
}
