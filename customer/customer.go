package customer

import "fmt"

var customers []Customer
var customerId int = 0

type Customer struct {
	Cid    int
	Name   string
	Age    int
	Gender string
	Email  string
}

// AddCustomer 添加客户
func AddCustomer() {
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
		newCustomer := Customer{
			Cid:    customerId,
			Name:   name,
			Age:    age,
			Gender: gender,
			Email:  email,
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

// ShowCustomer 查看客户
func ShowCustomer() {
	fmt.Printf("\033[1;30;42m%s\033[0m\n", `--------查看客户--------`)

	// 打印表头
	fmt.Printf("%-4s | %-10s | %-4s | %-6s | %-20s\n", "ID", "姓名", "年龄", "性别", "邮箱")
	fmt.Println("----------------------------------------------------")
	if len(customers) == 0 {
		fmt.Println("暂无客户信息")
		return
	} else {
		// 打印客户信息
		for _, customer := range customers {
			fmt.Printf("%-4d | %-10s | %-4d | %-6s | %-20s\n",
				customer.Cid,
				customer.Name,
				customer.Age,
				customer.Gender,
				customer.Email)
		}
	}
	fmt.Printf("\033[1;30;42m%s\033[0m\n", `--------查看结束--------`)
}

// UpdateCustomer 更新客户
func UpdateCustomer() {
	fmt.Printf("\033[1;30;42m%s\033[0m\n", `--------更新客户--------`)

	fmt.Println("请输入要更新的客户id：")
	var cid int
	fmt.Scan(&cid)
	index := -1
	for i, v := range customers {
		if v.Cid == cid {
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
	fmt.Printf("请输入新的客户姓名（当前：%s）：\n", customers[index].Name)
	fmt.Scan(&name)

	var age int
	fmt.Printf("请输入新的客户年龄（当前：%d）：\n", customers[index].Age)
	fmt.Scan(&age)

	var gender string
	fmt.Printf("请输入新的客户性别（当前：%s）：\n", customers[index].Gender)
	fmt.Scan(&gender)

	var email string
	fmt.Printf("请输入新的客户邮箱（当前：%s）：\n", customers[index].Email)
	fmt.Scan(&email)

	// 更新客户信息
	customers[index].Name = name
	customers[index].Age = age
	customers[index].Gender = gender
	customers[index].Email = email

	fmt.Println("更新成功")
	fmt.Printf("更新后的用户信息：%v\n", customers[index])
	fmt.Printf("\033[1;30;42m%s\033[0m\n", `--------更新结束--------`)
}

// DeleteCustomer 删除客户
func DeleteCustomer() {
	for {
		fmt.Printf("\033[1;30;42m%s\033[0m\n", `--------删除客户--------`)

		fmt.Println("请输入要删除的客户id：")
		var cid int
		fmt.Scan(&cid)
		index := -1
		for i, v := range customers {
			if v.Cid == cid {
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

// SaveCustomer 保存客户
func SaveCustomer() {
	fmt.Println("保存客户:")
}
