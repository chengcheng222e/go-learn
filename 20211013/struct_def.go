package main

import (
	"encoding/json"
	"fmt"
)

type CommonResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
	Success bool   `json:"success"`
	TraceId string `json:"traceId"`
}

type User struct {
	UserName string
	Age      int
}

type Data interface {
}

func main() {
	resp := CommonResp{}
	resp.Code = "0000"
	resp.Message = "操作成功!"
	resp.Success = true
	resp.TraceId = "6ba7b810-9dad-11d1-80b4-00c04fd430c8"

	// 赋值数据
	user := User{}
	user.UserName = "陈袁"
	user.Age = 30
	resp.Data = user

	jsonStr, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("JSON转换错误")
	}
	fmt.Println("输出结果：" + string(jsonStr))
}
