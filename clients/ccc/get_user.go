package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetUserRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	UserId     string `position:"Query" name:"UserId"`
}

func (r GetUserRequest) Invoke(client *sdk.Client) (response *GetUserResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetUserRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CCC", "2017-07-05", "GetUser", "ccc", "")

	resp := struct {
		*responses.BaseResponse
		GetUserResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetUserResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetUserResponse struct {
	RequestId      string
	Success        bool
	Code           string
	Message        string
	HttpStatusCode int
	User           GetUserUser
}

type GetUserUser struct {
	UserId      string
	RamId       string
	InstanceId  string
	Roles       GetUserRoleList
	SkillLevels GetUserSkillLevelList
	Detail      GetUserDetail
}

type GetUserRole struct {
	RoleId          string
	InstanceId      string
	RoleName        string
	RoleDescription string
}

type GetUserSkillLevel struct {
	SkillLevelId string
	Level        int
	Skill        GetUserSkill
}

type GetUserSkill struct {
	SkillGroupId          string
	InstanceId            string
	SkillGroupName        string
	SkillGroupDescription string
}

type GetUserDetail struct {
	LoginName   string
	DisplayName string
	Phone       string
	Email       string
	Department  string
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
