package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetLoadBalancerTCPListenerAttributeRequest struct {
	Access_key_id             string `position:"Query" name:"Access_key_id"`
	HealthCheckConnectTimeout int    `position:"Query" name:"HealthCheckConnectTimeout"`
	ResourceOwnerId           int64  `position:"Query" name:"ResourceOwnerId"`
	HealthCheckURI            string `position:"Query" name:"HealthCheckURI"`
	UnhealthyThreshold        int    `position:"Query" name:"UnhealthyThreshold"`
	HealthyThreshold          int    `position:"Query" name:"HealthyThreshold"`
	Scheduler                 string `position:"Query" name:"Scheduler"`
	MasterSlaveServerGroup    string `position:"Query" name:"MasterSlaveServerGroup"`
	EstablishedTimeout        int    `position:"Query" name:"EstablishedTimeout"`
	MaxConnection             int    `position:"Query" name:"MaxConnection"`
	PersistenceTimeout        int    `position:"Query" name:"PersistenceTimeout"`
	VServerGroupId            string `position:"Query" name:"VServerGroupId"`
	ListenerPort              int    `position:"Query" name:"ListenerPort"`
	HealthCheckType           string `position:"Query" name:"HealthCheckType"`
	ResourceOwnerAccount      string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth                 int    `position:"Query" name:"Bandwidth"`
	HealthCheckDomain         string `position:"Query" name:"HealthCheckDomain"`
	OwnerAccount              string `position:"Query" name:"OwnerAccount"`
	SynProxy                  string `position:"Query" name:"SynProxy"`
	OwnerId                   int64  `position:"Query" name:"OwnerId"`
	Tags                      string `position:"Query" name:"Tags"`
	LoadBalancerId            string `position:"Query" name:"LoadBalancerId"`
	MasterSlaveServerGroupId  string `position:"Query" name:"MasterSlaveServerGroupId"`
	HealthCheckInterval       int    `position:"Query" name:"HealthCheckInterval"`
	HealthCheckConnectPort    int    `position:"Query" name:"HealthCheckConnectPort"`
	HealthCheckHttpCode       string `position:"Query" name:"HealthCheckHttpCode"`
	VServerGroup              string `position:"Query" name:"VServerGroup"`
}

func (r SetLoadBalancerTCPListenerAttributeRequest) Invoke(client *sdk.Client) (response *SetLoadBalancerTCPListenerAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetLoadBalancerTCPListenerAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "SetLoadBalancerTCPListenerAttribute", "slb", "")

	resp := struct {
		*responses.BaseResponse
		SetLoadBalancerTCPListenerAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetLoadBalancerTCPListenerAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetLoadBalancerTCPListenerAttributeResponse struct {
	RequestId string
}
