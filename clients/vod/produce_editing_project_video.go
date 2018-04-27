package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ProduceEditingProjectVideoRequest struct {
	requests.RpcRequest
	CoverURL             string `position:"Query" name:"CoverURL"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	MediaMetadata        string `position:"Query" name:"MediaMetadata"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Timeline             string `position:"Query" name:"Timeline"`
	Description          string `position:"Query" name:"Description"`
	ProduceConfig        string `position:"Query" name:"ProduceConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Title                string `position:"Query" name:"Title"`
	ProjectId            string `position:"Query" name:"ProjectId"`
}

func (req *ProduceEditingProjectVideoRequest) Invoke(client *sdk.Client) (resp *ProduceEditingProjectVideoResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "ProduceEditingProjectVideo", "vod", "")
	resp = &ProduceEditingProjectVideoResponse{}
	err = client.DoAction(req, resp)
	return
}

type ProduceEditingProjectVideoResponse struct {
	responses.BaseResponse
	RequestId string
	MediaId   string
	ProjectId string
}
