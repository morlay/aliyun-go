package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListSkillGroupsRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

func (req *ListSkillGroupsRequest) Invoke(client *sdk.Client) (resp *ListSkillGroupsResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "ListSkillGroups", "CCC", "")
	resp = &ListSkillGroupsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListSkillGroupsResponse struct {
	responses.BaseResponse
	RequestId      string
	Success        bool
	Code           string
	Message        string
	HttpStatusCode int
	SkillGroups    ListSkillGroupsSkillGroupList
}

type ListSkillGroupsSkillGroup struct {
	SkillGroupId          string
	InstanceId            string
	SkillGroupName        string
	AccSkillGroupName     string
	AccQueueName          string
	SkillGroupDescription string
	UserCount             int
	OutboundPhoneNumbers  ListSkillGroupsPhoneNumberList
}

type ListSkillGroupsPhoneNumber struct {
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

type ListSkillGroupsSkillGroupList []ListSkillGroupsSkillGroup

func (list *ListSkillGroupsSkillGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListSkillGroupsSkillGroup)
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

type ListSkillGroupsPhoneNumberList []ListSkillGroupsPhoneNumber

func (list *ListSkillGroupsPhoneNumberList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListSkillGroupsPhoneNumber)
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
