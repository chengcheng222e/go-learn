package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./README.md")
	if err != nil {
		fmt.Printf("open file failed, err is ", err)
		return
	}

	defer file.Close()

	//读取文件里面的内容
	var tmp = make([]byte, 128)
	n, err := file.Read(tmp)
	if err == io.EOF {
		fmt.Println("文件读取完了")
		return
	}

	fmt.Printf("读取了%d字节数据\n", n)
	fmt.Println(string(tmp[:n]))

}
