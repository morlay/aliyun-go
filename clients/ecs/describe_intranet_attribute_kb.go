package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeIntranetAttributeKbRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeIntranetAttributeKbRequest) Invoke(client *sdk.Client) (response *DescribeIntranetAttributeKbResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeIntranetAttributeKbRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeIntranetAttributeKb", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribeIntranetAttributeKbResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeIntranetAttributeKbResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeIntranetAttributeKbResponse struct {
	RequestId               string
	InstanceId              string
	VlanId                  string
	IntranetIpAddress       string
	IntranetMaxBandwidthIn  int
	IntranetMaxBandwidthOut int
}
