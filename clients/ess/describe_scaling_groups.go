package ess

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeScalingGroupsRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ScalingGroupId10     string `position:"Query" name:"ScalingGroupId.10"`
	ScalingGroupId12     string `position:"Query" name:"ScalingGroupId.12"`
	ScalingGroupId13     string `position:"Query" name:"ScalingGroupId.13"`
	ScalingGroupId14     string `position:"Query" name:"ScalingGroupId.14"`
	ScalingGroupId15     string `position:"Query" name:"ScalingGroupId.15"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	PageSize             int    `position:"Query" name:"PageSize"`
	ScalingGroupName20   string `position:"Query" name:"ScalingGroupName.20"`
	ScalingGroupName19   string `position:"Query" name:"ScalingGroupName.19"`
	ScalingGroupId20     string `position:"Query" name:"ScalingGroupId.20"`
	ScalingGroupName18   string `position:"Query" name:"ScalingGroupName.18"`
	ScalingGroupName17   string `position:"Query" name:"ScalingGroupName.17"`
	ScalingGroupName16   string `position:"Query" name:"ScalingGroupName.16"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupName     string `position:"Query" name:"ScalingGroupName"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ScalingGroupName1    string `position:"Query" name:"ScalingGroupName.1"`
	ScalingGroupName2    string `position:"Query" name:"ScalingGroupName.2"`
	ScalingGroupId2      string `position:"Query" name:"ScalingGroupId.2"`
	ScalingGroupId1      string `position:"Query" name:"ScalingGroupId.1"`
	ScalingGroupId6      string `position:"Query" name:"ScalingGroupId.6"`
	ScalingGroupId16     string `position:"Query" name:"ScalingGroupId.16"`
	ScalingGroupName7    string `position:"Query" name:"ScalingGroupName.7"`
	ScalingGroupName11   string `position:"Query" name:"ScalingGroupName.11"`
	ScalingGroupId5      string `position:"Query" name:"ScalingGroupId.5"`
	ScalingGroupId17     string `position:"Query" name:"ScalingGroupId.17"`
	ScalingGroupName8    string `position:"Query" name:"ScalingGroupName.8"`
	ScalingGroupName10   string `position:"Query" name:"ScalingGroupName.10"`
	ScalingGroupId4      string `position:"Query" name:"ScalingGroupId.4"`
	ScalingGroupId18     string `position:"Query" name:"ScalingGroupId.18"`
	ScalingGroupName9    string `position:"Query" name:"ScalingGroupName.9"`
	ScalingGroupId3      string `position:"Query" name:"ScalingGroupId.3"`
	ScalingGroupId19     string `position:"Query" name:"ScalingGroupId.19"`
	ScalingGroupName3    string `position:"Query" name:"ScalingGroupName.3"`
	ScalingGroupName15   string `position:"Query" name:"ScalingGroupName.15"`
	ScalingGroupId9      string `position:"Query" name:"ScalingGroupId.9"`
	ScalingGroupName4    string `position:"Query" name:"ScalingGroupName.4"`
	ScalingGroupName14   string `position:"Query" name:"ScalingGroupName.14"`
	ScalingGroupId8      string `position:"Query" name:"ScalingGroupId.8"`
	ScalingGroupName5    string `position:"Query" name:"ScalingGroupName.5"`
	ScalingGroupName13   string `position:"Query" name:"ScalingGroupName.13"`
	ScalingGroupId7      string `position:"Query" name:"ScalingGroupId.7"`
	ScalingGroupName6    string `position:"Query" name:"ScalingGroupName.6"`
	ScalingGroupName12   string `position:"Query" name:"ScalingGroupName.12"`
}

func (req *DescribeScalingGroupsRequest) Invoke(client *sdk.Client) (resp *DescribeScalingGroupsResponse, err error) {
	req.InitWithApiInfo("Ess", "2014-08-28", "DescribeScalingGroups", "ess", "")
	resp = &DescribeScalingGroupsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeScalingGroupsResponse struct {
	responses.BaseResponse
	TotalCount    int
	PageNumber    int
	PageSize      int
	RequestId     string
	ScalingGroups DescribeScalingGroupsScalingGroupList
}

type DescribeScalingGroupsScalingGroup struct {
	DefaultCooldown              int
	MaxSize                      int
	PendingCapacity              int
	RemovingCapacity             int
	ScalingGroupName             string
	ActiveCapacity               int
	StandbyCapacity              int
	ProtectedCapacity            int
	ActiveScalingConfigurationId string
	ScalingGroupId               string
	RegionId                     string
	TotalCapacity                int
	MinSize                      int
	LifecycleState               string
	CreationTime                 string
	ModificationTime             string
	VpcId                        string
	VSwitchId                    string
	MultiAZPolicy                string
	VSwitchIds                   DescribeScalingGroupsVSwitchIdList
	RemovalPolicies              DescribeScalingGroupsRemovalPolicyList
	DBInstanceIds                DescribeScalingGroupsDBInstanceIdList
	LoadBalancerIds              DescribeScalingGroupsLoadBalancerIdList
}

type DescribeScalingGroupsScalingGroupList []DescribeScalingGroupsScalingGroup

func (list *DescribeScalingGroupsScalingGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeScalingGroupsScalingGroup)
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

type DescribeScalingGroupsVSwitchIdList []string

func (list *DescribeScalingGroupsVSwitchIdList) UnmarshalJSON(data []byte) error {
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

type DescribeScalingGroupsRemovalPolicyList []string

func (list *DescribeScalingGroupsRemovalPolicyList) UnmarshalJSON(data []byte) error {
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

type DescribeScalingGroupsDBInstanceIdList []string

func (list *DescribeScalingGroupsDBInstanceIdList) UnmarshalJSON(data []byte) error {
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

type DescribeScalingGroupsLoadBalancerIdList []string

func (list *DescribeScalingGroupsLoadBalancerIdList) UnmarshalJSON(data []byte) error {
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
