package customer

import (
	"fmt"
)

// UpdateCustomer 更新客户
func UpdateCustomer() {
	fmt.Printf(MainColor+"%s"+ResetColor+"\n", `----------------------更新客户----------------------`)
	for {
		fmt.Println("请输入要更新的客户id（输入0取消）：")
		var cid int
		fmt.Scan(&cid)
		if cid == 0 {
			fmt.Printf(MainColor+"%s"+ResetColor+"\n", `----------------------更新取消----------------------`)
			return
		}

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
		fmt.Printf("请输入新的客户姓名（当前：%s）（输入q取消）：\n", customers[index].Name)
		fmt.Scan(&name)
		if name == "q" {
			fmt.Printf(MainColor+"%s"+ResetColor+"\n", `--------更新取消--------`)
			return
		}

		var ageStr string
		fmt.Printf("请输入新的客户年龄（当前：%d）（输入q取消）：\n", customers[index].Age)
		fmt.Scan(&ageStr)
		if ageStr == "q" {
			fmt.Printf(MainColor+"%s"+ResetColor+"\n", `--------更新取消--------`)
			return
		}
		age := 0
		fmt.Sscanf(ageStr, "%d", &age)

		var gender string
		fmt.Printf("请输入新的客户性别（当前：%s）（输入q取消）：\n", customers[index].Gender)
		fmt.Scan(&gender)
		if gender == "q" {
			fmt.Printf(MainColor+"%s"+ResetColor+"\n", `--------更新取消--------`)
			return
		}

		var email string
		fmt.Printf("请输入新的客户邮箱（当前：%s）（输入q取消）：\n", customers[index].Email)
		fmt.Scan(&email)
		if email == "q" {
			fmt.Printf(MainColor+"%s"+ResetColor+"\n", `--------更新取消--------`)
			return
		}

		// 更新客户信息
		customers[index].Name = name
		customers[index].Age = age
		customers[index].Gender = gender
		customers[index].Email = email

		fmt.Println("更新成功")
		fmt.Printf("更新后的用户信息：%v\n", customers[index])
		fmt.Printf(MainColor+"%s"+ResetColor+"\n", `---------------------------------------------------`)

		for {
			fmt.Println("是否继续更新？(y/n)")
			var choice string
			fmt.Scan(&choice)
			if choice == "n" {
				fmt.Printf(MainColor+"%s"+ResetColor+"\n", `----------------------更新结束----------------------`)
				return
			} else if choice == "y" {
				break
			} else {
				fmt.Println("输入有误")
			}
		}
	}
}
