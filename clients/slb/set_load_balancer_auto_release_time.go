package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetLoadBalancerAutoReleaseTimeRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	AutoReleaseTime      int64  `position:"Query" name:"AutoReleaseTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (req *SetLoadBalancerAutoReleaseTimeRequest) Invoke(client *sdk.Client) (resp *SetLoadBalancerAutoReleaseTimeResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "SetLoadBalancerAutoReleaseTime", "slb", "")
	resp = &SetLoadBalancerAutoReleaseTimeResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetLoadBalancerAutoReleaseTimeResponse struct {
	responses.BaseResponse
	RequestId common.String
}
