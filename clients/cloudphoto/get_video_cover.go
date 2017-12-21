package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetVideoCoverRequest struct {
	requests.RpcRequest
	LibraryId string `position:"Query" name:"LibraryId"`
	PhotoId   int64  `position:"Query" name:"PhotoId"`
	StoreName string `position:"Query" name:"StoreName"`
	ZoomType  string `position:"Query" name:"ZoomType"`
}

func (req *GetVideoCoverRequest) Invoke(client *sdk.Client) (resp *GetVideoCoverResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "GetVideoCover", "cloudphoto", "")
	resp = &GetVideoCoverResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetVideoCoverResponse struct {
	responses.BaseResponse
	Code          string
	Message       string
	VideoCoverUrl string
	RequestId     string
	Action        string
}
