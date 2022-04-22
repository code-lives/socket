package autoloading

import (
	"gopkg.in/ini.v1"
	"strings"
)

func GetEnv(s string, config interface{}) {
	var err error
	var conf *ini.File
	if conf, err = ini.Load("env/env.ini"); err != nil {
		panic("env: 读取redis配置失败" + err.Error())
	}
	if conf.Section(strings.ToUpper(s)).MapTo(config); err != nil {
		panic("redis结构体绑定失败" + err.Error())
	}
}
