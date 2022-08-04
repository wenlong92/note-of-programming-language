package main

import (
	"fmt"
)

func main() {
	// 切片创建时通常比相关数组小
	// slice1 := make([]tyep, start_length, capacity)
	// start_length作为切片初始长度而capacity作为相关数组的长度
	// 这样做的好处是切片在达到容量上限后可以扩容
	// 改变切片长度的过程称之为切片重组 reslicing
	// 做法如下： slice1 = slice1[0:end]，其中end是新的末尾索引(即长度)

	// 将切片扩展1位可以这样做： sl = sl[0:len(sl)+1]
	// 切片可以反复扩展知道占据整个相关数组
	slice1 := make([]int, 0, 10)
	for i := 0; i < cap(slice1); i++ {
		slice1 = slice1[0:i+1]
		slice1[i] = i
		fmt.Printf("The length of slice is %d\n", len(slice1))
	}
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	var ar = [10]int{0,1,2,3,4,5,6,7,8,9}
	var a = ar[5:7]
	fmt.Print(a,"\n")
	fmt.Print(len(a),"\n")
	fmt.Print(cap(a),"\n")
	a = a[0:4]
	fmt.Print(a,"\n")
	fmt.Print(len(a),"\n")
	fmt.Print(cap(a),"\n")

	slice2 := make([]int, 0, 10)
	fmt.Print(len(slice2[4:4]),"\n")
	fmt.Print(len(slice2[4:5]),"\n")
}
