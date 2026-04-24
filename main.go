package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/google/uuid"
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

// generateRandomString 生成指定长度的随机字符串
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
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
	showTime := flag.Bool("time", false, "显示当前时间")
	showDate := flag.Bool("date", false, "显示当前日期")
	showDateTime := flag.Bool("datetime", false, "显示当前日期和时间")
	showRandom := flag.Int("random", 0, "生成指定长度的随机字符串")
	showUUID := flag.Bool("uuid", false, "生成UUID")

	// 3. 解析参数
	flag.Parse()

	// 4. 检查是否请求版本信息
	if *showVersion || *showVersionLong {
		fmt.Printf("%s version is %s\n", config.App.Name, config.App.Version)
		return
	}

	// 5. 检查是否请求时间信息
	if *showTime {
		fmt.Println(time.Now().Format("15:04:05"))
		return
	}

	// 6. 检查是否请求日期信息
	if *showDate {
		fmt.Println(time.Now().Format("2006-01-02"))
		return
	}

	// 7. 检查是否请求日期时间信息
	if *showDateTime {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		return
	}

	// 8. 检查是否请求生成随机字符串
	if *showRandom > 0 {
		fmt.Println(generateRandomString(*showRandom))
		return
	}

	// 9. 检查是否请求生成UUID
	if *showUUID {
		fmt.Println(uuid.New().String())
		return
	}

	// 10. 业务逻辑
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