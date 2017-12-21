package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLoadBalancerHTTPListenerAttributeRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ListenerPort         int    `position:"Query" name:"ListenerPort"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (r DescribeLoadBalancerHTTPListenerAttributeRequest) Invoke(client *sdk.Client) (response *DescribeLoadBalancerHTTPListenerAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeLoadBalancerHTTPListenerAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeLoadBalancerHTTPListenerAttribute", "slb", "")

	resp := struct {
		*responses.BaseResponse
		DescribeLoadBalancerHTTPListenerAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeLoadBalancerHTTPListenerAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeLoadBalancerHTTPListenerAttributeResponse struct {
	RequestId              string
	ListenerPort           int
	BackendServerPort      int
	Bandwidth              int
	Status                 string
	SecurityStatus         string
	XForwardedFor          string
	Scheduler              string
	StickySession          string
	StickySessionType      string
	CookieTimeout          int
	Cookie                 string
	HealthCheck            string
	HealthCheckDomain      string
	HealthCheckURI         string
	HealthyThreshold       int
	UnhealthyThreshold     int
	HealthCheckTimeout     int
	HealthCheckInterval    int
	HealthCheckConnectPort int
	HealthCheckHttpCode    string
	MaxConnection          int
	VServerGroupId         string
	Gzip                   string
	XForwardedFor_SLBIP    string
	XForwardedFor_SLBID    string
	XForwardedFor_proto    string
}
