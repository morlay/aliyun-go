package imm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId  string
	NextMarker string
	Jobs       ListPornBatchDetectJobsJobsItemList
}

type ListPornBatchDetectJobsJobsItem struct {
	JobId           string
	SrcUri          string
	Status          string
	TgtUri          string
	NotifyTopicName int
	NotifyEndpoint  string
	ExternalID      string
	CreateTime      string
	FinishTime      string
	Percent         int
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
