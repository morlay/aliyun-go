package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateUploadVideoRequest struct {
	requests.RpcRequest
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
	CateId               int64  `position:"Query" name:"CateId"`
}

func (req *CreateUploadVideoRequest) Invoke(client *sdk.Client) (resp *CreateUploadVideoResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "CreateUploadVideo", "vod", "")
	resp = &CreateUploadVideoResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateUploadVideoResponse struct {
	responses.BaseResponse
	RequestId     common.String
	VideoId       common.String
	UploadAddress common.String
	UploadAuth    common.String
}
