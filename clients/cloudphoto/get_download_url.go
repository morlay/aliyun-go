package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetDownloadUrlRequest struct {
	requests.RpcRequest
	LibraryId string `position:"Query" name:"LibraryId"`
	PhotoId   int64  `position:"Query" name:"PhotoId"`
	StoreName string `position:"Query" name:"StoreName"`
}

func (req *GetDownloadUrlRequest) Invoke(client *sdk.Client) (resp *GetDownloadUrlResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "GetDownloadUrl", "cloudphoto", "")
	resp = &GetDownloadUrlResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetDownloadUrlResponse struct {
	responses.BaseResponse
	Code        string
	Message     string
	DownloadUrl string
	RequestId   string
	Action      string
}
