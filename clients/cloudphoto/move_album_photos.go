package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type MoveAlbumPhotosRequest struct {
	requests.RpcRequest
	SourceAlbumId int64                       `position:"Query" name:"SourceAlbumId"`
	TargetAlbumId int64                       `position:"Query" name:"TargetAlbumId"`
	LibraryId     string                      `position:"Query" name:"LibraryId"`
	PhotoIds      *MoveAlbumPhotosPhotoIdList `position:"Query" type:"Repeated" name:"PhotoId"`
	StoreName     string                      `position:"Query" name:"StoreName"`
}

func (req *MoveAlbumPhotosRequest) Invoke(client *sdk.Client) (resp *MoveAlbumPhotosResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "MoveAlbumPhotos", "cloudphoto", "")
	resp = &MoveAlbumPhotosResponse{}
	err = client.DoAction(req, resp)
	return
}

type MoveAlbumPhotosResponse struct {
	responses.BaseResponse
	Code      common.String
	Message   common.String
	RequestId common.String
	Action    common.String
	Results   MoveAlbumPhotosResultList
}

type MoveAlbumPhotosResult struct {
	Id      common.Long
	IdStr   common.String
	Code    common.String
	Message common.String
}

type MoveAlbumPhotosPhotoIdList []int64

func (list *MoveAlbumPhotosPhotoIdList) UnmarshalJSON(data []byte) error {
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

type MoveAlbumPhotosResultList []MoveAlbumPhotosResult

func (list *MoveAlbumPhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]MoveAlbumPhotosResult)
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
