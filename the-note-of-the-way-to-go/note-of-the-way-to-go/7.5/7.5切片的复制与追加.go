package main

import(
	"fmt"
)

func main() {
	// 如果想增加切片的容量，必须创建一个新的更大的切片并把原分片的内容都拷贝过来
	// 下例描述了从拷贝切片的copy函数和向切片追加新元素的append函数
	slFrom := []int{1, 2, 3}
	slTo := make([]int, 10)
	n := copy(slTo, slFrom)
	fmt.Println(slTo)
	fmt.Printf("Copied %d elements\n",n)
	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
	// func append(s[]T, x...) []T
	// 其中append方法将0个或多个具有相同类型s的元素追加到切片后面并且返回新的切片
	// 追加的元素必须和原切片元素是相同类型
	// 如果s的容量不足以存储新增元素，append会分配新的切片，因此，返回的切片可能已经指向一个不同的相关数组
	// append方法总是返回成功，除非系统内存耗尽
}