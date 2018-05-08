package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId      common.String
	Success        bool
	Code           common.String
	Message        common.String
	HttpStatusCode common.Integer
	SkillLevels    ListSkillGroupsOfUserSkillLevelList
}

type ListSkillGroupsOfUserSkillLevel struct {
	SkillLevelId common.String
	Level        common.Integer
	Skill        ListSkillGroupsOfUserSkill
}

type ListSkillGroupsOfUserSkill struct {
	SkillGroupId          common.String
	InstanceId            common.String
	SkillGroupName        common.String
	SkillGroupDescription common.String
	OutboundPhoneNumbers  ListSkillGroupsOfUserPhoneNumberList
}

type ListSkillGroupsOfUserPhoneNumber struct {
	PhoneNumberId          common.String
	InstanceId             common.String
	Number                 common.String
	PhoneNumberDescription common.String
	TestOnly               bool
	RemainingTime          common.Integer
	AllowOutbound          bool
	Usage                  common.String
	Trunks                 common.Integer
	Province               common.String
	City                   common.String
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
