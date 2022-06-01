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

	// for i := 1; i <= 25; i++ {
	// 	for j := 1; j <= i; j++ {
	// 		print("G")
	// 	}
	// 	println()
	// }


	// for结构第二种形式是没有头部的条件判断迭代
	var int2 int = 5
	for int2 >= 0 {
		int2 = int2 -1
		fmt.Printf("The variable int2 is now: %d\n", int2)
	}

	// for-range结构 Go特有的迭代结构  可以迭代任何一个集合(包括数组和map)
	// 类似其他语言的foreach语句

	str5 := "Go is a beautiful language!"
	for pos, char := range str5 {
		fmt.Printf("Character on position %d is: %c \n", pos, char)
	}

	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v)
		v = 5
	}

	// for i := 0; ; i++ {
	// 	fmt.Println("Value of i is now:", i)
	// }

	// for i := 0; i < 3; {
	// 	fmt.Println("Value of i:", i)
	// }

	s := ""
	for ; s != "aaaaa"; {
		fmt.Println("Value of s:", s)
		s = s + "a"
	}

	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
		s = i+1, j+1, s + "a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}
}