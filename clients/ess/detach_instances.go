package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DetachInstancesRequest struct {
	requests.RpcRequest
	InstanceId10         string `position:"Query" name:"InstanceId.10"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId12         string `position:"Query" name:"InstanceId.12"`
	InstanceId11         string `position:"Query" name:"InstanceId.11"`
	ScalingGroupId       string `position:"Query" name:"ScalingGroupId"`
	InstanceId20         string `position:"Query" name:"InstanceId.20"`
	InstanceId1          string `position:"Query" name:"InstanceId.1"`
	InstanceId3          string `position:"Query" name:"InstanceId.3"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	InstanceId2          string `position:"Query" name:"InstanceId.2"`
	InstanceId5          string `position:"Query" name:"InstanceId.5"`
	InstanceId4          string `position:"Query" name:"InstanceId.4"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	InstanceId7          string `position:"Query" name:"InstanceId.7"`
	InstanceId6          string `position:"Query" name:"InstanceId.6"`
	InstanceId9          string `position:"Query" name:"InstanceId.9"`
	InstanceId8          string `position:"Query" name:"InstanceId.8"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	InstanceId18         string `position:"Query" name:"InstanceId.18"`
	InstanceId17         string `position:"Query" name:"InstanceId.17"`
	InstanceId19         string `position:"Query" name:"InstanceId.19"`
	InstanceId14         string `position:"Query" name:"InstanceId.14"`
	InstanceId13         string `position:"Query" name:"InstanceId.13"`
	InstanceId16         string `position:"Query" name:"InstanceId.16"`
	InstanceId15         string `position:"Query" name:"InstanceId.15"`
}

func (req *DetachInstancesRequest) Invoke(client *sdk.Client) (resp *DetachInstancesResponse, err error) {
	req.InitWithApiInfo("Ess", "2014-08-28", "DetachInstances", "ess", "")
	resp = &DetachInstancesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DetachInstancesResponse struct {
	responses.BaseResponse
	ScalingActivityId string
	RequestId         string
}
