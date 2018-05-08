package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetPhotosRequest struct {
	requests.RpcRequest
	LibraryId string                `position:"Query" name:"LibraryId"`
	PhotoIds  *GetPhotosPhotoIdList `position:"Query" type:"Repeated" name:"PhotoId"`
	StoreName string                `position:"Query" name:"StoreName"`
}

func (req *GetPhotosRequest) Invoke(client *sdk.Client) (resp *GetPhotosResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "GetPhotos", "cloudphoto", "")
	resp = &GetPhotosResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetPhotosResponse struct {
	responses.BaseResponse
	Code      common.String
	Message   common.String
	RequestId common.String
	Action    common.String
	Photos    GetPhotosPhotoList
}

type GetPhotosPhoto struct {
	Id              common.Long
	IdStr           common.String
	Title           common.String
	FileId          common.String
	Location        common.String
	State           common.String
	Md5             common.String
	IsVideo         bool
	Remark          common.String
	Width           common.Long
	Height          common.Long
	Size            common.Long
	Ctime           common.Long
	Mtime           common.Long
	TakenAt         common.Long
	InactiveTime    common.Long
	ShareExpireTime common.Long
	Like            common.Long
}

type GetPhotosPhotoIdList []int64

func (list *GetPhotosPhotoIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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

type GetPhotosPhotoList []GetPhotosPhoto

func (list *GetPhotosPhotoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetPhotosPhoto)
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
