package main

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
)

// 日志不在终端打印 将日志写入文件
var sugarLogger *zap.SugaredLogger

func InitZapLogger() {
	writerSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)
	//zap.AddCaller() 调用函数信息写入日志 哪个文件哪一行调用了该日志
	logger := zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
}

// 日志的格式控制
func getEncoder() zapcore.Encoder {
	config := zapcore.EncoderConfig{
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
	return zapcore.NewJSONEncoder(config)
}

// 日志的输出位置
func getLogWriter() zapcore.WriteSyncer {
	//文件追加的形式
	//file ,_:=os.OpenFile("./test.log",os.O_CREATE|os.O_APPEND,666)
	//return  zapcore.AddSync(file)
	// 日志切割 防止所有的日志全部都仿造一个文件里
	lumberjackLogger := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    10,    //M
		MaxBackups: 5,     //备份数量
		MaxAge:     30,    //最大分数天数
		Compress:   false, //是否压缩
	}
	return zapcore.AddSync(lumberjackLogger)
}
func SimpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Error(
			"Error fetching url",
			zap.String("url", url),
			zap.Error(err),
		)
	} else {
		sugarLogger.Info("Success...", zap.String("statusCode", resp.Status),
			zap.String("url", url),
		)
		resp.Body.Close()
	}

}

func main() {
	InitZapLogger()
	//程序结束后将缓冲区的数据都刷到磁盘里
	defer sugarLogger.Sync()
	//go SimpleHttpGet("www.baidu.com")
	//go SimpleHttpGet("https://www.baidu.com")
	//time.Sleep(time.Second*5)
	for i := 0; i < 100000; i++ {
		sugarLogger.Info("")
	}
}
