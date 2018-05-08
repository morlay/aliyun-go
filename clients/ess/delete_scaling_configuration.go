package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteScalingConfigurationRequest struct {
	requests.RpcRequest
	ScalingConfigurationId string `position:"Query" name:"ScalingConfigurationId"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteScalingConfigurationRequest) Invoke(client *sdk.Client) (resp *DeleteScalingConfigurationResponse, err error) {
	req.InitWithApiInfo("Ess", "2014-08-28", "DeleteScalingConfiguration", "ess", "")
	resp = &DeleteScalingConfigurationResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteScalingConfigurationResponse struct {
	responses.BaseResponse
	RequestId common.String
}
