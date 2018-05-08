package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListAlbumsRequest struct {
	requests.RpcRequest
	Cursor    string `position:"Query" name:"Cursor"`
	Size      int    `position:"Query" name:"Size"`
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	State     string `position:"Query" name:"State"`
	Direction string `position:"Query" name:"Direction"`
}

func (req *ListAlbumsRequest) Invoke(client *sdk.Client) (resp *ListAlbumsResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "ListAlbums", "cloudphoto", "")
	resp = &ListAlbumsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListAlbumsResponse struct {
	responses.BaseResponse
	Code       common.String
	Message    common.String
	NextCursor common.String
	TotalCount common.Integer
	RequestId  common.String
	Action     common.String
	Albums     ListAlbumsAlbumList
}

type ListAlbumsAlbum struct {
	Id          common.Long
	IdStr       common.String
	Name        common.String
	State       common.String
	Remark      common.String
	PhotosCount common.Long
	Ctime       common.Long
	Mtime       common.Long
	Cover       ListAlbumsCover
}

type ListAlbumsCover struct {
	Id      common.Long
	IdStr   common.String
	Title   common.String
	FileId  common.String
	State   common.String
	Md5     common.String
	IsVideo bool
	Remark  common.String
	Width   common.Long
	Height  common.Long
	Ctime   common.Long
	Mtime   common.Long
}

type ListAlbumsAlbumList []ListAlbumsAlbum

func (list *ListAlbumsAlbumList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAlbumsAlbum)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
