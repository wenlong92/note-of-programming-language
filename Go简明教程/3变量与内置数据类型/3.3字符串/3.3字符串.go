package main

import (
	"fmt"
	"reflect"
)

func main() {
	str1 := "Golang"
	str2 := "Go语言"
	fmt.Println(reflect.TypeOf(str2[2]).Kind()) //字符串已byte 数组形式保存，类型 uint8，占 1 个 byte，打印时需用 string 进行类型转换，否则是编码值
	fmt.Println(str1[2], string(str1[2]))
	fmt.Printf("%d %c\n", str2[2], str2[2])
	fmt.Println("len(str2):", len(str2))

	//正确的方式是将 string 转为rune 数组，转成[]rune类型后，字符串的每个字符，都用 int32表示
	str3 := "Go语言"
	runeArr := []rune(str3)
	fmt.Println(reflect.TypeOf(runeArr[2]).Kind())
	fmt.Println(runeArr[2], string(runeArr[2]))
	fmt.Println("len(runeArr):", len(runeArr))
}
