package main

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"wing.com/magic-golang/configuration/yamlfile/conf"
)

// yaml文件解析示例
/*
	yaml解析库：
		go get -u gopkg.in/conf.v2
*/

func main() {
	conf := new(conf.Yaml)
	yamlFile, err := ioutil.ReadFile("./conf/test.yaml")

	fmt.Println("yamlFile:", string(yamlFile))
	if err != nil {
		fmt.Printf("yamlFile.Get err #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		fmt.Printf("Unmarshal failed, error is: %v\n", err)
	}
	fmt.Printf("conf is %v \n", conf)
}
