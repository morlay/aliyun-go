package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeFlowRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	ProjectId       string `position:"Query" name:"ProjectId"`
}

func (req *DescribeFlowRequest) Invoke(client *sdk.Client) (resp *DescribeFlowResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DescribeFlow", "", "")
	resp = &DescribeFlowResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeFlowResponse struct {
	responses.BaseResponse
	RequestId     string
	Id            string
	GmtCreate     int64
	GmtModified   int64
	Name          string
	Description   string
	Type          string
	Status        string
	Periodic      bool
	StartSchedule int64
	EndSchedule   int64
	CronExpr      string
	CreateCluster bool
	ClusterId     string
	Graph         string
	CategoryId    string
}
