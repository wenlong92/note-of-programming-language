package main

import "fmt"

//参数可以有多个，返回值也支持有多个
//func funcName(param1 Type1, param2 Type2, ...) (reture1 Type3, ...){ }

func add(num1 int, num2 int) int {
	return num1 + num2
}

func div(num1 int, num2 int) (int, int) {
	return num1 / num2, num1 % num2
}

func main() {
	quo, rem := div(100, 17)
	fmt.Println(quo, rem)
	fmt.Println(add(100, 17))
	fmt.Println(add1(100, 17))
}

// 也可以给返回值命名，简化 return
func add1(num1 int, num2 int) (ans int) {
	ans = num1 + num2
	return
}
