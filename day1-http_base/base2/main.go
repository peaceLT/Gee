package main

import (
	"fmt"
	"log"
	"net/http"
)

// 定义一个空的结构体类型 Engine
type Engine struct{}

// 实现 http.Handler 接口的 ServeHTTP 方法
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		// 如果请求路径是根路径 "/", 返回请求的 URL 路径
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		// 如果请求路径是 "/hello"，遍历请求头部信息并写入响应
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		// 对于其他路径，返回 404 错误信息
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func main() {
	// 创建一个 Engine 实例
	engine := new(Engine)
	// 启动 HTTP 服务器，监听 9999 端口，并使用 engine 处理请求
	log.Fatal(http.ListenAndServe(":9999", engine))
}
