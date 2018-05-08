package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId    common.String
	JobInstances ListFailureJobExecutionInstancesJobInstanceList
}

type ListFailureJobExecutionInstancesJobInstance struct {
	Id        common.String
	JobName   common.String
	StartTime common.Long
	RunTime   common.Integer
	JobType   common.String
	JobId     common.String
	ClusterId common.String
	Status    common.String
	RetryInfo common.String
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
