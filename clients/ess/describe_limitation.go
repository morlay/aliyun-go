package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLimitationRequest struct {
	requests.RpcRequest
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeLimitationRequest) Invoke(client *sdk.Client) (resp *DescribeLimitationResponse, err error) {
	req.InitWithApiInfo("Ess", "2014-08-28", "DescribeLimitation", "ess", "")
	resp = &DescribeLimitationResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLimitationResponse struct {
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
