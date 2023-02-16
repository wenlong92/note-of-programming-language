package main

import "fmt"

func main() {
	// switch语句从上至下，知道找到匹配项，匹配项后面不需要再加break
	// 默认情况下case最后自带break语句，匹配成功后不会执行其他case，如果需要执行后面的case，可使用fallthrough

	//switch var1 {
	//	case val1:
	//		...
	//	case val2:
	//		...
	//	default:
	//		...
	//}
	// 其中var1可以为任何类型，但val1和val2需是同类型或者最终结果为相同类型的表达式

	var grade string = "B"
	marks := 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀！\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
	fmt.Printf("你的等级是 %s\n", grade)

	// fallthrough会强制执行后面的case语句，fallthrough不会判断下一条case表达式结果是否为true
	switch {
	case false:
		fmt.Printf("1.case条件语句为false\n")
		fallthrough
	case true:
		fmt.Printf("2.case条件语句为true\n")
		fallthrough
	case false:
		fmt.Printf("3.case条件语句为false\n")
		fallthrough
	case true:
		fmt.Printf("4.case条件语句为true\n")
	case false:
		fmt.Printf("5.case条件语句为false\n")
		fallthrough
	default:
		fmt.Printf("6.默认case")
	}

	// switch语句可以用于type switch来判断某个interface变量中实际存储的变量类型
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf("x的类型是%T", i)
	case int:
		fmt.Printf("x是int型")
	case float64:
		fmt.Printf("x是float64型")
	case func(int) float64:
		fmt.Printf("x是func(int)型")
	case bool, string:
		fmt.Printf("x是bool或string型")
	default:
		fmt.Printf("未知型")
	}
}
