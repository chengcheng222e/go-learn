package main

import (
	"encoding/json"
	"fmt"
)

type CommonResp2 struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    Data2  `json:"data"`
	Success string `json:"success"`
	TraceId string `json:"traceId"`
}

type Data2 struct {
	UserName string `json:"userName"`
	Age      int    `json:"age"`
}

func main() {
	// 赋值数据
	user := Data2{UserName: "陈袁", Age: 30}

	resp := CommonResp2{Code: "0000", Message: "操作成功", Data: user, Success: "true", TraceId: "6ba7b810-9dad-11d1-80b4-00c04fd430c8"}

	jsonStr, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("JSON转换错误")
	}
	fmt.Println("输出结果：" + string(jsonStr))
}
