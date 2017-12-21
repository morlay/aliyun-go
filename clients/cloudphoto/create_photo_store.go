package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreatePhotoStoreRequest struct {
	requests.RpcRequest
	BucketName   string `position:"Query" name:"BucketName"`
	StoreName    string `position:"Query" name:"StoreName"`
	Remark       string `position:"Query" name:"Remark"`
	DefaultQuota int64  `position:"Query" name:"DefaultQuota"`
}

func (req *CreatePhotoStoreRequest) Invoke(client *sdk.Client) (resp *CreatePhotoStoreResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "CreatePhotoStore", "cloudphoto", "")
	resp = &CreatePhotoStoreResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreatePhotoStoreResponse struct {
	responses.BaseResponse
	Code      string
	Message   string
	RequestId string
	Action    string
}
