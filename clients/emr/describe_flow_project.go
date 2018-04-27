package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId   string
	Id          string
	GmtCreate   int64
	GmtModified int64
	UserId      string
	Name        string
	Description string
}
