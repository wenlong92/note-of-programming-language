package main
import (
	"fmt"
)

func main() {
	// for、switch、select语句都可以配合标签形式的标识符使用
	// 标签一般大写
	LABEL1:
		for i := 0; i <= 5; i++ {
			for j := 0; j <=5; j++ {
				if j == 4 {
					break LABEL1
				}
				fmt.Printf("i is: %d, and j is: %d\n", i, j)
			}
		}

		// 可以使用goto语句和标签配合使用，模拟循环
		// 尽量不适用goto和标签，如果使用，则goto应当放在标签前面，且goto和标签之间不能出席那定义新变量的语句，否则编译失败



		i := 0
		for { //since there are no checks, this is an infinite loop
			if i >= 3 { break }
			//break out of this for loop when this condition is met
			fmt.Println("Value of i is:", i)
			i++
		}
		fmt.Println("A statement just after for loop.")
		

		for i := 0; i<7 ; i++ {
			if i%2 == 0 { continue }
			fmt.Println("Odd:", i)
		}
}