package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type KickAndClearPMKcacheRequest struct {
	Id int64 `position:"Query" name:"Id"`
}

func (r KickAndClearPMKcacheRequest) Invoke(client *sdk.Client) (response *KickAndClearPMKcacheResponse, err error) {
	req := struct {
		*requests.RpcRequest
		KickAndClearPMKcacheRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "KickAndClearPMKcache", "", "")

	resp := struct {
		*responses.BaseResponse
		KickAndClearPMKcacheResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.KickAndClearPMKcacheResponse

	err = client.DoAction(&req, &resp)
	return
}

type KickAndClearPMKcacheResponse struct {
	RequestId string
	Success   bool
	Message   string
	ErrorCode int
	ErrorMsg  string
}
