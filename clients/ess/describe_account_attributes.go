package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeAccountAttributesRequest struct {
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeAccountAttributesRequest) Invoke(client *sdk.Client) (response *DescribeAccountAttributesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeAccountAttributesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ess", "2014-08-28", "DescribeAccountAttributes", "ess", "")

	resp := struct {
		*responses.BaseResponse
		DescribeAccountAttributesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeAccountAttributesResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeAccountAttributesResponse struct {
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
