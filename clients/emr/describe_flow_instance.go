package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId    common.String
	Id           common.String
	GmtCreate    common.Long
	GmtModified  common.Long
	FlowId       common.String
	FlowName     common.String
	ProjectId    common.String
	Status       common.String
	ClusterId    common.String
	StartTime    common.Long
	EndTime      common.Long
	NodeInstance DescribeFlowInstanceNodeInstanceItemList
}

type DescribeFlowInstanceNodeInstanceItem struct {
	Id             common.String
	GmtCreate      common.Long
	GmtModified    common.Long
	Type           common.String
	Status         common.String
	JobId          common.String
	JobName        common.String
	JobType        common.String
	FailAct        common.String
	MaxRetry       common.String
	RetryInterval  common.String
	NodeName       common.String
	ClusterId      common.String
	HostName       common.String
	ProjectId      common.String
	StartTime      common.Long
	EndTime        common.Long
	Retries        common.Integer
	ExternalId     common.String
	ExternalStatus common.String
	ExternalInfo   common.String
	ParamConf      common.String
	EnvConf        common.String
	RunConf        common.String
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
