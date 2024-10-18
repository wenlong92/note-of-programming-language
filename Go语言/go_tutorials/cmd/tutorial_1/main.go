package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("hello world")
	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 5
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	var myString1 string = "hello \nworld"
	var myString2 string = `hello
world`
	fmt.Println(myString1)
	fmt.Println(myString2)

	var myString3 string = "hello" + " " + "world"
	fmt.Println(myString3)

	// 获取字符串长度
	fmt.Println(utf8.RuneCountInString("Y"))

	var myRune rune = 'a'
	fmt.Println(myRune)

	myVar := "text"
	fmt.Println(myVar)

	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	//常量
	const pi float32 = 3.1415
	fmt.Println(pi)

	var printValue string = "printMe"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 2
	var result1 int = intDivision(numerator, denominator)
	fmt.Println(result1)
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) int {
	var result int = numerator / denominator
	return result
}
