package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ActivatePhotosRequest struct {
	LibraryId string                     `position:"Query" name:"LibraryId"`
	PhotoIds  *ActivatePhotosPhotoIdList `position:"Query" type:"Repeated" name:"PhotoId"`
	StoreName string                     `position:"Query" name:"StoreName"`
}

func (r ActivatePhotosRequest) Invoke(client *sdk.Client) (response *ActivatePhotosResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ActivatePhotosRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "ActivatePhotos", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		ActivatePhotosResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ActivatePhotosResponse

	err = client.DoAction(&req, &resp)
	return
}

type ActivatePhotosResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Results   ActivatePhotosResultList
}

type ActivatePhotosResult struct {
	Id      int64
	Code    string
	Message string
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
