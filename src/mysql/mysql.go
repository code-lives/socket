package mysql

import (
	"chat/src/autoloading"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	DB  *gorm.DB
	err error
)

type Config struct {
	Host     string
	Table    string
	User     string
	Password string
	Port     string
	Prefix   string
	Singular bool
	Charset  string
}

func Start() {
	dns := &Config{}
	autoloading.GetEnv("Databases", dns)
	if err != nil {
		panic("数据库结构体绑定失败" + err.Error())
	}
	driver := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", dns.User, dns.Password, dns.Host, dns.Port, dns.Table, dns.Charset)
	DB, err = gorm.Open(mysql.Open(driver), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: dns.Singular,
			TablePrefix:   dns.Prefix,
		},
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("mysql连接成功")

}
func GetDB() *gorm.DB {
	return DB
}
