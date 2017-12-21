package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AllocateEipAddressRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth            string `position:"Query" name:"Bandwidth"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	InternetChargeType   string `position:"Query" name:"InternetChargeType"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r AllocateEipAddressRequest) Invoke(client *sdk.Client) (response *AllocateEipAddressResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AllocateEipAddressRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "AllocateEipAddress", "ecs", "")

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
}
