package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId      string
	Success        bool
	Code           string
	Message        string
	HttpStatusCode int
	Data           ListRealTimeAgentUserList
}

type ListRealTimeAgentUser struct {
	RamId       string
	DisplayName string
	Phone       string
	Dn          string
	State       string
	StateDesc   string
	SkillLevels ListRealTimeAgentSkillLevelList
}

type ListRealTimeAgentSkillLevel struct {
	SkillLevelId string
	Level        int
	Skill        ListRealTimeAgentSkill
}

type ListRealTimeAgentSkill struct {
	SkillGroupId          string
	InstanceId            string
	SkillGroupName        string
	SkillGroupDescription string
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
