package main

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"time"
)

var sugarlogger *zap.SugaredLogger

func getEncoder() zapcore.Encoder {
	//return zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())
	encodingConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder, //时间格式更改
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	return zapcore.NewJSONEncoder(encodingConfig)
}
func getWriteSyncer() zapcore.WriteSyncer {
	//file, _ := os.OpenFile("./go_zap_loglib/selflog.log", os.O_CREATE|os.O_APPEND, 666)
	//return zapcore.AddSync(file)
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./go_zap_loglib/selflog.log", //文件位置
		MaxSize:    10,                            //在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 5,                             //保留旧文件的最大个数
		MaxAge:     30,                            // 保留旧文件的最大天数
		Compress:   false,                         // 是否压缩/归档旧文件
	}
	return zapcore.AddSync(lumberJackLogger)

}
func InitLogger() {
	// zapcore.Core需要三个配置——Encoder，WriteSyncer，LogLevel
	encoder := getEncoder()
	writerSync := getWriteSyncer()
	core := zapcore.NewCore(encoder, writerSync, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	sugarlogger = logger.Sugar()

}
func SimpleHttpGet(url string) {
	//info  := zapcore.Field{}
	resp, err := http.Get(url)
	if err != nil {
		sugarlogger.Error(
			"Error fetching url",
			zap.String("url", url),
			zap.Error(err),
		)
	} else {

		sugarlogger.Info("Success...", "状态码", resp.StatusCode)
		resp.Body.Close()
	}

}
func main() {
	InitLogger()
	defer sugarlogger.Sync()
	//r := gin.Default()
	//r.GET("/sugerlogger", func(c *gin.Context) {
	//	sugarlogger.Info("自定义的log:访问了sugerlogger", zap.S("sss", "vaLue"))
	//})
	//_ = r.Run(":8888")
	go SimpleHttpGet("https://www.baidu.com")
	time.Sleep(time.Second)
}
