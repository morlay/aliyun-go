package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AssignUsersRequest struct {
	requests.RpcRequest
	UserRamIds    *AssignUsersUserRamIdList    `position:"Query" type:"Repeated" name:"UserRamId"`
	SkillLevels   *AssignUsersSkillLevelList   `position:"Query" type:"Repeated" name:"SkillLevel"`
	InstanceId    string                       `position:"Query" name:"InstanceId"`
	RoleIds       *AssignUsersRoleIdList       `position:"Query" type:"Repeated" name:"RoleId"`
	SkillGroupIds *AssignUsersSkillGroupIdList `position:"Query" type:"Repeated" name:"SkillGroupId"`
}

func (req *AssignUsersRequest) Invoke(client *sdk.Client) (resp *AssignUsersResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "AssignUsers", "ccc", "")
	resp = &AssignUsersResponse{}
	err = client.DoAction(req, resp)
	return
}

type AssignUsersResponse struct {
	responses.BaseResponse
	RequestId      string
	Success        bool
	Code           string
	Message        string
	HttpStatusCode int
}

type AssignUsersUserRamIdList []string

func (list *AssignUsersUserRamIdList) UnmarshalJSON(data []byte) error {
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

type AssignUsersSkillLevelList []int

func (list *AssignUsersSkillLevelList) UnmarshalJSON(data []byte) error {
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

type AssignUsersRoleIdList []string

func (list *AssignUsersRoleIdList) UnmarshalJSON(data []byte) error {
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

type AssignUsersSkillGroupIdList []string

func (list *AssignUsersSkillGroupIdList) UnmarshalJSON(data []byte) error {
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
