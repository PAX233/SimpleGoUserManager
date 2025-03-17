package customer

import (
	"fmt"
)

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
