package main

import (
	"fmt"
	"sort"
)

func main() {
	// map默认是无序的，无论是key还是value，默认都不排序
	// 若想为map排序，需要将key或value拷贝到一个切片，在对切片排序(使用sort包)
	// 再使用切片的for-range方法打印所有的key和value

	barVal := map[string]int{
		"alpha": 34, "bravo": 56, "charlie": 23, "delta": 87, "echo": 56, "kili": 43,
		"foxtrot": 12, "golf": 34, "hotel": 16, "indio": 87, "juliet": 65, "lima": 98}
		fmt.Println("unsorted:")
		for k, v := range barVal {
			fmt.Printf("Key: %v, Value: %v\n", k, v)
		}
		keys := make([]string, len(barVal))
		i := 0
		for k, _ := range barVal {
			keys[i] = k
			i++
		}
		sort.Strings(keys)
		fmt.Println("sorted:")
		for _, k := range keys {
			fmt.Printf("Key: %v, Value: %v\n", k, barVal[k])
		}

		// 但如果想要一个排序的列表，最好使用结构体切片，会更有效：
		// tyep name struct {
		// 	key string
		// 	value int
		// }
}