package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeSecurityGroupReferencesRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64                                               `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string                                              `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                                              `position:"Query" name:"OwnerAccount"`
	SecurityGroupIds     *DescribeSecurityGroupReferencesSecurityGroupIdList `position:"Query" type:"Repeated" name:"SecurityGroupId"`
	OwnerId              int64                                               `position:"Query" name:"OwnerId"`
}

func (req *DescribeSecurityGroupReferencesRequest) Invoke(client *sdk.Client) (resp *DescribeSecurityGroupReferencesResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeSecurityGroupReferences", "ecs", "")
	resp = &DescribeSecurityGroupReferencesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeSecurityGroupReferencesResponse struct {
	responses.BaseResponse
	RequestId               string
	SecurityGroupReferences DescribeSecurityGroupReferencesSecurityGroupReferenceList
}

type DescribeSecurityGroupReferencesSecurityGroupReference struct {
	SecurityGroupId           string
	ReferencingSecurityGroups DescribeSecurityGroupReferencesReferencingSecurityGroupList
}

type DescribeSecurityGroupReferencesReferencingSecurityGroup struct {
	AliUid          string
	SecurityGroupId string
}

type DescribeSecurityGroupReferencesSecurityGroupIdList []string

func (list *DescribeSecurityGroupReferencesSecurityGroupIdList) UnmarshalJSON(data []byte) error {
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

type DescribeSecurityGroupReferencesSecurityGroupReferenceList []DescribeSecurityGroupReferencesSecurityGroupReference

func (list *DescribeSecurityGroupReferencesSecurityGroupReferenceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSecurityGroupReferencesSecurityGroupReference)
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

type DescribeSecurityGroupReferencesReferencingSecurityGroupList []DescribeSecurityGroupReferencesReferencingSecurityGroup

func (list *DescribeSecurityGroupReferencesReferencingSecurityGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSecurityGroupReferencesReferencingSecurityGroup)
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
