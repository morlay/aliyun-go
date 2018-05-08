package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeInstancePhysicalAttributeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeInstancePhysicalAttributeRequest) Invoke(client *sdk.Client) (resp *DescribeInstancePhysicalAttributeResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInstancePhysicalAttribute", "ecs", "")
	resp = &DescribeInstancePhysicalAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeInstancePhysicalAttributeResponse struct {
	responses.BaseResponse
	RequestId        common.String
	InstanceId       common.String
	VlanId           common.String
	NodeControllerId common.String
	RackId           common.String
}
