package main

import (
	"fmt"
)

func main() {
	// 用for循环读取map
	// for key, value := range map1 { ... }

	// key是可选元素，只选取value，可这样使用:
	// for _, value := range map1 { ... }

	// 只获取key，可这样使用:
	// for key := range map1 { ... }

	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0
	for key, value := range map1 {
		fmt.Printf("key is: %d - value is: %f\n", key, value)
	}
	// map不是按照key顺序排列的，也不是按照value顺序排列的
	// (map本质是散列表，map的增长扩容会导致重新进行散列，这就是map的遍历结果在扩容前后变得不可靠，为了不让大家依赖遍历顺序，
	// 每次遍历的起点即起始bucket的位置不一样，即不让遍历都从bucket0开始，所以即使未扩容时，遍历出的map也是无序的)

	capitals := map[string] string {"France":"Paris", "Italy":"Rome","Japan":"Tokyo"}
	for key := range capitals {
		fmt.Println("Map item: Capital of", key, "is", capitals[key],"\n")
	}
}