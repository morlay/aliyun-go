package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId      common.String
	Success        common.String
	ErrCode        common.String
	ErrMsg         common.String
	TotalCount     common.Integer
	PageSize       common.Integer
	PageNumber     common.Integer
	HierarchyInfos ListJobExecutionPlanHierarchyHierarchyInfoList
}

type ListJobExecutionPlanHierarchyHierarchyInfo struct {
	Id                   common.Long
	NodeBizType          common.String
	NodeType             common.String
	RelateId             common.String
	Name                 common.String
	ParentId             common.String
	ResourceOwnerId      common.String
	UtcCreateTimestamp   common.Long
	UtcModifiedTimestamp common.Long
	NodeStatus           common.Integer
	ExecutionPlan        ListJobExecutionPlanHierarchyExecutionPlan
	Job                  ListJobExecutionPlanHierarchyJob
}

type ListJobExecutionPlanHierarchyExecutionPlan struct {
	BizId                common.String
	Name                 common.String
	Status               common.Integer
	LastExeStatus        common.Integer
	IsCreateCluster      bool
	IsInterruptWhenError bool
	IsCycle              bool
	ScheduleCycle        common.String
	RegionId             common.String
	CronExpr             common.String
	UtcCreateTimestamp   common.Long
	UtcModifiedTimestamp common.Long
	Version              common.String
	ClusterTemplateId    common.String
}

type ListJobExecutionPlanHierarchyJob struct {
	BizId         common.String
	Name          common.String
	JobFailAct    common.Integer
	JobType       common.Integer
	EnvParam      common.String
	MaxRetry      common.Integer
	RetryInterval common.Integer
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
