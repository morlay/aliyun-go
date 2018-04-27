package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListFailureJobExecutionInstancesRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
	Count           int   `position:"Query" name:"Count"`
}

func (req *ListFailureJobExecutionInstancesRequest) Invoke(client *sdk.Client) (resp *ListFailureJobExecutionInstancesResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListFailureJobExecutionInstances", "", "")
	resp = &ListFailureJobExecutionInstancesResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListFailureJobExecutionInstancesResponse struct {
	responses.BaseResponse
	RequestId    string
	JobInstances ListFailureJobExecutionInstancesJobInstanceList
}

type ListFailureJobExecutionInstancesJobInstance struct {
	Id        string
	JobName   string
	StartTime int64
	RunTime   int
	JobType   string
	JobId     string
	ClusterId string
	Status    string
	RetryInfo string
}

type ListFailureJobExecutionInstancesJobInstanceList []ListFailureJobExecutionInstancesJobInstance

func (list *ListFailureJobExecutionInstancesJobInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListFailureJobExecutionInstancesJobInstance)
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
