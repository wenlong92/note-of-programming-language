package main

import (
	"fmt"
)

func main() {
	// 如果map的值类型可以作为key且所有value是唯一的，可通过下面方法简单的做到键值对调：
	barVal := map[string]int{
		"alpha": 34, "bravo": 56, "charlie": 23, "delta": 87, "echo": 56, "kili": 43,
		"foxtrot": 12, "golf": 34, "hotel": 16, "indio": 87, "juliet": 65, "lima": 98}
	
	invMap := make(map[int]string, len(barVal))
	for k, v := range barVal {
		invMap[v] = k
	}
	fmt.Println("inverted:")
	for k, v := range invMap {
		fmt.Printf("Key: %v, Value: %v \n", k, v)
	}

	// 如果原始value不唯一，这样做会出问题，一种解决方法是仔细检查唯一性并使用多值map，例如： map[int][]string
}