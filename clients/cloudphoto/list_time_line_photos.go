package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListTimeLinePhotosRequest struct {
	requests.RpcRequest
	Size      int    `position:"Query" name:"Size"`
	LibraryId string `position:"Query" name:"LibraryId"`
	EndTime   int64  `position:"Query" name:"EndTime"`
	StoreName string `position:"Query" name:"StoreName"`
	Page      int    `position:"Query" name:"Page"`
	StartTime int64  `position:"Query" name:"StartTime"`
	FilterBy  string `position:"Query" name:"FilterBy"`
	Direction string `position:"Query" name:"Direction"`
	Order     string `position:"Query" name:"Order"`
}

func (req *ListTimeLinePhotosRequest) Invoke(client *sdk.Client) (resp *ListTimeLinePhotosResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "ListTimeLinePhotos", "cloudphoto", "")
	resp = &ListTimeLinePhotosResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListTimeLinePhotosResponse struct {
	responses.BaseResponse
	Code       string
	Message    string
	TotalCount int
	RequestId  string
	Action     string
	Photos     ListTimeLinePhotosPhotoList
}

type ListTimeLinePhotosPhoto struct {
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
	ShareExpireTime int64
	Like            int64
}

type ListTimeLinePhotosPhotoList []ListTimeLinePhotosPhoto

func (list *ListTimeLinePhotosPhotoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListTimeLinePhotosPhoto)
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
