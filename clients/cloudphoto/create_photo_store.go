package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreatePhotoStoreRequest struct {
	BucketName   string `position:"Query" name:"BucketName"`
	StoreName    string `position:"Query" name:"StoreName"`
	Remark       string `position:"Query" name:"Remark"`
	DefaultQuota int64  `position:"Query" name:"DefaultQuota"`
}

func (r CreatePhotoStoreRequest) Invoke(client *sdk.Client) (response *CreatePhotoStoreResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreatePhotoStoreRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "CreatePhotoStore", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		CreatePhotoStoreResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreatePhotoStoreResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreatePhotoStoreResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
}
