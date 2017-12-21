package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type LikePhotoRequest struct {
	requests.RpcRequest
	LibraryId string `position:"Query" name:"LibraryId"`
	PhotoId   int64  `position:"Query" name:"PhotoId"`
	StoreName string `position:"Query" name:"StoreName"`
}

func (req *LikePhotoRequest) Invoke(client *sdk.Client) (resp *LikePhotoResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "LikePhoto", "cloudphoto", "")
	resp = &LikePhotoResponse{}
	err = client.DoAction(req, resp)
	return
}

type LikePhotoResponse struct {
	responses.BaseResponse
	Code      string
	Message   string
	RequestId string
	Action    string
}
