package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetLoadBalancerStatusRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	LoadBalancerStatus   string `position:"Query" name:"LoadBalancerStatus"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (req *SetLoadBalancerStatusRequest) Invoke(client *sdk.Client) (resp *SetLoadBalancerStatusResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "SetLoadBalancerStatus", "slb", "")
	resp = &SetLoadBalancerStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetLoadBalancerStatusResponse struct {
	responses.BaseResponse
	RequestId common.String
}
