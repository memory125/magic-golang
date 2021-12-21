package main

import (
	"fmt"
	"io"
	"os"
)

func main()  {
	fileObj, err := os.Open("D:\\project\\go\\gopath\\src\\code.wing.com\\fundamental\\poem.txt")
	if err != nil {
		fmt.Printf("File open failed! error is %v.\n", err)
		return
	}

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
		if b == 0 {
			return
		}
	}

}