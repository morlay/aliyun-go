package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetUserRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	UserId     string `position:"Query" name:"UserId"`
}

func (req *GetUserRequest) Invoke(client *sdk.Client) (resp *GetUserResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "GetUser", "ccc", "")
	resp = &GetUserResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetUserResponse struct {
	responses.BaseResponse
	RequestId      common.String
	Success        bool
	Code           common.String
	Message        common.String
	HttpStatusCode common.Integer
	User           GetUserUser
}

type GetUserUser struct {
	UserId      common.String
	RamId       common.String
	InstanceId  common.String
	Roles       GetUserRoleList
	SkillLevels GetUserSkillLevelList
	Detail      GetUserDetail
}

type GetUserRole struct {
	RoleId          common.String
	InstanceId      common.String
	RoleName        common.String
	RoleDescription common.String
}

type GetUserSkillLevel struct {
	SkillLevelId common.String
	Level        common.Integer
	Skill        GetUserSkill
}

type GetUserSkill struct {
	SkillGroupId          common.String
	InstanceId            common.String
	SkillGroupName        common.String
	SkillGroupDescription common.String
}

type GetUserDetail struct {
	LoginName   common.String
	DisplayName common.String
	Phone       common.String
	Email       common.String
	Department  common.String
}

type GetUserRoleList []GetUserRole

func (list *GetUserRoleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetUserRole)
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

type GetUserSkillLevelList []GetUserSkillLevel

func (list *GetUserSkillLevelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetUserSkillLevel)
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
