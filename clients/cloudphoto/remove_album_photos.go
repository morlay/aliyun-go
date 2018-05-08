package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	Code      common.String
	Message   common.String
	RequestId common.String
	Action    common.String
	Results   RemoveAlbumPhotosResultList
}

type RemoveAlbumPhotosResult struct {
	Id      common.Long
	IdStr   common.String
	Code    common.String
	Message common.String
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
