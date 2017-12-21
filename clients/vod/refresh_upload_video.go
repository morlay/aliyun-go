package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RefreshUploadVideoRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VideoId              string `position:"Query" name:"VideoId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r RefreshUploadVideoRequest) Invoke(client *sdk.Client) (response *RefreshUploadVideoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RefreshUploadVideoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("vod", "2017-03-21", "RefreshUploadVideo", "", "")

	resp := struct {
		*responses.BaseResponse
		RefreshUploadVideoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RefreshUploadVideoResponse

	err = client.DoAction(&req, &resp)
	return
}

type RefreshUploadVideoResponse struct {
	RequestId     string
	UploadAuth    string
	UploadAddress string
}
