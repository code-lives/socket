package rabbitmq

import (
	"chat/src/msg"
	"chat/src/mysql"
	"encoding/json"
	"fmt"
)

var queueName = "chat"

type Message struct {
	SendAccount      string `json:"SendAccount"`
	RecipientAccount string `json:"RecipientAccount"`
	MessageType      int    `json:"MessageType"`
	Content          string `json:"Content"`
	Time             int64  `json:"Time"`
	TimeStamp        string `json:"TimeStamp"`
}

// ChatConsume 聊天消费
func ChatConsume() {
	ch, err := Mq.Channel()
	defer ch.Close()
	msg.FailOnError(err, "打开通道失败")
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	msg.FailOnError(err, "创建失败")
	msgs, err := ch.Consume(
		q.Name, // queue
		"Mc",   // consumer
		true,   // auto-ack,true消费了就消失
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	msg.FailOnError(err, "Failed to register a consumer")
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			var message Message
			if json.Unmarshal([]byte(d.Body), &message); err != nil {
				fmt.Println(err)
			}
			err = mysql.GetDB().Create(message).Error
			if err != nil {
				fmt.Println("写入数据库失败", err)
			}
			fmt.Println("数据库存入成功")

		}
	}()
	fmt.Println("[*] Waiting for messages. To exit press CTRL+C")
	fmt.Println("chan", <-forever)
}
