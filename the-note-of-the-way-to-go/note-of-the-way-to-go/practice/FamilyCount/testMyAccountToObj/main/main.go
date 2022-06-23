package main

import (
	"fmt"
	"../utils"
)

func main() {
	// 面向对象的实现方式
	fmt.Println("这是一个面向对象的实现方式")
	utils.NewFamilyAccount().MainMenu()
}