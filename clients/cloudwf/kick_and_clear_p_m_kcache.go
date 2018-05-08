package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type KickAndClearPMKcacheRequest struct {
	requests.RpcRequest
	Id int64 `position:"Query" name:"Id"`
}

func (req *KickAndClearPMKcacheRequest) Invoke(client *sdk.Client) (resp *KickAndClearPMKcacheResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "KickAndClearPMKcache", "", "")
	resp = &KickAndClearPMKcacheResponse{}
	err = client.DoAction(req, resp)
	return
}

type KickAndClearPMKcacheResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Message   common.String
	ErrorCode common.Integer
	ErrorMsg  common.String
}
