package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AllocateEipAddressRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Period               int    `position:"Query" name:"Period"`
	AutoPay              string `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth            string `position:"Query" name:"Bandwidth"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	ISP                  string `position:"Query" name:"ISP"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	InternetChargeType   string `position:"Query" name:"InternetChargeType"`
	Netmode              string `position:"Query" name:"Netmode"`
	PricingCycle         string `position:"Query" name:"PricingCycle"`
	InstanceChargeType   string `position:"Query" name:"InstanceChargeType"`
}

func (r AllocateEipAddressRequest) Invoke(client *sdk.Client) (response *AllocateEipAddressResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AllocateEipAddressRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "AllocateEipAddress", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		AllocateEipAddressResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AllocateEipAddressResponse

	err = client.DoAction(&req, &resp)
	return
}

type AllocateEipAddressResponse struct {
	RequestId    string
	AllocationId string
	EipAddress   string
	OrderId      int64
}
