package module

import (
	"net/http"
)

// Counts 代表用于汇集组件内部计数的类型。
type Counts struct {
	// CalledCount 代表调用计数。
	CalledCount uint64
	// AcceptedCount 代表接受计数。
	AcceptedCount uint64
	// CompletedCount 代表成功完成计数。
	CompletedCount uint64
	// HandlingNumber 代表实时处理数。
	HandlingNumber uint64
}

// SummaryStruct 代表组件摘要结构的类型。
type SummaryStruct struct {
	ID        MID         `json:"id"`
	Called    uint64      `json:"called"`
	Accepted  uint64      `json:"accepted"`
	Completed uint64      `json:"completed"`
	Handling  uint64      `json:"handling"`
	Extra     interface{} `json:"extra,omitempty"`
}

type Module interface {
	ID() MID                         //获取当前组件的ID
	Addr() string                    //网络地址的字符串形式
	Score() uint64                   //用于获取当前组件的评分
	ScoreCalculator() CalculateScore //用于获取评分计算器
	CalledCount() uint64             //当前组件被调用的计数
	AcceptedCount() uint64           //接受的调用的次数
	CompletedCount() uint64          //已完成的调用的次数
	HandlingNumber() uint64          //正在处理的调用的数量
	Counts() Counts                  //用于一次性获取所有计数
	Summary() SummaryStruct          //组件摘要
}

//下载器的接口类型
//Downkload 根据事先类型获取内容并返回响应
type Downloader interface {
	Module
	Download(req *Request) (*Response, error)
}

//Analyzer代表下载器的接口类型 并发安全
//唯一标识自己
type Analyzer interface {
	Module
	RespParsers() []ParseResponse
	// 会根据规则分析响应并返回请求和条目。
	// 响应需要分别经过若干响应解析函数的处理，然后合并结果。
	Analyze(resp *Response) ([]Data, []error)
}

// 代表用于解析http响应的函数的类型
// 目的是使用者自定义响应分析规则
type ParseResponse func(httpResp *http.Response, respDepth uint32) ([]Data, []error)
