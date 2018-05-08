package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetVideoPlayInfoRequest struct {
	requests.RpcRequest
	SignVersion          string `position:"Query" name:"SignVersion"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ClientVersion        string `position:"Query" name:"ClientVersion"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Channel              string `position:"Query" name:"Channel"`
	PlaySign             string `position:"Query" name:"PlaySign"`
	VideoId              string `position:"Query" name:"VideoId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ClientTS             int64  `position:"Query" name:"ClientTS"`
}

func (req *GetVideoPlayInfoRequest) Invoke(client *sdk.Client) (resp *GetVideoPlayInfoResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "GetVideoPlayInfo", "vod", "")
	resp = &GetVideoPlayInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetVideoPlayInfoResponse struct {
	responses.BaseResponse
	RequestId common.String
	PlayInfo  GetVideoPlayInfoPlayInfo
	VideoInfo GetVideoPlayInfoVideoInfo
}

type GetVideoPlayInfoPlayInfo struct {
	AccessKeyId     common.String
	AccessKeySecret common.String
	AuthInfo        common.String
	SecurityToken   common.String
	Region          common.String
	PlayDomain      common.String
}

type GetVideoPlayInfoVideoInfo struct {
	CoverURL   common.String
	CustomerId common.Long
	Duration   common.Float
	Status     common.String
	Title      common.String
	VideoId    common.String
}
