package main

import (
	"example.com/hello/service"
	"fmt"
)

func main() {
	// 模拟http调用
	user ,err:= service.SearchUser(1)
	fmt.Println(user,err)
}
