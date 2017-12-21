package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteScalingConfigurationRequest struct {
	ScalingConfigurationId string `position:"Query" name:"ScalingConfigurationId"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteScalingConfigurationRequest) Invoke(client *sdk.Client) (response *DeleteScalingConfigurationResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteScalingConfigurationRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ess", "2014-08-28", "DeleteScalingConfiguration", "ess", "")

	resp := struct {
		*responses.BaseResponse
		DeleteScalingConfigurationResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteScalingConfigurationResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteScalingConfigurationResponse struct {
	RequestId string
}
