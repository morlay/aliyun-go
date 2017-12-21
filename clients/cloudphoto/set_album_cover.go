package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetAlbumCoverRequest struct {
	LibraryId string `position:"Query" name:"LibraryId"`
	AlbumId   int64  `position:"Query" name:"AlbumId"`
	PhotoId   int64  `position:"Query" name:"PhotoId"`
	StoreName string `position:"Query" name:"StoreName"`
}

func (r SetAlbumCoverRequest) Invoke(client *sdk.Client) (response *SetAlbumCoverResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetAlbumCoverRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "SetAlbumCover", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		SetAlbumCoverResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetAlbumCoverResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetAlbumCoverResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
}
