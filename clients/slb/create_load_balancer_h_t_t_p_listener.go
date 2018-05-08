package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateLoadBalancerHTTPListenerRequest struct {
	requests.RpcRequest
	Access_key_id          string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId        int64  `position:"Query" name:"ResourceOwnerId"`
	HealthCheckTimeout     int    `position:"Query" name:"HealthCheckTimeout"`
	ListenerForward        string `position:"Query" name:"ListenerForward"`
	XForwardedFor          string `position:"Query" name:"XForwardedFor"`
	HealthCheckURI         string `position:"Query" name:"HealthCheckURI"`
	UnhealthyThreshold     int    `position:"Query" name:"UnhealthyThreshold"`
	HealthyThreshold       int    `position:"Query" name:"HealthyThreshold"`
	AclStatus              string `position:"Query" name:"AclStatus"`
	Scheduler              string `position:"Query" name:"Scheduler"`
	AclType                string `position:"Query" name:"AclType"`
	HealthCheck            string `position:"Query" name:"HealthCheck"`
	ForwardPort            int    `position:"Query" name:"ForwardPort"`
	MaxConnection          int    `position:"Query" name:"MaxConnection"`
	CookieTimeout          int    `position:"Query" name:"CookieTimeout"`
	StickySessionType      string `position:"Query" name:"StickySessionType"`
	VpcIds                 string `position:"Query" name:"VpcIds"`
	VServerGroupId         string `position:"Query" name:"VServerGroupId"`
	AclId                  string `position:"Query" name:"AclId"`
	ListenerPort           int    `position:"Query" name:"ListenerPort"`
	Cookie                 string `position:"Query" name:"Cookie"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth              int    `position:"Query" name:"Bandwidth"`
	StickySession          string `position:"Query" name:"StickySession"`
	HealthCheckDomain      string `position:"Query" name:"HealthCheckDomain"`
	RequestTimeout         int    `position:"Query" name:"RequestTimeout"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	Gzip                   string `position:"Query" name:"Gzip"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
	Tags                   string `position:"Query" name:"Tags"`
	IdleTimeout            int    `position:"Query" name:"IdleTimeout"`
	LoadBalancerId         string `position:"Query" name:"LoadBalancerId"`
	XForwardedFor_SLBIP    string `position:"Query" name:"XForwardedFor_SLBIP"`
	BackendServerPort      int    `position:"Query" name:"BackendServerPort"`
	HealthCheckInterval    int    `position:"Query" name:"HealthCheckInterval"`
	XForwardedFor_proto    string `position:"Query" name:"XForwardedFor_proto"`
	XForwardedFor_SLBID    string `position:"Query" name:"XForwardedFor_SLBID"`
	HealthCheckConnectPort int    `position:"Query" name:"HealthCheckConnectPort"`
	HealthCheckHttpCode    string `position:"Query" name:"HealthCheckHttpCode"`
}

func (req *CreateLoadBalancerHTTPListenerRequest) Invoke(client *sdk.Client) (resp *CreateLoadBalancerHTTPListenerResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "CreateLoadBalancerHTTPListener", "slb", "")
	resp = &CreateLoadBalancerHTTPListenerResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateLoadBalancerHTTPListenerResponse struct {
	responses.BaseResponse
	RequestId common.String
}
