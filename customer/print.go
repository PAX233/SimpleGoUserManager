package customer

import "fmt"

// PrintCustomerInfo 打印客户信息
func PrintCustomerInfo(customer Customer) {
	fmt.Printf("用户信息：\nCid：%d\n姓名：%s\n年龄：%d\n性别：%s\n邮箱：%s\n",
		customer.Cid,
		customer.Name,
		customer.Age,
		customer.Gender,
		customer.Email)
}
