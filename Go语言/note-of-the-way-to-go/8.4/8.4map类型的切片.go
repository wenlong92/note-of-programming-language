package main

import (
	"fmt"
)

func main() {
	// 若想获取一个map类型的切片，要使用两次make()函数
	// 第一次分配切片，第二次分配切片中每个map元素
	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}
	fmt.Printf("Version A: Value of items: %v\n", items)
}