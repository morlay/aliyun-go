package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateMediaPublishStateRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Publish              string `position:"Query" name:"Publish"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
}

func (req *UpdateMediaPublishStateRequest) Invoke(client *sdk.Client) (resp *UpdateMediaPublishStateResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "UpdateMediaPublishState", "", "")
	resp = &UpdateMediaPublishStateResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateMediaPublishStateResponse struct {
	responses.BaseResponse
	RequestId string
}
