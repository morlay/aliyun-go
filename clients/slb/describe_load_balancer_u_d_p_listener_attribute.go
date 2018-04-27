package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLoadBalancerUDPListenerAttributeRequest struct {
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

func (req *DescribeLoadBalancerUDPListenerAttributeRequest) Invoke(client *sdk.Client) (resp *DescribeLoadBalancerUDPListenerAttributeResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeLoadBalancerUDPListenerAttribute", "slb", "")
	resp = &DescribeLoadBalancerUDPListenerAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLoadBalancerUDPListenerAttributeResponse struct {
	responses.BaseResponse
	RequestId                 string
	ListenerPort              int
	BackendServerPort         int
	Status                    string
	Bandwidth                 int
	Scheduler                 string
	PersistenceTimeout        int
	HealthCheck               string
	HealthyThreshold          int
	UnhealthyThreshold        int
	HealthCheckConnectTimeout int
	HealthCheckConnectPort    int
	HealthCheckInterval       int
	HealthCheckReq            string
	HealthCheckExp            string
	MaxConnection             int
	VServerGroupId            string
	MasterSlaveServerGroupId  string
	AclId                     string
	AclType                   string
	AclStatus                 string
	VpcIds                    string
}
