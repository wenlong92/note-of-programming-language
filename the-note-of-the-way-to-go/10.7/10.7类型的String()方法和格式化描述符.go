package main

import (
	"fmt"
	"strconv"
)

type TwoInts struct {
	a int
	b int
}

func main() {
	// 当定义了一个很多方法的类型时，一般会使用String()方法来控制类型的字符串形式的输出。
	// 如果类型定义了String()方法，它会被用在fmt.Printf()中生成默认的输出：等同于使用格式化描述符 %v 产生的输出
	// 还有fmt.Print()和fmt.Println()也会自动使用String()方法
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10
	fmt.Printf("two1 is: %v\n", two1)
	fmt.Println("two1 is:", two1)
	fmt.Printf("two1 is: %T\n", two1)
	fmt.Printf("two1 is: %#v\n", two1)
	// 当你广泛使用一个自定义类型时，最好为它定义String()方法

	// 不要在String()方法里调用涉及String()方法的方法，会导致意料之外的错误。
}

func (tn *TwoInts) string() string {
	return "(" + strconv.Itoa(tn.a) + "/" + strconv.Itoa(tn.b) + ")"
}