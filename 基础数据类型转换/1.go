package main

import (
	"fmt"
)

func main() {
	var n1 uint8 = 8
	var n2 float32 = float32(n1) //自动转换无效必须显示转换
	fmt.Println(n2)
	fmt.Printf("/'%T/'", n2)
	//var n9 int8 = int8(n2) + 1278 //数值溢出  报错
	fmt.Println(n2)
}
