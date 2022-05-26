package main
import (
	"fmt"
	"runtime"
	"os"
)
var a, b = 2, 3
// 因式分解关键字，一般用于声明全局变量
var (
	c int
	d bool
	str string = ""
)
// 变量声明后，系统自动赋予该类型的初始值，int 0,float 0.0,bool false, string空字符串,指针 nil
func main() {
	fmt.Println(a, b, c, d, str)
	// 变量命名规则遵循骆驼命名法，如numShips
	// 如果希望被外部包使用，首个单词首字母也大写

	// 函数体内声明局部变量时，可使用简短声明语法 :=
	e := 1
	fmt.Println(e)

	goos := runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	var path string = os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
	demo1 := 1
	var demo2 int
	demo2 = demo1
	fmt.Println(demo2)
	fmt.Println(demo1)
	demo1 = 2
	fmt.Println(demo2)
	fmt.Println(demo1)
	fmt.Print("Hello:", 23)

	// 基本类型和数组、结构等复合类型属于值类型，这些类型的变量直接指向内存中的值
	// 指针、slices、maps、channel属于引用类型

	// 同类型多个变量可以声明在同一行
	// var as, ta, taa int

	// 多变量可以在同一行赋值
	// as, ta, taa = 3, 4, "jkl"
	// 也可以写成
	// as, ta, taa := 3, 4, "jkl"

	// 交换两个变量的值
	a, b = b, a
	fmt.Println(a, b)

	// 空白标识符 _ 用于抛弃值，_ 实际上是一个只写变量，不能得到它的值
	_, b = 5, 7

	// 并行赋值也用于当函数返回多个返回值时
	// val, err = Func1(val1)

	// init函数不能被人为调用，是在每个包完成初始化后自动执行的，优先级比main函数高
	// 每个源文件只能包含一个init函数。初始化总是以单线程执行，按照包的依赖关系顺序执行。
	// 一个可能的用途是在开始执行程序之前对数据进行检验或修复，保证程序状态的正确性
	
}
