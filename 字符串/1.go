package main

import (
	"fmt"
)

func main() {
	var n1 string = "你好 学习golang很开心" //字符是不可变的一旦定义好，字符内容无法改变
	fmt.Println(n1)
	// var n2 string = "abc"
	// n2[1] = 't'
	// fmt.Println(n2)
	var n2 string = "当前产生的字符串为" //没有特殊符号一般用双引号
	fmt.Println(n2)             //输出含有特殊符号的用反引号   ``
	var n3 string = `    
package main

import (
	"fmt"
)

func main() {
	fmt.Println("aaaaaaa\naaaaaaa") //换行
	fmt.Println("aaaaaaa\taaaaaaa") //制表位
	fmt.Println("aaaa\baaaaaaa")    //退格
	fmt.Println("\"golang\"")       //输出引号
}
	`
	fmt.Println(n3)
	var n4 string = "abc"
	n4 += "bcd"
	fmt.Println(n4)
	var n5 = "abc" + "bcd" + "cdf" + "jkl" + "hng" //字符串可以拼接
	fmt.Println(n5)

}
