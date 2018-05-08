package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type MoveFacePhotosRequest struct {
	requests.RpcRequest
	LibraryId    string                     `position:"Query" name:"LibraryId"`
	TargetFaceId int64                      `position:"Query" name:"TargetFaceId"`
	PhotoIds     *MoveFacePhotosPhotoIdList `position:"Query" type:"Repeated" name:"PhotoId"`
	StoreName    string                     `position:"Query" name:"StoreName"`
	SourceFaceId int64                      `position:"Query" name:"SourceFaceId"`
}

func (req *MoveFacePhotosRequest) Invoke(client *sdk.Client) (resp *MoveFacePhotosResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "MoveFacePhotos", "cloudphoto", "")
	resp = &MoveFacePhotosResponse{}
	err = client.DoAction(req, resp)
	return
}

type MoveFacePhotosResponse struct {
	responses.BaseResponse
	Code      common.String
	Message   common.String
	RequestId common.String
	Action    common.String
	Results   MoveFacePhotosResultList
}

type MoveFacePhotosResult struct {
	Id      common.Long
	IdStr   common.String
	Code    common.String
	Message common.String
}

type MoveFacePhotosPhotoIdList []int64

func (list *MoveFacePhotosPhotoIdList) UnmarshalJSON(data []byte) error {
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

type MoveFacePhotosResultList []MoveFacePhotosResult

func (list *MoveFacePhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]MoveFacePhotosResult)
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
