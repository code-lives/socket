# GO Socket 简易Demo

## 1.安装go依赖

```

go mod tidy
```
# 2.docker 安装mysql redis rabbitmq （进入location文件 执行，执行后把当前sql文件导入mysql进行聊天记录的存储）

```
docker-compose up -d
```
# 3.mysql redis rabbitmq 配置文件在（env）


# 4.启动go服务

```
go run src/main.go
```

# 5.测试（打开html 中的两个文件测试）
![image](/html/1.png)

### 自定义 连接用户uid ws://127.0.0.1:8888/wss?UserId=1234567

```

index.html

indes.html

```
## Web发送参数说明
```
SendAccount 当前人id
RecipientAccount 接受人id
MessageType 消息类型 自定义【src/msg/msg.go】
Content 内容
{"Content":"Test","SendAccount":"123456","RecipientAccount":"1234567","MessageType":1}
```
## Web返回参数说明
```
SendAccount 发送人id
RecipientAccount 接受人id
MessageType 消息类型  自定义【src/msg/msg.go】
Content 内容
Time 时间戳
TimeStamp 格式化时间
Info 发送人的信息【NickName、Image】
{"SendAccount":"123456","RecipientAccount":"1234567","MessageType":404,"Content":"对方不在线","Time":1650275061,"TimeStamp":"2022-4-18 17:44:21","Info":{"NickName":"张三","Image":"https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fb-ssl.duitang.com%2Fuploads%2Fitem%2F201412%2F28%2F20141228171331_L5T2n.jpeg\u0026refer=http%3A%2F%2Fb-ssl.duitang.com\u0026app=2002\u0026size=f9999,10000\u0026q=a80\u0026n=0\u0026g=0n\u0026fmt=auto?sec=1652863955\u0026t=e81f65e13d7be8ae4e15ef86cd7cc43a"}}
```

### 自定义获取用户信息 【src/redis/user.go】

```
 // uid 是 ws 连接上面的当前用户id
func GetUser(uid string) *Info {
	return &Info{
		Image:    "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fb-ssl.duitang.com%2Fuploads%2Fitem%2F201412%2F28%2F20141228171331_L5T2n.jpeg&refer=http%3A%2F%2Fb-ssl.duitang.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1652863955&t=e81f65e13d7be8ae4e15ef86cd7cc43a",
		NickName: "张三",
	}
}
```