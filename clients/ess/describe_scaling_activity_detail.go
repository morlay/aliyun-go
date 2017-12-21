package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeScalingActivityDetailRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ScalingActivityId    string `position:"Query" name:"ScalingActivityId"`
}

func (req *DescribeScalingActivityDetailRequest) Invoke(client *sdk.Client) (resp *DescribeScalingActivityDetailResponse, err error) {
	req.InitWithApiInfo("Ess", "2014-08-28", "DescribeScalingActivityDetail", "ess", "")
	resp = &DescribeScalingActivityDetailResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeScalingActivityDetailResponse struct {
	responses.BaseResponse
	ScalingActivityId string
	Detail            string
}
