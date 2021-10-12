package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Open("./README.md")
	if err != nil {
		fmt.Println("读取文件失败", err)
		return
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Print(line)
			}
			fmt.Println("\n")
			fmt.Println("文件读完了")
			break
		}

		if err != nil {
			fmt.Println("文件读取出错了", err)
			return
		}
		fmt.Print(line)
	}

}
