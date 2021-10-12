package main

import (
	"fmt"
	"os"
)

func main() {
	//os.O_WRONLY	只写
	//os.O_CREATE	创建文件
	//os.O_RDONLY	只读
	//os.O_RDWR		读写
	//os.O_TRUNC	清空
	//os.O_APPEND	追加
	file, err := os.OpenFile("xx.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}

	defer file.Close()

	var str = "Hello World"
	file.Write([]byte(str))

	file.WriteString("，陈袁")

	fmt.Println("操作完成，结束。")

}
