package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeFlowInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	ProjectId       string `position:"Query" name:"ProjectId"`
}

func (req *DescribeFlowInstanceRequest) Invoke(client *sdk.Client) (resp *DescribeFlowInstanceResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DescribeFlowInstance", "", "")
	resp = &DescribeFlowInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeFlowInstanceResponse struct {
	responses.BaseResponse
	RequestId    string
	Id           string
	GmtCreate    int64
	GmtModified  int64
	FlowId       string
	FlowName     string
	ProjectId    string
	Status       string
	ClusterId    string
	StartTime    int64
	EndTime      int64
	NodeInstance DescribeFlowInstanceNodeInstanceItemList
}

type DescribeFlowInstanceNodeInstanceItem struct {
	Id             string
	GmtCreate      int64
	GmtModified    int64
	Type           string
	Status         string
	JobId          string
	JobName        string
	JobType        string
	FailAct        string
	MaxRetry       string
	RetryInterval  string
	NodeName       string
	ClusterId      string
	HostName       string
	ProjectId      string
	StartTime      int64
	EndTime        int64
	Retries        int
	ExternalId     string
	ExternalStatus string
	ExternalInfo   string
	ParamConf      string
	EnvConf        string
	RunConf        string
}

type DescribeFlowInstanceNodeInstanceItemList []DescribeFlowInstanceNodeInstanceItem

func (list *DescribeFlowInstanceNodeInstanceItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeFlowInstanceNodeInstanceItem)
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
