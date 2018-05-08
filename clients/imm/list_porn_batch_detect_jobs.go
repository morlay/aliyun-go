package imm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListPornBatchDetectJobsRequest struct {
	requests.RpcRequest
	MaxKeys int    `position:"Query" name:"MaxKeys"`
	Marker  string `position:"Query" name:"Marker"`
	Project string `position:"Query" name:"Project"`
}

func (req *ListPornBatchDetectJobsRequest) Invoke(client *sdk.Client) (resp *ListPornBatchDetectJobsResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "ListPornBatchDetectJobs", "imm", "")
	resp = &ListPornBatchDetectJobsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListPornBatchDetectJobsResponse struct {
	responses.BaseResponse
	RequestId  common.String
	NextMarker common.String
	Jobs       ListPornBatchDetectJobsJobsItemList
}

type ListPornBatchDetectJobsJobsItem struct {
	JobId           common.String
	SrcUri          common.String
	Status          common.String
	TgtUri          common.String
	NotifyTopicName common.Integer
	NotifyEndpoint  common.String
	ExternalID      common.String
	CreateTime      common.String
	FinishTime      common.String
	Percent         common.Integer
}

type ListPornBatchDetectJobsJobsItemList []ListPornBatchDetectJobsJobsItem

func (list *ListPornBatchDetectJobsJobsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPornBatchDetectJobsJobsItem)
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
