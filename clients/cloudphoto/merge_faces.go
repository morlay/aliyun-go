package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type MergeFacesRequest struct {
	LibraryId    string                `position:"Query" name:"LibraryId"`
	TargetFaceId int64                 `position:"Query" name:"TargetFaceId"`
	StoreName    string                `position:"Query" name:"StoreName"`
	FaceIds      *MergeFacesFaceIdList `position:"Query" type:"Repeated" name:"FaceId"`
}

func (r MergeFacesRequest) Invoke(client *sdk.Client) (response *MergeFacesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		MergeFacesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "MergeFaces", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		MergeFacesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.MergeFacesResponse

	err = client.DoAction(&req, &resp)
	return
}

type MergeFacesResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Results   MergeFacesResultList
}

type MergeFacesResult struct {
	Id      int64
	Code    string
	Message string
}

type MergeFacesFaceIdList []int64

func (list *MergeFacesFaceIdList) UnmarshalJSON(data []byte) error {
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

type MergeFacesResultList []MergeFacesResult

func (list *MergeFacesResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]MergeFacesResult)
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
