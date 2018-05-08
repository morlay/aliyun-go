package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateUploadImageRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ImageType            string `position:"Query" name:"ImageType"`
	OriginalFileName     string `position:"Query" name:"OriginalFileName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ImageExt             string `position:"Query" name:"ImageExt"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Title                string `position:"Query" name:"Title"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (req *CreateUploadImageRequest) Invoke(client *sdk.Client) (resp *CreateUploadImageResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "CreateUploadImage", "vod", "")
	resp = &CreateUploadImageResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateUploadImageResponse struct {
	responses.BaseResponse
	RequestId     common.String
	ImageId       common.String
	ImageURL      common.String
	UploadAddress common.String
	UploadAuth    common.String
}
