package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetVideoConfigRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VideoId              string `position:"Query" name:"VideoId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	AuthInfo             string `position:"Query" name:"AuthInfo"`
}

func (r GetVideoConfigRequest) Invoke(client *sdk.Client) (response *GetVideoConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetVideoConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("vod", "2017-03-21", "GetVideoConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		GetVideoConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetVideoConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetVideoConfigResponse struct {
	RequestId      string
	DownloadSwitch string
}
