package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReactivatePhotosRequest struct {
	LibraryId string                       `position:"Query" name:"LibraryId"`
	PhotoIds  *ReactivatePhotosPhotoIdList `position:"Query" type:"Repeated" name:"PhotoId"`
	StoreName string                       `position:"Query" name:"StoreName"`
}

func (r ReactivatePhotosRequest) Invoke(client *sdk.Client) (response *ReactivatePhotosResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ReactivatePhotosRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "ReactivatePhotos", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		ReactivatePhotosResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ReactivatePhotosResponse

	err = client.DoAction(&req, &resp)
	return
}

type ReactivatePhotosResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Results   ReactivatePhotosResultList
}

type ReactivatePhotosResult struct {
	Id      int64
	Code    string
	Message string
}

type ReactivatePhotosPhotoIdList []int64

func (list *ReactivatePhotosPhotoIdList) UnmarshalJSON(data []byte) error {
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

type ReactivatePhotosResultList []ReactivatePhotosResult

func (list *ReactivatePhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ReactivatePhotosResult)
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
