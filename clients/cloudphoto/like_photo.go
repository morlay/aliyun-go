package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type LikePhotoRequest struct {
	LibraryId string `position:"Query" name:"LibraryId"`
	PhotoId   int64  `position:"Query" name:"PhotoId"`
	StoreName string `position:"Query" name:"StoreName"`
}

func (r LikePhotoRequest) Invoke(client *sdk.Client) (response *LikePhotoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		LikePhotoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "LikePhoto", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		LikePhotoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.LikePhotoResponse

	err = client.DoAction(&req, &resp)
	return
}

type LikePhotoResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
}
