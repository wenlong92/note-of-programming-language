package main

import "fmt"

func main() {
	// 引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数
	// 引用传递指针参数传递到函数内

	var a, b int = 100, 200
	fmt.Printf("交换前a的值为：%d \n", a)
	fmt.Printf("交换前b的值为：%d \n", b)
	/*调用函数swap
	&a指向a指针，a变量的地址
	&b指向b指针，b变量的地址
	*/
	swap(&a, &b)
	fmt.Printf("交换后a的值为：%d \n", a)
	fmt.Printf("交换后b的值为：%d \n", b)
}

func swap(x, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

// 其他运算符： & 返回变量存储地址，&a将给出变量的实际地址； * 指针变量，*a是一个指针变量；
