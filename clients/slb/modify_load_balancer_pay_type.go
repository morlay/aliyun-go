package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyLoadBalancerPayTypeRequest struct {
	requests.RpcRequest
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

func (req *ModifyLoadBalancerPayTypeRequest) Invoke(client *sdk.Client) (resp *ModifyLoadBalancerPayTypeResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "ModifyLoadBalancerPayType", "slb", "")
	resp = &ModifyLoadBalancerPayTypeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyLoadBalancerPayTypeResponse struct {
	responses.BaseResponse
	RequestId common.String
	OrderId   common.Long
}
