package gee

import (
	"log"
	"net/http"
)

// router 结构体用于存储路由信息
type router struct {
	handlers map[string]HandlerFunc // 存储路由和处理函数的映射
}

// newRouter 创建一个新的 router 实例
func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

// addRoute 添加路由到路由表中
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern) // 打印路由信息
	key := method + "-" + pattern                 // 生成路由的唯一键
	r.handlers[key] = handler                     // 存储处理函数
}

// handle 处理请求
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path // 根据请求方法和路径生成键
	if handler, ok := r.handlers[key]; ok {
		handler(c) // 如果找到处理函数，调用它
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path) // 否则返回 404 错误
	}
}
