package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeAccountAttributesRequest struct {
	requests.RpcRequest
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeAccountAttributesRequest) Invoke(client *sdk.Client) (resp *DescribeAccountAttributesResponse, err error) {
	req.InitWithApiInfo("Ess", "2014-08-28", "DescribeAccountAttributes", "ess", "")
	resp = &DescribeAccountAttributesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeAccountAttributesResponse struct {
	responses.BaseResponse
	MaxNumberOfScalingGroups         int
	MaxNumberOfScalingConfigurations int
	MaxNumberOfScalingRules          int
	MaxNumberOfScheduledTasks        int
	MaxNumberOfScalingInstances      int
	MaxNumberOfDBInstances           int
	MaxNumberOfLoadBalancers         int
	MaxNumberOfMinSize               int
	MaxNumberOfMaxSize               int
}
