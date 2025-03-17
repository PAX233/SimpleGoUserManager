package main

import (
	"basicSystem/customer"
	"basicSystem/menu"
)

func main() {
	// 加载客户数据
	customer.LoadCustomer()
	// 启动菜单
	menu.HandleUserChoice()
}
