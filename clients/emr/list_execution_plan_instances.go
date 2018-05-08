package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId              common.String
	TotalCount             common.Integer
	PageNumber             common.Integer
	PageSize               common.Integer
	ExecutionPlanInstances ListExecutionPlanInstancesExecutionPlanInstanceList
}

type ListExecutionPlanInstancesExecutionPlanInstance struct {
	Id                common.String
	ExecutionPlanId   common.String
	ExecutionPlanName common.String
	StartTime         common.Long
	RunTime           common.Integer
	ClusterId         common.String
	ClusterName       common.String
	ClusterType       common.String
	Status            common.String
	LogEnable         bool
	LogPath           common.String
	WorkflowApp       common.String
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
