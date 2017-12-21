package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLoadBalancerAutoReleaseTimeRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (req *DescribeLoadBalancerAutoReleaseTimeRequest) Invoke(client *sdk.Client) (resp *DescribeLoadBalancerAutoReleaseTimeResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeLoadBalancerAutoReleaseTime", "slb", "")
	resp = &DescribeLoadBalancerAutoReleaseTimeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLoadBalancerAutoReleaseTimeResponse struct {
	responses.BaseResponse
	RequestId       string
	AutoReleaseTime int64
}
