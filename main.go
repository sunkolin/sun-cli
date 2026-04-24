package main

import (
	"flag"
	"fmt"
)

func main() {
	// 1. 定义命令行参数
	name := flag.String("name", "陌生人", "你的名字")
	age := flag.Int("age", 0, "你的年龄")
	verbose := flag.Bool("verbose", false, "显示详细信息")

	// 2. 解析参数
	flag.Parse()

	// 3. 业务逻辑
	fmt.Println("========================")
	fmt.Println("   我的通用 CLI 工具")
	fmt.Println("========================")

	if *verbose {
		fmt.Println("[调试模式] 已开启")
	}

	fmt.Printf("你好：%s\n", *name)
	if *age > 0 {
		fmt.Printf("年龄：%d 岁\n", *age)
	}

	fmt.Println("\n✅ CLI 运行成功！")
}