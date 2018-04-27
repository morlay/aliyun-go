package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	Code       string
	Message    string
	TotalCount int
	RequestId  string
	Action     string
	Photos     FetchPhotosPhotoList
}

type FetchPhotosPhoto struct {
	Id              int64
	IdStr           string
	Title           string
	FileId          string
	Location        string
	State           string
	Md5             string
	IsVideo         bool
	Remark          string
	Size            int64
	Width           int64
	Height          int64
	Ctime           int64
	Mtime           int64
	TakenAt         int64
	InactiveTime    int64
	ShareExpireTime int64
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
