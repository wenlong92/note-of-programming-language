package main

import "fmt"

func main() {
	// Go语言中切片是对数组的抽象
	// 数组的长度不可改变，在特定场景不太适用，切片（动态数组）长度不固定，可以追加元素，追加时可能使切片的容量增大

	// 可以声明一个未指定大小的数组来定义切片
	// var identifier []type

	// 或使用make()函数创建切片
	// var slice1 []type = make([]type, len)
	// 简写 slice1 := make([]type,len)

	// 也可以指定容量，为可选参数
	// make([]T, length, capacity)

	// 切片初始化
	// s := []int{1, 2, 3}

	// len()  cap()函数，分别是获取长度和容量的方法
	var numbers = make([]int, 3, 5)
	printSlice(numbers)

	// 空切片nil
	var numbers1 []int
	printSlice(numbers1)
	if numbers1 == nil {
		fmt.Printf("切片是空的")
	}

	// 切片截取
	numbers2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers2)
	// 打印原始切片
	fmt.Println("numbers2 == ", numbers2)
	// 打印子切片从索引1到索引4(不包含)
	fmt.Println("numbers2[1:4]", numbers2[1:4])
	// 默认下限为0
	fmt.Println("numbers2[:3]", numbers2[:3])
	// 默认上限为len(s)
	fmt.Println("numbers2[4:]", numbers2[4:])

	numbers3 := make([]int, 0, 5)
	printSlice(numbers3)

	numbers4 := numbers2[:2]
	printSlice(numbers4)

	numbers5 := numbers2[2:5]
	printSlice(numbers5)

	// append()和copy()函数
	var nb []int
	printSlice(nb)

	// 允许追加空切片
	nb = append(nb, 0)
	printSlice(nb)

	// 向切片添加一个元素
	nb = append(nb, 1)
	printSlice(nb)

	// 同时添加多个元素
	nb = append(nb, 2, 3, 4)
	printSlice(nb)

	// 创建新切片，是之前切片容量的两倍
	nb1 := make([]int, len(nb), (cap(nb))*2)

	// 拷贝nb到新切片中
	copy(nb1, nb)
	printSlice(nb1)

}
func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
