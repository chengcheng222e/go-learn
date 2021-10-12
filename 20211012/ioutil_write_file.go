package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	str := "我是陈袁，通过ioutil.WriteFile写入"
	err := ioutil.WriteFile("xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("读取文件失败", err)
		return
	}
	fmt.Println("文件写入完成")
}
