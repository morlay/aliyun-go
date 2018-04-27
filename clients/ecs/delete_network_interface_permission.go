package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteNetworkInterfacePermissionRequest struct {
	requests.RpcRequest
	ResourceOwnerId              int64  `position:"Query" name:"ResourceOwnerId"`
	NetworkInterfacePermissionId string `position:"Query" name:"NetworkInterfacePermissionId"`
	ResourceOwnerAccount         string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                 string `position:"Query" name:"OwnerAccount"`
	OwnerId                      int64  `position:"Query" name:"OwnerId"`
	Force                        string `position:"Query" name:"Force"`
}

func (req *DeleteNetworkInterfacePermissionRequest) Invoke(client *sdk.Client) (resp *DeleteNetworkInterfacePermissionResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DeleteNetworkInterfacePermission", "ecs", "")
	resp = &DeleteNetworkInterfacePermissionResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteNetworkInterfacePermissionResponse struct {
	responses.BaseResponse
	RequestId string
}
