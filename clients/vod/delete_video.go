package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteVideoRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	VideoIds             string `position:"Query" name:"VideoIds"`
}

func (r DeleteVideoRequest) Invoke(client *sdk.Client) (response *DeleteVideoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteVideoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("vod", "2017-03-21", "DeleteVideo", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteVideoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteVideoResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteVideoResponse struct {
	RequestId string
}
