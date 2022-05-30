package main
import (
	"fmt"
)

func main() {
	// Go语言的函数经常使用两个返回值表示执行是否成功，返回true 返回零值(或nil)和false
	// 不适用true或false时，也可以使用一个error类型的变量代替，作为第二个返回值，成功则error为nil，否则会包含相应的错误信息
	// 需要用一个if语句测试执行结果，由于其符号的原因，这种形式又被称之为comma,ok模式
	// 习惯用法：
	// value, err := pack1.Function1(param1)
	// if err != nil {
	// 	fmt.Printf("An error occured in pack1.Function1 with paramter %v", param1)
	// 	return err
	// }
	

}