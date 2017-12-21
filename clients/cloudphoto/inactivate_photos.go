package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type InactivatePhotosRequest struct {
	requests.RpcRequest
	LibraryId    string                       `position:"Query" name:"LibraryId"`
	PhotoIds     *InactivatePhotosPhotoIdList `position:"Query" type:"Repeated" name:"PhotoId"`
	StoreName    string                       `position:"Query" name:"StoreName"`
	InactiveTime int64                        `position:"Query" name:"InactiveTime"`
}

func (req *InactivatePhotosRequest) Invoke(client *sdk.Client) (resp *InactivatePhotosResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "InactivatePhotos", "cloudphoto", "")
	resp = &InactivatePhotosResponse{}
	err = client.DoAction(req, resp)
	return
}

type InactivatePhotosResponse struct {
	responses.BaseResponse
	Code      string
	Message   string
	RequestId string
	Action    string
	Results   InactivatePhotosResultList
}

type InactivatePhotosResult struct {
	Id      int64
	Code    string
	Message string
}

type InactivatePhotosPhotoIdList []int64

func (list *InactivatePhotosPhotoIdList) UnmarshalJSON(data []byte) error {
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

type InactivatePhotosResultList []InactivatePhotosResult

func (list *InactivatePhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]InactivatePhotosResult)
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
