package main

import (
	"fmt"
)

func main() {
	// 这种构建方法可应用于数组和切片,例如：
	// for ix, value := range slice1 {
	// 	...
	// }
	// ix是数组或切片的索引，value是该索引位置的值
	// value只是slice1某个索引位置的一个拷贝，不能用来修改slice1该索引位置的值
	var slice1 []int = make([]int, 4)
	slice1[0] = 1
	slice1[1] = 2
	slice1[2] = 3
	slice1[3] = 4
	for ix, value := range slice1 {
		fmt.Printf("Slice at %d is: %d\n",ix, value)
	}


	seasons := []string{"Spring", "Summer", "Autumn", "Winter"}
	for ix, season := range seasons {
		fmt.Printf("Season %d is: %s\n", ix, season)
	}
	var season string
	for _, season = range seasons {  //可用 _ 忽略索引或value值
		fmt.Printf("%s\n", season)
	}

	// 多维切片下的for-range:
	// for row := range screen {
	// 	for column := range screen[row] {
	// 		screen[row][column] = 1
	// 	}
	// }
	items := [...]int{10, 20, 30, 40, 50}
	for _, item := range items {
		item *= 2
	}
	fmt.Print("items:",items) //未改变数组items中的值
}