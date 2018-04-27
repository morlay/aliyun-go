package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type PushObjectCacheRequest struct {
	requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ObjectPath           string `position:"Query" name:"ObjectPath"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

func (req *PushObjectCacheRequest) Invoke(client *sdk.Client) (resp *PushObjectCacheResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "PushObjectCache", "vod", "")
	resp = &PushObjectCacheResponse{}
	err = client.DoAction(req, resp)
	return
}

type PushObjectCacheResponse struct {
	responses.BaseResponse
	RequestId  string
	PushTaskId string
}
