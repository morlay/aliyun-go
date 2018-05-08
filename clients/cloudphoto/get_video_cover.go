package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	Code          common.String
	Message       common.String
	VideoCoverUrl common.String
	RequestId     common.String
	Action        common.String
}
