package customer

import (
	"fmt"
)

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
