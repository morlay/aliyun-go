package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SearchPhotosRequest struct {
	requests.RpcRequest
	Size      int    `position:"Query" name:"Size"`
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	Page      int    `position:"Query" name:"Page"`
	Keyword   string `position:"Query" name:"Keyword"`
}

func (req *SearchPhotosRequest) Invoke(client *sdk.Client) (resp *SearchPhotosResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "SearchPhotos", "cloudphoto", "")
	resp = &SearchPhotosResponse{}
	err = client.DoAction(req, resp)
	return
}

type SearchPhotosResponse struct {
	responses.BaseResponse
	Code       common.String
	Message    common.String
	TotalCount common.Integer
	RequestId  common.String
	Action     common.String
	Photos     SearchPhotosPhotoList
}

type SearchPhotosPhoto struct {
	Id              common.Long
	IdStr           common.String
	Title           common.String
	FileId          common.String
	Location        common.String
	State           common.String
	Md5             common.String
	IsVideo         bool
	Size            common.Long
	Width           common.Long
	Height          common.Long
	Ctime           common.Long
	Mtime           common.Long
	TakenAt         common.Long
	ShareExpireTime common.Long
}

type SearchPhotosPhotoList []SearchPhotosPhoto

func (list *SearchPhotosPhotoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SearchPhotosPhoto)
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
