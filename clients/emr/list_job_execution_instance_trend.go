package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	Trends    ListJobExecutionInstanceTrendTrendList
}

type ListJobExecutionInstanceTrendTrend struct {
	Day    common.String
	Count  common.Integer
	Status common.String
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
