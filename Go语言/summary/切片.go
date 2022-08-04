package main

import (
	"fmt"
)

func main() {
	// Go语言切片是对数组的抽象
	// 数组的长度是不可改变的，在特定场景中不太适用
	// 内置类型切片(“动态数组”)，长度不固定，可以追加元素，追加时可能使切片容量增大

	// 可以声明一个未指定大小的数组来定义切片：
	// var identifier []type

	// 或者用make()函数创建切片
	// var slice1 []type = make([]type, len)
	// 简写：slice := make([]type, len)

	// 指定容量， capacity为可选参数
	// make([]T, length, capacity)

	// 初始化：
	s := []int {1, 2, 3}  //这里 len=cap=3
	fmt.Println(s)

	arr := [3]int{0, 1, 2}
	s1 := arr[:]  //初始化切片s1 是数组arr的引用
	fmt.Println(s1)


	// len()和cap()函数
	// 切片是可索引的，并且可由len()方法获取长度，切片通了cap()方法计算容量，可测量切片最长可达到多少，下例：
	var numbers = make([]int, 3, 5)
	printSlice(numbers)


	// 空(nil)切片
	var numbers1 []int
	printSlice(numbers1)
	if(numbers1 == nil) {
		fmt.Printf("切片为空\n")
	}


	// 切片截取
	// 可以通过设置下限及上限设置截取切片，例如：
	numbers2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers2)
	fmt.Println("numbers2[1:4] == ", numbers2[1:4])
	fmt.Println("numbers2[:3] == ", numbers2[:3])
	fmt.Println("numbers2[4:] == ", numbers2[4:])

	numbers3 := make([]int, 0, 5)
	printSlice(numbers3)

	numbers4 := numbers2[:2]
	printSlice(numbers4)
	printSlice(numbers2)
	
	numbers5 := numbers2[2:5]
	printSlice(numbers5)
	printSlice(numbers2)


	// append()和copy()函数
	var num []int
	printSlice(num)

	num = append(num, 0)  //添加一个元素
	printSlice(num)

	num = append(num, 2, 3, 4)  //添加多个元素
	printSlice(num)

	num1 := make([]int, len(num), (cap(num) * 2))  //创建num1，切片容量翻倍

	copy(num1, num)  //拷贝num的内容到num1
	printSlice(num1)
}
func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}