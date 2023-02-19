package main

import "fmt"

func main() {
	// for循环有三种形式：
	// 1. for init; condition; post { }  一般为赋值循环表达式，给控制变量赋初始值
	// 2. for condition { }              关系表达式或逻辑表达式，循环控制条件
	// 3. for { }                        一般为赋值表达式，给控制变量增量或减量

	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum1 := 1
	for sum1 <= 10 {
		sum1 += sum1
	}
	fmt.Println(sum1)

	// For-each range 循环
	strings := []string{"gogle", "runoob"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	numbers := [6]int{1, 2, 3, 5}
	for i, x := range numbers {
		fmt.Printf("第%d位x的值%d\n\n", i, x)
	}

	// for循环的range格式可以省略key和value
	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0
	for key, value := range map1 {
		fmt.Printf("key is: %d -value is: %f\n", key, value)
	}
	for key := range map1 {
		fmt.Printf("key is: %d\n", key)
	}
	for _, value := range map1 {
		fmt.Printf("value is: %f\n", value)
	}

}
