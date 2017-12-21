package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetVideoPlayAuthRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ReAuthInfo           string `position:"Query" name:"ReAuthInfo"`
	AuthInfoTimeout      int64  `position:"Query" name:"AuthInfoTimeout"`
	VideoId              string `position:"Query" name:"VideoId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r GetVideoPlayAuthRequest) Invoke(client *sdk.Client) (response *GetVideoPlayAuthResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetVideoPlayAuthRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("vod", "2017-03-21", "GetVideoPlayAuth", "", "")

	resp := struct {
		*responses.BaseResponse
		GetVideoPlayAuthResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetVideoPlayAuthResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetVideoPlayAuthResponse struct {
	RequestId string
	PlayAuth  string
	VideoMeta GetVideoPlayAuthVideoMeta
}

type GetVideoPlayAuthVideoMeta struct {
	CoverURL string
	Duration float32
	Status   string
	Title    string
	VideoId  string
}
