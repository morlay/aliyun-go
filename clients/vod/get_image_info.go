package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetImageInfoRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ImageId              string `position:"Query" name:"ImageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	AuthTimeout          int64  `position:"Query" name:"AuthTimeout"`
}

func (req *GetImageInfoRequest) Invoke(client *sdk.Client) (resp *GetImageInfoResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "GetImageInfo", "vod", "")
	resp = &GetImageInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetImageInfoResponse struct {
	responses.BaseResponse
	RequestId common.String
	ImageInfo GetImageInfoImageInfo
}

type GetImageInfoImageInfo struct {
	ImageId      common.String
	Title        common.String
	CreationTime common.String
	ImageType    common.String
	Tags         common.String
	URL          common.String
	Mezzanine    GetImageInfoMezzanine
}

type GetImageInfoMezzanine struct {
	OriginalFileName common.String
}
