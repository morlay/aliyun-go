package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	Code      common.String
	Message   common.String
	RequestId common.String
	Action    common.String
	Albums    GetAlbumsByNamesAlbumList
}

type GetAlbumsByNamesAlbum struct {
	Id          common.Long
	IdStr       common.String
	Name        common.String
	State       common.String
	PhotosCount common.Long
	Ctime       common.Long
	Mtime       common.Long
	Cover       GetAlbumsByNamesCover
}

type GetAlbumsByNamesCover struct {
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
