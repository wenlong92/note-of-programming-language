package main

import (
	"fmt"
	"math"
)

func main() {
	// 一个接口类型的变量varI中可以包含任何类型的值，必须有一种方式来检测它的动态类型，即运行时在变量中存储的值的实际类型
	// 在执行过程中动态类型可能会有所不同，但是它总可以分配给接口变量本身的类型。
	// 通常可以用类型断言来测试在某个时刻varI是否包含类型T的值：
	// v := varI.(T)
	// varI必须是一个接口变量，否则编译器会报错

	// 类型断言可能会无效，因为他不可能预见所有的可能性。
	// 如果转换在程序运行时失败会导致错误发生。更安全的类型断言方式如下：

	// if v, ok := varI.(T); ok {
	// 	Process(v)
	// 	return 
	// }

	// 应该总是使用上面的方式来进行类型断言
	
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5

	areaIntf = sq1
	if t, ok := areaIntf.(*Square); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}
	if u, ok := areaIntf.(*Circle); ok {
		fmt.Printf("The type of areaIntf is: %T\n", u)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}

}

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

type Shaper interface {
	Area() float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}