package main

import "fmt"

func main() {
	// continue有点像break，但continue不是跳出循环，而是跳过当前循环执行下一次循环语句
	// for循环中，执行continue语句会触发for增量语句的执行
	// 在多重循环中，可用标号label标出想continue的循环
	var a int = 10
	for a < 20 {
		if a == 15 {
			a = a + 1
			continue
		}
		fmt.Printf("a 的值为: %d \n", a)
		a++
	}

	// 多重循环：使用标记和不使用标记的区别
	fmt.Println("----continue----")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d \n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d \n", i2)
			continue
		}
	}

	// 使用标记
	fmt.Println("----continue label----")
re:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d \n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d \n", i2)
			continue re
		}
	}
}
