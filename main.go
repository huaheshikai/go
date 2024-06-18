package main

import "fmt"

var (
	nx1, nx2, nx3 = 100, 200, 300
)

func main() {
	var num = 18.55 //将变量赋值，自己确定变量类型
	fmt.Println(num)

	var num1 int //声明变量类型再赋值
	num1 = 18
	fmt.Println(num1)

	var nmu2 int //直接使用变量默认值
	fmt.Println(nmu2)

	num3 := "真" //省略var
	fmt.Println(num3)
	fmt.Println("----------------------------------------------------------")

	var sex1, sex2, sex3 int //创建多个变量
	fmt.Println(sex1, sex2, sex3)

	var sex4, sex5, sex6 int = 1, 2, 3 //创建多个变量
	fmt.Println(sex4, sex5, sex6)

	n1, n2, n3 := 1, 2, 3
	fmt.Println(n1, n2, n3)

	fmt.Println(nx1, nx2, nx3)
}
