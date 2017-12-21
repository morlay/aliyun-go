package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyLoadBalancerPayTypeRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AutoPay              string `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
	Duration             int    `position:"Query" name:"Duration"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	PayType              string `position:"Query" name:"PayType"`
	PricingCycle         string `position:"Query" name:"PricingCycle"`
}

func (r ModifyLoadBalancerPayTypeRequest) Invoke(client *sdk.Client) (response *ModifyLoadBalancerPayTypeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyLoadBalancerPayTypeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "ModifyLoadBalancerPayType", "slb", "")

	resp := struct {
		*responses.BaseResponse
		ModifyLoadBalancerPayTypeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyLoadBalancerPayTypeResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyLoadBalancerPayTypeResponse struct {
	RequestId string
	OrderId   int64
}
