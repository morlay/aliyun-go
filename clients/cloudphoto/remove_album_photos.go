package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RemoveAlbumPhotosRequest struct {
	requests.RpcRequest
	LibraryId string                        `position:"Query" name:"LibraryId"`
	AlbumId   int64                         `position:"Query" name:"AlbumId"`
	PhotoIds  *RemoveAlbumPhotosPhotoIdList `position:"Query" type:"Repeated" name:"PhotoId"`
	StoreName string                        `position:"Query" name:"StoreName"`
}

func (req *RemoveAlbumPhotosRequest) Invoke(client *sdk.Client) (resp *RemoveAlbumPhotosResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "RemoveAlbumPhotos", "cloudphoto", "")
	resp = &RemoveAlbumPhotosResponse{}
	err = client.DoAction(req, resp)
	return
}

type RemoveAlbumPhotosResponse struct {
	responses.BaseResponse
	Code      string
	Message   string
	RequestId string
	Action    string
	Results   RemoveAlbumPhotosResultList
}

type RemoveAlbumPhotosResult struct {
	Id      int64
	IdStr   string
	Code    string
	Message string
}

type RemoveAlbumPhotosPhotoIdList []int64

func (list *RemoveAlbumPhotosPhotoIdList) UnmarshalJSON(data []byte) error {
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

type RemoveAlbumPhotosResultList []RemoveAlbumPhotosResult

func (list *RemoveAlbumPhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]RemoveAlbumPhotosResult)
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
