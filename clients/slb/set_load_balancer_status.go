package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetLoadBalancerStatusRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	LoadBalancerStatus   string `position:"Query" name:"LoadBalancerStatus"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (r SetLoadBalancerStatusRequest) Invoke(client *sdk.Client) (response *SetLoadBalancerStatusResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetLoadBalancerStatusRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "SetLoadBalancerStatus", "slb", "")

	resp := struct {
		*responses.BaseResponse
		SetLoadBalancerStatusResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetLoadBalancerStatusResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetLoadBalancerStatusResponse struct {
	RequestId string
}
