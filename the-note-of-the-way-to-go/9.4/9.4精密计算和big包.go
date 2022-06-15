package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	// 有些时候通过编程方式进行计算是不精确的
	// 使用Go语言float64类型进行浮点运算，返回结果精确到15位，对精度有严格要求时，不能使用浮点数，在内存中它们只能被近似表示

	// 整数的高精度计算，Go中提供了big包，包含在math包中
	// big.Int表示大整数  big.Rat表示大有理数(可以表示2/5或3.1416这样的分数，而不是无理数或π)
	// 这些类型可以实现任意位类型的数字，只要内存足够大。缺点是更大的内存和处理开销，处理速度比内置数字类型慢很多

	// 大的整型数字通过big.NewInt(n)来构造，n 为int64类型整数
	// 大有理数通过big.NewRat(n, d)方法构造 n 为分子  d为分母  都是int64型整数
	// Go语言不支持运算符重载(对已有的运算符重新进行定义，赋予另一种功能，以适应不同的数据类型)，
	// 所以大数字类型都有Add() Mul()这种方法。他们作用于作为receiver的整数和有理数
	// 大多数情况下它们修改receiver并以receiver作为返回结果
	// 因为没有必要闯将big.Int类型的临时变量来存放中间结果，所以运算可以被链式的调用，节省内存

	im := big.NewInt(math.MaxInt64)
	in := im
	io := big.NewInt(1956)
	ip := big.NewInt(1)
	ip.Mul(im, in).Add(ip, im).Div(ip, io)
	fmt.Printf("Big Int: %v\n", ip)
	
	demo1 := big.NewInt(12)
	demo2 := big.NewInt(4)
	demo1.Div(demo1, demo2)
	fmt.Printf("Big Int: %v\n", demo1)
}