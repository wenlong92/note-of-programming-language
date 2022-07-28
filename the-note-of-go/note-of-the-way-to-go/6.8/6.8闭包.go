package main
import (
	"fmt"
)

func main() {
	// 不希望给函数起名字时，可以使用匿名函数
	// 匿名函数不能独立存在，但可以被赋值某个变量，即保存函数地址到变量中
	// fplus := func(x, y int) int { return x + y }，然后通过变量名对函数进行调用：fplus(3, 4)
	// 也可以直接通过匿名函数进行调用：func(x, y int) int { return x + y }(3, 4)
	func() {
		sum := 0
		for i := 0; i <= 1e6; i++ {
			sum ++
		}
	}()
	
	// 下例展示通过匿名函数赋值给变量并对其进行调用：
	f()

	// defer语言经常配合匿名函数使用，用于改变函数的命名返回值 
	fmt.Println(f1())

	// 匿名函数同样被称之为闭包(函数式语言的术语)：它们被允许调用定义在其他环境下的变量。
	// 闭包可使得某个函数捕捉到一些外部状态
}

func f() {
	for i := 0; i < 4; i++ {
		g := func(i int) { fmt.Printf("%d ", i)}
		g(i)
		fmt.Printf(" - g is of type %T and has value %v\n", g, g)
	}
}

func f1() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}