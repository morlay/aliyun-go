package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLimitationRequest struct {
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeLimitationRequest) Invoke(client *sdk.Client) (response *DescribeLimitationResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeLimitationRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ess", "2014-08-28", "DescribeLimitation", "ess", "")

	resp := struct {
		*responses.BaseResponse
		DescribeLimitationResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeLimitationResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeLimitationResponse struct {
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
