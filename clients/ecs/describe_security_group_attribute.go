package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId         string
	RegionId          string
	SecurityGroupId   string
	Description       string
	SecurityGroupName string
	VpcId             string
	InnerAccessPolicy string
	Permissions       DescribeSecurityGroupAttributePermissionList
}

type DescribeSecurityGroupAttributePermission struct {
	IpProtocol              string
	PortRange               string
	SourceGroupId           string
	SourceGroupName         string
	SourceCidrIp            string
	Policy                  string
	NicType                 string
	SourceGroupOwnerAccount string
	DestGroupId             string
	DestGroupName           string
	DestCidrIp              string
	DestGroupOwnerAccount   string
	Priority                string
	Direction               string
	Description             string
	CreateTime              string
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
