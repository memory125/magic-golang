package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"wing.com/magic-golang/configuration/inifile/conf"
)

// ini文件解析示例
/*
	ini解析库：
		go get -u gopkg.in/ini.v1
*/

func main() {
	cfg, err := ini.Load("./conf/cfg.ini")
	if err != nil {
		log.Fatal("Fail to read file: ", err)
	}

	fmt.Println("App Name:", cfg.Section("").Key("app_name").String())
	fmt.Println("Log Level:", cfg.Section("").Key("log_level").String())

	fmt.Println("MySQL IP:", cfg.Section("mysql").Key("ip").String())
	mysqlPort, err := cfg.Section("mysql").Key("port").Int()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MySQL Port:", mysqlPort)
	fmt.Println("MySQL User:", cfg.Section("mysql").Key("user").String())
	fmt.Println("MySQL Password:", cfg.Section("mysql").Key("password").String())
	fmt.Println("MySQL Database:", cfg.Section("mysql").Key("database").String())

	fmt.Println("Redis IP:", cfg.Section("redis").Key("ip").String())
	redisPort, err := cfg.Section("redis").Key("port").Int()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Redis Port:", redisPort)

	var proCfg = new(conf.ProjectConf)
	err = ini.MapTo(proCfg, "./conf/cfg.ini")
	if err != nil {
		fmt.Printf("load cfg.ini failed, error is %v.\n", err)
		return
	}
	fmt.Println("=======MySQL==========")
	fmt.Println("host: ", proCfg.MySQLConf.IP)
	fmt.Println("port: ", proCfg.MySQLConf.Port)
	fmt.Println("user: ", proCfg.MySQLConf.User)
	fmt.Println("password: ", proCfg.MySQLConf.Password)
	fmt.Println("database: ", proCfg.MySQLConf.Database)

	fmt.Println("=======Redis==========")
	fmt.Println("host: ", proCfg.RedisConf.IP)
	fmt.Println("port: ", proCfg.RedisConf.Port)
}
