package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListRealTimeAgentRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

func (req *ListRealTimeAgentRequest) Invoke(client *sdk.Client) (resp *ListRealTimeAgentResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "ListRealTimeAgent", "ccc", "")
	resp = &ListRealTimeAgentResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListRealTimeAgentResponse struct {
	responses.BaseResponse
	RequestId      common.String
	Success        bool
	Code           common.String
	Message        common.String
	HttpStatusCode common.Integer
	Data           ListRealTimeAgentUserList
}

type ListRealTimeAgentUser struct {
	RamId       common.String
	DisplayName common.String
	Phone       common.String
	Dn          common.String
	State       common.String
	StateDesc   common.String
	SkillLevels ListRealTimeAgentSkillLevelList
}

type ListRealTimeAgentSkillLevel struct {
	SkillLevelId common.String
	Level        common.Integer
	Skill        ListRealTimeAgentSkill
}

type ListRealTimeAgentSkill struct {
	SkillGroupId          common.String
	InstanceId            common.String
	SkillGroupName        common.String
	SkillGroupDescription common.String
}

type ListRealTimeAgentUserList []ListRealTimeAgentUser

func (list *ListRealTimeAgentUserList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListRealTimeAgentUser)
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

type ListRealTimeAgentSkillLevelList []ListRealTimeAgentSkillLevel

func (list *ListRealTimeAgentSkillLevelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListRealTimeAgentSkillLevel)
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
