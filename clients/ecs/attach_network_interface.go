package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AttachNetworkInterfaceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	NetworkInterfaceId   string `position:"Query" name:"NetworkInterfaceId"`
}

func (r AttachNetworkInterfaceRequest) Invoke(client *sdk.Client) (response *AttachNetworkInterfaceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AttachNetworkInterfaceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "AttachNetworkInterface", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		AttachNetworkInterfaceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AttachNetworkInterfaceResponse

	err = client.DoAction(&req, &resp)
	return
}

type AttachNetworkInterfaceResponse struct {
	RequestId string
}
