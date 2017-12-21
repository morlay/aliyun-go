package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RenameFaceRequest struct {
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	FaceId    int64  `position:"Query" name:"FaceId"`
	FaceName  string `position:"Query" name:"FaceName"`
}

func (r RenameFaceRequest) Invoke(client *sdk.Client) (response *RenameFaceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RenameFaceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "RenameFace", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		RenameFaceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RenameFaceResponse

	err = client.DoAction(&req, &resp)
	return
}

type RenameFaceResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
}
