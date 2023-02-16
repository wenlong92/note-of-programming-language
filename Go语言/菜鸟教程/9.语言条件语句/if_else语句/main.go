package main

import "fmt"

func main() {
	a := 100

	if a < 20 {
		fmt.Printf("a小于20\n")
	} else {
		fmt.Printf("a不小于20\n")
	}
	fmt.Printf("a的值为：%d", a)
}
