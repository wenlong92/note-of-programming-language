package main
import (
	"fmt"
)

func main() {
	// 函数最后一个参数若是...type形式，则此函数可以处理一个变长的参数
	// 长度可以为0，这样的函数成为变参函数
	// func myFunc(a, b, arg ...int) {}

	// 如果参数存储在一个slice类型变量中，可以通过slice...形式传递参数
	x := min(1, 3, 2, 0)
	fmt.Printf("The minimum is: %d\n", x)
	slice := []int{7,9,3,5,1}
	x = min(slice...)
	fmt.Printf("The minimum in the slice is: %d\n", x)
	print(slice)

	// 如果变长参数类型不相同，有两个方案解决，1.使用结构；2.使用空接口
	// 1.使用结构
	// type Options struct {
	// 	par1 type1,
	// 	par2 type2,
	// 	...
	// }

	// 2.使用空接口：如果变长参数类型没有指定，可以使用默认的空接口 interface{}，这样可以接受任何类型的参数
}

func min(s ...int) int {
	if len(s)==0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}