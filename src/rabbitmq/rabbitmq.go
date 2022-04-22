package rabbitmq

import (
	"chat/src/autoloading"
	"chat/src/msg"
	"fmt"
	"github.com/streadway/amqp"
)

type Rabbitmq struct {
	Host     string
	Port     int
	Name     string
	Password string
}

var (
	Mq  *amqp.Connection
	err error
)

// Start 创建Mq
func Start() {
	config := &Rabbitmq{}
	autoloading.GetEnv("Rabbitmq", config)
	Mq, err = amqp.Dial(fmt.Sprintf("amqp://%v:%v@%v:%d/", config.Name, config.Password, config.Host, config.Port))
	if err != nil {
		panic("连接rabbitmq出错" + err.Error())
	}
	msg.FailOnError(err, "连接rabbitmq出错")
	fmt.Println("Rabbitmq：成功启程")
	//启动消费队列
	QueueList()
}

// Queue 写入Mq 【queue 队列名称 content 内容】
func Queue(queue string, content string) {
	ch, err := Mq.Channel()
	defer ch.Close()
	msg.FailOnError(err, "打开通道失败")
	q, err := ch.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	msg.FailOnError(err, "创建失败")
	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(content),
	}
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		message)
	msg.FailOnError(err, "rabbitmq 写入失败")
	fmt.Println("rabbitmq 写入成功")
}
