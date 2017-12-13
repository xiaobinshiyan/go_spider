package module

import (
	"net/http"
)

// valid 用于判断数据是否有效
// **作用更多的是作为某一类的类型的标签而不是定义这类类型的行为
type Data interface {
	Valid() bool
}

//Rquest代表数据请求的类型
//@httpReq 代表http请求  是一个结构体类型 它的零值不是nil
//@depth 请求的深度
//**在一个Module.request类型的值被创建和初始化后 当前代码包之外的代码不能对它的任何字段的值进行更改
type Request struct {
	//1:小写包级私有
	//2:编写一个创建和初始化的函数
	//3:编写获取字段值的方法
	httpReq *http.Request
	depth   uint32
}

//创建新的请求
func NewRequest(httpReq *http.Request, depth uint32) *Request {
	return &Request{httpReq: httpReq, depth: depth}
}

//获取http请求
func (req *Request) HTTPReq() *http.Request {
	return req.httpReq
}

//获取请求深度
func (req *Request) Depth() uint32 {
	return req.depth
}

//判断请求是否有效
func (req *Request) Valid() bool {
	return req.httpReq != nil && req.httpReq.URL != nil
}

//resbonse代表响应数据的类型
type Response struct {
	httpResp *http.Response
	depth    uint32
}

func NewResponse(httpResp *http.Response, depth uint32) *Response {
	return &Response{httpResp: httpResp, depth: depth}
}

func (resp *Response) HTTPResp() *http.Response {
	return resp.httpResp
}

func (resp *Response) Depth() uint32 {
	return resp.depth
}

//判断响应是否有效
func (resp *Response) Valid() bool {
	return resp.httpResp != nil && resp.httpResp.Body != nil
}

//代表条目的类型
//应该足够灵活，可以容纳所有从响应内容中筛选出的数据,字典类型的别名类型
//**处理条目器是由使用者提供
type Item map[string]interface{}

func (item *Item) Valid() bool {
	return item != nil
}
