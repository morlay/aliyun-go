package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	Code      common.String
	Message   common.String
	RequestId common.String
	Action    common.String
	Results   DeletePhotosResultList
}

type DeletePhotosResult struct {
	Id      common.Long
	IdStr   common.String
	Code    common.String
	Message common.String
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
