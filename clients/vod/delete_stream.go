package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteStreamRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	JobIds               string `position:"Query" name:"JobIds"`
	VideoId              string `position:"Query" name:"VideoId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteStreamRequest) Invoke(client *sdk.Client) (response *DeleteStreamResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteStreamRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("vod", "2017-03-21", "DeleteStream", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteStreamResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteStreamResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteStreamResponse struct {
	RequestId string
}
