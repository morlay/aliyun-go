package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetPhotosRequest struct {
	LibraryId string                `position:"Query" name:"LibraryId"`
	PhotoIds  *GetPhotosPhotoIdList `position:"Query" type:"Repeated" name:"PhotoId"`
	StoreName string                `position:"Query" name:"StoreName"`
}

func (r GetPhotosRequest) Invoke(client *sdk.Client) (response *GetPhotosResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetPhotosRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "GetPhotos", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		GetPhotosResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetPhotosResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetPhotosResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Photos    GetPhotosPhotoList
}

type GetPhotosPhoto struct {
	Id              int64
	Title           string
	FileId          string
	State           string
	Md5             string
	IsVideo         bool
	Remark          string
	Width           int64
	Height          int64
	Ctime           int64
	Mtime           int64
	TakenAt         int64
	InactiveTime    int64
	ShareExpireTime int64
	Like            int64
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
