package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteVideoRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	VideoIds             string `position:"Query" name:"VideoIds"`
}

func (req *DeleteVideoRequest) Invoke(client *sdk.Client) (resp *DeleteVideoResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "DeleteVideo", "", "")
	resp = &DeleteVideoResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteVideoResponse struct {
	responses.BaseResponse
	RequestId string
}
