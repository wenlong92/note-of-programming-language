package main

import "fmt"

func main() {
	// map 类似于 python 的字典，是一种存储键值对的数据结构，使用方式和其他语言几乎没有区别

	//仅声明
	m1 := make(map[string]int)
	fmt.Println(m1)

	//声明时初始化
	m2 := map[string]string{
		"Sam":   "Male",
		"Alice": "Female",
	}
	fmt.Println(m2)

	// 赋值/修改
	m1["Tom"] = 18
	fmt.Println(m1)
	m2["Sam"] = "Animals"
	fmt.Println(m2)
}
