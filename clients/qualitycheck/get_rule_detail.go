package qualitycheck

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetRuleDetailRequest struct {
	requests.RpcRequest
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (req *GetRuleDetailRequest) Invoke(client *sdk.Client) (resp *GetRuleDetailResponse, err error) {
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "GetRuleDetail", "", "")
	resp = &GetRuleDetailResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetRuleDetailResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Code      common.String
	Message   common.String
	Data      GetRuleDetailData
}

type GetRuleDetailData struct {
	Conditions GetRuleDetailConditionBasicInfoList
	Rules      GetRuleDetailRuleBasicInfoList
}

type GetRuleDetailConditionBasicInfo struct {
	ConditionInfoCid common.String
	OperLambda       common.String
	Operators        GetRuleDetailOperatorBasicInfoList
	CheckRange       GetRuleDetailCheckRange
}

type GetRuleDetailOperatorBasicInfo struct {
	Oid      common.String
	Type     common.String
	OperName common.String
	Param    GetRuleDetailParam
}

type GetRuleDetailParam struct {
	Regex         common.String
	Phrase        common.String
	Interval      common.Integer
	Threshold     common.Float
	InSentence    bool
	Target        common.Integer
	FromEnd       bool
	DifferentRole bool
	TargetRole    common.String
	OperKeyWords  GetRuleDetailOperKeyWordList
	References    GetRuleDetailReferenceList
}

type GetRuleDetailCheckRange struct {
	Role   common.String
	Anchor GetRuleDetailAnchor
	Range  GetRuleDetailRange
}

type GetRuleDetailAnchor struct {
	AnchorCid common.String
	Location  common.String
	HitTime   common.Integer
}

type GetRuleDetailRange struct {
	From common.Integer
	To   common.Integer
}

type GetRuleDetailRuleBasicInfo struct {
	Rid                common.String
	RuleLambda         common.String
	BusinessCategories GetRuleDetailBusinessCategoryBasicInfoList
	Triggers           GetRuleDetailTriggerList
}

type GetRuleDetailBusinessCategoryBasicInfo struct {
	Bid          common.Integer
	ServiceType  common.Integer
	BusinessName common.String
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

type GetRuleDetailOperKeyWordList []common.String

func (list *GetRuleDetailOperKeyWordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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

type GetRuleDetailReferenceList []common.String

func (list *GetRuleDetailReferenceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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

type GetRuleDetailBusinessCategoryBasicInfoList []GetRuleDetailBusinessCategoryBasicInfo

func (list *GetRuleDetailBusinessCategoryBasicInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetRuleDetailBusinessCategoryBasicInfo)
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

type GetRuleDetailTriggerList []common.String

func (list *GetRuleDetailTriggerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
