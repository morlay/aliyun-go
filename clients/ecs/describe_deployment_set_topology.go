package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDeploymentSetTopologyRequest struct {
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

func (r DescribeDeploymentSetTopologyRequest) Invoke(client *sdk.Client) (response *DescribeDeploymentSetTopologyResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDeploymentSetTopologyRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeDeploymentSetTopology", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDeploymentSetTopologyResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDeploymentSetTopologyResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDeploymentSetTopologyResponse struct {
	RequestId string
	Switchs   DescribeDeploymentSetTopology_SwitchList
	Racks     DescribeDeploymentSetTopologyRackList
}

type DescribeDeploymentSetTopology_Switch struct {
	SwitchId string
	Hosts    DescribeDeploymentSetTopologyHostList
}

type DescribeDeploymentSetTopologyHost struct {
	HostId      string
	InstanceIds DescribeDeploymentSetTopologyInstanceIdList
}

type DescribeDeploymentSetTopologyRack struct {
	RackId string
	Hosts1 DescribeDeploymentSetTopologyHost2List
}

type DescribeDeploymentSetTopologyHost2 struct {
	HostId       string
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

type DescribeDeploymentSetTopologyInstanceIdList []string

func (list *DescribeDeploymentSetTopologyInstanceIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type DescribeDeploymentSetTopologyInstanceIds3List []string

func (list *DescribeDeploymentSetTopologyInstanceIds3List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
