package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetAlbumsByNamesRequest struct {
	requests.RpcRequest
	LibraryId string                    `position:"Query" name:"LibraryId"`
	Names     *GetAlbumsByNamesNameList `position:"Query" type:"Repeated" name:"Name"`
	StoreName string                    `position:"Query" name:"StoreName"`
}

func (req *GetAlbumsByNamesRequest) Invoke(client *sdk.Client) (resp *GetAlbumsByNamesResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "GetAlbumsByNames", "cloudphoto", "")
	resp = &GetAlbumsByNamesResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetAlbumsByNamesResponse struct {
	responses.BaseResponse
	Code      string
	Message   string
	RequestId string
	Action    string
	Albums    GetAlbumsByNamesAlbumList
}

type GetAlbumsByNamesAlbum struct {
	Id          int64
	IdStr       string
	Name        string
	State       string
	PhotosCount int64
	Ctime       int64
	Mtime       int64
	Cover       GetAlbumsByNamesCover
}

type GetAlbumsByNamesCover struct {
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

type GetAlbumsByNamesNameList []string

func (list *GetAlbumsByNamesNameList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type GetAlbumsByNamesAlbumList []GetAlbumsByNamesAlbum

func (list *GetAlbumsByNamesAlbumList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetAlbumsByNamesAlbum)
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
