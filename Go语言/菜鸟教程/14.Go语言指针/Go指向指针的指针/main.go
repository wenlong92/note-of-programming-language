package main

import "fmt"

func main() {
	// 如果一个指针变量存放的是另一个指针变量的地址，则这个指针变量为指向指针的指针变量
	// 当顶一个一个指向指针的指针变量时，第一个指针存放第二指针的地址，第二个指针存放变量的地址

	// 访问指向指针的指针变量值需要使用两个**号
	var a int
	var ptr *int
	var pptr **int

	a = 3000
	ptr = &a
	pptr = &ptr
	fmt.Printf("a = %d\n", a)
	fmt.Printf("*ptr = %d\n", *ptr)
	fmt.Printf("**ptr = %d\n", **pptr)
}
