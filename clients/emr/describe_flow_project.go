package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeFlowProjectRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ProjectId       string `position:"Query" name:"ProjectId"`
}

func (req *DescribeFlowProjectRequest) Invoke(client *sdk.Client) (resp *DescribeFlowProjectResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DescribeFlowProject", "", "")
	resp = &DescribeFlowProjectResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeFlowProjectResponse struct {
	responses.BaseResponse
	RequestId   common.String
	Id          common.String
	GmtCreate   common.Long
	GmtModified common.Long
	UserId      common.String
	Name        common.String
	Description common.String
}
