package gee

import "net/http"

// HandlerFunc 定义了请求处理函数的类型
type HandlerFunc func(*Context)

// Engine 是框架的核心结构，包含一个路由器
type Engine struct {
	router *router // 路由器，用于管理路由和处理请求
}

// New 创建一个新的 Engine 实例
func New() *Engine {
	return &Engine{router: newRouter()} // 初始化路由器
}

// addRoute 添加路由到路由器
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler) // 调用路由器的 addRoute 方法
}

// GET 定义 GET 请求的路由
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler) // 添加 GET 请求的路由
}

// POST 定义 POST 请求的路由
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler) // 添加 POST 请求的路由
}

// Run 启动 HTTP 服务器
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine) // 监听并服务 HTTP 请求
}

// ServeHTTP 实现 http.Handler 接口，处理所有的 HTTP 请求
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req) // 创建新的请求上下文
	engine.router.handle(c) // 通过路由器处理请求
}
