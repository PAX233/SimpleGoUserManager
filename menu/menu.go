package menu

import (
	"basicSystem/customer"
	"fmt"
	"os"
	"time"
)

// ShowMenu 显示主菜单
func ShowMenu() {
	fmt.Printf("\033[1;30;42m%s\033[0m\n", `
════════════════════════════
    客户信息管理系统       
════════════════════════════
    1、添加客户           
    2、查看客户           
    3、更新客户           
    4、删除客户           
    5、保存               
    6、退出               
════════════════════════════
`)
}

// HandleUserChoice 处理用户选择
func HandleUserChoice() {
	for {
		ShowMenu()
		userChoice := -1

		fmt.Print("请输入选择(1-6)：")
		fmt.Scanln(&userChoice)

		// 调用对应的函数
		switch userChoice {
		case 1:
			customer.AddCustomer()
		case 2:
			customer.ShowCustomer()
		case 3:
			customer.UpdateCustomer()
		case 4:
			customer.DeleteCustomer()
		case 5:
			customer.SaveCustomer()
		case 6:
			os.Exit(0)
		default:
			fmt.Println("输入有误")
		}
		time.Sleep(2 * time.Second) // 延时返回菜单
	}
}
