package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateNetworkInterfaceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	SecurityGroupId      string `position:"Query" name:"SecurityGroupId"`
	Description          string `position:"Query" name:"Description"`
	NetworkInterfaceName string `position:"Query" name:"NetworkInterfaceName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	PrimaryIpAddress     string `position:"Query" name:"PrimaryIpAddress"`
}

func (r CreateNetworkInterfaceRequest) Invoke(client *sdk.Client) (response *CreateNetworkInterfaceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateNetworkInterfaceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "CreateNetworkInterface", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		CreateNetworkInterfaceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateNetworkInterfaceResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateNetworkInterfaceResponse struct {
	RequestId          string
	NetworkInterfaceId string
}
