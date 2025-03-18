package main

import (
	"SimpleGoUserManager/customer"
	"SimpleGoUserManager/menu"
)

func main() {
	// 加载客户数据
	customer.LoadCustomer()
	// 启动菜单
	menu.HandleUserChoice()
}
