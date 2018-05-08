package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeSecurityGroupAttributeRequest struct {
	requests.RpcRequest
	NicType              string `position:"Query" name:"NicType"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	SecurityGroupId      string `position:"Query" name:"SecurityGroupId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Direction            string `position:"Query" name:"Direction"`
}

func (req *DescribeSecurityGroupAttributeRequest) Invoke(client *sdk.Client) (resp *DescribeSecurityGroupAttributeResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeSecurityGroupAttribute", "ecs", "")
	resp = &DescribeSecurityGroupAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeSecurityGroupAttributeResponse struct {
	responses.BaseResponse
	RequestId         common.String
	RegionId          common.String
	SecurityGroupId   common.String
	Description       common.String
	SecurityGroupName common.String
	VpcId             common.String
	InnerAccessPolicy common.String
	Permissions       DescribeSecurityGroupAttributePermissionList
}

type DescribeSecurityGroupAttributePermission struct {
	IpProtocol              common.String
	PortRange               common.String
	SourceGroupId           common.String
	SourceGroupName         common.String
	SourceCidrIp            common.String
	Policy                  common.String
	NicType                 common.String
	SourceGroupOwnerAccount common.String
	DestGroupId             common.String
	DestGroupName           common.String
	DestCidrIp              common.String
	DestGroupOwnerAccount   common.String
	Priority                common.String
	Direction               common.String
	Description             common.String
	CreateTime              common.String
}

type DescribeSecurityGroupAttributePermissionList []DescribeSecurityGroupAttributePermission

func (list *DescribeSecurityGroupAttributePermissionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSecurityGroupAttributePermission)
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
