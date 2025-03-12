# basicSystem

一个简单的用户管理系统示例，使用Go语言实现基本的增删改查(CRUD)功能。

## 功能特性

- 使用map存储用户信息（姓名、年龄）
- 支持添加新用户
- 支持删除指定索引的用户
- 支持修改用户信息
- 支持查询用户列表及单个用户信息

## 运行要求

- Go 1.24 或更高版本

## 快速开始

1. 克隆仓库
```bash
git clone https://github.com/yourusername/basicSystem.git
```

2. 进入项目目录
```bash
cd basicSystem
```

3. 运行程序
```bash
go run main.go
```

## 代码结构

```
├── main.go        # 主程序入口
├── go.mod         # Go模块配置文件
└── README.md      # 项目说明文档
```

## 示例输出

程序运行后会演示：
- 添加3个初始用户
- 新增用户
- 删除中间用户
- 修改用户信息
- 最终用户列表输出

## 许可证
MIT License