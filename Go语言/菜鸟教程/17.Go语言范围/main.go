package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// range关键字用于for循环中迭代数组、切片、通道或集合的元素。在数组和切片中它返回元素的索引和索引对应的值，在集合中返回key-value对
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// for循环的range格式可以省略key和value
	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0
	for key, value := range map1 {
		fmt.Printf("key is %d -value is %f\n", key, value)
	}
	for _, value := range map1 {
		fmt.Printf("value is %f\n", value)
	}
	for key := range map1 {
		fmt.Printf("key is %d\n", key)
	}

	// range遍历其他数据结构
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
