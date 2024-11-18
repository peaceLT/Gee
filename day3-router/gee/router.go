package gee

import (
	"net/http"
	"strings"
)

// router 结构体用于存储路由信息
type router struct {
	roots    map[string]*node       // 每种请求方法的根节点
	handlers map[string]HandlerFunc // 路由对应的处理函数
}

// roots key eg, roots['GET'] roots['POST']
// handlers key eg, handlers['GET-/p/:lang/doc'], handlers['POST-/p/book']
// newRouter 创建并返回一个新的路由器实例
func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

// parsePattern 解析路由模式，将其分解为各个部分
// 例如，将 "/p/:lang/doc" 解析为 ["p", ":lang", "doc"]
// Only one * is allowed
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			// 如果遇到通配符 *，则停止解析
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

// addRoute 添加路由到路由器中
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern) // 解析路由模式

	key := method + "-" + pattern // 生成路由键
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{} // 如果没有根节点，则创建一个新的
	}
	r.roots[method].insert(pattern, parts, 0) // 插入路由到前缀树中
	r.handlers[key] = handler                 // 存储路由对应的处理函数
}

// getRoute 查找与请求路径匹配的路由节点和参数
func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path) // 解析请求路径
	params := make(map[string]string) // 存储路径参数
	root, ok := r.roots[method]

	if !ok {
		return nil, nil // 如果没有对应方法的根节点，返回 nil
	}

	n := root.search(searchParts, 0) // 在前缀树中查找匹配的节点
	if n != nil {
		parts := parsePattern(n.pattern) // 解析匹配节点的完整路径
		for index, part := range parts {
			if part[0] == ':' {
				// 动态参数，存储到 params 中
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				// 通配符参数，存储路径的剩余部分
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params // 返回匹配的节点和参数
	}
	return nil, nil // 如果没有匹配的节点，返回 nil
}

// handle 处理请求
func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path) // 查找匹配的路由节点和参数
	if n != nil {
		c.Params = params // 设置请求上下文的参数
		key := c.Method + "-" + n.pattern
		r.handlers[key](c) // 调用匹配的处理函数
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path) // 返回 404 错误
	}
}
