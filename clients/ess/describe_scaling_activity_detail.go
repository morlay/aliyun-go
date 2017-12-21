package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeScalingActivityDetailRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ScalingActivityId    string `position:"Query" name:"ScalingActivityId"`
}

func (r DescribeScalingActivityDetailRequest) Invoke(client *sdk.Client) (response *DescribeScalingActivityDetailResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeScalingActivityDetailRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ess", "2014-08-28", "DescribeScalingActivityDetail", "ess", "")

	resp := struct {
		*responses.BaseResponse
		DescribeScalingActivityDetailResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeScalingActivityDetailResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeScalingActivityDetailResponse struct {
	ScalingActivityId string
	Detail            string
}
