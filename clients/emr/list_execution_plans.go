package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListExecutionPlansRequest struct {
	requests.RpcRequest
	JobId           string                            `position:"Query" name:"JobId"`
	ResourceOwnerId int64                             `position:"Query" name:"ResourceOwnerId"`
	StatusLists     *ListExecutionPlansStatusListList `position:"Query" type:"Repeated" name:"StatusList"`
	PageSize        int                               `position:"Query" name:"PageSize"`
	QueryString     string                            `position:"Query" name:"QueryString"`
	ClusterId       string                            `position:"Query" name:"ClusterId"`
	IsDesc          string                            `position:"Query" name:"IsDesc"`
	Strategy        string                            `position:"Query" name:"Strategy"`
	PageNumber      int                               `position:"Query" name:"PageNumber"`
	QueryType       string                            `position:"Query" name:"QueryType"`
}

func (req *ListExecutionPlansRequest) Invoke(client *sdk.Client) (resp *ListExecutionPlansResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListExecutionPlans", "", "")
	resp = &ListExecutionPlansResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListExecutionPlansResponse struct {
	responses.BaseResponse
	RequestId      common.String
	TotalCount     common.Integer
	PageNumber     common.Integer
	PageSize       common.Integer
	ExecutionPlans ListExecutionPlansExecutionPlanInfoList
}

type ListExecutionPlansExecutionPlanInfo struct {
	Id                    common.String
	Name                  common.String
	CreateClusterOnDemand bool
	Stragety              common.String
	Status                common.String
	TimeInterval          common.Integer
	StartTime             common.Long
	TimeUnit              common.String
}

type ListExecutionPlansStatusListList []string

func (list *ListExecutionPlansStatusListList) UnmarshalJSON(data []byte) error {
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

type ListExecutionPlansExecutionPlanInfoList []ListExecutionPlansExecutionPlanInfo

func (list *ListExecutionPlansExecutionPlanInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListExecutionPlansExecutionPlanInfo)
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
