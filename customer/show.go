package customer

import (
	"fmt"
)

// ShowCustomer 查看客户
func ShowCustomer() {
	fmt.Printf(MainColor+"%s"+ResetColor+"\n", `--------查看客户--------`)

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
	fmt.Printf(MainColor+"%s"+ResetColor+"\n", `--------查看结束--------`)
}
