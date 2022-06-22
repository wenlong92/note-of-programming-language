package main

import (
	"fmt"
)

func main() {

	key := ""
	loop := true

	// 收支说明
	details := "收支\t账户金额\t收支金额\t说   明"
	// 余额
	balance := 0.0
	// 金额
	var money float64
	// 说明
	var note string

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
				fmt.Println(details)
			case "2":
				fmt.Println("收入金额：")
				fmt.Scanln(&money)
				fmt.Println("收入说明：")
				fmt.Scanln(&note)
				balance += money
				details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)
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