package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListUsersOfSkillGroupRequest struct {
	requests.RpcRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	SkillGroupId string `position:"Query" name:"SkillGroupId"`
	PageSize     int    `position:"Query" name:"PageSize"`
	PageNumber   int    `position:"Query" name:"PageNumber"`
}

func (req *ListUsersOfSkillGroupRequest) Invoke(client *sdk.Client) (resp *ListUsersOfSkillGroupResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "ListUsersOfSkillGroup", "ccc", "")
	resp = &ListUsersOfSkillGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListUsersOfSkillGroupResponse struct {
	responses.BaseResponse
	RequestId      common.String
	Success        bool
	Code           common.String
	Message        common.String
	HttpStatusCode common.Integer
	Users          ListUsersOfSkillGroupUsers
}

type ListUsersOfSkillGroupUsers struct {
	TotalCount common.Integer
	PageNumber common.Integer
	PageSize   common.Integer
	List       ListUsersOfSkillGroupUserList
}

type ListUsersOfSkillGroupUser struct {
	UserId      common.String
	RamId       common.String
	InstanceId  common.String
	Roles       ListUsersOfSkillGroupRoleList
	SkillLevels ListUsersOfSkillGroupSkillLevelList
	Detail      ListUsersOfSkillGroupDetail
}

type ListUsersOfSkillGroupRole struct {
	RoleId          common.String
	InstanceId      common.String
	RoleName        common.String
	RoleDescription common.String
	UserCount       common.Integer
	Privileges      ListUsersOfSkillGroupPrivilegeList
}

type ListUsersOfSkillGroupPrivilege struct {
	PrivilegeId          common.String
	PrivilegeName        common.String
	PrivilegeDescription common.String
}

type ListUsersOfSkillGroupSkillLevel struct {
	SkillLevelId common.String
	Level        common.Integer
	Skill        ListUsersOfSkillGroupSkill
}

type ListUsersOfSkillGroupSkill struct {
	SkillGroupId          common.String
	InstanceId            common.String
	SkillGroupName        common.String
	SkillGroupDescription common.String
}

type ListUsersOfSkillGroupDetail struct {
	LoginName   common.String
	DisplayName common.String
	Phone       common.String
	Email       common.String
	Department  common.String
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
