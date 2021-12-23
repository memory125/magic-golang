package main

import (
	"fmt"
	"io"
	"os"
	"path"
)

// 在文件的任意位置插入任意内容，并保存
func insertContentToFile(file string, pos int64, content string)  {
	// 首先保证源文件存在，打开源文件
	fileObj, err := os.OpenFile(file, os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("====1====File %s open failed! error is %v.\n", file, err)
		return
	}

	var bytesRead []byte
	bytesRead = make([]byte, pos)
	bytes, err := fileObj.Read(bytesRead[:])
	if err != nil {
		fmt.Printf("====2====Read from file %s failed, error is %v.\n", file, err)
		return
	}

	fmt.Println("Read Content is ", string(bytesRead[:]))

	// 输出读到的内容
	fmt.Println(string(bytesRead[:]))

	// 临时文件,
	filePath := path.Dir(file)
	tmpfile := filePath + "\\tmp.tmp"
	tmpFileObj, err := os.OpenFile(tmpfile, os.O_CREATE | os.O_WRONLY | os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("====3====File %s open failed! error is %v.\n", tmpfile, err)
		return
	}

	// 将源文件中读到的内容写入临时文件
	tmpFileObj.Write(bytesRead[:bytes])
	// 将要插入的内容写入文件
	tmpFileObj.Write([]byte(content))

	// 将源文件位置偏移到pos处
	fileObj.Seek(pos, 0)

	var buf [256]byte
	for  {
		bRead, err := fileObj.Read(buf[:])
		if err == io.EOF {
			tmpFileObj.Write(buf[:bRead])
			break
		}
		if err != nil {
			fmt.Printf("====4====Read from file %s failed, error is %v.\n", file, err)
			return
		}
		tmpFileObj.Write(buf[:bRead])
	}

	// 源文件后续的也写入了临时文件
	fileObj.Close()

	// 关闭临时文件
	tmpFileObj.Close()

	// 覆盖源文件
	os.Rename(tmpfile, file)
}

func main()  {
	insertContentToFile("D:\\project\\go\\gopath\\src\\code.wing.com\\fundamental\\test1.txt", 3, "****")
}
