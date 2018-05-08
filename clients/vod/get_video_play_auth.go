package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetVideoPlayAuthRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ReAuthInfo           string `position:"Query" name:"ReAuthInfo"`
	AuthInfoTimeout      int64  `position:"Query" name:"AuthInfoTimeout"`
	VideoId              string `position:"Query" name:"VideoId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *GetVideoPlayAuthRequest) Invoke(client *sdk.Client) (resp *GetVideoPlayAuthResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "GetVideoPlayAuth", "vod", "")
	resp = &GetVideoPlayAuthResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetVideoPlayAuthResponse struct {
	responses.BaseResponse
	RequestId common.String
	PlayAuth  common.String
	VideoMeta GetVideoPlayAuthVideoMeta
}

type GetVideoPlayAuthVideoMeta struct {
	CoverURL common.String
	Duration common.Float
	Status   common.String
	Title    common.String
	VideoId  common.String
}
