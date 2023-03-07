package main

import "fmt"

func main() {
	// Go语言中取地址符是&，放到一个变量前使用就会返回相应变量的内存地址，例如
	var aa int = 10
	fmt.Printf("变量地址：%x\n", &aa)

	// 指针：一个指针变量指向了一个值的内存地址。使用指针前需要声明
	//var ip *int       指向整型
	//var fp *float32   指向浮点型
	var a int = 10
	var ipp *int
	ipp = &a
	fmt.Printf("a的变量地址：%x\n", &a)
	fmt.Printf("ipp变量存储的指针地址：%x\n", ipp)
	fmt.Printf("*ipp变量的值：%d\n", *ipp)

	// 空指针：指针被定义后没分配到变量时，值为null
	var ptr *int
	fmt.Printf("ptr的值为：%x\n", ptr)
	if ptr != nil {
		fmt.Printf("不是空值")
	} else {
		fmt.Printf("是空值")
	}

}
