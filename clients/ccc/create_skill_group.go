package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateSkillGroupRequest struct {
	requests.RpcRequest
	SkillLevels            *CreateSkillGroupSkillLevelList            `position:"Query" type:"Repeated" name:"SkillLevel"`
	InstanceId             string                                     `position:"Query" name:"InstanceId"`
	OutboundPhoneNumberIds *CreateSkillGroupOutboundPhoneNumberIdList `position:"Query" type:"Repeated" name:"OutboundPhoneNumberId"`
	Name                   string                                     `position:"Query" name:"Name"`
	Description            string                                     `position:"Query" name:"Description"`
	UserIds                *CreateSkillGroupUserIdList                `position:"Query" type:"Repeated" name:"UserId"`
}

func (req *CreateSkillGroupRequest) Invoke(client *sdk.Client) (resp *CreateSkillGroupResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "CreateSkillGroup", "CCC", "")
	resp = &CreateSkillGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateSkillGroupResponse struct {
	responses.BaseResponse
	RequestId      string
	Success        bool
	Code           string
	Message        string
	HttpStatusCode int
	SkillGroupId   string
}

type CreateSkillGroupSkillLevelList []int

func (list *CreateSkillGroupSkillLevelList) UnmarshalJSON(data []byte) error {
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

type CreateSkillGroupOutboundPhoneNumberIdList []string

func (list *CreateSkillGroupOutboundPhoneNumberIdList) UnmarshalJSON(data []byte) error {
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

type CreateSkillGroupUserIdList []string

func (list *CreateSkillGroupUserIdList) UnmarshalJSON(data []byte) error {
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
