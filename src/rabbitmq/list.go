package rabbitmq

func QueueList() {
	//启动聊天消费队列
	go ChatConsume()
	//继续添加其他消费队列
}
