package main

import "fmt"

func main() {
	// 传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，不会影响到实际参数
	// 默认情况下，Go语言使用的是值传递

	// 使用值传递调用swap函数
	var a, b int = 100, 200
	fmt.Printf("交换前a的值为：%d \n", a)
	fmt.Printf("交换前b的值为：%d \n", b)

	swap(a, b)
	fmt.Printf("交换后a的值为：%d \n", a)
	fmt.Printf("交换后b的值为：%d \n", b)
	// 程序中使用的是值传递，所以两个值并没有实现交换，可以使用引用传递来实现交换效果

}
func swap(x, y int) int {
	var temp int
	temp = x
	x = y
	y = temp
	return temp
}
