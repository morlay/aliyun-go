package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateLoadBalancerUDPListenerRequest struct {
	requests.RpcRequest
	Access_key_id             string `position:"Query" name:"Access_key_id"`
	HealthCheckConnectTimeout int    `position:"Query" name:"HealthCheckConnectTimeout"`
	ResourceOwnerId           int64  `position:"Query" name:"ResourceOwnerId"`
	UnhealthyThreshold        int    `position:"Query" name:"UnhealthyThreshold"`
	HealthyThreshold          int    `position:"Query" name:"HealthyThreshold"`
	AclStatus                 string `position:"Query" name:"AclStatus"`
	Scheduler                 string `position:"Query" name:"Scheduler"`
	AclType                   string `position:"Query" name:"AclType"`
	MaxConnection             int    `position:"Query" name:"MaxConnection"`
	PersistenceTimeout        int    `position:"Query" name:"PersistenceTimeout"`
	VpcIds                    string `position:"Query" name:"VpcIds"`
	VServerGroupId            string `position:"Query" name:"VServerGroupId"`
	AclId                     string `position:"Query" name:"AclId"`
	ListenerPort              int    `position:"Query" name:"ListenerPort"`
	ResourceOwnerAccount      string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth                 int    `position:"Query" name:"Bandwidth"`
	OwnerAccount              string `position:"Query" name:"OwnerAccount"`
	OwnerId                   int64  `position:"Query" name:"OwnerId"`
	Tags                      string `position:"Query" name:"Tags"`
	LoadBalancerId            string `position:"Query" name:"LoadBalancerId"`
	MasterSlaveServerGroupId  string `position:"Query" name:"MasterSlaveServerGroupId"`
	HealthCheckReq            string `position:"Query" name:"HealthCheckReq"`
	BackendServerPort         int    `position:"Query" name:"BackendServerPort"`
	HealthCheckInterval       int    `position:"Query" name:"HealthCheckInterval"`
	HealthCheckExp            string `position:"Query" name:"HealthCheckExp"`
	HealthCheckConnectPort    int    `position:"Query" name:"HealthCheckConnectPort"`
}

func (req *CreateLoadBalancerUDPListenerRequest) Invoke(client *sdk.Client) (resp *CreateLoadBalancerUDPListenerResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "CreateLoadBalancerUDPListener", "slb", "")
	resp = &CreateLoadBalancerUDPListenerResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateLoadBalancerUDPListenerResponse struct {
	responses.BaseResponse
	RequestId string
}
