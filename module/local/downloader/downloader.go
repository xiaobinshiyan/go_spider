package downloader

import (
	"net/http"

	"go_spider/module"
	"go_spider/module/stub"
	"gopcp.v2/helper/log"
)

var logger = log.DLogger()

// /下载器的实现模型
type myDownloader struct {
	//组件基础实例  包含 module/base.go
	stub.ModuleInternal
	//http客户端
	httpClient http.Client
}

func New(
	mid module.MID,
	client *http.Client,
	scoreCalculator module.CalculateScore) (module.Downloader, error) {
	moduleBase, err := stub.NewModuleInternal(mid, scoreCalculator)
	if err != nil {
		return nil, err
	}

	if client == nil {
		return nil, genParameterError("nil http client")
	}
	return &myDownloader{
		ModuleInternal: moduleBase,
		httpClient:     *client,
	}, nil
}

func (downloader *myDownloader) Download(req *module.Request) (*module.Response, error) {
	downloader.ModuleInternal.IncrHandlingNumber()
	defer downloader.ModuleInternal.DecrHandlingNumber()
	downloader.ModuleInternal.IncrCalledCount()
	if req == nil {
		return nil, genParameterError("nil request")
	}
	//data.go 获取http请求
	httpReq := req.httpReq()
	downloader.ModuleInternal.IncrAcceptedCount()
	logger.Infof("Do the request (URL: %s, depth: %d)... \n", httpReq.URL, req.Depth())
	httpResp, err := downloader.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	downloader.ModuleInternal.IncrCompletedCount()
	return module.NewResponse(httpResp, req.Depth()), nil
}
