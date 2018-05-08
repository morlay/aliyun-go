package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AllocateEipAddressRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth            string `position:"Query" name:"Bandwidth"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	InternetChargeType   string `position:"Query" name:"InternetChargeType"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *AllocateEipAddressRequest) Invoke(client *sdk.Client) (resp *AllocateEipAddressResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "AllocateEipAddress", "ecs", "")
	resp = &AllocateEipAddressResponse{}
	err = client.DoAction(req, resp)
	return
}

type AllocateEipAddressResponse struct {
	responses.BaseResponse
	RequestId    common.String
	AllocationId common.String
	EipAddress   common.String
}
