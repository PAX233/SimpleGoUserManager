package customer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// SaveCustomer 保存客户
func SaveCustomer() {
	fmt.Printf(MainColor+"%s"+ResetColor+"\n", `----------------------保存客户----------------------`)

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

	// 获取当前工作目录
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("获取当前目录失败：%v\n", err)
		return
	}

	// 创建data目录（如果不存在）
	dataDir := filepath.Join(wd, "data")
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		fmt.Printf("创建数据目录失败：%v\n", err)
		return
	}

	// 保存到文件
	filePath := filepath.Join(dataDir, dataFile)
	if err := os.WriteFile(filePath, jsonData, 0644); err != nil {
		fmt.Printf("保存数据失败：%v\n", err)
		return
	}

	fmt.Printf("数据已保存到：%s\n", filePath)
	fmt.Printf(MainColor+"%s"+ResetColor+"\n", `----------------------保存结束----------------------`)
}
