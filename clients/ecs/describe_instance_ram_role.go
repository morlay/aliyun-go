package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeInstanceRamRoleRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	InstanceIds          string `position:"Query" name:"InstanceIds"`
	PageSize             int    `position:"Query" name:"PageSize"`
	RamRoleName          string `position:"Query" name:"RamRoleName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeInstanceRamRoleRequest) Invoke(client *sdk.Client) (resp *DescribeInstanceRamRoleResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInstanceRamRole", "ecs", "")
	resp = &DescribeInstanceRamRoleResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeInstanceRamRoleResponse struct {
	responses.BaseResponse
	RequestId           common.String
	RegionId            common.String
	TotalCount          common.Integer
	InstanceRamRoleSets DescribeInstanceRamRoleInstanceRamRoleSetList
}

type DescribeInstanceRamRoleInstanceRamRoleSet struct {
	InstanceId  common.String
	RamRoleName common.String
}

type DescribeInstanceRamRoleInstanceRamRoleSetList []DescribeInstanceRamRoleInstanceRamRoleSet

func (list *DescribeInstanceRamRoleInstanceRamRoleSetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstanceRamRoleInstanceRamRoleSet)
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
