package main

import "fmt"

func main() {
	// 使用递归时，需要设置退出条件，否则递归将陷入无限循环中
	var i int = 15
	fmt.Printf("%d的阶乘是 %d\n", i, Factorial(uint64(i)))

	// 斐波那切数列
	var n int
	for n = 0; n < 10; n++ {
		fmt.Printf("%d\t", fibonacci(n))
	}
}

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}
