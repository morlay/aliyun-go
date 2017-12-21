package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyLoadBalancerInstanceSpecRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	LoadBalancerSpec     string `position:"Query" name:"LoadBalancerSpec"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	AutoPay              string `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (req *ModifyLoadBalancerInstanceSpecRequest) Invoke(client *sdk.Client) (resp *ModifyLoadBalancerInstanceSpecResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "ModifyLoadBalancerInstanceSpec", "slb", "")
	resp = &ModifyLoadBalancerInstanceSpecResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyLoadBalancerInstanceSpecResponse struct {
	responses.BaseResponse
	RequestId string
	OrderId   int64
}
