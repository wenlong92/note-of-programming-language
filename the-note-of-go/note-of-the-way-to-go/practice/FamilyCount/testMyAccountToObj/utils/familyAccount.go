package utils

import "fmt"

type FamilyAccount struct {
	key string
	loop bool
	// 收支说明
	details string
	// 余额
	balance float64
	// 金额
	money float64
	// 说明
	note string
	// 标记收支记录
	flag bool
}

// 给该结构体绑定相应的方法

// 将显示明细写成一个方法
func (this *FamilyAccount) ShowDetails() {
	if this.flag {
		fmt.Println("----------当前收支明细记录----------")
		fmt.Println(this.details)
	} else {
		fmt.Println("当前还没有收支记录，来记录一笔吧")
	}
}

// 将登记收入写成一个方法
func (this *FamilyAccount) income() {
	fmt.Println("收入金额：")
	fmt.Scanln(&this.money)
	fmt.Println("收入说明：")
	fmt.Scanln(&this.note)
	this.balance += this.money
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true
}

// 将登记支出写成一个方法
func (this *FamilyAccount) pay() {
	fmt.Println("支出金额：")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("余额中的金额不足")
		// break
	}
	fmt.Println("支出说明：")
	fmt.Scanln(&this.note)
	this.balance -= this.money
	this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true
}

// 将退出系统写成一个方法
func (this *FamilyAccount) exit() {
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
		this.loop = false
	}
}

// 编写一个工厂模式的构造方法，返回一个*FamilyAccount的实例
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key: "",
		loop: true,
		balance: 0.0,
		money: 0.0,
		note: "",
		flag: false,
		details: "收支\t账户金额\t收支金额\t说   明",
	}
}


// 显示主菜单
func (this *FamilyAccount) MainMenu() {  //10.6
	for {
		fmt.Println("\n----------家庭收支记账软件----------")
		fmt.Println("             1 收支明细")
		fmt.Println("             2 登记收入")
		fmt.Println("             3 登记支出")
		fmt.Println("             4 退出软件")
		fmt.Println("请选择(1-4):")
	
		fmt.Scanln(&this.key)
	
		switch this.key {
			case "1":
				this.ShowDetails()
			case "2":
				this.income()
			case "3":
				this.pay()
			case "4":
				this.exit()
			default:
				fmt.Println("输入错误，请输入正确的选项")
		}
	
		if !this.loop {
			break
		}
	}
}