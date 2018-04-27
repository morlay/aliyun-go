package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyUserRequest struct {
	requests.RpcRequest
	SkillLevels   *ModifyUserSkillLevelList   `position:"Query" type:"Repeated" name:"SkillLevel"`
	InstanceId    string                      `position:"Query" name:"InstanceId"`
	Phone         string                      `position:"Query" name:"Phone"`
	RoleIds       *ModifyUserRoleIdList       `position:"Query" type:"Repeated" name:"RoleId"`
	DisplayName   string                      `position:"Query" name:"DisplayName"`
	SkillGroupIds *ModifyUserSkillGroupIdList `position:"Query" type:"Repeated" name:"SkillGroupId"`
	UserId        string                      `position:"Query" name:"UserId"`
	Email         string                      `position:"Query" name:"Email"`
}

func (req *ModifyUserRequest) Invoke(client *sdk.Client) (resp *ModifyUserResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "ModifyUser", "ccc", "")
	resp = &ModifyUserResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyUserResponse struct {
	responses.BaseResponse
	RequestId      string
	Success        bool
	Code           string
	Message        string
	HttpStatusCode int
}

type ModifyUserSkillLevelList []int

func (list *ModifyUserSkillLevelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int)
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

type ModifyUserRoleIdList []string

func (list *ModifyUserRoleIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type ModifyUserSkillGroupIdList []string

func (list *ModifyUserSkillGroupIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
