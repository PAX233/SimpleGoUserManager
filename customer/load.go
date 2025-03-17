package customer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// LoadCustomer 加载客户数据
func LoadCustomer() {
	// 获取可执行文件所在目录
	exePath, err := os.Executable()
	if err != nil {
		fmt.Printf("获取程序路径失败：%v\n", err)
		return
	}
	exeDir := filepath.Dir(exePath)

	// 尝试读取文件
	filePath := filepath.Join(exeDir, dataFile)
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
