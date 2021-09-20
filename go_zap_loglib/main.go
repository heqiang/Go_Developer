package main

import (
	"go.uber.org/zap"
	"net/http"
	"time"
)

var logger *zap.Logger

func main() {
	InitLogger()
	//程序结束后将缓冲区的数据都刷到磁盘里
	defer logger.Sync()
	go SimpleHttpGet("www.baidu.com")
	go SimpleHttpGet("https://www.baidu.com")
	time.Sleep(time.Second * 5)
}

func InitLogger() {
	logger, _ = zap.NewProduction()
}

func SimpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(
			"Error fetching url",
			zap.String("url", url),
			zap.Error(err),
		)
	} else {
		logger.Info("Success...", zap.String("statusCode", resp.Status),
			zap.String("url", url),
		)
		resp.Body.Close()
	}

}
