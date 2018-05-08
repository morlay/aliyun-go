package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListSkillGroupsRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

func (req *ListSkillGroupsRequest) Invoke(client *sdk.Client) (resp *ListSkillGroupsResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "ListSkillGroups", "ccc", "")
	resp = &ListSkillGroupsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListSkillGroupsResponse struct {
	responses.BaseResponse
	RequestId      common.String
	Success        bool
	Code           common.String
	Message        common.String
	HttpStatusCode common.Integer
	SkillGroups    ListSkillGroupsSkillGroupList
}

type ListSkillGroupsSkillGroup struct {
	SkillGroupId          common.String
	InstanceId            common.String
	SkillGroupName        common.String
	AccSkillGroupName     common.String
	AccQueueName          common.String
	SkillGroupDescription common.String
	UserCount             common.Integer
	OutboundPhoneNumbers  ListSkillGroupsPhoneNumberList
}

type ListSkillGroupsPhoneNumber struct {
	PhoneNumberId          common.String
	InstanceId             common.String
	Number                 common.String
	PhoneNumberDescription common.String
	TestOnly               bool
	RemainingTime          common.Integer
	AllowOutbound          bool
	Usage                  common.String
	Trunks                 common.Integer
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
