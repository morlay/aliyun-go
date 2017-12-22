package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetVideoConfigRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VideoId              string `position:"Query" name:"VideoId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	AuthInfo             string `position:"Query" name:"AuthInfo"`
}

func (req *GetVideoConfigRequest) Invoke(client *sdk.Client) (resp *GetVideoConfigResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "GetVideoConfig", "vod", "")
	resp = &GetVideoConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetVideoConfigResponse struct {
	responses.BaseResponse
	RequestId      string
	DownloadSwitch string
}
