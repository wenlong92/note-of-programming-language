package main

import "fmt"

const MAX int = 3

func main() {
	// 先看实例
	a := []int{10, 100, 200}
	var i int
	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, a[i])
	}
	// 如果需要保存数组，需要使用到指针
	var ptr [MAX]*int
	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i]
	}
	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}
}
