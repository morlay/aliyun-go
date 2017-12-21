package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetPhotosByMd5sRequest struct {
	requests.RpcRequest
	LibraryId string                  `position:"Query" name:"LibraryId"`
	StoreName string                  `position:"Query" name:"StoreName"`
	State     string                  `position:"Query" name:"State"`
	Md5s      *GetPhotosByMd5sMd5List `position:"Query" type:"Repeated" name:"Md.5"`
}

func (req *GetPhotosByMd5sRequest) Invoke(client *sdk.Client) (resp *GetPhotosByMd5sResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "GetPhotosByMd5s", "cloudphoto", "")
	resp = &GetPhotosByMd5sResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetPhotosByMd5sResponse struct {
	responses.BaseResponse
	Code      string
	Message   string
	RequestId string
	Action    string
	Photos    GetPhotosByMd5sPhotoList
}

type GetPhotosByMd5sPhoto struct {
	Id              int64
	Title           string
	FileId          string
	State           string
	Md5             string
	IsVideo         bool
	Remark          string
	Width           int64
	Height          int64
	Ctime           int64
	Mtime           int64
	TakenAt         int64
	ShareExpireTime int64
}

type GetPhotosByMd5sMd5List []string

func (list *GetPhotosByMd5sMd5List) UnmarshalJSON(data []byte) error {
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

type GetPhotosByMd5sPhotoList []GetPhotosByMd5sPhoto

func (list *GetPhotosByMd5sPhotoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetPhotosByMd5sPhoto)
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
