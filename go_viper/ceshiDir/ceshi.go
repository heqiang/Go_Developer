package ceshiDir

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"runtime"
)

func GetConfig() {
	//viper.SetDefault("fileDir", "./")
	//viper.SetConfigFile("config.yaml")    // 指定配置文件路径
	//viper.SetConfigName("config")         // 配置文件名称(无扩展名)
	//viper.SetConfigType("yaml")           // 如果配置文件的名称中没有扩展名，则需要配置此项
	//viper.AddConfigPath("/etc/appname/")  // 查找配置文件所在的路径
	//viper.AddConfigPath("$HOME/.appname") // 多次调用以添加多个搜索路径
	viper.AddConfigPath("./go_viper") // 还可以在工作目录中查找配置
	err := viper.ReadInConfig()       // 查找并读取配置文件
	if err != nil {                   // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	//r := gin.Default()
	//r.GET("/config", func(c *gin.Context) {
	//	c.String(http.StatusOK, viper.GetString("mysql.host"))
	//})
	//r.Run(":8888")
	_, fileStr, _, _ := runtime.Caller(1)
	fmt.Println(fileStr)
	dir, _ := os.Getwd()
	fmt.Println(dir)
}
