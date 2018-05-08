package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateUserRequest struct {
	requests.RpcRequest
	SkillLevels   *CreateUserSkillLevelList   `position:"Query" type:"Repeated" name:"SkillLevel"`
	InstanceId    string                      `position:"Query" name:"InstanceId"`
	LoginName     string                      `position:"Query" name:"LoginName"`
	Phone         string                      `position:"Query" name:"Phone"`
	RoleIds       *CreateUserRoleIdList       `position:"Query" type:"Repeated" name:"RoleId"`
	DisplayName   string                      `position:"Query" name:"DisplayName"`
	SkillGroupIds *CreateUserSkillGroupIdList `position:"Query" type:"Repeated" name:"SkillGroupId"`
	Email         string                      `position:"Query" name:"Email"`
}

func (req *CreateUserRequest) Invoke(client *sdk.Client) (resp *CreateUserResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "CreateUser", "ccc", "")
	resp = &CreateUserResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateUserResponse struct {
	responses.BaseResponse
	RequestId      common.String
	Success        bool
	Code           common.String
	Message        common.String
	HttpStatusCode common.Integer
	UserId         common.String
}

type CreateUserSkillLevelList []int

func (list *CreateUserSkillLevelList) UnmarshalJSON(data []byte) error {
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

type CreateUserRoleIdList []string

func (list *CreateUserRoleIdList) UnmarshalJSON(data []byte) error {
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

type CreateUserSkillGroupIdList []string

func (list *CreateUserSkillGroupIdList) UnmarshalJSON(data []byte) error {
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
