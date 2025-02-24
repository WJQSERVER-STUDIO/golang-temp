package main

import (
	"flag"
	"fmt"
	"io"
	"log"

	"go/api"
	"go/config"

	"github.com/WJQSERVER-STUDIO/go-utils/logger"

	"github.com/gin-gonic/gin"
)

var (
	cfg     *config.Config
	router  *gin.Engine
	version string
)

var (
	cfgfile string
)

// 日志模块
var (
	logw       = logger.Logw
	LogDump    = logger.LogDump
	logDebug   = logger.LogDebug
	logInfo    = logger.LogInfo
	logWarning = logger.LogWarning
	logError   = logger.LogError
)

func ReadFlag() {
	cfgfilePtr := flag.String("cfg", "./config/config.toml", "config file path")
	flag.Parse()
	cfgfile = *cfgfilePtr
}

func loadConfig() {
	var err error
	// 初始化配置
	cfg, err = config.LoadConfig(cfgfile)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	fmt.Printf("Loaded config: %v\n", cfg)
}

func setupLogger() {
	// 初始化日志模块
	var err error
	err = logger.Init(cfg.Log.LogFilePath, cfg.Log.MaxLogSize) // 传递日志文件路径
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	logInfo("Logger initialized")
	logInfo("Init Completed")
}

func setupApi(cfg *config.Config, router *gin.Engine, version string) {
	api.InitHandleRouter(cfg, router, version)
}

func init() {
	ReadFlag()
	flag.Parse()
	loadConfig()
	setupLogger()

	//gin.SetMode(gin.ReleaseMode)

	gin.LoggerWithWriter(io.Discard)
	router = gin.New()
	router.Use(gin.Recovery())
	router.UseH2C = false
	setupApi(cfg, router, version)
}

func main() {
	err := router.Run(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port))
	if err != nil {
		logError("Failed to start server: %v\n", err)
	}
	defer logger.Close() // 确保在退出时关闭日志文件
	fmt.Println("Program Exit")
}
