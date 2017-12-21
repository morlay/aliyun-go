package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateUploadImageRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ImageType            string `position:"Query" name:"ImageType"`
	OriginalFileName     string `position:"Query" name:"OriginalFileName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ImageExt             string `position:"Query" name:"ImageExt"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Title                string `position:"Query" name:"Title"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (r CreateUploadImageRequest) Invoke(client *sdk.Client) (response *CreateUploadImageResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateUploadImageRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("vod", "2017-03-21", "CreateUploadImage", "", "")

	resp := struct {
		*responses.BaseResponse
		CreateUploadImageResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateUploadImageResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateUploadImageResponse struct {
	RequestId     string
	ImageId       string
	ImageURL      string
	UploadAddress string
	UploadAuth    string
}
