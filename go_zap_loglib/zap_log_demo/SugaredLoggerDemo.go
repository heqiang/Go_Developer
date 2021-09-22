package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var sugarlogger *zap.SugaredLogger
var logger *zap.Logger

func main() {
	InitLogger()
	defer sugarlogger.Sync()
	r := gin.Default()
	r.GET("/sugerlogger", func(c *gin.Context) {
		sugarlogger.Info("sugerlogger:访问了sugerlogger")
		logger.Info("logger:访问了sugerlogger")
	})
	_ = r.Run(":8888")
}
func InitLogger() {
	// zap.NewExample() 用于测试
	//{"level":"info","msg":"sugerlogger:访问了sugerlogger"}
	//zap.NewProduction()
	//{"level":"info","ts":"2021-09-22T20:17:15.979+0800","caller":"zap_log_demo/SugaredLoggerDemo.go:16","msg":"sugerlogger:访问了sugerlogger"}
	//zap.NewDevelopment()
	//2021-09-22T20:17:53.011+0800	INFO	zap_log_demo/SugaredLoggerDemo.go:16	sugerlogger:访问了sugerlogger
	// 该方法
	logger, _ = zap.NewDevelopment()
	sugarlogger = logger.Sugar()
}
