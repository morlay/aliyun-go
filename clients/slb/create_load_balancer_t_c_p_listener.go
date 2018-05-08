package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateLoadBalancerTCPListenerRequest struct {
	requests.RpcRequest
	Access_key_id             string `position:"Query" name:"Access_key_id"`
	HealthCheckConnectTimeout int    `position:"Query" name:"HealthCheckConnectTimeout"`
	ResourceOwnerId           int64  `position:"Query" name:"ResourceOwnerId"`
	HealthCheckURI            string `position:"Query" name:"HealthCheckURI"`
	UnhealthyThreshold        int    `position:"Query" name:"UnhealthyThreshold"`
	HealthyThreshold          int    `position:"Query" name:"HealthyThreshold"`
	AclStatus                 string `position:"Query" name:"AclStatus"`
	Scheduler                 string `position:"Query" name:"Scheduler"`
	AclType                   string `position:"Query" name:"AclType"`
	EstablishedTimeout        int    `position:"Query" name:"EstablishedTimeout"`
	MaxConnection             int    `position:"Query" name:"MaxConnection"`
	PersistenceTimeout        int    `position:"Query" name:"PersistenceTimeout"`
	VpcIds                    string `position:"Query" name:"VpcIds"`
	VServerGroupId            string `position:"Query" name:"VServerGroupId"`
	AclId                     string `position:"Query" name:"AclId"`
	ListenerPort              int    `position:"Query" name:"ListenerPort"`
	HealthCheckType           string `position:"Query" name:"HealthCheckType"`
	ResourceOwnerAccount      string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth                 int    `position:"Query" name:"Bandwidth"`
	HealthCheckDomain         string `position:"Query" name:"HealthCheckDomain"`
	OwnerAccount              string `position:"Query" name:"OwnerAccount"`
	OwnerId                   int64  `position:"Query" name:"OwnerId"`
	Tags                      string `position:"Query" name:"Tags"`
	LoadBalancerId            string `position:"Query" name:"LoadBalancerId"`
	MasterSlaveServerGroupId  string `position:"Query" name:"MasterSlaveServerGroupId"`
	BackendServerPort         int    `position:"Query" name:"BackendServerPort"`
	HealthCheckInterval       int    `position:"Query" name:"HealthCheckInterval"`
	HealthCheckConnectPort    int    `position:"Query" name:"HealthCheckConnectPort"`
	HealthCheckHttpCode       string `position:"Query" name:"HealthCheckHttpCode"`
}

func (req *CreateLoadBalancerTCPListenerRequest) Invoke(client *sdk.Client) (resp *CreateLoadBalancerTCPListenerResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "CreateLoadBalancerTCPListener", "slb", "")
	resp = &CreateLoadBalancerTCPListenerResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateLoadBalancerTCPListenerResponse struct {
	responses.BaseResponse
	RequestId common.String
}
