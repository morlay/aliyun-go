package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeactivateScalingConfigurationRequest struct {
	requests.RpcRequest
	ScalingConfigurationId string `position:"Query" name:"ScalingConfigurationId"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
}

func (req *DeactivateScalingConfigurationRequest) Invoke(client *sdk.Client) (resp *DeactivateScalingConfigurationResponse, err error) {
	req.InitWithApiInfo("Ess", "2014-08-28", "DeactivateScalingConfiguration", "ess", "")
	resp = &DeactivateScalingConfigurationResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeactivateScalingConfigurationResponse struct {
	responses.BaseResponse
	RequestId string
}
