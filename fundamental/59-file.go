package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 利用os包和file读取文件
func readFileByOsRead()  {
	fileObj, err := os.Open("D:\\project\\go\\gopath\\src\\code.wing.com\\fundamental\\poem.txt")
	if err != nil {
		fmt.Printf("File open failed! error is %v.\n", err)
		return
	}

	// 关闭文件
	defer fileObj.Close()

	var tmp = make([]byte, 128)
	for  {
		b, err := fileObj.Read(tmp)
		if err == io.EOF {
			fmt.Println("All file has read out!!!!")
			return
		}

		if err != nil {
			fmt.Printf("File read failed! error is %v.\n", err)
			return
		}
		fmt.Printf("%d bytes has read out!\n", b)
		fmt.Println(string(tmp[:b]))
		if b < 128 {
			return
		}
	}
}

// 利用bufio包读取文件
func readFileByBufio()  {
	fileObj, err := os.Open("D:\\project\\go\\gopath\\src\\code.wing.com\\fundamental\\poem.txt")
	if err != nil {
		fmt.Printf("File open failed! error is %v.\n", err)
		return
	}

	// 关闭文件
	defer fileObj.Close()

	// 创建用来从文件中读内容的对象
	reader := bufio.NewReader(fileObj)
	for  {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("All file has read out!!!!")
			return
		}
		if err != nil {
			fmt.Printf("File read failed! error is %v.\n", err)
			return
		}

		fmt.Print(str)
	}
}

// 使用ioutil读取文件
func readFileByIOutils()  {
	b, err := ioutil.ReadFile("D:\\project\\go\\gopath\\src\\code.wing.com\\fundamental\\poem.txt")
	if err != nil {
		fmt.Printf("File read failed! error is %v.\n", err)
		return
	}
	fmt.Print(string(b))
}

func main()  {
	// readFileByOsRead()
	//readFileByBufio()
	readFileByIOutils()
}