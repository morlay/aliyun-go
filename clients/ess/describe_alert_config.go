package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeAlertConfigRequest struct {
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupId       string `position:"Query" name:"ScalingGroupId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeAlertConfigRequest) Invoke(client *sdk.Client) (response *DescribeAlertConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeAlertConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ess", "2014-08-28", "DescribeAlertConfig", "ess", "")

	resp := struct {
		*responses.BaseResponse
		DescribeAlertConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeAlertConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeAlertConfigResponse struct {
	SuccessConfig int
	FailConfig    int
	RejectConfig  int
	RequestId     string
}
