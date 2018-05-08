package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetAlbumCoverRequest struct {
	requests.RpcRequest
	LibraryId string `position:"Query" name:"LibraryId"`
	AlbumId   int64  `position:"Query" name:"AlbumId"`
	PhotoId   int64  `position:"Query" name:"PhotoId"`
	StoreName string `position:"Query" name:"StoreName"`
}

func (req *SetAlbumCoverRequest) Invoke(client *sdk.Client) (resp *SetAlbumCoverResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "SetAlbumCover", "cloudphoto", "")
	resp = &SetAlbumCoverResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetAlbumCoverResponse struct {
	responses.BaseResponse
	Code      common.String
	Message   common.String
	RequestId common.String
	Action    common.String
}
