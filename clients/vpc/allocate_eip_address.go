package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AllocateEipAddressRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Period               int    `position:"Query" name:"Period"`
	AutoPay              string `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth            string `position:"Query" name:"Bandwidth"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	ISP                  string `position:"Query" name:"ISP"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceGroupId      string `position:"Query" name:"ResourceGroupId"`
	InternetChargeType   string `position:"Query" name:"InternetChargeType"`
	Netmode              string `position:"Query" name:"Netmode"`
	PricingCycle         string `position:"Query" name:"PricingCycle"`
	InstanceChargeType   string `position:"Query" name:"InstanceChargeType"`
}

func (req *AllocateEipAddressRequest) Invoke(client *sdk.Client) (resp *AllocateEipAddressResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "AllocateEipAddress", "vpc", "")
	resp = &AllocateEipAddressResponse{}
	err = client.DoAction(req, resp)
	return
}

type AllocateEipAddressResponse struct {
	responses.BaseResponse
	RequestId       string
	AllocationId    string
	EipAddress      string
	OrderId         int64
	ResourceGroupId string
}
