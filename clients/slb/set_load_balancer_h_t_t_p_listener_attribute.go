package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetLoadBalancerHTTPListenerAttributeRequest struct {
	Access_key_id          string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId        int64  `position:"Query" name:"ResourceOwnerId"`
	HealthCheckTimeout     int    `position:"Query" name:"HealthCheckTimeout"`
	XForwardedFor          string `position:"Query" name:"XForwardedFor"`
	HealthCheckURI         string `position:"Query" name:"HealthCheckURI"`
	UnhealthyThreshold     int    `position:"Query" name:"UnhealthyThreshold"`
	HealthyThreshold       int    `position:"Query" name:"HealthyThreshold"`
	Scheduler              string `position:"Query" name:"Scheduler"`
	HealthCheck            string `position:"Query" name:"HealthCheck"`
	MaxConnection          int    `position:"Query" name:"MaxConnection"`
	CookieTimeout          int    `position:"Query" name:"CookieTimeout"`
	StickySessionType      string `position:"Query" name:"StickySessionType"`
	VServerGroupId         string `position:"Query" name:"VServerGroupId"`
	ListenerPort           int    `position:"Query" name:"ListenerPort"`
	Cookie                 string `position:"Query" name:"Cookie"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth              int    `position:"Query" name:"Bandwidth"`
	StickySession          string `position:"Query" name:"StickySession"`
	HealthCheckDomain      string `position:"Query" name:"HealthCheckDomain"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	Gzip                   string `position:"Query" name:"Gzip"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
	Tags                   string `position:"Query" name:"Tags"`
	LoadBalancerId         string `position:"Query" name:"LoadBalancerId"`
	XForwardedFor_SLBIP    string `position:"Query" name:"XForwardedFor_SLBIP"`
	HealthCheckInterval    int    `position:"Query" name:"HealthCheckInterval"`
	XForwardedFor_proto    string `position:"Query" name:"XForwardedFor_proto"`
	XForwardedFor_SLBID    string `position:"Query" name:"XForwardedFor_SLBID"`
	HealthCheckConnectPort int    `position:"Query" name:"HealthCheckConnectPort"`
	HealthCheckHttpCode    string `position:"Query" name:"HealthCheckHttpCode"`
	VServerGroup           string `position:"Query" name:"VServerGroup"`
}

func (r SetLoadBalancerHTTPListenerAttributeRequest) Invoke(client *sdk.Client) (response *SetLoadBalancerHTTPListenerAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetLoadBalancerHTTPListenerAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "SetLoadBalancerHTTPListenerAttribute", "slb", "")

	resp := struct {
		*responses.BaseResponse
		SetLoadBalancerHTTPListenerAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetLoadBalancerHTTPListenerAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetLoadBalancerHTTPListenerAttributeResponse struct {
	RequestId string
}
