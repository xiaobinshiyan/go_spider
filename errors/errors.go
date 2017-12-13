package errors

import (
	"bytes"
	"fmt"
	"strings"
)

//别名 错误类型
type ErrorType string

const (
	ERROR_TYPE_DOWNLOADER ErrorType = "downloder error"
	ERROR_TYPE_ANALYZER   ErrorType = "analyzer error"
	ERROR_TYPE_PIPELINE   ErrorType = "pipeline error"
	ERROR_TYPE_SCHEDULER  ErrorType = "scheduler error"
)

type CrawlerError interface {
	Type() ErrorType //获得错误类型
	Error() string   //用于获得错误提示信息
}

//错误的实现类型
type myCrawlerError struct {
	errType    ErrorType //错误类型
	errMsg     string    //错误信息
	fullErrMsg string    //完整错误信息
}

//创建并初始化
//*myCrawlerError 是CrawlerError 的一个实现类型
func NewCrawlerError(errType ErrorType, errMsg string) CrawlerError {
	return &myCrawlerError{
		errType: errType,
		errMsg:  strings.TrimSpace(errMsg),
	}
}

func (ce *myCrawlerError) Type() ErrorType {
	return ce.errType
}

//*myCrawerError类型的方法集合中包含以上接口类型的Type方法和Error方法
func (ce *myCrawlerError) Error() string {
	if ce.fullErrMsg == "" {
		ce.genFullErrMsg()
	}
	return ce.fullErrMsg
}

//生成错误提示信息 并给响应的字段赋值
func (ce *myCrawlerError) genFullErrMsg() {
	var buffer bytes.Buffer
	buffer.WriteString("Crawler error:")
	//为“”直接返回 频繁拼接字符串引发性能影响
	if ce.errType != "" {
		buffer.WriteString(string(ce.errType))
		buffer.WriteString(": ")
	}
	buffer.WriteString(ce.errMsg)
	ce.fullErrMsg = fmt.Sprintf("%s", buffer.String())
	return
}

//代表非法的参数的错误类型
type IllegalParameterError struct {
	msg string
}

func NewIllegalParameterError(errMsg string) IllegalParameterError {
	return IllegalParameterError{
		msg: fmt.Sprintf("illegal parameter: %s", strings.TrimSpace(errMsg)),
	}
}

func (ipe IllegalParameterError) Error() string {
	return ipe.msg
}
