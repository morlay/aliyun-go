package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SearchPhotosRequest struct {
	Size      int    `position:"Query" name:"Size"`
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	Page      int    `position:"Query" name:"Page"`
	Keyword   string `position:"Query" name:"Keyword"`
}

func (r SearchPhotosRequest) Invoke(client *sdk.Client) (response *SearchPhotosResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SearchPhotosRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "SearchPhotos", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		SearchPhotosResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SearchPhotosResponse

	err = client.DoAction(&req, &resp)
	return
}

type SearchPhotosResponse struct {
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
