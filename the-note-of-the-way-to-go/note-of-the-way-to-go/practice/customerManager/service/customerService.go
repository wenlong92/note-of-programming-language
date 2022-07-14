package service

import (
	"./model"
)

// 完成对Customer的操作，包括增删改查

type CustomerService struct {
	customers []model.customer
}
