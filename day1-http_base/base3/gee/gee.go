package gee

import (
	"fmt"
	"net/http"
)

// 定义一个处理函数类型，接收http.ResponseWriter和*http.Request作为参数
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine结构体，包含一个路由映射表
type Engine struct {
	router map[string]HandlerFunc
}

// New函数，创建并返回一个新的Engine实例
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// addRoute方法，向路由映射表中添加路由
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern // 生成路由的唯一键
	engine.router[key] = handler  // 将处理函数与路由键关联
}

// GET方法，注册GET请求的路由
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST方法，注册POST请求的路由
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run方法，启动HTTP服务器，监听指定的地址
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// ServeHTTP方法，实现http.Handler接口，处理所有的HTTP请求
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path // 根据请求方法和路径生成路由键
	if handler, ok := engine.router[key]; ok {
		handler(w, req) // 如果找到处理函数，调用它
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL) // 否则返回404错误
	}
}
