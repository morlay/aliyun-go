package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListUsersRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	PageSize   int    `position:"Query" name:"PageSize"`
	PageNumber int    `position:"Query" name:"PageNumber"`
}

func (req *ListUsersRequest) Invoke(client *sdk.Client) (resp *ListUsersResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "ListUsers", "CCC", "")
	resp = &ListUsersResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListUsersResponse struct {
	responses.BaseResponse
	RequestId      string
	Success        bool
	Code           string
	Message        string
	HttpStatusCode int
	Users          ListUsersUsers
}

type ListUsersUsers struct {
	TotalCount int
	PageNumber int
	PageSize   int
	List       ListUsersUserList
}

type ListUsersUser struct {
	UserId      string
	RamId       string
	InstanceId  string
	Primary     bool
	Roles       ListUsersRoleList
	SkillLevels ListUsersSkillLevelList
	Detail      ListUsersDetail
}

type ListUsersRole struct {
	RoleId          string
	InstanceId      string
	RoleName        string
	RoleDescription string
}

type ListUsersSkillLevel struct {
	SkillLevelId string
	Level        int
	Skill        ListUsersSkill
}

type ListUsersSkill struct {
	SkillGroupId          string
	InstanceId            string
	SkillGroupName        string
	SkillGroupDescription string
}

type ListUsersDetail struct {
	LoginName   string
	DisplayName string
	Phone       string
	Email       string
	Department  string
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
