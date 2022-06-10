package main

import (
	"fmt"
	"time"
)

const LIM = 41

var fibs [LIM]uint64

func main() {
	// 进行大量计算时，提升性能最直接有效的一种方式就是避免重复计算
	// 通过在内存中缓存和重复利用相同计算的结果，称之为内存缓存。
	// 最明显的例子就是生成斐波那契数列的程序(6.6/6.11小节)

	// 还可以应用到其他类型计算中，例如使用map而不是数组或切片
	var result uint64 = 0
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longcalculation took this amount of time: %s\n", delta)

	// 内存缓存技术在使用计算成本相对昂贵的函数时非常有用(不仅限于例子中的递归)，譬如大量进行相同参数的运算
	// 这种技术还可以应用于纯函数中，即相同输入必定获得相同输出的函数
}

func fibonacci(n int) (res uint64) {
	if fibs[n] != 0 {
		res = fibs[n]
		return 
	}
	if n <= 1{
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	fibs[n] = res
	return 
}