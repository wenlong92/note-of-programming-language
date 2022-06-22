package main
import (
	"fmt"
)

func main() {
	// 函数在其函数体内调用自身，称之为递归。
	// 例如斐波那契数列,前两个数分别为1，从第三个数开始每个数均为前两个数之和
	result := 0
	for i := 0; i <= 10; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n",i, result)
	}

	// result1 := 0
	// index1 := 0
	// for i := 0; i <= 10; i++ {
	// 	result1, index1 = fibonacci1(i)
	// 	fmt.Printf("fibonacci(%d) is: %d\n",index1, result1)
	// }

	// 很多问题都可以用递归优雅的解决，例如快速排序算法

	// 使用递归时经常会遇到栈溢出的问题：一般出现在大量的递归调用导致的程序栈内存分配耗尽
	// Go语言中，可以使用管道(channel)和goroutine实现

	// Go语言中可以使用相互调用的递归函数：多个函数之间相互调用形成闭环。
	// 因为Go语言编译器的特殊性，这些函数的声明顺序可以任意。
	// 下例展示od和even函数之间相互调用
	fmt.Printf("%d is even: is %t\n", 16, even(16))
	fmt.Printf("%d is odd: is %t\n", 17, odd(17))
	fmt.Printf("%d is odd: is %t\n", 18, odd(18))
	fibonacci2(10)
}

func fibonacci2(n int) {
	if(n == 1) {
		fmt.Printf("%d",n)
		} else {
			fmt.Printf("%d\n",n)
			fibonacci2(n - 1)
		}
	return
}

func fibonacci1(n int) (res, resIndex int) {
	temIndex := n
	if n <= 1 {
		res = 1
	} else {
		res1, _ := fibonacci1(n - 1)
		res2, _ := fibonacci1(n - 2)
		res = res1 + res2
	}
	resIndex = temIndex
	return
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n - 1) + fibonacci(n - 2)
	}
	return
}

func even(nr int) bool {
	if nr == 0 {
		return true
	}
	return odd(revSign(nr) - 1)
}

func odd(nr int) bool {
	if nr == 0 {
		return false
	}
	return even(revSign(nr) - 1)
}

func revSign(nr int) int {
	if nr < 0 {
		return -nr
	}
	return nr
}