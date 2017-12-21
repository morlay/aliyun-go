package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AllocatePublicIpAddressRequest struct {
	requests.RpcRequest
	IpAddress            string `position:"Query" name:"IpAddress"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VlanId               string `position:"Query" name:"VlanId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *AllocatePublicIpAddressRequest) Invoke(client *sdk.Client) (resp *AllocatePublicIpAddressResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "AllocatePublicIpAddress", "ecs", "")
	resp = &AllocatePublicIpAddressResponse{}
	err = client.DoAction(req, resp)
	return
}

type AllocatePublicIpAddressResponse struct {
	responses.BaseResponse
	RequestId string
	IpAddress string
}
