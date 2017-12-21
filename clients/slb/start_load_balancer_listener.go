package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type StartLoadBalancerListenerRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ListenerPort         int    `position:"Query" name:"ListenerPort"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (r StartLoadBalancerListenerRequest) Invoke(client *sdk.Client) (response *StartLoadBalancerListenerResponse, err error) {
	req := struct {
		*requests.RpcRequest
		StartLoadBalancerListenerRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "StartLoadBalancerListener", "slb", "")

	resp := struct {
		*responses.BaseResponse
		StartLoadBalancerListenerResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.StartLoadBalancerListenerResponse

	err = client.DoAction(&req, &resp)
	return
}

type StartLoadBalancerListenerResponse struct {
	RequestId string
}
