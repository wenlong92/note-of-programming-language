package main

import (
	"fmt"
)

type List []int

func (l List) Len() int {
	return len(l)
}

func (l *List) Append(val int) {
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

func CountInto(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}


func main() {
	// 在10.6.3中可以看出，作用于变量上的方法实际上是不区分变量到底是指针还是值
	// 当碰到接口类型值时，这回变得有点复杂，原因是接口变量中存储的具体值是不可寻址的
	// 幸运的是，如果使用不当编译器会给出错误，例如：
	var lst List
	if LongEnough(lst) {
		fmt.Printf("- lst is long enough\n")
	}

	plst := new(List)
	CountInto(plst, 1, 10)
	if LongEnough(plst) {
		fmt.Printf("- plst is long enough\n")
	}
}