package main

import (
	"fmt"
)

func main() {
	// 接口变量的类型也可以使用一种特殊形式switch来检测：type-switch
	// 例如：
	switch t := areaIntf.(type) {
	case *Square:
		fmt.Printf("Type Square %T with value %v\n", t, t)
	case *Circle:
		fmt.Printf("Type Circle %T with value %v\n", t, t)
	case nil:
		fmt.Printf("nil value: nothing to check?\n")
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}

	// 可以用type-switch进行运行时类型分析，但是在type-switch不允许有fallthrough
	// 如果仅仅是测试变量类型，不用它的值，name可以不需要赋值语句，例如：
	switch areaIntf.(type) {
	case *Square:
		//TODO
	case *Circle:
		//TODO
	default:
		//TODO
	}
}