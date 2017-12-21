package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListUsersOfSkillGroupRequest struct {
	InstanceId   string `position:"Query" name:"InstanceId"`
	SkillGroupId string `position:"Query" name:"SkillGroupId"`
	PageSize     int    `position:"Query" name:"PageSize"`
	PageNumber   int    `position:"Query" name:"PageNumber"`
}

func (r ListUsersOfSkillGroupRequest) Invoke(client *sdk.Client) (response *ListUsersOfSkillGroupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListUsersOfSkillGroupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CCC", "2017-07-05", "ListUsersOfSkillGroup", "ccc", "")

	resp := struct {
		*responses.BaseResponse
		ListUsersOfSkillGroupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListUsersOfSkillGroupResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListUsersOfSkillGroupResponse struct {
	RequestId      string
	Success        bool
	Code           string
	Message        string
	HttpStatusCode int
	Users          ListUsersOfSkillGroupUsers
}

type ListUsersOfSkillGroupUsers struct {
	TotalCount int
	PageNumber int
	PageSize   int
	List       ListUsersOfSkillGroupUserList
}

type ListUsersOfSkillGroupUser struct {
	UserId      string
	RamId       string
	InstanceId  string
	Roles       ListUsersOfSkillGroupRoleList
	SkillLevels ListUsersOfSkillGroupSkillLevelList
	Detail      ListUsersOfSkillGroupDetail
}

type ListUsersOfSkillGroupRole struct {
	RoleId          string
	InstanceId      string
	RoleName        string
	RoleDescription string
	UserCount       int
	Privileges      ListUsersOfSkillGroupPrivilegeList
}

type ListUsersOfSkillGroupPrivilege struct {
	PrivilegeId          string
	PrivilegeName        string
	PrivilegeDescription string
}

type ListUsersOfSkillGroupSkillLevel struct {
	SkillLevelId string
	Level        int
	Skill        ListUsersOfSkillGroupSkill
}

type ListUsersOfSkillGroupSkill struct {
	SkillGroupId          string
	InstanceId            string
	SkillGroupName        string
	SkillGroupDescription string
}

type ListUsersOfSkillGroupDetail struct {
	LoginName   string
	DisplayName string
	Phone       string
	Email       string
	Department  string
}

type ListUsersOfSkillGroupUserList []ListUsersOfSkillGroupUser

func (list *ListUsersOfSkillGroupUserList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListUsersOfSkillGroupUser)
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

type ListUsersOfSkillGroupRoleList []ListUsersOfSkillGroupRole

func (list *ListUsersOfSkillGroupRoleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListUsersOfSkillGroupRole)
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

type ListUsersOfSkillGroupSkillLevelList []ListUsersOfSkillGroupSkillLevel

func (list *ListUsersOfSkillGroupSkillLevelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListUsersOfSkillGroupSkillLevel)
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

type ListUsersOfSkillGroupPrivilegeList []ListUsersOfSkillGroupPrivilege

func (list *ListUsersOfSkillGroupPrivilegeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListUsersOfSkillGroupPrivilege)
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
