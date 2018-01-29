package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReActivateInstancesRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ReActivateInstancesRequest) Invoke(client *sdk.Client) (resp *ReActivateInstancesResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ReActivateInstances", "ecs", "")
	resp = &ReActivateInstancesResponse{}
	err = client.DoAction(req, resp)
	return
}

type ReActivateInstancesResponse struct {
	responses.BaseResponse
	RequestId string
}
