package main

import (
	"fmt"
	"gee"      // 导入自定义的 gee 包
	"net/http" // 导入 net/http 包用于处理 HTTP 请求和响应
)

func main() {
	r := gee.New() // 创建一个新的 gee.Engine 实例

	// 注册一个处理函数，处理根路径 "/" 的 GET 请求
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		// 将请求的 URL 路径写入响应
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	// 注册一个处理函数，处理 "/hello" 路径的 GET 请求
	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		// 遍历请求头部信息，并将每个键值对写入响应
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	// 启动 HTTP 服务器，监听 9999 端口
	r.Run(":9999")
}
