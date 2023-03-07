package main

import "fmt"

func main() {
	// 数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，这种类型可以是任意的原始类型（整型、字符串等）或自定义类型
	// 数组元素可以通过索引（位置）来读取或修改

	// 数组声明：需要指定元素类型及元素个数
	// var balance [10] float32

	// 初始化数组
	var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Printf("balance:", balance, "\n")

	// 也可通过字面量在声明数组的同时快速初始化
	balance1 := [5]float32{10000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Printf("balance1:", balance1, "\n")

	// 如果长度不确定，可以使用...代替数组长度，编译器会根据元素个数自行推断数组长度
	var balance2 = [...]float32{100.0, 3.4, 5.0, 6.0}
	fmt.Printf("balance2:", balance2, "\n")
	// 或
	balance3 := [...]float32{1.9, 2.3, 6.7}
	fmt.Printf("balance3:", balance3, "\n")

	// 如果设置了数组长度，还可以指定下标初始化元素
	balance4 := [5]float32{1: 2.1, 3: 7.8}
	fmt.Printf("balance4:", balance4, "\n")

	// {}中元素个数不能大于[]数字
	// 如果忽略[]中的数字不设置数组大小，Go语言会根据元素的个数来设置数组的大小
	balance4[4] = 50.1
	fmt.Printf("balance4:", balance4, "\n")

	// 数组元素可以通过索引来读取
	var n [10]int
	var i, j int
	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}

	var i1, j1, k1 int
	balance5 := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	for i1 = 0; i1 < 5; i1++ {
		fmt.Printf("balance5[%d] = %f\n", i1, balance5[i1])
	}
	balance6 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	for j1 = 0; j1 < 5; j1++ {
		fmt.Printf("balance6[%d] = %f\n", j1, balance6[j1])
	}
	balance7 := [5]float32{1: 2.0, 3: 7.0}
	for k1 = 0; k1 < 5; k1++ {
		fmt.Printf("balance7[%d] = %f\n", k1, balance7[k1])
	}
}
