package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListJobExecutionInstancesRequest struct {
	requests.RpcRequest
	ResourceOwnerId         int64  `position:"Query" name:"ResourceOwnerId"`
	ExecutionPlanInstanceId string `position:"Query" name:"ExecutionPlanInstanceId"`
	PageSize                int    `position:"Query" name:"PageSize"`
	IsDesc                  string `position:"Query" name:"IsDesc"`
	PageNumber              int    `position:"Query" name:"PageNumber"`
}

func (req *ListJobExecutionInstancesRequest) Invoke(client *sdk.Client) (resp *ListJobExecutionInstancesResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListJobExecutionInstances", "", "")
	resp = &ListJobExecutionInstancesResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListJobExecutionInstancesResponse struct {
	responses.BaseResponse
	RequestId    common.String
	TotalCount   common.Integer
	PageNumber   common.Integer
	PageSize     common.Integer
	JobInstances ListJobExecutionInstancesJobInstanceList
}

type ListJobExecutionInstancesJobInstance struct {
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

type ListJobExecutionInstancesJobInstanceList []ListJobExecutionInstancesJobInstance

func (list *ListJobExecutionInstancesJobInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobExecutionInstancesJobInstance)
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
