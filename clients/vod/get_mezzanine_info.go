package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	Mezzanine GetMezzanineInfoMezzanine
}

type GetMezzanineInfoMezzanine struct {
	VideoId          common.String
	Bitrate          common.String
	CreationTime     common.String
	Duration         common.String
	Fps              common.String
	Height           common.Long
	Width            common.Long
	Size             common.Long
	Status           common.String
	FileURL          common.String
	FileName         common.String
	CRC64            common.String
	PreprocessStatus common.String
}
