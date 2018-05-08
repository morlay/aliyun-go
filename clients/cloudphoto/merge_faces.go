package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type MergeFacesRequest struct {
	requests.RpcRequest
	LibraryId    string                `position:"Query" name:"LibraryId"`
	TargetFaceId int64                 `position:"Query" name:"TargetFaceId"`
	StoreName    string                `position:"Query" name:"StoreName"`
	FaceIds      *MergeFacesFaceIdList `position:"Query" type:"Repeated" name:"FaceId"`
}

func (req *MergeFacesRequest) Invoke(client *sdk.Client) (resp *MergeFacesResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "MergeFaces", "cloudphoto", "")
	resp = &MergeFacesResponse{}
	err = client.DoAction(req, resp)
	return
}

type MergeFacesResponse struct {
	responses.BaseResponse
	Code      common.String
	Message   common.String
	RequestId common.String
	Action    common.String
	Results   MergeFacesResultList
}

type MergeFacesResult struct {
	Id      common.Long
	IdStr   common.String
	Code    common.String
	Message common.String
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
