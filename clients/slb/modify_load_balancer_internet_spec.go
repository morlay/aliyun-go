package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyLoadBalancerInternetSpecRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	AutoPay              string `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth            int    `position:"Query" name:"Bandwidth"`
	InternetChargeType   string `position:"Query" name:"InternetChargeType"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (req *ModifyLoadBalancerInternetSpecRequest) Invoke(client *sdk.Client) (resp *ModifyLoadBalancerInternetSpecResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "ModifyLoadBalancerInternetSpec", "slb", "")
	resp = &ModifyLoadBalancerInternetSpecResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyLoadBalancerInternetSpecResponse struct {
	responses.BaseResponse
	RequestId string
	OrderId   int64
}
