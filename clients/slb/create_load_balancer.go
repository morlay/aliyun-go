package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateLoadBalancerRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	MasterZoneId         string `position:"Query" name:"MasterZoneId"`
	Duration             int    `position:"Query" name:"Duration"`
	ResourceGroupId      string `position:"Query" name:"ResourceGroupId"`
	LoadBalancerName     string `position:"Query" name:"LoadBalancerName"`
	AddressType          string `position:"Query" name:"AddressType"`
	SlaveZoneId          string `position:"Query" name:"SlaveZoneId"`
	LoadBalancerSpec     string `position:"Query" name:"LoadBalancerSpec"`
	AutoPay              string `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth            int    `position:"Query" name:"Bandwidth"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	EnableVpcVipFlow     string `position:"Query" name:"EnableVpcVipFlow"`
	InternetChargeType   string `position:"Query" name:"InternetChargeType"`
	VpcId                string `position:"Query" name:"VpcId"`
	PayType              string `position:"Query" name:"PayType"`
	PricingCycle         string `position:"Query" name:"PricingCycle"`
}

func (req *CreateLoadBalancerRequest) Invoke(client *sdk.Client) (resp *CreateLoadBalancerResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "CreateLoadBalancer", "slb", "")
	resp = &CreateLoadBalancerResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateLoadBalancerResponse struct {
	responses.BaseResponse
	RequestId        string
	LoadBalancerId   string
	ResourceGroupId  string
	Address          string
	LoadBalancerName string
	VpcId            string
	VSwitchId        string
	NetworkType      string
	OrderId          int64
}
