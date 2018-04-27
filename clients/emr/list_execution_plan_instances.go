package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListExecutionPlanInstancesRequest struct {
	requests.RpcRequest
	OnlyLastInstance     string                                             `position:"Query" name:"OnlyLastInstance"`
	ResourceOwnerId      int64                                              `position:"Query" name:"ResourceOwnerId"`
	ExecutionPlanIdLists *ListExecutionPlanInstancesExecutionPlanIdListList `position:"Query" type:"Repeated" name:"ExecutionPlanIdList"`
	StatusLists          *ListExecutionPlanInstancesStatusListList          `position:"Query" type:"Repeated" name:"StatusList"`
	PageSize             int                                                `position:"Query" name:"PageSize"`
	IsDesc               string                                             `position:"Query" name:"IsDesc"`
	PageNumber           int                                                `position:"Query" name:"PageNumber"`
}

func (req *ListExecutionPlanInstancesRequest) Invoke(client *sdk.Client) (resp *ListExecutionPlanInstancesResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListExecutionPlanInstances", "", "")
	resp = &ListExecutionPlanInstancesResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListExecutionPlanInstancesResponse struct {
	responses.BaseResponse
	RequestId              string
	TotalCount             int
	PageNumber             int
	PageSize               int
	ExecutionPlanInstances ListExecutionPlanInstancesExecutionPlanInstanceList
}

type ListExecutionPlanInstancesExecutionPlanInstance struct {
	Id                string
	ExecutionPlanId   string
	ExecutionPlanName string
	StartTime         int64
	RunTime           int
	ClusterId         string
	ClusterName       string
	ClusterType       string
	Status            string
	LogEnable         bool
	LogPath           string
	WorkflowApp       string
}

type ListExecutionPlanInstancesExecutionPlanIdListList []string

func (list *ListExecutionPlanInstancesExecutionPlanIdListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type ListExecutionPlanInstancesStatusListList []string

func (list *ListExecutionPlanInstancesStatusListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type ListExecutionPlanInstancesExecutionPlanInstanceList []ListExecutionPlanInstancesExecutionPlanInstance

func (list *ListExecutionPlanInstancesExecutionPlanInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListExecutionPlanInstancesExecutionPlanInstance)
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
