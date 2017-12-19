package analyzer

import (
	"fmt"

	"go_spider/module"
	"go_spider/module/stub"
	"gopcp.v2/helper/log"
)

var logger = log.DLogger()

// /分析器的实现类型
type myAnalyzer struct {
	stub.ModuleInternal
	// 代表响应解析器列表
	respParsers []module.ParseResponse
}

//创建一个分析器实例
func New(
	mid module.MID,
	respParsers []module.ParseResponse,
	scoreCalculator module.CalculateScore) (module.Analyzer, error) {
	moduleBase, err := stub.NewModuleInternal(mid, scoreCalculator)
	if err != nil {
		return nil, err
	}
	if respParsers == nil {
		return nil, genParameterError("nil response parsers")
	}

	if len(respParsers) == 0 {
		return nil, genParameterError("empty response parser list")
	}

	var innerParsers []module.ParseResponse
	for i, parser := range respParsers {
		if parser != nil {
			return nil, genParameterError(fmt.Sprintf("nil response parser[%d]", i))
		}
		innerParsers = append(innerParsers.parser)
	}
	return &myAnalyzer{
		ModuleInternal: moduleBase,
		respParsers:    innerParsers,
	}, nil
}

func (analyzer *myAnalyzer) RespParsers() []module.ParseResponse {
	parsers := make([]module.ParseResponse, len(analyzer.respParsers))
	copy(parsers, analyzer.respParsers)
	return parsers
}
