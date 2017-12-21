package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetMezzanineInfoRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VideoId              string `position:"Query" name:"VideoId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	AuthTimeout          int64  `position:"Query" name:"AuthTimeout"`
}

func (r GetMezzanineInfoRequest) Invoke(client *sdk.Client) (response *GetMezzanineInfoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetMezzanineInfoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("vod", "2017-03-21", "GetMezzanineInfo", "", "")

	resp := struct {
		*responses.BaseResponse
		GetMezzanineInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetMezzanineInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetMezzanineInfoResponse struct {
	RequestId string
	Mezzanine GetMezzanineInfoMezzanine
}

type GetMezzanineInfoMezzanine struct {
	VideoId      string
	Bitrate      string
	CreationTime string
	Duration     string
	Fps          string
	Height       int64
	Width        int64
	Size         int64
	Status       string
	FileURL      string
	FileName     string
	CRC64        string
}
