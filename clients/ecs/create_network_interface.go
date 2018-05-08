package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateNetworkInterfaceRequest struct {
	requests.RpcRequest
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

func (req *CreateNetworkInterfaceRequest) Invoke(client *sdk.Client) (resp *CreateNetworkInterfaceResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "CreateNetworkInterface", "ecs", "")
	resp = &CreateNetworkInterfaceResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateNetworkInterfaceResponse struct {
	responses.BaseResponse
	RequestId          common.String
	NetworkInterfaceId common.String
}
