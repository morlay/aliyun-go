package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListRolesRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
}

func (r ListRolesRequest) Invoke(client *sdk.Client) (response *ListRolesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListRolesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CCC", "2017-07-05", "ListRoles", "ccc", "")

	resp := struct {
		*responses.BaseResponse
		ListRolesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListRolesResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListRolesResponse struct {
	RequestId      string
	Success        bool
	Code           string
	Message        string
	HttpStatusCode int
	Roles          ListRolesRoleList
}

type ListRolesRole struct {
	RoleId          string
	InstanceId      string
	RoleName        string
	RoleDescription string
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
