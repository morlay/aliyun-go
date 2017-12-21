package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RefreshUploadVideoRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VideoId              string `position:"Query" name:"VideoId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *RefreshUploadVideoRequest) Invoke(client *sdk.Client) (resp *RefreshUploadVideoResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "RefreshUploadVideo", "", "")
	resp = &RefreshUploadVideoResponse{}
	err = client.DoAction(req, resp)
	return
}

type RefreshUploadVideoResponse struct {
	responses.BaseResponse
	RequestId     string
	UploadAuth    string
	UploadAddress string
}
