package main

import (
	"gee"      // 导入自定义的 gee 包
	"net/http" // 导入 net/http 包，用于 HTTP 状态码
)

func main() {
	r := gee.New() // 创建一个新的 gee.Engine 实例

	// 定义一个 GET 请求的路由，路径为 "/"
	r.GET("/", func(c *gee.Context) {
		// 返回一个 HTML 响应，状态码为 200
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	// 定义一个 GET 请求的路由，路径为 "/hello"
	r.GET("/hello", func(c *gee.Context) {
		// 期望请求格式为 /hello?name=geektutu
		// 返回一个格式化的字符串响应，包含查询参数 "name" 和请求路径
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	// 定义一个 POST 请求的路由，路径为 "/login"
	r.POST("/login", func(c *gee.Context) {
		// 返回一个 JSON 响应，包含从 POST 表单中获取的 "username" 和 "password"
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	// 启动 HTTP 服务器，监听端口 9999
	r.Run(":9999")
}
