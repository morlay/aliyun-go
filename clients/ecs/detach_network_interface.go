package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DetachNetworkInterfaceRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	NetworkInterfaceId   string `position:"Query" name:"NetworkInterfaceId"`
}

func (req *DetachNetworkInterfaceRequest) Invoke(client *sdk.Client) (resp *DetachNetworkInterfaceResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DetachNetworkInterface", "ecs", "")
	resp = &DetachNetworkInterfaceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DetachNetworkInterfaceResponse struct {
	responses.BaseResponse
	RequestId common.String
}
