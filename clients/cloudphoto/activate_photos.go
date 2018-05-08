package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ActivatePhotosRequest struct {
	requests.RpcRequest
	LibraryId string                     `position:"Query" name:"LibraryId"`
	PhotoIds  *ActivatePhotosPhotoIdList `position:"Query" type:"Repeated" name:"PhotoId"`
	StoreName string                     `position:"Query" name:"StoreName"`
}

func (req *ActivatePhotosRequest) Invoke(client *sdk.Client) (resp *ActivatePhotosResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "ActivatePhotos", "cloudphoto", "")
	resp = &ActivatePhotosResponse{}
	err = client.DoAction(req, resp)
	return
}

type ActivatePhotosResponse struct {
	responses.BaseResponse
	Code      common.String
	Message   common.String
	RequestId common.String
	Action    common.String
	Results   ActivatePhotosResultList
}

type ActivatePhotosResult struct {
	Id      common.Long
	IdStr   common.String
	Code    common.String
	Message common.String
}

type ActivatePhotosPhotoIdList []int64

func (list *ActivatePhotosPhotoIdList) UnmarshalJSON(data []byte) error {
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

type ActivatePhotosResultList []ActivatePhotosResult

func (list *ActivatePhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ActivatePhotosResult)
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
