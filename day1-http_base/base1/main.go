package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)           // 注册根路径的处理函数
	http.HandleFunc("/hello", helloHandler)      // 注册/hello路径的处理函数
	log.Fatal(http.ListenAndServe(":9999", nil)) // 启动HTTP服务器，监听9999端口
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path) // 输出请求的URL路径
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header { // 遍历请求头
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v) // 输出每个请求头的键值对
	}
}
