package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListJobExecutionInstanceTrendRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
}

func (req *ListJobExecutionInstanceTrendRequest) Invoke(client *sdk.Client) (resp *ListJobExecutionInstanceTrendResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListJobExecutionInstanceTrend", "", "")
	resp = &ListJobExecutionInstanceTrendResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListJobExecutionInstanceTrendResponse struct {
	responses.BaseResponse
	RequestId string
	Trends    ListJobExecutionInstanceTrendTrendList
}

type ListJobExecutionInstanceTrendTrend struct {
	Day    string
	Count  int
	Status string
}

type ListJobExecutionInstanceTrendTrendList []ListJobExecutionInstanceTrendTrend

func (list *ListJobExecutionInstanceTrendTrendList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobExecutionInstanceTrendTrend)
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
