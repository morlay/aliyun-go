package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListPhotosRequest struct {
	requests.RpcRequest
	Cursor    string `position:"Query" name:"Cursor"`
	Size      int    `position:"Query" name:"Size"`
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	State     string `position:"Query" name:"State"`
	Direction string `position:"Query" name:"Direction"`
}

func (req *ListPhotosRequest) Invoke(client *sdk.Client) (resp *ListPhotosResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "ListPhotos", "cloudphoto", "")
	resp = &ListPhotosResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListPhotosResponse struct {
	responses.BaseResponse
	Code       common.String
	Message    common.String
	NextCursor common.String
	TotalCount common.Integer
	RequestId  common.String
	Action     common.String
	Photos     ListPhotosPhotoList
}

type ListPhotosPhoto struct {
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
	InactiveTime    common.Long
	ShareExpireTime common.Long
}

type ListPhotosPhotoList []ListPhotosPhoto

func (list *ListPhotosPhotoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPhotosPhoto)
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
