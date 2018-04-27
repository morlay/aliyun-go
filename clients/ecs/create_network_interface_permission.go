package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateNetworkInterfacePermissionRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountId            int64  `position:"Query" name:"AccountId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Permission           string `position:"Query" name:"Permission"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	NetworkInterfaceId   string `position:"Query" name:"NetworkInterfaceId"`
}

func (req *CreateNetworkInterfacePermissionRequest) Invoke(client *sdk.Client) (resp *CreateNetworkInterfacePermissionResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "CreateNetworkInterfacePermission", "ecs", "")
	resp = &CreateNetworkInterfacePermissionResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateNetworkInterfacePermissionResponse struct {
	responses.BaseResponse
	RequestId                  string
	NetworkInterfacePermission CreateNetworkInterfacePermissionNetworkInterfacePermission
}

type CreateNetworkInterfacePermissionNetworkInterfacePermission struct {
	AccountId                    int64
	ServiceName                  string
	NetworkInterfaceId           string
	NetworkInterfacePermissionId string
	Permission                   string
	PermissionState              string
}
