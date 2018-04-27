package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId      string
	Id             string
	GmtCreate      int64
	GmtModified    int64
	Type           string
	Status         string
	JobId          string
	JobName        string
	JobType        string
	JobParams      string
	FailAct        string
	MaxRetry       string
	RetryInterval  string
	NodeName       string
	FlowId         string
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
	ClusterName    string
}
