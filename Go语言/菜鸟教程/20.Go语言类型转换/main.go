package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 类型转换用于将一种数据类型的变量转换为另一种类型的变量

	// 数值类型转换：整型转换为浮点型
	var a int = 10
	var b float64 = float64(a)
	fmt.Println(b)

	var sum int = 17
	var count int = 5
	var mean float32
	mean = float32(sum) / float32(count)
	fmt.Printf("mean=%f\n", mean)

	// 字符串类型转换
	var str string = "10"
	var num1 int
	num1, _ = strconv.Atoi(str) // strconv返回两个值，一个是转换后的整型值，第二个是可能发生的错误，可用_来忽略这个错误
	fmt.Println(num1)

	str1 := "123"
	num2, err := strconv.Atoi(str1)
	if err != nil {
		fmt.Println("转换错误：", err)
	} else {
		fmt.Printf("str1转换为整数后，为:%d\n", num2)
	}

	// 整数转换为字符串
	num3 := 1212
	str2 := strconv.Itoa(num3)
	fmt.Printf("%s\n", str2)

	// 字符串转换为浮点数
	str3 := "3.14"
	num4, err := strconv.ParseFloat(str3, 64)
	if err != nil {
		fmt.Println("转换错误：", err)
	} else {
		fmt.Printf("%f \n\n", num4)
	}

	// 浮点数转换为字符串
	num5 := 3.14
	str4 := strconv.FormatFloat(num5, 'f', 2, 64)
	fmt.Printf("%s\n", str4)

	// 接口类型转换有两种情况：类型断言和类型转换
	// 类型断言用于将接口类型转换为指定类型
	var i interface{} = "Hello, World"
	str5, ok := i.(string)
	if ok {
		fmt.Printf("%s\n", str5)
	} else {
		fmt.Println("failed")
	}
	// 类型转换必须保证要转换的值喝目标接口类型之间时兼容的
	var w Writer = &StringWriter{}
	sw := w.(*StringWriter)
	sw.str = "Hello, World"
	fmt.Println(sw.str)
}

type Writer interface {
	Write([]byte) (int, error)
}
type StringWriter struct {
	str string
}

func (sw *StringWriter) Write(data []byte) (int, error) {
	sw.str += string(data)
	return len(data), nil
}
