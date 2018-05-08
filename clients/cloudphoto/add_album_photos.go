package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AddAlbumPhotosRequest struct {
	requests.RpcRequest
	LibraryId string                     `position:"Query" name:"LibraryId"`
	AlbumId   int64                      `position:"Query" name:"AlbumId"`
	PhotoIds  *AddAlbumPhotosPhotoIdList `position:"Query" type:"Repeated" name:"PhotoId"`
	StoreName string                     `position:"Query" name:"StoreName"`
}

func (req *AddAlbumPhotosRequest) Invoke(client *sdk.Client) (resp *AddAlbumPhotosResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "AddAlbumPhotos", "cloudphoto", "")
	resp = &AddAlbumPhotosResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddAlbumPhotosResponse struct {
	responses.BaseResponse
	Code      common.String
	Message   common.String
	RequestId common.String
	Action    common.String
	Results   AddAlbumPhotosResultList
}

type AddAlbumPhotosResult struct {
	Id      common.Long
	IdStr   common.String
	Code    common.String
	Message common.String
}

type AddAlbumPhotosPhotoIdList []int64

func (list *AddAlbumPhotosPhotoIdList) UnmarshalJSON(data []byte) error {
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

type AddAlbumPhotosResultList []AddAlbumPhotosResult

func (list *AddAlbumPhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AddAlbumPhotosResult)
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
