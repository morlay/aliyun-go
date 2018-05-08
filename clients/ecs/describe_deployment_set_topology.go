package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDeploymentSetTopologyRequest struct {
	requests.RpcRequest
	DeploymentSetId      string `position:"Query" name:"DeploymentSetId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Granularity          string `position:"Query" name:"Granularity"`
	Domain               string `position:"Query" name:"Domain"`
	NetworkType          string `position:"Query" name:"NetworkType"`
	DeploymentSetName    string `position:"Query" name:"DeploymentSetName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Strategy             string `position:"Query" name:"Strategy"`
}

func (req *DescribeDeploymentSetTopologyRequest) Invoke(client *sdk.Client) (resp *DescribeDeploymentSetTopologyResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeDeploymentSetTopology", "ecs", "")
	resp = &DescribeDeploymentSetTopologyResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDeploymentSetTopologyResponse struct {
	responses.BaseResponse
	RequestId common.String
	Switchs   DescribeDeploymentSetTopology_SwitchList
	Racks     DescribeDeploymentSetTopologyRackList
}

type DescribeDeploymentSetTopology_Switch struct {
	SwitchId common.String
	Hosts    DescribeDeploymentSetTopologyHostList
}

type DescribeDeploymentSetTopologyHost struct {
	HostId      common.String
	InstanceIds DescribeDeploymentSetTopologyInstanceIdList
}

type DescribeDeploymentSetTopologyRack struct {
	RackId common.String
	Hosts1 DescribeDeploymentSetTopologyHost2List
}

type DescribeDeploymentSetTopologyHost2 struct {
	HostId       common.String
	InstanceIds3 DescribeDeploymentSetTopologyInstanceIds3List
}

type DescribeDeploymentSetTopology_SwitchList []DescribeDeploymentSetTopology_Switch

func (list *DescribeDeploymentSetTopology_SwitchList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDeploymentSetTopology_Switch)
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

type DescribeDeploymentSetTopologyRackList []DescribeDeploymentSetTopologyRack

func (list *DescribeDeploymentSetTopologyRackList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDeploymentSetTopologyRack)
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

type DescribeDeploymentSetTopologyHostList []DescribeDeploymentSetTopologyHost

func (list *DescribeDeploymentSetTopologyHostList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDeploymentSetTopologyHost)
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

type DescribeDeploymentSetTopologyInstanceIdList []common.String

func (list *DescribeDeploymentSetTopologyInstanceIdList) UnmarshalJSON(data []byte) error {
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

type DescribeDeploymentSetTopologyHost2List []DescribeDeploymentSetTopologyHost2

func (list *DescribeDeploymentSetTopologyHost2List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDeploymentSetTopologyHost2)
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

type DescribeDeploymentSetTopologyInstanceIds3List []common.String

func (list *DescribeDeploymentSetTopologyInstanceIds3List) UnmarshalJSON(data []byte) error {
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
