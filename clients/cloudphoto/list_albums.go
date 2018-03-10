package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	Code       string
	Message    string
	NextCursor string
	TotalCount int
	RequestId  string
	Action     string
	Albums     ListAlbumsAlbumList
}

type ListAlbumsAlbum struct {
	Id          int64
	IdStr       string
	Name        string
	State       string
	PhotosCount int64
	Ctime       int64
	Mtime       int64
	Cover       ListAlbumsCover
}

type ListAlbumsCover struct {
	Id      int64
	IdStr   string
	Title   string
	FileId  string
	State   string
	Md5     string
	IsVideo bool
	Remark  string
	Width   int64
	Height  int64
	Ctime   int64
	Mtime   int64
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
