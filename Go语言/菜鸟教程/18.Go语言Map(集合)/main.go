package main

import "fmt"

func main() {
	// map是一种无序的键值对集合
	// 可通过key快速检索数据
	// 可以像迭代数组和切片那样迭代，但map是无序的，遍历返回的键值对顺序不确定
	// 获取map的值时，如果键不存在，返回该类型的零值，比如int类型零值是0，string类型零值是""
	// map是引用类型，如果将一个map传递给一个函数或赋值给另一个变量，它们都指向同一个底层数据结构。

	// 创建一个空map
	m := make(map[string]int)
	fmt.Println(m)
	// 创建一个初始容量为10的map
	m1 := make(map[string]int, 10)
	fmt.Println(m1)
	// 使用字面量创建map
	m2 := map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}
	fmt.Println(m2)

	// 获取元素
	v1 := m2["apple"]
	v2, ok := m2["peer"]
	v3, o := m2["banana"]
	fmt.Println(v1)
	fmt.Println(v2, ok)
	fmt.Println(v3, o)

	// 修改元素
	m2["apple"] = 5
	fmt.Println(m2)

	// 获取map长度
	fmt.Println(len(m2))

	// 遍历map
	for k, v := range m2 {
		fmt.Printf("key=%s, value=%d\n", k, v)
	}

	// 删除键值对
	delete(m2, "banana")
	fmt.Println(m2)
}
