package main
import (
	"fmt"
)

func main() {
	// 函数可以作为其他函数的参数进行传递，然后在其他函数内调用执行，一般称之为回调
	callback(1, add)

	// 将函数作为参数最好的例子是函数strings.IndexFunc()
}

func add(a, b int) {
	fmt.Printf("the sum of %d and %d is: %d\n", a, b, a + b)
}

func callback(y int, f func(int, int)) {
	f(y, 2)
}