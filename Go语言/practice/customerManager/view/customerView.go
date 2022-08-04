package main

import (
	"fmt"
	"../service"
	"../model"
)

type customerView struct {
	key             string  // 接收用户输入
	loop            bool    // 表示是否循环显示主菜单
	customerService *service.CustomerService
}

// 显示所有客户信息
func (this *customerView) list() {
	// 获取当前所有客户信息(在切片中)
	customers := this.customerService.List()
	// 显示
	fmt.Println("------------客户列表------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("----------客户列表完成----------\n\n")
}

// 得到用户输入，信息构建新的客户，并添加
func (this *customerView) add() {
	fmt.Println("------------添加客户------------")

	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)

	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)

	fmt.Println("年龄：")
	var age int
	fmt.Scanln(&age)

	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)

	fmt.Println("邮箱：")
	email := ""
	fmt.Scanln(&email)

    // 构建一个新的Customer实例，ID号没有让用户输入，id唯一，需要系统分配
	customer := model.NewCustomer2(name, gender, age, phone, email)
	// 调用
	if this.customerService.Add(customer) {
		fmt.Println("----------添加完成----------\n")
	} else {
		fmt.Println("----------添加失败----------\n")
	}
}

// 得到用户输入的id，删除该id对应的客户
func (this *customerView) delete() {
	fmt.Println("------------删除客户------------")
	fmt.Println("请选择待删除客户编号(-1退出)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return // 放弃删除操作
	}
	fmt.Println("确认是否删除(Y/N)：")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "Y" || choice == "y" {
		if this.customerService.Delete(id) {
			fmt.Println("----------删除完成----------\n")
		} else {
			fmt.Println("----------删除失败，输入的id不存在----------\n")
		}
	}
}

// 退出软件
func (this *customerView) exit() {
	fmt.Println("确认是否推出(Y/N):")
	for {
		fmt.Scanln(&this.key)
		if this.key == "y" || this.key == "Y" || this.key == "n" || this.key == "N" {
			break
		} else {
			fmt.Println("你的输入有误，确认是否退出(Y/N):")
		}
	}
	if this.key == "y" || this.key == "Y" {
		this.loop = false
	}
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
				this.add()
			case "2":
				fmt.Println("2")
			case "3":
				this.delete()
			case "4":
				 this.list()
			case "5":
				this.exit()
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
	customerView.customerService = service.NewCustomerService()
	customerView.mainMenu()

}