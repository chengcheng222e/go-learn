package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("./README.md")
	if err != nil {
		fmt.Println("读取文件失败", err)
		return
	}
	fmt.Println(string(content))
}
