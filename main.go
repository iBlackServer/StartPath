package main

import (
	"fmt"
	"StartPath/show2"
)


func init() {
	fmt.Println("在Main 函数执行之前")
}

func main() {
	show2.Show2()
	fmt.Println("测试import导入包")
}