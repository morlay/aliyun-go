package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteFacesRequest struct {
	requests.RpcRequest
	LibraryId string                 `position:"Query" name:"LibraryId"`
	StoreName string                 `position:"Query" name:"StoreName"`
	FaceIds   *DeleteFacesFaceIdList `position:"Query" type:"Repeated" name:"FaceId"`
}

func (req *DeleteFacesRequest) Invoke(client *sdk.Client) (resp *DeleteFacesResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "DeleteFaces", "cloudphoto", "")
	resp = &DeleteFacesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteFacesResponse struct {
	responses.BaseResponse
	Code      common.String
	Message   common.String
	RequestId common.String
	Action    common.String
	Results   DeleteFacesResultList
}

type DeleteFacesResult struct {
	Id      common.Long
	IdStr   common.String
	Code    common.String
	Message common.String
}

type DeleteFacesFaceIdList []int64

func (list *DeleteFacesFaceIdList) UnmarshalJSON(data []byte) error {
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

type DeleteFacesResultList []DeleteFacesResult

func (list *DeleteFacesResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DeleteFacesResult)
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
