package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetSimilarPhotosRequest struct {
	requests.RpcRequest
	LibraryId string `position:"Query" name:"LibraryId"`
	PhotoId   int64  `position:"Query" name:"PhotoId"`
	StoreName string `position:"Query" name:"StoreName"`
}

func (req *GetSimilarPhotosRequest) Invoke(client *sdk.Client) (resp *GetSimilarPhotosResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "GetSimilarPhotos", "cloudphoto", "")
	resp = &GetSimilarPhotosResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetSimilarPhotosResponse struct {
	responses.BaseResponse
	Code      string
	Message   string
	RequestId string
	Action    string
	Photos    GetSimilarPhotosPhotoList
}

type GetSimilarPhotosPhoto struct {
	Id              int64
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
	Like            int64
}

type GetSimilarPhotosPhotoList []GetSimilarPhotosPhoto

func (list *GetSimilarPhotosPhotoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetSimilarPhotosPhoto)
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
