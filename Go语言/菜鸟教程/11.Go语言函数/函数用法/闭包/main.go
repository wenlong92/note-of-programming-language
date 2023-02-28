package main

import "fmt"

func main() {
	// Go语言支持匿名函数，可作为闭包。
	// 匿名函数是一个"内联"语句或表达式。优越性在于可以直接使用函数内的变量，不必申明。
	nextNumber := getSequence()
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber()) // 此时nextNumber的值并没有被丢弃 这里为4
	fmt.Println(nextNumber1())

}
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
