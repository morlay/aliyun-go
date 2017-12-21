package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLoadBalancerHTTPSListenerAttributeRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ListenerPort         int    `position:"Query" name:"ListenerPort"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (r DescribeLoadBalancerHTTPSListenerAttributeRequest) Invoke(client *sdk.Client) (response *DescribeLoadBalancerHTTPSListenerAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeLoadBalancerHTTPSListenerAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeLoadBalancerHTTPSListenerAttribute", "slb", "")

	resp := struct {
		*responses.BaseResponse
		DescribeLoadBalancerHTTPSListenerAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeLoadBalancerHTTPSListenerAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeLoadBalancerHTTPSListenerAttributeResponse struct {
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
	ServerCertificateId    string
	CACertificateId        string
	MaxConnection          int
	VServerGroupId         string
	Gzip                   string
	XForwardedFor_SLBIP    string
	XForwardedFor_SLBID    string
	XForwardedFor_proto    string
}
