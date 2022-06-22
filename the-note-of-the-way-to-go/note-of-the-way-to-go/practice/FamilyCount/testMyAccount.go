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
	// 标记收支记录
	flag := false

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
				if flag {
					fmt.Println("----------当前收支明细记录----------")
					fmt.Println(details)
				} else {
					fmt.Println("当前还没有收支记录，来记录一笔吧")
				}
			case "2":
				fmt.Println("收入金额：")
				fmt.Scanln(&money)
				fmt.Println("收入说明：")
				fmt.Scanln(&note)
				balance += money
				details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)
				flag = true
			case "3":
				fmt.Println("支出金额：")
				fmt.Scanln(&money)
				if money > balance {
					fmt.Println("余额中的金额不足")
					break
				}
				fmt.Println("支出说明：")
				fmt.Scanln(&note)
				balance -= money
				details += fmt.Sprintf("\n支出\t%v\t%v\t%v", balance, money, note)
				flag = true
			case "4":
				fmt.Println("确定要退出软件吗？Y/N")
				quit := ""
				for {
					fmt.Scanln(&quit)
					if quit == "Y" || quit == "N" {
						break
					}
					fmt.Println("你的输入有误，请重新输入 Y/N")
				}
				if quit == "Y" {
					fmt.Println("已退出软件")
					loop = false
				}
			default:
				fmt.Println("输入错误，请输入正确的选项")
		}
	
		if !loop {
			break
		}
	}


}