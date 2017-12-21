package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetVideoCoverRequest struct {
	LibraryId string `position:"Query" name:"LibraryId"`
	PhotoId   int64  `position:"Query" name:"PhotoId"`
	StoreName string `position:"Query" name:"StoreName"`
	ZoomType  string `position:"Query" name:"ZoomType"`
}

func (r GetVideoCoverRequest) Invoke(client *sdk.Client) (response *GetVideoCoverResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetVideoCoverRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "GetVideoCover", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		GetVideoCoverResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetVideoCoverResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetVideoCoverResponse struct {
	Code          string
	Message       string
	VideoCoverUrl string
	RequestId     string
	Action        string
}
