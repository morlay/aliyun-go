package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeletePhotosRequest struct {
	LibraryId string                   `position:"Query" name:"LibraryId"`
	PhotoIds  *DeletePhotosPhotoIdList `position:"Query" type:"Repeated" name:"PhotoId"`
	StoreName string                   `position:"Query" name:"StoreName"`
}

func (r DeletePhotosRequest) Invoke(client *sdk.Client) (response *DeletePhotosResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeletePhotosRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "DeletePhotos", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		DeletePhotosResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeletePhotosResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeletePhotosResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Results   DeletePhotosResultList
}

type DeletePhotosResult struct {
	Id      int64
	Code    string
	Message string
}

type DeletePhotosPhotoIdList []int64

func (list *DeletePhotosPhotoIdList) UnmarshalJSON(data []byte) error {
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

type DeletePhotosResultList []DeletePhotosResult

func (list *DeletePhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DeletePhotosResult)
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
