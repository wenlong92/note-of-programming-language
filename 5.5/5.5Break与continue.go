package main
import (
	"fmt"
)

func main() {
	i := 5
	for {
		i = i - 1
		fmt.Printf("The variable i is now: %d\n", i)
		if i < 0 {
			break
		}
	}
	// break在for循环中的作用范围是该语句出现后的最内部的结构
	// 在switch或select语句中，break语句的作用是跳过整个代码块，执行后续代码
	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			if j > 5 {
				break
			}
			print(j)
		}
	}

	// continue忽略剩余的循环体，直接进入下一次循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		print(i)
		print(" ")
	}
	// continue只能用于for循环中
}