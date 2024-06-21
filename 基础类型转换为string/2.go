// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	var n1 int = 8
// 	var n2 float32 = 3.14
// 	var n3 bool = false
// 	var s1 string = strconv.FormatInt(int64(n1), 10) //第一个参数转为int64位，第二个参数为指定进制位
// 	fmt.Printf("%T s1的值为%v  %q \n", s1, s1, s1)
// 	var s2 string = strconv.FormatFloat(float64(n2), 'f', 9, 64) /*第一个参数为转换为float64，第二个参数为指数，
// 	第三个参数为保留9位，第四个方式为float64的类型 */
// 	fmt.Printf("%T s2的值为%q %v \n", s2, s2, s2)
// 	var s3 string = strconv.FormatBool(n3)
// 	fmt.Printf("%T, %v , %q", s3, s3, s3)
// }