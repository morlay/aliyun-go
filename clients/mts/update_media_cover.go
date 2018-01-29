package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateMediaCoverRequest struct {
	requests.RpcRequest
	CoverURL             string `position:"Query" name:"CoverURL"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
}

func (req *UpdateMediaCoverRequest) Invoke(client *sdk.Client) (resp *UpdateMediaCoverResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "UpdateMediaCover", "mts", "")
	resp = &UpdateMediaCoverResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateMediaCoverResponse struct {
	responses.BaseResponse
	RequestId string
}
