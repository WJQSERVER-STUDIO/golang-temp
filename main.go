package main

import (
	"fmt"
	"log"

	"go/config"
	"go/logger"
)

var cfg *config.Config
var logw = logger.Logw

func loadConfig() {
	var err error
	// 初始化配置
	cfg, err = config.LoadConfig("/data/go/config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	fmt.Printf("Loaded config: %v\n", cfg)
}

func setupLogger() {
	// 初始化日志模块
	var err error
	err = logger.Init(cfg.LogFilePath) // 传递日志文件路径
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	logw("Logger initialized")
	logw("Init Completed")
}

func init() {
	loadConfig()
	setupLogger()
}

func main() {
	defer logger.Close() // 确保在退出时关闭日志文件
}
