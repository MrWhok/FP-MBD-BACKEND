package restclient

import (
	"context"
	"github.com/MrWhok/FP-MBD-BACKEND/client"
	"github.com/MrWhok/FP-MBD-BACKEND/common"
	"github.com/MrWhok/FP-MBD-BACKEND/exception"
	"github.com/MrWhok/FP-MBD-BACKEND/model"
)

func NewHttpBinRestClient() client.HttpBinClient {
	return &HttpBinRestClient{}
}

type HttpBinRestClient struct {
}

func (h HttpBinRestClient) PostMethod(ctx context.Context, requestBody *model.HttpBin, response *map[string]interface{}) {
	var headers []common.HttpHeader
	headers = append(headers, common.HttpHeader{Key: "X-Key", Value: "123456"})

	httpClient := common.ClientComponent[model.HttpBin, map[string]interface{}]{
		HttpMethod:     "POST",
		UrlApi:         "https://httpbin.org/post",
		RequestBody:    requestBody,
		ResponseBody:   response,
		Headers:        headers,
		ConnectTimeout: 30000,
		ActiveTimeout:  30000,
	}
	err := httpClient.Execute(ctx)
	exception.PanicLogging(err)
}
