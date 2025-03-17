package customer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// SaveCustomer 保存客户
func SaveCustomer() {
	fmt.Printf("\033[1;30;42m%s\033[0m\n", `--------保存客户--------`)

	// 准备要保存的数据
	data := customerData{
		Customers:  customers,
		CustomerId: customerId,
	}

	// 将数据转换为JSON格式
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Printf("转换数据失败：%v\n", err)
		return
	}

	// 获取可执行文件所在目录
	exePath, err := os.Executable()
	if err != nil {
		fmt.Printf("获取程序路径失败：%v\n", err)
		return
	}
	exeDir := filepath.Dir(exePath)

	// 保存到文件
	filePath := filepath.Join(exeDir, dataFile)
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		fmt.Printf("保存数据失败：%v\n", err)
		return
	}

	fmt.Println("保存成功！")
	fmt.Printf("\033[1;30;42m%s\033[0m\n", `--------保存结束--------`)
}
