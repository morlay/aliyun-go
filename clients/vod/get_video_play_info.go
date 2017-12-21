package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetVideoPlayInfoRequest struct {
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

func (r GetVideoPlayInfoRequest) Invoke(client *sdk.Client) (response *GetVideoPlayInfoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetVideoPlayInfoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("vod", "2017-03-21", "GetVideoPlayInfo", "", "")

	resp := struct {
		*responses.BaseResponse
		GetVideoPlayInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetVideoPlayInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetVideoPlayInfoResponse struct {
	RequestId string
	PlayInfo  GetVideoPlayInfoPlayInfo
	VideoInfo GetVideoPlayInfoVideoInfo
}

type GetVideoPlayInfoPlayInfo struct {
	AccessKeyId     string
	AccessKeySecret string
	AuthInfo        string
	SecurityToken   string
	Region          string
	PlayDomain      string
}

type GetVideoPlayInfoVideoInfo struct {
	CoverURL   string
	CustomerId int64
	Duration   float32
	Status     string
	Title      string
	VideoId    string
}
