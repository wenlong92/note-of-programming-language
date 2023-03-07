package main

import "fmt"

func main() {
	// Go 语言支持多维数组
	//var threedim [5][10][4] int //声明了三维的整型数组

	// 创建数组
	values := [][]int{}
	// 使用append添加两行一维数组
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	values = append(values, row1)
	values = append(values, row2)
	// 显示两行数据
	fmt.Println("Row 1:")
	fmt.Println(values[0])
	fmt.Println("Row 2:")
	fmt.Println(values[1])
	// 访问第一个元素
	fmt.Println("第一个元素为：")
	fmt.Println(values[0][0])

	// 初始化二维数组
	a := [3][4]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 10, 11}}
	fmt.Println(a)

	// 初始化一个二维数组
	sites := [2][2]string{}
	sites[0][0] = "Google"
	sites[0][1] = "Runoob"
	sites[1][0] = "Taobao"
	sites[1][1] = "Weibo"
	fmt.Println(sites)

	// 访问二维数组
	var a1 = [5][2]int{{0, 0}, {1, 2}, {3, 4}, {5, 6}, {7, 8}}
	fmt.Println(a1[1][1])

	// 创建各个维度元素数量不一致的多维数组
	// 创建空的二维数组
	animals := [][]string{}
	// 创建三个一维数组，数组长度不同
	row10 := []string{"fish", "shark", "ell"}
	row20 := []string{"bird"}
	row30 := []string{"lizard", "salamander"}
	animals = append(animals, row10)
	animals = append(animals, row20)
	animals = append(animals, row30)
	for i := range animals {
		fmt.Printf("Row: %v\n", i)
		fmt.Println(animals[i])
	}
}
