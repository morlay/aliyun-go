package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetLoadBalancerUDPListenerAttributeRequest struct {
	Access_key_id             string `position:"Query" name:"Access_key_id"`
	HealthCheckConnectTimeout int    `position:"Query" name:"HealthCheckConnectTimeout"`
	ResourceOwnerId           int64  `position:"Query" name:"ResourceOwnerId"`
	UnhealthyThreshold        int    `position:"Query" name:"UnhealthyThreshold"`
	HealthyThreshold          int    `position:"Query" name:"HealthyThreshold"`
	Scheduler                 string `position:"Query" name:"Scheduler"`
	MasterSlaveServerGroup    string `position:"Query" name:"MasterSlaveServerGroup"`
	MaxConnection             int    `position:"Query" name:"MaxConnection"`
	PersistenceTimeout        int    `position:"Query" name:"PersistenceTimeout"`
	VServerGroupId            string `position:"Query" name:"VServerGroupId"`
	ListenerPort              int    `position:"Query" name:"ListenerPort"`
	ResourceOwnerAccount      string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth                 int    `position:"Query" name:"Bandwidth"`
	OwnerAccount              string `position:"Query" name:"OwnerAccount"`
	OwnerId                   int64  `position:"Query" name:"OwnerId"`
	Tags                      string `position:"Query" name:"Tags"`
	LoadBalancerId            string `position:"Query" name:"LoadBalancerId"`
	MasterSlaveServerGroupId  string `position:"Query" name:"MasterSlaveServerGroupId"`
	HealthCheckReq            string `position:"Query" name:"HealthCheckReq"`
	HealthCheckInterval       int    `position:"Query" name:"HealthCheckInterval"`
	HealthCheckExp            string `position:"Query" name:"HealthCheckExp"`
	HealthCheckConnectPort    int    `position:"Query" name:"HealthCheckConnectPort"`
	VServerGroup              string `position:"Query" name:"VServerGroup"`
}

func (r SetLoadBalancerUDPListenerAttributeRequest) Invoke(client *sdk.Client) (response *SetLoadBalancerUDPListenerAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetLoadBalancerUDPListenerAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "SetLoadBalancerUDPListenerAttribute", "slb", "")

	resp := struct {
		*responses.BaseResponse
		SetLoadBalancerUDPListenerAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetLoadBalancerUDPListenerAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetLoadBalancerUDPListenerAttributeResponse struct {
	RequestId string
}
