package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RebalanceInstancesRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupId       string `position:"Query" name:"ScalingGroupId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *RebalanceInstancesRequest) Invoke(client *sdk.Client) (resp *RebalanceInstancesResponse, err error) {
	req.InitWithApiInfo("Ess", "2014-08-28", "RebalanceInstances", "ess", "")
	resp = &RebalanceInstancesResponse{}
	err = client.DoAction(req, resp)
	return
}

type RebalanceInstancesResponse struct {
	responses.BaseResponse
	ScalingActivityId string
	RequestId         string
}
