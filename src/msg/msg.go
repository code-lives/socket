package msg

import (
	"chat/src/redis"
	"encoding/json"
	"log"
)

var err error

const (
	CONTENTS = 1   //发送内容
	IMAGES   = 2   //发送图片
	LEAVE    = 404 //对方不在线不在线
)

var StatusText = map[int]string{
	CONTENTS: "发送内容",
	IMAGES:   "发送图片",
	LEAVE:    "对方不在线",
}

// Msg 自定义封装用户信息
type Msg struct {
	SendAccount      string      `json:"SendAccount"`
	RecipientAccount string      `json:"RecipientAccount"`
	MessageType      int         `json:"MessageType"`
	Content          string      `json:"Content"`
	Time             int64       `json:"Time"`
	TimeStamp        string      `json:"TimeStamp"`
	Info             *redis.Info `json:"Info"`
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// GetMsgContent 自定义描述
func GetMsgContent(code int) string {
	return StatusText[code]
}

// GetMessage 获取发送信息Json格式
func GetMessage(message []byte) Msg {
	var msg Msg
	err = json.Unmarshal(message, &msg)
	if err != nil {
		panic("消息结构体绑定失败" + err.Error())
	}
	return msg
}
