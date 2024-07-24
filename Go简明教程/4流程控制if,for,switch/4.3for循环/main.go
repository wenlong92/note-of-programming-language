package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		if sum > 50 {
			break
		}
		sum += i
	}
	fmt.Println(sum)

	//数组遍历
	nums := []int{10, 20, 30, 40}
	for i, num := range nums {
		fmt.Println(i, num)
	}

	//字典遍历
	m2 := map[string]string{
		"Sam":   "Male",
		"Alice": "Female",
	}
	for key, value := range m2 {
		fmt.Println(key, value)
	}

}
