package main

import "fmt"

func main() {
	// 数组声明
	var arr [5]int     //一维
	var arr2 [5][5]int //二维

	//数组声明时初始化
	var arr3 = [5]int{1, 2, 3, 4, 5}
	arr4 := [5]int{1, 2, 3, 4, 5}

	//使用[]索引修改数组
	arr5 := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		arr[i] += 100
	}
	fmt.Println(arr5)

	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)

	//数组长度无法改变，如果想拼接两个数组，或获取子数组，需要使用切片。
	//切片是数组的抽象。切片使用数组作为底层结构。
	//切片包含三个组件：容量、长度和指向底层数组的指针。切片随时可以进行扩展

	//声明切片
	slice1 := make([]float32, 0)    //长度为 0 的切片
	slice2 := make([]float32, 3, 5) //[0 0 0]长度为 3，容量为 5 的切片
	fmt.Println(len(slice1), cap(slice1))
	fmt.Println(len(slice2), cap(slice2))

	//使用切片
	slice2 = append(slice2, 1, 2, 3, 4) //切片容量可以根据需要自动扩展
	fmt.Println(len(slice2), cap(slice2))

	//子切片
	sub1 := slice2[3:]
	fmt.Println(sub1)
	sub2 := slice2[:3]
	fmt.Println(sub2)
	sub3 := slice2[1:4]
	fmt.Println(sub3)

	//合并切片
	combined := append(sub1, sub2...)
	fmt.Println(combined)

	//声明切片时可以为切片设置容量大小，为切片预分配空间。
	//在实际使用过程中，如果容量不够，切片容量会自动扩展

	//sub2...是切片的解构写法，将切片解构为 N 个独立的元素
}
