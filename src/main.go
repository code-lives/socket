package main

import (
	"chat/src/mysql"
	"chat/src/rabbitmq"
	"chat/src/redis"
	"chat/src/socket"
)

func main() {
	redis.Start() //开启 redis
	mysql.Start()
	rabbitmq.Start() //开启rabbitmq
	socket.Start()   //开启socket 监听

}
