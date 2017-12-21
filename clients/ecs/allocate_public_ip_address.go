package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AllocatePublicIpAddressRequest struct {
	IpAddress            string `position:"Query" name:"IpAddress"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VlanId               string `position:"Query" name:"VlanId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r AllocatePublicIpAddressRequest) Invoke(client *sdk.Client) (response *AllocatePublicIpAddressResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AllocatePublicIpAddressRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "AllocatePublicIpAddress", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		AllocatePublicIpAddressResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AllocatePublicIpAddressResponse

	err = client.DoAction(&req, &resp)
	return
}

type AllocatePublicIpAddressResponse struct {
	RequestId string
	IpAddress string
}
