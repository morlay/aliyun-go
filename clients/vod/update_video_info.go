package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateVideoInfoRequest struct {
	requests.RpcRequest
	CoverURL             string `position:"Query" name:"CoverURL"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CateId               int    `position:"Query" name:"CateId"`
	Description          string `position:"Query" name:"Description"`
	VideoId              string `position:"Query" name:"VideoId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Title                string `position:"Query" name:"Title"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (req *UpdateVideoInfoRequest) Invoke(client *sdk.Client) (resp *UpdateVideoInfoResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "UpdateVideoInfo", "", "")
	resp = &UpdateVideoInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateVideoInfoResponse struct {
	responses.BaseResponse
	RequestId string
}
