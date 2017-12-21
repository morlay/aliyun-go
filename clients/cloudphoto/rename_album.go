package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RenameAlbumRequest struct {
	requests.RpcRequest
	AlbumName string `position:"Query" name:"AlbumName"`
	LibraryId string `position:"Query" name:"LibraryId"`
	AlbumId   int64  `position:"Query" name:"AlbumId"`
	StoreName string `position:"Query" name:"StoreName"`
}

func (req *RenameAlbumRequest) Invoke(client *sdk.Client) (resp *RenameAlbumResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "RenameAlbum", "cloudphoto", "")
	resp = &RenameAlbumResponse{}
	err = client.DoAction(req, resp)
	return
}

type RenameAlbumResponse struct {
	responses.BaseResponse
	Code      string
	Message   string
	RequestId string
	Action    string
}
