package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLoadBalancerTCPListenerAttributeRequest struct {
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

func (req *DescribeLoadBalancerTCPListenerAttributeRequest) Invoke(client *sdk.Client) (resp *DescribeLoadBalancerTCPListenerAttributeResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeLoadBalancerTCPListenerAttribute", "slb", "")
	resp = &DescribeLoadBalancerTCPListenerAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLoadBalancerTCPListenerAttributeResponse struct {
	responses.BaseResponse
	RequestId                 string
	ListenerPort              int
	BackendServerPort         int
	Status                    string
	Bandwidth                 int
	Scheduler                 string
	SynProxy                  string
	PersistenceTimeout        int
	EstablishedTimeout        int
	HealthCheck               string
	HealthyThreshold          int
	UnhealthyThreshold        int
	HealthCheckConnectTimeout int
	HealthCheckConnectPort    int
	HealthCheckInterval       int
	HealthCheckHttpCode       string
	HealthCheckDomain         string
	HealthCheckURI            string
	HealthCheckType           string
	MaxConnection             int
	VServerGroupId            string
	MasterSlaveServerGroupId  string
	AclId                     string
	AclType                   string
	AclStatus                 string
	VpcIds                    string
}
