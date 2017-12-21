package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateAlbumRequest struct {
	AlbumName string `position:"Query" name:"AlbumName"`
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	Remark    string `position:"Query" name:"Remark"`
}

func (r CreateAlbumRequest) Invoke(client *sdk.Client) (response *CreateAlbumResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateAlbumRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "CreateAlbum", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		CreateAlbumResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateAlbumResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateAlbumResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Album     CreateAlbumAlbum
}

type CreateAlbumAlbum struct {
	Id          int64
	Name        string
	State       string
	Remark      string
	PhotosCount int64
	Ctime       int64
	Mtime       int64
	Cover       CreateAlbumCover
}

type CreateAlbumCover struct {
	Id      int64
	Title   string
	FileId  string
	State   string
	Md5     string
	IsVideo bool
	Width   int64
	Height  int64
	Ctime   int64
	Mtime   int64
	Remark  string
}
