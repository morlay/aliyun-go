package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListJobExecutionInstancesRequest struct {
	ResourceOwnerId         int64  `position:"Query" name:"ResourceOwnerId"`
	ExecutionPlanInstanceId string `position:"Query" name:"ExecutionPlanInstanceId"`
	IsDesc                  string `position:"Query" name:"IsDesc"`
	PageNumber              int    `position:"Query" name:"PageNumber"`
	PageSize                int    `position:"Query" name:"PageSize"`
}

func (r ListJobExecutionInstancesRequest) Invoke(client *sdk.Client) (response *ListJobExecutionInstancesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListJobExecutionInstancesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Emr", "2016-04-08", "ListJobExecutionInstances", "", "")

	resp := struct {
		*responses.BaseResponse
		ListJobExecutionInstancesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListJobExecutionInstancesResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListJobExecutionInstancesResponse struct {
	RequestId    string
	TotalCount   int
	PageNumber   int
	PageSize     int
	JobInstances ListJobExecutionInstancesJobInstanceList
}

type ListJobExecutionInstancesJobInstance struct {
	Id        string
	JobName   string
	StartTime int64
	RunTime   int
	JobType   string
	JobId     string
	ClusterId string
	Status    string
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
