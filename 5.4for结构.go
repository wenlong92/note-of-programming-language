package main

import (
	"fmt"
)

func main() {
	// 想重复某些语句，Go语言中只有for结构可以使用
	// 其他许多语言中没有发现与do...while完全对象的for结构，可能这种需求不强烈

	// 最简单的基于计数器的迭代，基本形式为：
	// for 初始化语句；条件语句；修饰语句 {}
	for i := 0; i < 5; i++ {
		fmt.Printf("this is the %d iteration\n", i)
	}
	// 不要在循环体内修改计数器

	// 还可以在循环中同时使用多个计数器
	// for i, j := 0, N; i < j; i, j = i+1, j-1 {}
	// 得益于Go语言具有平行赋值的特性

	// for循环迭代一个Unicode编码的字符串
	str1 := "日本語"
	fmt.Printf("The length of str1 is : %d\n", len(str1))
	for ix := 0; ix < len(str1); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str1[ix])
	}

	for iy := 1; iy < 25; iy++ {
		fmt.Printf("demo")
		var str2 = ""
		var str3 string = ""
		for jy := 1; jy < iy; jy++ {
			str3 = str2 + "G"
		}
		fmt.Printf("%f\n", str3)
	}
}