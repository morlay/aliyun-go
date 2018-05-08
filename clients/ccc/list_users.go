package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListUsersRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	PageSize   int    `position:"Query" name:"PageSize"`
	PageNumber int    `position:"Query" name:"PageNumber"`
}

func (req *ListUsersRequest) Invoke(client *sdk.Client) (resp *ListUsersResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "ListUsers", "ccc", "")
	resp = &ListUsersResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListUsersResponse struct {
	responses.BaseResponse
	RequestId      common.String
	Success        bool
	Code           common.String
	Message        common.String
	HttpStatusCode common.Integer
	Users          ListUsersUsers
}

type ListUsersUsers struct {
	TotalCount common.Integer
	PageNumber common.Integer
	PageSize   common.Integer
	List       ListUsersUserList
}

type ListUsersUser struct {
	UserId      common.String
	RamId       common.String
	InstanceId  common.String
	Primary     bool
	Roles       ListUsersRoleList
	SkillLevels ListUsersSkillLevelList
	Detail      ListUsersDetail
}

type ListUsersRole struct {
	RoleId          common.String
	InstanceId      common.String
	RoleName        common.String
	RoleDescription common.String
}

type ListUsersSkillLevel struct {
	SkillLevelId common.String
	Level        common.Integer
	Skill        ListUsersSkill
}

type ListUsersSkill struct {
	SkillGroupId          common.String
	InstanceId            common.String
	SkillGroupName        common.String
	SkillGroupDescription common.String
}

type ListUsersDetail struct {
	LoginName   common.String
	DisplayName common.String
	Phone       common.String
	Email       common.String
	Department  common.String
}

type ListUsersUserList []ListUsersUser

func (list *ListUsersUserList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListUsersUser)
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

type ListUsersRoleList []ListUsersRole

func (list *ListUsersRoleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListUsersRole)
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

type ListUsersSkillLevelList []ListUsersSkillLevel

func (list *ListUsersSkillLevelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListUsersSkillLevel)
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
