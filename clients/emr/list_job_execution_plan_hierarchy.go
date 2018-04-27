package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListJobExecutionPlanHierarchyRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
	CurrentId       int64 `position:"Query" name:"CurrentId"`
	PageSize        int   `position:"Query" name:"PageSize"`
	PageNumber      int   `position:"Query" name:"PageNumber"`
}

func (req *ListJobExecutionPlanHierarchyRequest) Invoke(client *sdk.Client) (resp *ListJobExecutionPlanHierarchyResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListJobExecutionPlanHierarchy", "", "")
	resp = &ListJobExecutionPlanHierarchyResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListJobExecutionPlanHierarchyResponse struct {
	responses.BaseResponse
	RequestId      string
	Success        string
	ErrCode        string
	ErrMsg         string
	TotalCount     int
	PageSize       int
	PageNumber     int
	HierarchyInfos ListJobExecutionPlanHierarchyHierarchyInfoList
}

type ListJobExecutionPlanHierarchyHierarchyInfo struct {
	Id                   int64
	NodeBizType          string
	NodeType             string
	RelateId             string
	Name                 string
	ParentId             string
	ResourceOwnerId      string
	UtcCreateTimestamp   int64
	UtcModifiedTimestamp int64
	NodeStatus           int
	ExecutionPlan        ListJobExecutionPlanHierarchyExecutionPlan
	Job                  ListJobExecutionPlanHierarchyJob
}

type ListJobExecutionPlanHierarchyExecutionPlan struct {
	BizId                string
	Name                 string
	Status               int
	LastExeStatus        int
	IsCreateCluster      bool
	IsInterruptWhenError bool
	IsCycle              bool
	ScheduleCycle        string
	RegionId             string
	CronExpr             string
	UtcCreateTimestamp   int64
	UtcModifiedTimestamp int64
	Version              string
	ClusterTemplateId    string
}

type ListJobExecutionPlanHierarchyJob struct {
	BizId         string
	Name          string
	JobFailAct    int
	JobType       int
	EnvParam      string
	MaxRetry      int
	RetryInterval int
}

type ListJobExecutionPlanHierarchyHierarchyInfoList []ListJobExecutionPlanHierarchyHierarchyInfo

func (list *ListJobExecutionPlanHierarchyHierarchyInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobExecutionPlanHierarchyHierarchyInfo)
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
