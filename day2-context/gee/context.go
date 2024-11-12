package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// H 是一个简化的 map 类型，用于 JSON 数据
type H map[string]interface{}

// Context 封装了 HTTP 请求和响应的上下文
type Context struct {
	Writer http.ResponseWriter // HTTP 响应写入器
	Req    *http.Request       // HTTP 请求
	// 请求信息
	Path   string // 请求路径
	Method string // 请求方法
	// 响应信息
	StatusCode int // 响应状态码
}

// newContext 创建一个新的 Context 实例
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// PostForm 获取 HTTP 请求的POST 表单数据
func (c *Context) PostForm(key string) string {
	// 从 HTTP 请求中获取指定 key 的表单值，这是 Go 的 http.Request 提供的方法，用于获取 POST 表单数据。
	return c.Req.FormValue(key)
}

// Query 获取 URL 查询参数
func (c *Context) Query(key string) string {
	// 从 URL 查询字符串中获取指定 key 的值， Go 的 http.Request 提供的方法，用于获取 URL 查询参数。
	return c.Req.URL.Query().Get(key)
}

// Status 设置 HTTP 响应状态码, 这是个封装方法。
func (c *Context) Status(code int) {
	c.StatusCode = code
	//这是一个低级别的操作，用于指示请求的处理结果（例如，200 表示成功，404 表示未找到，500 表示服务器错误等）。
	//WriteHeader 只能调用一次，并且必须在写入响应体之前调用。
	c.Writer.WriteHeader(code)
}

// SetHeader 设置 HTTP 响应头，这是一个封装方法，简化了设置响应头的过程。它内部调用了 c.Writer.Header().Set(key, value)。
func (c *Context) SetHeader(key string, value string) {
	//设置 HTTP 响应头的键值对。
	//用于设置响应的元数据，例如 Content-Type、Content-Length、Set-Cookie 等。
	//可以在 WriteHeader 调用之前多次设置。
	c.Writer.Header().Set(key, value)
}

// String 返回纯文本响应
func (c *Context) String(code int, format string, value ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	//写入响应体。将字节数组写入响应体中，这是实际返回给客户端的数据部分。
	//写入响应体时，通常需要先设置状态码和响应头。
	c.Writer.Write([]byte(fmt.Sprintf(format, value...)))
}

// JSON 返回 JSON 格式的响应
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	// 创建一个 JSON 编码器，将响应写入 c.Writer
	encoder := json.NewEncoder(c.Writer)
	// 尝试将 obj 编码为 JSON 格式并写入响应
	if err := encoder.Encode(obj); err != nil {
		// 如果编码失败，返回 500 状态码和错误信息
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data 返回二进制数据响应
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// HTML 返回 HTML 格式的响应
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
