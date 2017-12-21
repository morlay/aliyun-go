package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteNetworkInterfaceRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	NetworkInterfaceId   string `position:"Query" name:"NetworkInterfaceId"`
}

func (req *DeleteNetworkInterfaceRequest) Invoke(client *sdk.Client) (resp *DeleteNetworkInterfaceResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DeleteNetworkInterface", "ecs", "")
	resp = &DeleteNetworkInterfaceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteNetworkInterfaceResponse struct {
	responses.BaseResponse
	RequestId string
}
