package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	Code      common.String
	Message   common.String
	RequestId common.String
	Action    common.String
	Photos    GetPhotosByMd5sPhotoList
}

type GetPhotosByMd5sPhoto struct {
	Id              common.Long
	IdStr           common.String
	Title           common.String
	FileId          common.String
	Location        common.String
	State           common.String
	Md5             common.String
	IsVideo         bool
	Remark          common.String
	Size            common.Long
	Width           common.Long
	Height          common.Long
	Ctime           common.Long
	Mtime           common.Long
	TakenAt         common.Long
	ShareExpireTime common.Long
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
