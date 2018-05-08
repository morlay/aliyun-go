package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeFlowNodeInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	ProjectId       string `position:"Query" name:"ProjectId"`
}

func (req *DescribeFlowNodeInstanceRequest) Invoke(client *sdk.Client) (resp *DescribeFlowNodeInstanceResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DescribeFlowNodeInstance", "", "")
	resp = &DescribeFlowNodeInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeFlowNodeInstanceResponse struct {
	responses.BaseResponse
	RequestId      common.String
	Id             common.String
	GmtCreate      common.Long
	GmtModified    common.Long
	Type           common.String
	Status         common.String
	JobId          common.String
	JobName        common.String
	JobType        common.String
	JobParams      common.String
	FailAct        common.String
	MaxRetry       common.String
	RetryInterval  common.String
	NodeName       common.String
	FlowId         common.String
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
	ClusterName    common.String
}
