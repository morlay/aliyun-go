package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeAlertConfigRequest struct {
	requests.RpcRequest
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupId       string `position:"Query" name:"ScalingGroupId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeAlertConfigRequest) Invoke(client *sdk.Client) (resp *DescribeAlertConfigResponse, err error) {
	req.InitWithApiInfo("Ess", "2014-08-28", "DescribeAlertConfig", "ess", "")
	resp = &DescribeAlertConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeAlertConfigResponse struct {
	responses.BaseResponse
	SuccessConfig common.Integer
	FailConfig    common.Integer
	RejectConfig  common.Integer
	RequestId     common.String
}
