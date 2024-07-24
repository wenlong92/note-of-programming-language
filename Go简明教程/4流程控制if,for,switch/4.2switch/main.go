package main

import "fmt"

func main() {
	//type关键字定义一个新的类型 Gender
	type Gender int8
	//const 定义 MALE 和 FEMALE 两个常量
	const (
		MALE   Gender = 1
		FEMALE Gender = 2
	)

	gender := MALE

	//Go 语言中 switch 不需要 break，匹配到某个结果后，默认不会继续执行
	switch gender {
	case FEMALE:
		fmt.Println("female")
	case MALE:
		fmt.Println("male")
	default:
		fmt.Println("unknown")
	}

	//如果需要继续执行，使用 fallthrough
	switch gender {
	case FEMALE:
		fmt.Println("female")
		fallthrough
	case MALE:
		fmt.Println("male")
		fallthrough
	default:
		fmt.Println("unknown")
	}

}
