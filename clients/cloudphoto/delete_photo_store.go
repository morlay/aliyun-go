package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeletePhotoStoreRequest struct {
	StoreName string `position:"Query" name:"StoreName"`
}

func (r DeletePhotoStoreRequest) Invoke(client *sdk.Client) (response *DeletePhotoStoreResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeletePhotoStoreRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "DeletePhotoStore", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		DeletePhotoStoreResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeletePhotoStoreResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeletePhotoStoreResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
}
