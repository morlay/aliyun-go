package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeLoadBalancerAttributeRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (req *DescribeLoadBalancerAttributeRequest) Invoke(client *sdk.Client) (resp *DescribeLoadBalancerAttributeResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeLoadBalancerAttribute", "slb", "")
	resp = &DescribeLoadBalancerAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLoadBalancerAttributeResponse struct {
	responses.BaseResponse
	RequestId                common.String
	LoadBalancerId           common.String
	ResourceGroupId          common.String
	LoadBalancerName         common.String
	LoadBalancerStatus       common.String
	RegionId                 common.String
	RegionIdAlias            common.String
	Address                  common.String
	AddressType              common.String
	VpcId                    common.String
	VSwitchId                common.String
	NetworkType              common.String
	InternetChargeType       common.String
	AutoReleaseTime          common.Long
	Bandwidth                common.Integer
	LoadBalancerSpec         common.String
	CreateTime               common.String
	CreateTimeStamp          common.Long
	EndTime                  common.String
	EndTimeStamp             common.Long
	PayType                  common.String
	MasterZoneId             common.String
	SlaveZoneId              common.String
	ListenerPortsAndProtocal DescribeLoadBalancerAttributeListenerPortAndProtocalList
	ListenerPortsAndProtocol DescribeLoadBalancerAttributeListenerPortAndProtocolList
	BackendServers           DescribeLoadBalancerAttributeBackendServerList
	ListenerPorts            DescribeLoadBalancerAttributeListenerPortList
}

type DescribeLoadBalancerAttributeListenerPortAndProtocal struct {
	ListenerPort     common.Integer
	ListenerProtocal common.String
}

type DescribeLoadBalancerAttributeListenerPortAndProtocol struct {
	ListenerPort     common.Integer
	ListenerProtocol common.String
	ListenerForward  common.String
	ForwardPort      common.Integer
}

type DescribeLoadBalancerAttributeBackendServer struct {
	ServerId common.String
	Weight   common.Integer
	Type     common.String
	ServerIp common.String
	VpcId    common.String
}

type DescribeLoadBalancerAttributeListenerPortAndProtocalList []DescribeLoadBalancerAttributeListenerPortAndProtocal

func (list *DescribeLoadBalancerAttributeListenerPortAndProtocalList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLoadBalancerAttributeListenerPortAndProtocal)
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

type DescribeLoadBalancerAttributeListenerPortAndProtocolList []DescribeLoadBalancerAttributeListenerPortAndProtocol

func (list *DescribeLoadBalancerAttributeListenerPortAndProtocolList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLoadBalancerAttributeListenerPortAndProtocol)
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

type DescribeLoadBalancerAttributeBackendServerList []DescribeLoadBalancerAttributeBackendServer

func (list *DescribeLoadBalancerAttributeBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLoadBalancerAttributeBackendServer)
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

type DescribeLoadBalancerAttributeListenerPortList []common.String

func (list *DescribeLoadBalancerAttributeListenerPortList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
