package qualitycheck

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetRuleDetailRequest struct {
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (r GetRuleDetailRequest) Invoke(client *sdk.Client) (response *GetRuleDetailResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetRuleDetailRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "GetRuleDetail", "", "")

	resp := struct {
		*responses.BaseResponse
		GetRuleDetailResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetRuleDetailResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetRuleDetailResponse struct {
	RequestId string
	Success   bool
	Code      string
	Message   string
	Data      GetRuleDetailData
}

type GetRuleDetailData struct {
	Conditions GetRuleDetailConditionBasicInfoList
	Rules      GetRuleDetailRuleBasicInfoList
}

type GetRuleDetailConditionBasicInfo struct {
	ConditionInfoCid string
	OperLambda       string
	Operators        GetRuleDetailOperatorBasicInfoList
	CheckRange       GetRuleDetailCheckRange
}

type GetRuleDetailOperatorBasicInfo struct {
	Oid      string
	Type     string
	OperName string
	Param    GetRuleDetailParam
}

type GetRuleDetailParam struct {
	Regex         string
	Phrase        string
	Interval      int
	Threshold     float32
	InSentence    bool
	Target        int
	FromEnd       bool
	DifferentRole bool
	TargetRole    string
	OperKeyWords  GetRuleDetailOperKeyWordList
	References    GetRuleDetailReferenceList
}

type GetRuleDetailCheckRange struct {
	Role   string
	Anchor GetRuleDetailAnchor
	Range  GetRuleDetailRange
}

type GetRuleDetailAnchor struct {
	AnchorCid string
	Location  string
	HitTime   int
}

type GetRuleDetailRange struct {
	From int
	To   int
}

type GetRuleDetailRuleBasicInfo struct {
	Rid        string
	RuleLambda string
	Triggers   GetRuleDetailTriggerList
}

type GetRuleDetailConditionBasicInfoList []GetRuleDetailConditionBasicInfo

func (list *GetRuleDetailConditionBasicInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetRuleDetailConditionBasicInfo)
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

type GetRuleDetailRuleBasicInfoList []GetRuleDetailRuleBasicInfo

func (list *GetRuleDetailRuleBasicInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetRuleDetailRuleBasicInfo)
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

type GetRuleDetailOperatorBasicInfoList []GetRuleDetailOperatorBasicInfo

func (list *GetRuleDetailOperatorBasicInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetRuleDetailOperatorBasicInfo)
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

type GetRuleDetailOperKeyWordList []string

func (list *GetRuleDetailOperKeyWordList) UnmarshalJSON(data []byte) error {
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

type GetRuleDetailReferenceList []string

func (list *GetRuleDetailReferenceList) UnmarshalJSON(data []byte) error {
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

type GetRuleDetailTriggerList []string

func (list *GetRuleDetailTriggerList) UnmarshalJSON(data []byte) error {
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
