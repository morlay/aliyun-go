package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeLoadBalancerHTTPSListenerAttributeRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ListenerPort         int    `position:"Query" name:"ListenerPort"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (req *DescribeLoadBalancerHTTPSListenerAttributeRequest) Invoke(client *sdk.Client) (resp *DescribeLoadBalancerHTTPSListenerAttributeResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeLoadBalancerHTTPSListenerAttribute", "slb", "")
	resp = &DescribeLoadBalancerHTTPSListenerAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLoadBalancerHTTPSListenerAttributeResponse struct {
	responses.BaseResponse
	RequestId              common.String
	ListenerPort           common.Integer
	BackendServerPort      common.Integer
	Bandwidth              common.Integer
	Status                 common.String
	SecurityStatus         common.String
	XForwardedFor          common.String
	Scheduler              common.String
	StickySession          common.String
	StickySessionType      common.String
	CookieTimeout          common.Integer
	Cookie                 common.String
	HealthCheck            common.String
	HealthCheckDomain      common.String
	HealthCheckURI         common.String
	HealthyThreshold       common.Integer
	UnhealthyThreshold     common.Integer
	HealthCheckTimeout     common.Integer
	HealthCheckInterval    common.Integer
	HealthCheckConnectPort common.Integer
	HealthCheckHttpCode    common.String
	ServerCertificateId    common.String
	CACertificateId        common.String
	MaxConnection          common.Integer
	VServerGroupId         common.String
	Gzip                   common.String
	XForwardedFor_SLBIP    common.String
	XForwardedFor_SLBID    common.String
	XForwardedFor_proto    common.String
	AclId                  common.String
	AclType                common.String
	AclStatus              common.String
	VpcIds                 common.String
	RequestTimeout         common.Integer
	IdleTimeout            common.Integer
}
