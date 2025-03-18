package customer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// LoadCustomer 加载客户数据
func LoadCustomer() {
	// 获取当前工作目录
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("获取当前目录失败：%v\n", err)
		return
	}

	// 构建data目录路径
	dataDir := filepath.Join(wd, "data")

	// 尝试读取文件
	filePath := filepath.Join(dataDir, dataFile)
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// 文件不存在，使用默认值
			customers = []Customer{}
			customerId = 0
			return
		}
		fmt.Printf("读取数据失败：%v\n", err)
		return
	}

	// 解析JSON数据
	var data customerData
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Printf("解析数据失败：%v\n", err)
		return
	}

	// 更新全局变量
	customers = data.Customers
	customerId = data.CustomerId
}
