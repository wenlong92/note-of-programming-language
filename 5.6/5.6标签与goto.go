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
		
}