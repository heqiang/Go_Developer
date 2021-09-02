package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//设置访问路由
	http.HandleFunc("/", printfFunc)
	//设置访问的ip和端口
	err := http.ListenAndServe("127.0.0.1:8882", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

func printfFunc(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "POST" {
		upload, _, _ := r.FormFile("file")
		ss := make([]byte, 6666)
		upload.Read(ss)
		fmt.Println("********************************")
		fmt.Println(string(ss))
		fmt.Println(ss)
		fmt.Println("********************************")
		//errorHandle(err, w)
		//
		//fmt.Println(handle.Filename)
		//os.Mkdir("./uploaded/", 0777)
		//saveFile, err := os.OpenFile("./uploaded/" + handle.Filename, os.O_WRONLY|os.O_CREATE, 0666);
		//errorHandle(err, w)
		//io.Copy(saveFile, upload)
	}
	if r.Method == "GET" {
		w.Write([]byte("这是get请求"))
	}
}

// 统一错误输出接口
func errorHandle(err error, w http.ResponseWriter) {
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}
