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

	var content []byte
	//读取文件里面的内容
	var tmp = make([]byte, 128)

	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读取完了")
			break
		}
		if err != nil {
			fmt.Println("读取文件失败，err ：", err)
			return
		}
		content = append(content, tmp[:n]...)
	}

	fmt.Println(string(content))
}
