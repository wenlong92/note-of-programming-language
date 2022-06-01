package main
import (
	"errors"
	"fmt"
	"math"
)

func main() {

	// 函数可以返回零个或多个值
	// 任何有返回值(单个或多个)的函数，都必须以return或panic结尾

	// Go默认使用按值传递参数，也就是传递参数的副本。使用变量的过程中，对副本的更改不会影响原来的变量
	// 也可以传递参数地址(引用传递)，函数可以直接修改参数的值，此时传递的是一个指针

	// 几乎任何情况下，传递指针的消耗都比副本小

	// 切片(slice)，字典(map)、接口(interface)、通道(channel)这样的引用类型都是默认使用引用传递
	demo1(1, 2)
	MySqrt(2.8)
}

func demo1(num1, num2 int) (num3, num4, num5 int) {
	num3 = num1 + num2
	num4 = num1 * num2
	num5 = num2 - num1
	fmt.Printf("%d,%d,%d\n",num3,num4,num5)
	return
}

func MySqrt(num1 float64) (float64, error) {
	if(num1 < 0) {
		return float64(math.NaN()), errors.New("I won't be able to do a sqrt of negative number!")
	}
	return math.Sqrt(num1), nil
}