package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetLoadBalancerNameRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerName     string `position:"Query" name:"LoadBalancerName"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (req *SetLoadBalancerNameRequest) Invoke(client *sdk.Client) (resp *SetLoadBalancerNameResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "SetLoadBalancerName", "slb", "")
	resp = &SetLoadBalancerNameResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetLoadBalancerNameResponse struct {
	responses.BaseResponse
	RequestId string
}
