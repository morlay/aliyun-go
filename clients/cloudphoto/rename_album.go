package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RenameAlbumRequest struct {
	AlbumName string `position:"Query" name:"AlbumName"`
	LibraryId string `position:"Query" name:"LibraryId"`
	AlbumId   int64  `position:"Query" name:"AlbumId"`
	StoreName string `position:"Query" name:"StoreName"`
}

func (r RenameAlbumRequest) Invoke(client *sdk.Client) (response *RenameAlbumResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RenameAlbumRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "RenameAlbum", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		RenameAlbumResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RenameAlbumResponse

	err = client.DoAction(&req, &resp)
	return
}

type RenameAlbumResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
}
