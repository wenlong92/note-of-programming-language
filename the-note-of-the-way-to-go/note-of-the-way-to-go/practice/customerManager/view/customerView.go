package main

import (
	"fmt"
)

type customerView struct {
	key string  // 接收用户输入
	loop bool   // 表示是否循环显示主菜单
}

// 显示主菜单
func (this *customerView) mainMenu() {
	for {
		fmt.Println("----------客户信息管理软件----------")
		fmt.Println("          1 添 加 客 户")
		fmt.Println("          2 修 改 客 户")
		fmt.Println("          3 删 除 客 户")
		fmt.Println("          4 客 户 列 表")
		fmt.Println("          5 退      出")
		fmt.Print("请选择(1-5): ")

		fmt.Scanln(&this.key)
		
		switch this.key {
			case "1":
				fmt.Println("1")
			case "2":
				fmt.Println("2")
			case "3":
				fmt.Println("3")
			case "4":
				fmt.Println("4")
			case "5":
				this.loop = false
			default:
				fmt.Println("你的输入有误，请重新输入...")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("已退出软件系统")
}

func main() {
	// 在main函数中，创建一个CustomerView，并运行显示主菜单
	customerView := customerView{
		key: "",
		loop: true,
	}
	customerView.mainMenu()

}