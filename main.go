package main

import (
	"fmt"
	"os"
	"strconv"
)

var customers []map[string]string
var customerId int = 0

type Customer struct {
	cid    int
	name   string
	age    int
	gender string
	email  string
}

func showMenu() {
	fmt.Printf("\033[1;30;42m%s\033[0m\n", `
--------客户信息管理系统--------
	1、添加客户
	2、查看客户
	3、更新客户
	4、删除客户
	5、保存
	6、退出
----------------------------
	`)

}

// 添加客户
func addCustomer() {
	for {
		fmt.Printf("\033[1;30;42m%s\033[0m\n", `--------添加客户--------`)

		var name string
		fmt.Println("请输入客户姓名：")
		fmt.Scan(&name)

		var age int
		fmt.Println("请输入客户年龄：")
		fmt.Scan(&age)

		var gender string
		fmt.Println("请输入客户性别：")
		fmt.Scan(&gender)

		var email string
		fmt.Println("请输入客户邮箱：")
		fmt.Scan(&email)

		customerId++
		cid := customerId
		newCustomer := map[string]string{
			"cid":    strconv.Itoa(cid),
			"name":   name,
			"age":    strconv.Itoa(age),
			"gender": gender,
			"email":  email,
		}

		customers = append(customers, newCustomer)

		fmt.Printf("用户信息：%v\n", newCustomer)

		fmt.Println("是否继续添加？(y/n)")
		var choice string
		fmt.Scan(&choice)
		if choice == "n" {
			fmt.Printf("\033[1;30;42m%s\033[0m\n",`--------添加结束--------`)
			break
		}
	}
}

// 查看客户
func showCustomer() {
	fmt.Println("查看客户:")
}

// 更新客户
func updateCustomer() {
	fmt.Println("更新客户:")
}

// 删除客户
func deleteCustomer() {
	fmt.Println("删除客户:")
}

// 保存客户
func saveCustomer() {
	fmt.Println("保存客户:")
}

func main() {
	for {
		showMenu()
		userChoice := -1

		fmt.Print("请输入选择(1-6)：")
		fmt.Scanln(&userChoice)

		fmt.Println(userChoice)

		// 调用对应的函数
		switch userChoice {
		case 1:
			addCustomer()
		case 2:
			showCustomer()
		case 3:
			updateCustomer()
		case 4:
			deleteCustomer()
		case 5:
			saveCustomer()
		case 6:
			os.Exit(0)
		default:
			fmt.Println("输入有误")
		}
	}
}
