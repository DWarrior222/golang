package service

import(
	"../model"
	"fmt"
)

type CustomerService struct {
	customers []model.CustomerInformation
	customerNum int
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "zhangsan", "男", 20, "112", "l@q.com") 
	fmt.Println(customer)
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

func (this *CustomerService) List() []model.CustomerInformation {
	return this.customers
}

func (this *CustomerService) Add(customer model.CustomerInformation) bool {
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

func (this *CustomerService) Del(delNum int) bool {
	index := this.FindById(delNum)
	if index == -1 {
		return false
	}

	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true
}

func (this *CustomerService) FindById(id int) int {
	index := -1
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			index = i
		}
	} 
	return index
}
