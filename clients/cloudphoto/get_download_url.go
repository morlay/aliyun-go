package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetDownloadUrlRequest struct {
	LibraryId string `position:"Query" name:"LibraryId"`
	PhotoId   int64  `position:"Query" name:"PhotoId"`
	StoreName string `position:"Query" name:"StoreName"`
}

func (r GetDownloadUrlRequest) Invoke(client *sdk.Client) (response *GetDownloadUrlResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetDownloadUrlRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "GetDownloadUrl", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		GetDownloadUrlResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetDownloadUrlResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetDownloadUrlResponse struct {
	Code        string
	Message     string
	DownloadUrl string
	RequestId   string
	Action      string
}
