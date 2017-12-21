package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type StopLoadBalancerListenerRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ListenerPort         int    `position:"Query" name:"ListenerPort"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (r StopLoadBalancerListenerRequest) Invoke(client *sdk.Client) (response *StopLoadBalancerListenerResponse, err error) {
	req := struct {
		*requests.RpcRequest
		StopLoadBalancerListenerRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "StopLoadBalancerListener", "slb", "")

	resp := struct {
		*responses.BaseResponse
		StopLoadBalancerListenerResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.StopLoadBalancerListenerResponse

	err = client.DoAction(&req, &resp)
	return
}

type StopLoadBalancerListenerResponse struct {
	RequestId string
}
