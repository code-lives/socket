package redis

import (
	"chat/src/autoloading"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type Config struct {
	Host     string
	Password string
	Port     int
	Prefix   string
	Online   string
}

var (
	config = Config{}
	Rdb    *redis.Client
	err    error
)

func Start() {
	autoloading.GetEnv("REDIS", &config)
	Rdb = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%v:%d", config.Host, config.Port),
		Password:     config.Password,
		DB:           0,
		ReadTimeout:  time.Duration(20000) * time.Millisecond,
		WriteTimeout: time.Duration(20000) * time.Millisecond,
	})
	if _, err = Rdb.Ping().Result(); err != nil {
		panic("redis 出问题啦" + err.Error())
	}
	fmt.Println("redis 启动成功")
}
func Redis() *redis.Client {
	return Rdb
}
func GetOnlineKey() string {
	return config.Online
}
func Close() {
	Rdb.Close()
}
