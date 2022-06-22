package main
import (
	"strings"
	"fmt"
)

func main() {
	// Add2和Adder均返回签名为func(b int) int 的函数：
	p2 := Add2()
	fmt.Printf("Call Add2 for 3 gives: %v\n",p2(3))

	TwoAdder := Adder(2)
	fmt.Printf("The result is: %v\n", TwoAdder(3))

	// 下例是一个略微不同的实现：
	var f = Adder1()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300))
	// 在多次调用中，变量x的值被保留
	// 闭包函数保存并积累其中的变量的值，不管外部函数退出与否，都能继续操作外部函数中的局部变量

	// 闭包中使用到的变量既可以在闭包函数体内声明，也可在外部函数声明：
	var g int
	func(i int) {
		s := 0
		for j := 0; j < i; j++ { s += j }
		g = s
	}(1000)
	fmt.Printf("\ng %d", g)

	// 不使用递归但使用闭包改写斐波那契数列函数
	f2 := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Print("\n", f2())
	}
	fmt.Println("...")

	// 一个返回值为另一个函数的函数称之为工厂函数
	// 在需要创建一系列相似的函数时非常有用
	// 书写一个工厂函数而不是针对没中情况都书写一个函数
	// 下例为动态返回追加后缀的函数
	addBmp := MakeAddSuffix(".bmp")
	addJpeg := MakeAddSuffix(".jpeg")

	bmpName := addBmp("file")
	jpegName := addJpeg("file")
	fmt.Printf(bmpName,"\n")
	fmt.Printf(jpegName,"\n")

	// 可以返回其他函数的函数和接受其他函数作为参数的函数被称之为高阶函数，是函数式语言的特点
}

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

func Adder1() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}

func fibonacci() func() int {
	x1, x2 := 0, 1
	
	return func() int {
		sum := x2
		x2 = x1 + x2
		x1 = sum
		return sum 
	}
}

func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}