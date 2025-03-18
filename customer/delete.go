package customer

import (
	"fmt"
)

// DeleteCustomer 删除客户
func DeleteCustomer() {
	for {
		fmt.Printf(MainColor+"%s"+ResetColor+"\n", `----------------------删除客户----------------------`)

		fmt.Println("请输入要删除的客户id（输入0取消）：")
		var cid int
		fmt.Scan(&cid)
		if cid == 0 {
			fmt.Printf(MainColor+"%s"+ResetColor+"\n", `----------------------删除取消----------------------`)
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
		} else {
			fmt.Printf(MainColor+"%s"+ResetColor+"\n", `----------------------客户信息----------------------`)
			PrintCustomerInfo(customers[index])
			fmt.Printf(MainColor+"%s"+ResetColor+"\n", `---------------------------------------------------`)

			for {
				fmt.Println("确认删除该客户？(y/n)")
				var confirm string
				fmt.Scan(&confirm)
				if confirm == "y" {
					customers = append(customers[:index], customers[index+1:]...)
					fmt.Println("删除成功")
					break
				} else if confirm == "n" {
					fmt.Println("取消删除")
					break
				} else {
					fmt.Println("输入有误，请重新输入")
				}
			}
		}
		for {
			fmt.Println("是否继续删除？(y/n)")
			var choice string
			fmt.Scan(&choice)
			if choice == "n" {
				fmt.Printf(MainColor+"%s"+ResetColor+"\n", `----------------------删除结束----------------------`)
				return
			} else if choice == "y" {
				break
			} else {
				fmt.Println("输入有误")
			}
		}
	}
}
