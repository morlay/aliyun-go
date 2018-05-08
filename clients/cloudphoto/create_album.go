package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateAlbumRequest struct {
	requests.RpcRequest
	AlbumName string `position:"Query" name:"AlbumName"`
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	Remark    string `position:"Query" name:"Remark"`
}

func (req *CreateAlbumRequest) Invoke(client *sdk.Client) (resp *CreateAlbumResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "CreateAlbum", "cloudphoto", "")
	resp = &CreateAlbumResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateAlbumResponse struct {
	responses.BaseResponse
	Code      common.String
	Message   common.String
	RequestId common.String
	Action    common.String
	Album     CreateAlbumAlbum
}

type CreateAlbumAlbum struct {
	Id          common.Long
	IdStr       common.String
	Name        common.String
	State       common.String
	Remark      common.String
	PhotosCount common.Long
	Ctime       common.Long
	Mtime       common.Long
	Cover       CreateAlbumCover
}

type CreateAlbumCover struct {
	Id      common.Long
	IdStr   common.String
	Title   common.String
	FileId  common.String
	State   common.String
	Md5     common.String
	IsVideo bool
	Width   common.Long
	Height  common.Long
	Ctime   common.Long
	Mtime   common.Long
	Remark  common.String
}
