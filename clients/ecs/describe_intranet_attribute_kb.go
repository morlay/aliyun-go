package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeIntranetAttributeKbRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeIntranetAttributeKbRequest) Invoke(client *sdk.Client) (resp *DescribeIntranetAttributeKbResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeIntranetAttributeKb", "ecs", "")
	resp = &DescribeIntranetAttributeKbResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeIntranetAttributeKbResponse struct {
	responses.BaseResponse
	RequestId               common.String
	InstanceId              common.String
	VlanId                  common.String
	IntranetIpAddress       common.String
	IntranetMaxBandwidthIn  common.Integer
	IntranetMaxBandwidthOut common.Integer
}
