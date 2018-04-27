package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetMezzanineInfoRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VideoId              string `position:"Query" name:"VideoId"`
	PreviewSegment       string `position:"Query" name:"PreviewSegment"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	AuthTimeout          int64  `position:"Query" name:"AuthTimeout"`
}

func (req *GetMezzanineInfoRequest) Invoke(client *sdk.Client) (resp *GetMezzanineInfoResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "GetMezzanineInfo", "vod", "")
	resp = &GetMezzanineInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetMezzanineInfoResponse struct {
	responses.BaseResponse
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
