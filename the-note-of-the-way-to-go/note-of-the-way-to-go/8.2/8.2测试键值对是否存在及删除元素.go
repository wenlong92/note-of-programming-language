package main

import (
	"fmt"
)

func main() {
	// 使用 val1 = map1[key1]方法获取key1对应的值val1时，如果map中不存在key1，val1就是一个值类型的空值
	// 但此时无法判断key1不存在还是对应的value就是空值
	// 未解决这个问题，可以用 val1, isPresent = map1[key1]
	// isPresent返回bool值，如果key1存在，val1就是对应的value值且isPresent为true，反之，val1是空值，isPresent为false

	// 如果指向判断key是否存在股关心对应的值是多少，可以这样：
	// _, ok := map1[key1]
	// 或者与if混合使用
	// if _, ok := map1[key1]; ok { ... }

	// 从map1中删除key1：delete(map1, key1)，如果key1不存在，该操作不会产生错误

	var value int
	var isPresent bool

	map1 := make(map[string]int)
	map1["New Delhi"] = 55
	map1["Beijing"] = 20
	map1["Washington"] = 25
	value, isPresent = map1["Beijing"]
	if isPresent {
		fmt.Printf("The value of \"Beijing\" in map1 is: %d\n", value)
	} else {
		fmt.Printf("map1 does not caotain Beijing")
	}

	value, isPresent = map1["Paris"]
	fmt.Printf("Is \"Paris\" in map1 ?: %t\n", isPresent)
	fmt.Printf("Value is: %d\n", value)

	// delete item
	delete(map1, "Washington")
	value, isPresent = map1["Washington"]
	if isPresent {
		fmt.Printf("The value of \"Wahshington\" in map1 is: %d\n", value)
	} else {
		fmt.Println("map1 does not contain Wahshington")
	}
}