package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeletePhotosRequest struct {
	requests.RpcRequest
	LibraryId string                   `position:"Query" name:"LibraryId"`
	StoreName string                   `position:"Query" name:"StoreName"`
	PhotoIds  *DeletePhotosPhotoIdList `position:"Query" type:"Repeated" name:"PhotoId"`
}

func (req *DeletePhotosRequest) Invoke(client *sdk.Client) (resp *DeletePhotosResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "DeletePhotos", "cloudphoto", "")
	resp = &DeletePhotosResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeletePhotosResponse struct {
	responses.BaseResponse
	Code      string
	Message   string
	RequestId string
	Action    string
	Results   DeletePhotosResultList
}

type DeletePhotosResult struct {
	Id      int64
	IdStr   string
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
