package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 文件写操作
// os
func writeFileByOs() {
	fileObj, err := os.OpenFile("D:\\project\\go\\gopath\\src\\code.wing.com\\fundamental\\test.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("Open file failed, error is %v.\n", err)
		return
	}

	defer fileObj.Close()

	// 写文件
	// 按字节
	fileObj.Write([]byte("Write byte to file!!!!\n"))

	// 按字符串
	fileObj.WriteString("Write string to file!!!!")
}

// Bufio
func writeFileByBufio() {
	fileObj, err := os.OpenFile("D:\\project\\go\\gopath\\src\\code.wing.com\\fundamental\\test.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("Open file failed, error is %v.\n", err)
		return
	}

	defer fileObj.Close()

	// 创建writer
	writer := bufio.NewWriter(fileObj)

	// 写文件
	// 按字节
	writer.Write([]byte("Write byte to file!!!!\n"))

	// 按字符串
	writer.WriteString("Write string to file!!!!")

	// 将缓存写入文件
	writer.Flush()
}

// ioutil
func writeFileByIoutil() {
	str := "Write string to file by Ioutil\n"

	err := ioutil.WriteFile("D:\\project\\go\\gopath\\src\\code.wing.com\\fundamental\\test.txt", []byte(str), 0666)
	if err != nil {
		fmt.Printf("Write file failed, error is %v.\n", err)
		return
	}
}


func main()  {
	// writeFileByOs()
	// writeFileByBufio()
	writeFileByIoutil()
}
