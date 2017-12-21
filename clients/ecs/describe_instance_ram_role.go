package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeInstanceRamRoleRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	InstanceIds          string `position:"Query" name:"InstanceIds"`
	PageSize             int    `position:"Query" name:"PageSize"`
	RamRoleName          string `position:"Query" name:"RamRoleName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (r DescribeInstanceRamRoleRequest) Invoke(client *sdk.Client) (response *DescribeInstanceRamRoleResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeInstanceRamRoleRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInstanceRamRole", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribeInstanceRamRoleResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeInstanceRamRoleResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeInstanceRamRoleResponse struct {
	RequestId           string
	RegionId            string
	TotalCount          int
	InstanceRamRoleSets DescribeInstanceRamRoleInstanceRamRoleSetList
}

type DescribeInstanceRamRoleInstanceRamRoleSet struct {
	InstanceId  string
	RamRoleName string
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
