package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	Code       string
	Message    string
	TotalCount int
	RequestId  string
	Action     string
	Photos     SearchPhotosPhotoList
}

type SearchPhotosPhoto struct {
	Id              int64
	Title           string
	FileId          string
	State           string
	Md5             string
	IsVideo         bool
	Width           int64
	Height          int64
	Ctime           int64
	Mtime           int64
	TakenAt         int64
	ShareExpireTime int64
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
