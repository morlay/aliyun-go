package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AttachNetworkInterfaceRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	NetworkInterfaceId   string `position:"Query" name:"NetworkInterfaceId"`
}

func (req *AttachNetworkInterfaceRequest) Invoke(client *sdk.Client) (resp *AttachNetworkInterfaceResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "AttachNetworkInterface", "ecs", "")
	resp = &AttachNetworkInterfaceResponse{}
	err = client.DoAction(req, resp)
	return
}

type AttachNetworkInterfaceResponse struct {
	responses.BaseResponse
	RequestId common.String
}
