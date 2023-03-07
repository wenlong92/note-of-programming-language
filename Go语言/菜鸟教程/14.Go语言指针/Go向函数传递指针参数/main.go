package main

import "fmt"

func main() {
	// Go语言允许向函数传递指针，只需要在函数定义的参数上设置为指针类型即可
	var a int = 100
	var b int = 200
	fmt.Printf("交换前a = %d\n", a)
	fmt.Printf("交换前b = %d\n", b)
	swap(&a, &b)
	fmt.Printf("交换后a = %d\n", a)
	fmt.Printf("交换后b = %d\n", b)

}

// 交换函数
func swap(x *int, y *int) {
	fmt.Printf("\nx", x)
	fmt.Printf("\n*x", *x)
	fmt.Printf("\ny", y)
	fmt.Printf("\n*x", *y)
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
