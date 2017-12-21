package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateVideoInfoRequest struct {
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

func (r UpdateVideoInfoRequest) Invoke(client *sdk.Client) (response *UpdateVideoInfoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UpdateVideoInfoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("vod", "2017-03-21", "UpdateVideoInfo", "", "")

	resp := struct {
		*responses.BaseResponse
		UpdateVideoInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpdateVideoInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpdateVideoInfoResponse struct {
	RequestId string
}
