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

		for {
			fmt.Println("是否继续添加？(y/n)")
			var choice string
			fmt.Scan(&choice)
			if choice == "n" {
				fmt.Printf("\033[1;30;42m%s\033[0m\n", `--------添加结束--------`)
				return
			} else if choice == "y" {
				break
			} else {
				fmt.Println("输入有误")
			}

		}
	}
}

// 查看客户
func showCustomer() {
	fmt.Printf("\033[1;30;42m%s\033[0m\n", `--------查看客户--------`)

	// 打印表头
	fmt.Printf("%-4s | %-10s | %-4s | %-6s | %-20s\n", "ID", "姓名", "年龄", "性别", "邮箱")
	fmt.Println("----------------------------------------------------")
	if len(customers) == 0 {
		fmt.Println("暂无客户信息")
		return
	} else
	// 打印客户信息
	{
		for _, customer := range customers {
			fmt.Printf("%-4s | %-10s | %-4s | %-6s | %-20s\n",
				customer["cid"],
				customer["name"],
				customer["age"],
				customer["gender"],
				customer["email"])
		}
	}
	fmt.Printf("\033[1;30;42m%s\033[0m\n", `--------查看结束--------`)
}

// 更新客户
func updateCustomer() {
	fmt.Printf("\033[1;30;42m%s\033[0m\n", `--------更新客户--------`)

	fmt.Println("请输入要更新的客户id：")
	var cid string
	fmt.Scan(&cid)
	index := -1
	for i, v := range customers {
		if v["cid"] == cid {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("没有找到该客户")
		return
	}

	// 更新客户信息
	var name string
	fmt.Println("请输入新的客户姓名（当前：" + customers[index]["name"] + "）：")
	fmt.Scan(&name)

	var age int
	fmt.Println("请输入新的客户年龄（当前：" + customers[index]["age"] + "）：")
	fmt.Scan(&age)

	var gender string
	fmt.Println("请输入新的客户性别（当前：" + customers[index]["gender"] + "）：")
	fmt.Scan(&gender)

	var email string
	fmt.Println("请输入新的客户邮箱（当前：" + customers[index]["email"] + "）：")
	fmt.Scan(&email)

	// 更新客户信息
	customers[index]["name"] = name
	customers[index]["age"] = strconv.Itoa(age)
	customers[index]["gender"] = gender
	customers[index]["email"] = email

	fmt.Println("更新成功")
	fmt.Printf("更新后的用户信息：%v\n", customers[index])
	fmt.Printf("\033[1;30;42m%s\033[0m\n", `--------更新结束--------`)
}

// 删除客户
func deleteCustomer() {
	for {
		fmt.Printf("\033[1;30;42m%s\033[0m\n", `--------删除客户--------`)

		fmt.Println("请输入要删除的客户id：")
		var cid string
		fmt.Scan(&cid)
		index := -1
		for i, v := range customers {
			if v["cid"] == cid {
				index = i
				break
			}
		}

		if index == -1 {
			fmt.Println("没有找到该客户")
		} else {
			customers = append(customers[:index], customers[index+1:]...)
			fmt.Println("删除成功")
		}
		for {
			fmt.Println("是否继续删除？(y/n)")
			var choice string
			fmt.Scan(&choice)
			if choice == "n" {
				fmt.Printf("\033[1;30;42m%s\033[0m\n", `--------删除结束--------`)
				break
			} else if choice == "y" {
				break
			} else {
				fmt.Println("输入有误")
			}
		}
	}
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

		// fmt.Println(userChoice)

		// 调用对应的函数
		switch userChoice {
		case 1:
			addCustomer()
			break
		case 2:
			showCustomer()
			break
		case 3:
			updateCustomer()
			break
		case 4:
			deleteCustomer()
			break
		case 5:
			saveCustomer()
			break
		case 6:
			os.Exit(0)
		default:
			fmt.Println("输入有误")
		}
	}
}
