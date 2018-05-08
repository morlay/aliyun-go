package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId     common.String
	Id            common.String
	GmtCreate     common.Long
	GmtModified   common.Long
	Name          common.String
	Description   common.String
	Type          common.String
	Status        common.String
	Periodic      bool
	StartSchedule common.Long
	EndSchedule   common.Long
	CronExpr      common.String
	CreateCluster bool
	ClusterId     common.String
	Graph         common.String
	CategoryId    common.String
}
