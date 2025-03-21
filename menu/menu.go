package menu

import (
	"SimpleGoUserManager/customer"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// ShowMenu 显示主菜单
func ShowMenu() {
	fmt.Printf(customer.MainColor+"%s"+customer.ResetColor+"\n", `
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

func ExitAndSave() {
	fmt.Println("\n检测到退出信号，正在保存数据...")
	customer.SaveCustomer()
	fmt.Println("数据已保存，程序退出")
	time.Sleep(1 * time.Second)
	os.Exit(0)
}

// HandleUserChoice 处理用户选择
func HandleUserChoice() {
	// 设置信号处理
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// 启动goroutine监听信号
	go func() {
		<-sigChan
		ExitAndSave()
	}()

	// 设置defer，确保函数返回时保存数据
	defer func() {
		customer.SaveCustomer()
	}()

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
			ExitAndSave()
		default:
			fmt.Println("输入有误")
		}
		fmt.Println("自动返回菜单...")
		time.Sleep(1 * time.Second) // 延时返回菜单
	}
}
