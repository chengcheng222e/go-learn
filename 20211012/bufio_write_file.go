package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		fmt.Println("读取文件失败", err)
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	for i := 0; i < 10; i++ {
		writer.WriteString("我是陈袁" + strconv.Itoa(i) + "\n")
	}
	writer.Flush()
}
