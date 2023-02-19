package main

import "fmt"

func main() {
	// Go中break语句用于以下两个方面：
	// 1.用于循环语句中跳出循环，并开始执行循环之后的语句
	// 2.break在switch中执行一条case后跳出语句的作用
	// 3.在多重循环中，可以用标号label标出想break的循环

	var a int = 10
	for a < 20 {
		fmt.Printf("a的值为：%d \n", a)
		a++
		if a > 15 {
			break
		}
	}

	// 多重循环：使用标记和不适用标记的区别
	fmt.Println("----break----")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break
		}
		fmt.Println("hello")
	}
	fmt.Println("----break label ----")
re:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d \n", i2)
			break re
		}
	}
}
