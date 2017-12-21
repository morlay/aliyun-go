package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLoadBalancersRelatedEcsRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (req *DescribeLoadBalancersRelatedEcsRequest) Invoke(client *sdk.Client) (resp *DescribeLoadBalancersRelatedEcsResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeLoadBalancersRelatedEcs", "slb", "")
	resp = &DescribeLoadBalancersRelatedEcsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLoadBalancersRelatedEcsResponse struct {
	responses.BaseResponse
	Message       string
	Success       bool
	RequestId     string
	LoadBalancers DescribeLoadBalancersRelatedEcsLoadBalancerList
}

type DescribeLoadBalancersRelatedEcsLoadBalancer struct {
	LoadBalancerId           string
	Count                    int
	MasterSlaveVServerGroups DescribeLoadBalancersRelatedEcsMasterSlaveVServerGroupList
	VServerGroups            DescribeLoadBalancersRelatedEcsVServerGroupList
	BackendServers           DescribeLoadBalancersRelatedEcsBackendServer4List
}

type DescribeLoadBalancersRelatedEcsMasterSlaveVServerGroup struct {
	GroupId         string
	GroupName       string
	BackendServers1 DescribeLoadBalancersRelatedEcsBackendServerList
}

type DescribeLoadBalancersRelatedEcsBackendServer struct {
	VmName      string
	Weight      int
	Port        int
	NetworkType string
}

type DescribeLoadBalancersRelatedEcsVServerGroup struct {
	GroupId         string
	GroupName       string
	BackendServers2 DescribeLoadBalancersRelatedEcsBackendServer3List
}

type DescribeLoadBalancersRelatedEcsBackendServer3 struct {
	VmName      string
	Weight      int
	Port        int
	NetworkType string
}

type DescribeLoadBalancersRelatedEcsBackendServer4 struct {
	VmName      string
	Weight      int
	Port        int
	NetworkType string
}

type DescribeLoadBalancersRelatedEcsLoadBalancerList []DescribeLoadBalancersRelatedEcsLoadBalancer

func (list *DescribeLoadBalancersRelatedEcsLoadBalancerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLoadBalancersRelatedEcsLoadBalancer)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

type DescribeLoadBalancersRelatedEcsMasterSlaveVServerGroupList []DescribeLoadBalancersRelatedEcsMasterSlaveVServerGroup

func (list *DescribeLoadBalancersRelatedEcsMasterSlaveVServerGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLoadBalancersRelatedEcsMasterSlaveVServerGroup)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

type DescribeLoadBalancersRelatedEcsVServerGroupList []DescribeLoadBalancersRelatedEcsVServerGroup

func (list *DescribeLoadBalancersRelatedEcsVServerGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLoadBalancersRelatedEcsVServerGroup)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

type DescribeLoadBalancersRelatedEcsBackendServer4List []DescribeLoadBalancersRelatedEcsBackendServer4

func (list *DescribeLoadBalancersRelatedEcsBackendServer4List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLoadBalancersRelatedEcsBackendServer4)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

type DescribeLoadBalancersRelatedEcsBackendServerList []DescribeLoadBalancersRelatedEcsBackendServer

func (list *DescribeLoadBalancersRelatedEcsBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLoadBalancersRelatedEcsBackendServer)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

type DescribeLoadBalancersRelatedEcsBackendServer3List []DescribeLoadBalancersRelatedEcsBackendServer3

func (list *DescribeLoadBalancersRelatedEcsBackendServer3List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLoadBalancersRelatedEcsBackendServer3)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
