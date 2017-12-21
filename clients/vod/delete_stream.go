package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteStreamRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	JobIds               string `position:"Query" name:"JobIds"`
	VideoId              string `position:"Query" name:"VideoId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteStreamRequest) Invoke(client *sdk.Client) (resp *DeleteStreamResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "DeleteStream", "", "")
	resp = &DeleteStreamResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteStreamResponse struct {
	responses.BaseResponse
	RequestId string
}
