package customer

import (
	"fmt"
)

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
