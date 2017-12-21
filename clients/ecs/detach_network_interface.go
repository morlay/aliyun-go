package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DetachNetworkInterfaceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	NetworkInterfaceId   string `position:"Query" name:"NetworkInterfaceId"`
}

func (r DetachNetworkInterfaceRequest) Invoke(client *sdk.Client) (response *DetachNetworkInterfaceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DetachNetworkInterfaceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DetachNetworkInterface", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DetachNetworkInterfaceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DetachNetworkInterfaceResponse

	err = client.DoAction(&req, &resp)
	return
}

type DetachNetworkInterfaceResponse struct {
	RequestId string
}
