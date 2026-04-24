package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config 定义配置文件结构
type Config struct {
	App struct {
		Name        string `yaml:"name"`
		Description string `yaml:"description"`
		Version     string `yaml:"version"`
	} `yaml:"app"`
}

// loadConfig 从配置文件加载配置
func loadConfig() (*Config, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %v", err)
	}

	return &config, nil
}

func main() {
	// 1. 加载配置文件
	config, err := loadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误: %v\n", err)
		os.Exit(1)
	}

	// 2. 定义命令行参数
	name := flag.String("name", "陌生人", "你的名字")
	age := flag.Int("age", 0, "你的年龄")
	verbose := flag.Bool("verbose", false, "显示详细信息")
	showVersion := flag.Bool("v", false, "显示版本信息")
	showVersionLong := flag.Bool("version", false, "显示版本信息")

	// 3. 解析参数
	flag.Parse()

	// 4. 检查是否请求版本信息
	if *showVersion || *showVersionLong {
		fmt.Printf("%s version is %s\n", config.App.Name, config.App.Version)
		return
	}

	// 5. 业务逻辑
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