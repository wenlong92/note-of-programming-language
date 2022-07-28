package service

import (
	"../model"
)

// 完成对Customer的操作，包括增删改查

type CustomerService struct {
	customers []model.Customer

	// 声明一个字段，表示当前切片含有多少个客户，该字段后面，还可以作为新客户的id+1
	customerNum int
}

// 编写一个方法，可以返回 *CustomerService
func NewCustomerService() *CustomerService {
	// 为了能够看到有客户在切片中，先初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "Bill", "男", 20, "1111111", "Bill@163.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

// 返回客户切片
func (this *CustomerService) List() []model.Customer {
	return this.customers
}

// 添加客户到customers切片
func (this *CustomerService) Add(customer model.Customer) bool {

	// 确定一个分配id的规则，添加顺序为id
	this.customerNum++
	customer.Id = this.customerNum

	this.customers = append(this.customers, customer)
	return true
}

// 根据id查找客户在切片中对应下标，如果没有该客户，返回-1
func (this *CustomerService) FindById(id int) int {
	index := -1
	// 遍历
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			index = i
		}
	}
	return index
}

// 根据id删除客户(从切中删除)
func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		return false
	} else {
		this.customers = append(this.customers[:index], this.customers[index+1:]...)
		return true
	}
}
