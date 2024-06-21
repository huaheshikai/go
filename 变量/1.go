package main

import (
	"fmt"
)

func main() {
	// var number = 128
	// fmt.Printf("%T \n", number)  //输出当前变量的数据类型
	// fmt.Println(unsafe.Sizeof(number))  //查看当前占用的字节
	// var pai float32 = 3.14
	// fmt.Println("圆周率为 ", pai)
	// var pai1 float32 = 3e+1
	// fmt.Println("输出为", pai1)
	// var pai2 float32 = 3.1415699999
	// var pai3 float64 = 3.1415926888
	// fmt.Println("float32 \n float64 \n", pai2, pai3)
	// fmt.Printf("%T\n", pai2)
	// fmt.Printf("%T\n", pai3)
	var number byte = 'a'
	fmt.Println(number)
	var n1 byte = '6'
	fmt.Println(n1)
	var n2 int = '中'
	fmt.Printf("%c", n2)
}
