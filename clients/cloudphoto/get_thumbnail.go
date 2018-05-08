package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetThumbnailRequest struct {
	requests.RpcRequest
	LibraryId string `position:"Query" name:"LibraryId"`
	PhotoId   int64  `position:"Query" name:"PhotoId"`
	StoreName string `position:"Query" name:"StoreName"`
	ZoomType  string `position:"Query" name:"ZoomType"`
}

func (req *GetThumbnailRequest) Invoke(client *sdk.Client) (resp *GetThumbnailResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "GetThumbnail", "cloudphoto", "")
	resp = &GetThumbnailResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetThumbnailResponse struct {
	responses.BaseResponse
	Code         common.String
	Message      common.String
	ThumbnailUrl common.String
	RequestId    common.String
	Action       common.String
}
