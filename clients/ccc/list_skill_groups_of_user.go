package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListSkillGroupsOfUserRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	UserId     string `position:"Query" name:"UserId"`
}

func (req *ListSkillGroupsOfUserRequest) Invoke(client *sdk.Client) (resp *ListSkillGroupsOfUserResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "ListSkillGroupsOfUser", "ccc", "")
	resp = &ListSkillGroupsOfUserResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListSkillGroupsOfUserResponse struct {
	responses.BaseResponse
	RequestId      string
	Success        bool
	Code           string
	Message        string
	HttpStatusCode int
	SkillLevels    ListSkillGroupsOfUserSkillLevelList
}

type ListSkillGroupsOfUserSkillLevel struct {
	SkillLevelId string
	Level        int
	Skill        ListSkillGroupsOfUserSkill
}

type ListSkillGroupsOfUserSkill struct {
	SkillGroupId          string
	InstanceId            string
	SkillGroupName        string
	SkillGroupDescription string
	OutboundPhoneNumbers  ListSkillGroupsOfUserPhoneNumberList
}

type ListSkillGroupsOfUserPhoneNumber struct {
	PhoneNumberId          string
	InstanceId             string
	Number                 string
	PhoneNumberDescription string
	TestOnly               bool
	RemainingTime          int
	AllowOutbound          bool
	Usage                  string
	Trunks                 int
}

type ListSkillGroupsOfUserSkillLevelList []ListSkillGroupsOfUserSkillLevel

func (list *ListSkillGroupsOfUserSkillLevelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListSkillGroupsOfUserSkillLevel)
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

type ListSkillGroupsOfUserPhoneNumberList []ListSkillGroupsOfUserPhoneNumber

func (list *ListSkillGroupsOfUserPhoneNumberList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListSkillGroupsOfUserPhoneNumber)
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
