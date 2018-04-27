package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListExecutionPlanInstanceTrendRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
}

func (req *ListExecutionPlanInstanceTrendRequest) Invoke(client *sdk.Client) (resp *ListExecutionPlanInstanceTrendResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListExecutionPlanInstanceTrend", "", "")
	resp = &ListExecutionPlanInstanceTrendResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListExecutionPlanInstanceTrendResponse struct {
	responses.BaseResponse
	RequestId string
	Trends    ListExecutionPlanInstanceTrendTrendList
}

type ListExecutionPlanInstanceTrendTrend struct {
	Day    string
	Count  int
	Status string
}

type ListExecutionPlanInstanceTrendTrendList []ListExecutionPlanInstanceTrendTrend

func (list *ListExecutionPlanInstanceTrendTrendList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListExecutionPlanInstanceTrendTrend)
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
