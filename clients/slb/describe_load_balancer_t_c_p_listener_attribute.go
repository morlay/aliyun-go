package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId                 common.String
	ListenerPort              common.Integer
	BackendServerPort         common.Integer
	Status                    common.String
	Bandwidth                 common.Integer
	Scheduler                 common.String
	SynProxy                  common.String
	PersistenceTimeout        common.Integer
	EstablishedTimeout        common.Integer
	HealthCheck               common.String
	HealthyThreshold          common.Integer
	UnhealthyThreshold        common.Integer
	HealthCheckConnectTimeout common.Integer
	HealthCheckConnectPort    common.Integer
	HealthCheckInterval       common.Integer
	HealthCheckHttpCode       common.String
	HealthCheckDomain         common.String
	HealthCheckURI            common.String
	HealthCheckType           common.String
	MaxConnection             common.Integer
	VServerGroupId            common.String
	MasterSlaveServerGroupId  common.String
	AclId                     common.String
	AclType                   common.String
	AclStatus                 common.String
	VpcIds                    common.String
}
