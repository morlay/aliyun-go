package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyLoadBalancerInternetSpecRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AutoPay              string `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth            int    `position:"Query" name:"Bandwidth"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	InternetChargeType   string `position:"Query" name:"InternetChargeType"`
}

func (r ModifyLoadBalancerInternetSpecRequest) Invoke(client *sdk.Client) (response *ModifyLoadBalancerInternetSpecResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyLoadBalancerInternetSpecRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "ModifyLoadBalancerInternetSpec", "slb", "")

	resp := struct {
		*responses.BaseResponse
		ModifyLoadBalancerInternetSpecResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyLoadBalancerInternetSpecResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyLoadBalancerInternetSpecResponse struct {
	RequestId string
	OrderId   int64
}
