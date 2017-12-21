package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetThumbnailRequest struct {
	LibraryId string `position:"Query" name:"LibraryId"`
	PhotoId   int64  `position:"Query" name:"PhotoId"`
	StoreName string `position:"Query" name:"StoreName"`
	ZoomType  string `position:"Query" name:"ZoomType"`
}

func (r GetThumbnailRequest) Invoke(client *sdk.Client) (response *GetThumbnailResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetThumbnailRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "GetThumbnail", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		GetThumbnailResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetThumbnailResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetThumbnailResponse struct {
	Code         string
	Message      string
	ThumbnailUrl string
	RequestId    string
	Action       string
}
