package socket

import (
	"chat/src/autoloading"
	"chat/src/msg"
	"chat/src/user"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var err error
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Server struct {
	Host string
	Port int
}

var OnlineMap = make(map[string]*user.User)

func NewServer() *Server {
	config := &Server{}
	autoloading.GetEnv("SOCKET", config)
	return config
}
func Start() {
	config := &Server{}
	autoloading.GetEnv("SOCKET", config)
	http.HandleFunc("/wss", Init)
	http.ListenAndServe(fmt.Sprintf("%v:%d", config.Host, config.Port), nil)
}

func Init(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	uid := r.FormValue("UserId") //获取wss链接参数 比如用户uid
	CkToken(r.FormValue("token"))
	if err != nil {
		panic(err)
	}
	user.NewUser(uid, OnlineMap, conn)
	defer conn.Close()
	for {
		//接收消息
		_, message, err := conn.ReadMessage()
		if err != nil {
			//用户断开了
			user.Offline(uid, OnlineMap)
			break
		}
		messages := msg.GetMessage(message)
		user.SendMsg(messages, OnlineMap)
	}
}

// CkToken 解密token 自定义去吧
func CkToken(t string) {
	fmt.Println("token:" + t)
	if t == "" {
		panic("非法连接")
	}
}
