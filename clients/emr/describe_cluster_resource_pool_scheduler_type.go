package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeClusterResourcePoolSchedulerTypeRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
}

func (req *DescribeClusterResourcePoolSchedulerTypeRequest) Invoke(client *sdk.Client) (resp *DescribeClusterResourcePoolSchedulerTypeResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DescribeClusterResourcePoolSchedulerType", "", "")
	resp = &DescribeClusterResourcePoolSchedulerTypeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeClusterResourcePoolSchedulerTypeResponse struct {
	responses.BaseResponse
	RequestId            common.String
	CurrentSchedulerType common.String
	SupportSchedulerType common.String
}
