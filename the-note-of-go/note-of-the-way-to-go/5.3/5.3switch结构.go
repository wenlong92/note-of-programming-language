package main
import (
	"fmt"
)

func main() {
	// 变量var1可以是任何类型，而val1和val2则可以是同类型的任意值。
	// 类型不局限于常量或整数，但必须是同类型；或者最终结果为相同类型的表达式
	// 一旦成功匹配到某个分支，执行完相应代码后悔退出整个switch代码块，不需要特别使用break表示结束
	// switch var1 {
	// 	case val1:
	// 		...
	// 	case val2:
	// 		...
	// 	default:
	// 		...
	// }

	// 如果在执行完每个分支代码后，还希望继续执行后续分支代码，使用fallthrough关键字
	// switch i {
	// case 0: fallthrough
	// case 1: 
	// 	f()  //当 i == 0时函数也会被调用
	// }

	var num1 int = 100

	switch num1 {
		case 98, 99:
			fmt.Println("equal to 98 or 99\n")
		case 100:
			fmt.Println("equal to 100\n")
		default:
			fmt.Println("not equal to 98 or 100\n")
	}

	// switch第二种形式是不提供任何被判断的值(默认为判断是否为true)
	// 然后在每个case分支中进行测试不同的条件，任一分支测试结果为true时，执行该分支代码
	// 向链式if-else语句，但在测试条件很多时，提供了更好的书写方式
	// switch {
	// 	case i < 0:
	// 		f1()
	// 	case i == 0:
	// 		f2()
	// 	case i > 0:
	// 		f3()
	// }
	// 任何支持进行相等判断的类型都可以作为测试表达式的条件，像int/string/指针等
	var num2 int = 7
	switch {
		case num2 < 0:
			fmt.Println("num2 is negative\n")
		case num2 > 0 && num2 < 10:
			fmt.Println("num2 is between 0 and 10\n")
		default:
			fmt.Println("num2 is 10 or greater\n")
	}

	// switch第三种形式是包含一个初始化语句
	// 例如下面代码中，变量a和b被平行初始化，然后作为判断条件
	// switch a, b := x[i], y[i]; {
	// 	case a < b: t = -1
	// 	case a == b: t = 0
	// 	case a > b: t = 1
	// }
	
	
	
}