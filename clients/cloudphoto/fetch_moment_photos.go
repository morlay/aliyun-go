package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type FetchMomentPhotosRequest struct {
	requests.RpcRequest
	Size      int    `position:"Query" name:"Size"`
	LibraryId string `position:"Query" name:"LibraryId"`
	OrderBy   string `position:"Query" name:"OrderBy"`
	StoreName string `position:"Query" name:"StoreName"`
	Page      int    `position:"Query" name:"Page"`
	MomentId  int64  `position:"Query" name:"MomentId"`
	Order     string `position:"Query" name:"Order"`
}

func (req *FetchMomentPhotosRequest) Invoke(client *sdk.Client) (resp *FetchMomentPhotosResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "FetchMomentPhotos", "cloudphoto", "")
	resp = &FetchMomentPhotosResponse{}
	err = client.DoAction(req, resp)
	return
}

type FetchMomentPhotosResponse struct {
	responses.BaseResponse
	Code       string
	Message    string
	TotalCount int
	RequestId  string
	Action     string
	Photos     FetchMomentPhotosPhotoList
}

type FetchMomentPhotosPhoto struct {
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

type FetchMomentPhotosPhotoList []FetchMomentPhotosPhoto

func (list *FetchMomentPhotosPhotoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FetchMomentPhotosPhoto)
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
