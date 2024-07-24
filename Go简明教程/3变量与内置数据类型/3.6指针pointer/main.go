package main

import "fmt"

func main() {
	//指针：某个值的地址，类型定义时使用符号*，
	//对一个已经存在的变量，使用&获取该变量的地址

	str := "Golang"
	var p *string = &str
	fmt.Println(p)

	*p = "hello"
	fmt.Println(str)

	num := 100
	add(num)
	fmt.Println(num) // 100 ，num 没有变化

	readAdd(&num)
	fmt.Println(num) // 101，指针传递，num 被修改

}

//一般来说，指针通常在函数传递参数或给某给类型定义新的方法时使用。
//Go 语言中，参数是按值传递的，如果不使用指针，函数内部会拷贝一份参数的副本，对参数的修改不会影响到外部变量的值。
//如果参数使用指针，对参数的传递将会影响到外部变量

func add(num int) {
	num += 1
}

func readAdd(num *int) {
	*num += 1
}

//  var a int = 10
//	var ipp *int
//	ipp = &a
// 此时 ipp 等于 &a
// *ipp 等于 a
