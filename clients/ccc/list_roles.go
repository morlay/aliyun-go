package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListRolesRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

func (req *ListRolesRequest) Invoke(client *sdk.Client) (resp *ListRolesResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "ListRoles", "ccc", "")
	resp = &ListRolesResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListRolesResponse struct {
	responses.BaseResponse
	RequestId      common.String
	Success        bool
	Code           common.String
	Message        common.String
	HttpStatusCode common.Integer
	Roles          ListRolesRoleList
}

type ListRolesRole struct {
	RoleId          common.String
	InstanceId      common.String
	RoleName        common.String
	RoleDescription common.String
}

type ListRolesRoleList []ListRolesRole

func (list *ListRolesRoleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListRolesRole)
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
