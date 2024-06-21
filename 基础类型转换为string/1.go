// package main

// import (
// 	"fmt"
// )

//	func main() {
//		// var n1 int8 = 8
//		// var n2 float32 = 3.14
//		var n3 bool = false
//		var n4 byte = 'a'
//		// var s1 string = fmt.Sprintf("%d", n1)
//		// fmt.Printf("%T\n s1的值为 %v \n", s1, s1)
//		// fmt.Println(s1)   //将int类型转换为string类型
//		// var s2 string = fmt.Sprintf("%f", n2)
//		// fmt.Printf("%T \n s2的值为 %v \n s2等于%q", s2, s2, s2)
//		var s3 string = fmt.Sprintf("%t", n3)
//		var s4 string = fmt.Sprintf("%c", n4)
//		fmt.Printf("%T s3的值为%q s3的值为%v \n", s3, s3, s3)
//		fmt.Printf("%T s4的值为%q s4的值为%v \n", s4, s4, s4)
//	}
package main

import (
	"fmt"
)

func main() {
	var n1 int = 8
	var n2 float32 = 3.1414
	var s1 string = fmt.Sprintf("%d", n1)
	fmt.Printf("%T %v \n", s1, s1)
	var s2 string = fmt.Sprintf("%f", n2)
	fmt.Printf("%T %v", s2, s2)
}
