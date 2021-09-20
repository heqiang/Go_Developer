package main

import (
	"Go_Developer/go_zap_loglib/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
}
