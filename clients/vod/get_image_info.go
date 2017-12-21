package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetImageInfoRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ImageId              string `position:"Query" name:"ImageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	AuthTimeout          int64  `position:"Query" name:"AuthTimeout"`
}

func (r GetImageInfoRequest) Invoke(client *sdk.Client) (response *GetImageInfoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetImageInfoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("vod", "2017-03-21", "GetImageInfo", "", "")

	resp := struct {
		*responses.BaseResponse
		GetImageInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetImageInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetImageInfoResponse struct {
	RequestId string
	ImageInfo GetImageInfoImageInfo
}

type GetImageInfoImageInfo struct {
	ImageId      string
	Title        string
	CreationTime string
	ImageType    string
	Tags         string
	URL          string
	Mezzanine    GetImageInfoMezzanine
}

type GetImageInfoMezzanine struct {
	OriginalFileName string
}
