# basicSystem

一个简单的用户管理系统示例，使用Go语言实现基本的增删改查(CRUD)功能。

## 功能特性

- 使用结构体存储用户信息（姓名、年龄、性别、邮箱）
- 使用JSON文件实现数据持久化存储
- 支持添加新用户
- 支持查看用户列表及单个用户信息
- 支持更新用户信息
- 支持删除指定ID的用户

## 运行要求

- Go 1.24 或更高版本

## 快速开始

1. 克隆仓库
```bash
git clone https://github.com/PAX233/basicSystem.git
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
├── customer/      # 客户相关功能实现
│   ├── add.go
│   ├── delete.go
│   ├── load.go
│   ├── save.go
│   ├── show.go
│   ├── types.go
│   └── update.go
├── menu/          # 菜单相关功能实现
│   └── menu.go
├── go.mod         # Go模块配置文件
└── README.md      # 项目说明文档
```

## 示例输出

程序运行后会演示：
- 从持久化文件加载用户数据
- 添加新用户
- 查看用户列表
- 更新用户信息
- 删除用户
- 自动保存数据到文件

## 许可证
MIT License

### 解释
- **功能特性**：更新了功能特性以反映使用结构体存储用户信息。
- **代码结构**：添加了 `customer/` 和 `menu/` 目录以反映项目的代码组织。
- **示例输出**：更新了示例输出以更准确地描述程序的功能。
