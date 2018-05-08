package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	MaxNumberOfScalingGroups         common.Integer
	MaxNumberOfScalingConfigurations common.Integer
	MaxNumberOfScalingRules          common.Integer
	MaxNumberOfScheduledTasks        common.Integer
	MaxNumberOfScalingInstances      common.Integer
	MaxNumberOfDBInstances           common.Integer
	MaxNumberOfLoadBalancers         common.Integer
	MaxNumberOfMinSize               common.Integer
	MaxNumberOfMaxSize               common.Integer
}
