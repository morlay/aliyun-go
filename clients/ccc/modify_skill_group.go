package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifySkillGroupRequest struct {
	requests.RpcRequest
	SkillLevels            *ModifySkillGroupSkillLevelList            `position:"Query" type:"Repeated" name:"SkillLevel"`
	InstanceId             string                                     `position:"Query" name:"InstanceId"`
	OutboundPhoneNumberIds *ModifySkillGroupOutboundPhoneNumberIdList `position:"Query" type:"Repeated" name:"OutboundPhoneNumberId"`
	SkillGroupId           string                                     `position:"Query" name:"SkillGroupId"`
	Name                   string                                     `position:"Query" name:"Name"`
	Description            string                                     `position:"Query" name:"Description"`
	UserIds                *ModifySkillGroupUserIdList                `position:"Query" type:"Repeated" name:"UserId"`
}

func (req *ModifySkillGroupRequest) Invoke(client *sdk.Client) (resp *ModifySkillGroupResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "ModifySkillGroup", "ccc", "")
	resp = &ModifySkillGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifySkillGroupResponse struct {
	responses.BaseResponse
	RequestId      string
	Success        bool
	Code           string
	Message        string
	HttpStatusCode int
}

type ModifySkillGroupSkillLevelList []int

func (list *ModifySkillGroupSkillLevelList) UnmarshalJSON(data []byte) error {
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

type ModifySkillGroupOutboundPhoneNumberIdList []string

func (list *ModifySkillGroupOutboundPhoneNumberIdList) UnmarshalJSON(data []byte) error {
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

type ModifySkillGroupUserIdList []string

func (list *ModifySkillGroupUserIdList) UnmarshalJSON(data []byte) error {
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
