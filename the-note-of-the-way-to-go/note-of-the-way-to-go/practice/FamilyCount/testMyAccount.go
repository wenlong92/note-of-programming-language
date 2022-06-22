package main

import (
	"fmt"
)

func main() {

	key := ""
	loop := true

	for {
		fmt.Println("\n----------家庭收支记账软件----------")
		fmt.Println("             1 收支明细")
		fmt.Println("             2 登记收入")
		fmt.Println("             3 登记支出")
		fmt.Println("             4 退出软件")
		fmt.Println("请选择(1-4):")
	
		fmt.Scanln(&key)
	
		switch key {
			case "1":
				fmt.Println("已选择收支明细")
			case "2":
				fmt.Println("已选择登记收入")
			case "3":
				fmt.Println("已选择登记支出")
			case "4":
				fmt.Println("已退出软件")
				loop = false
			default:
				fmt.Println("输入错误，请输入正确的选项")
		}
	
		if !loop {
			break
		}
	}


}