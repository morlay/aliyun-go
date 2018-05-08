package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type FetchPhotosRequest struct {
	requests.RpcRequest
	Size      int    `position:"Query" name:"Size"`
	LibraryId string `position:"Query" name:"LibraryId"`
	OrderBy   string `position:"Query" name:"OrderBy"`
	StoreName string `position:"Query" name:"StoreName"`
	State     string `position:"Query" name:"State"`
	Page      int    `position:"Query" name:"Page"`
	Order     string `position:"Query" name:"Order"`
}

func (req *FetchPhotosRequest) Invoke(client *sdk.Client) (resp *FetchPhotosResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "FetchPhotos", "cloudphoto", "")
	resp = &FetchPhotosResponse{}
	err = client.DoAction(req, resp)
	return
}

type FetchPhotosResponse struct {
	responses.BaseResponse
	Code       common.String
	Message    common.String
	TotalCount common.Integer
	RequestId  common.String
	Action     common.String
	Photos     FetchPhotosPhotoList
}

type FetchPhotosPhoto struct {
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

type FetchPhotosPhotoList []FetchPhotosPhoto

func (list *FetchPhotosPhotoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FetchPhotosPhoto)
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
