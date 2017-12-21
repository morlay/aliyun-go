package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteLoadBalancerListenerRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ListenerPort         int    `position:"Query" name:"ListenerPort"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (r DeleteLoadBalancerListenerRequest) Invoke(client *sdk.Client) (response *DeleteLoadBalancerListenerResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteLoadBalancerListenerRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "DeleteLoadBalancerListener", "slb", "")

	resp := struct {
		*responses.BaseResponse
		DeleteLoadBalancerListenerResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteLoadBalancerListenerResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteLoadBalancerListenerResponse struct {
	RequestId string
}
