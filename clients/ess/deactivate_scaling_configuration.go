package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeactivateScalingConfigurationRequest struct {
	ScalingConfigurationId string `position:"Query" name:"ScalingConfigurationId"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
}

func (r DeactivateScalingConfigurationRequest) Invoke(client *sdk.Client) (response *DeactivateScalingConfigurationResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeactivateScalingConfigurationRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ess", "2014-08-28", "DeactivateScalingConfiguration", "ess", "")

	resp := struct {
		*responses.BaseResponse
		DeactivateScalingConfigurationResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeactivateScalingConfigurationResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeactivateScalingConfigurationResponse struct {
	RequestId string
}
