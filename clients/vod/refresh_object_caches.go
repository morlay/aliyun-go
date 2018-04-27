package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RefreshObjectCachesRequest struct {
	requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ObjectPath           string `position:"Query" name:"ObjectPath"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	ObjectType           string `position:"Query" name:"ObjectType"`
}

func (req *RefreshObjectCachesRequest) Invoke(client *sdk.Client) (resp *RefreshObjectCachesResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "RefreshObjectCaches", "vod", "")
	resp = &RefreshObjectCachesResponse{}
	err = client.DoAction(req, resp)
	return
}

type RefreshObjectCachesResponse struct {
	responses.BaseResponse
	RequestId     string
	RefreshTaskId string
}
