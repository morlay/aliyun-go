package ram

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListRolesRequest struct {
	Marker   string `position:"Query" name:"Marker"`
	MaxItems int    `position:"Query" name:"MaxItems"`
}

func (r ListRolesRequest) Invoke(client *sdk.Client) (response *ListRolesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListRolesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "ListRoles", "", "")

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
	RequestId   string
	IsTruncated bool
	Marker      string
	Roles       ListRolesRoleList
}

type ListRolesRole struct {
	RoleId      string
	RoleName    string
	Arn         string
	Description string
	CreateDate  string
	UpdateDate  string
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
