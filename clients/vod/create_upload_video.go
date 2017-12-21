package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateUploadVideoRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	TranscodeMode        string `position:"Query" name:"TranscodeMode"`
	IP                   string `position:"Query" name:"IP"`
	Description          string `position:"Query" name:"Description"`
	FileSize             int64  `position:"Query" name:"FileSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Title                string `position:"Query" name:"Title"`
	Tags                 string `position:"Query" name:"Tags"`
	CoverURL             string `position:"Query" name:"CoverURL"`
	UserData             string `position:"Query" name:"UserData"`
	FileName             string `position:"Query" name:"FileName"`
	TemplateGroupId      string `position:"Query" name:"TemplateGroupId"`
	CateId               int    `position:"Query" name:"CateId"`
}

func (r CreateUploadVideoRequest) Invoke(client *sdk.Client) (response *CreateUploadVideoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateUploadVideoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("vod", "2017-03-21", "CreateUploadVideo", "", "")

	resp := struct {
		*responses.BaseResponse
		CreateUploadVideoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateUploadVideoResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateUploadVideoResponse struct {
	RequestId     string
	VideoId       string
	UploadAddress string
	UploadAuth    string
}
