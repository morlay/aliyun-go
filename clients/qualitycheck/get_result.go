package qualitycheck

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetResultRequest struct {
	requests.RpcRequest
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (req *GetResultRequest) Invoke(client *sdk.Client) (resp *GetResultResponse, err error) {
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "GetResult", "", "")
	resp = &GetResultResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetResultResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Code      common.String
	Message   common.String
	Count     common.Integer
	Data      GetResultResultInfoList
}

type GetResultResultInfo struct {
	Tid             common.String
	AsrMsg          common.String
	Score           common.Integer
	ReviewStatus    common.Integer
	HitId           common.String
	Rules           GetResultRuleHitInfoList
	HandScoreIdList GetResultHandScoreIdListList
}

type GetResultRuleHitInfo struct {
	HitStatus     common.Integer
	Rid           common.String
	Hit           GetResultConditionHitInfoList
	ConditionInfo GetResultConditionBasicInfoList
}

type GetResultConditionHitInfo struct {
	HitKeyWords GetResultHitKeyWordList
	HitCids     GetResultHitCidList
	Phrase      GetResultPhrase
}

type GetResultHitKeyWord struct {
	Val  common.String
	Pid  common.Integer
	From common.Integer
	To   common.Integer
	Tid  common.String
}

type GetResultPhrase struct {
	Role       common.String
	Identity   common.String
	Words      common.String
	Begin      common.Long
	End        common.Long
	BeginTime  common.String
	HourMinSec common.String
}

type GetResultConditionBasicInfo struct {
	ConditionInfoCid common.String
	Lambda           common.String
	Operators        GetResultOperatorBasicInfoList
	CheckRange       GetResultCheckRange
}

type GetResultOperatorBasicInfo struct {
	Oid   common.String
	Type  common.String
	Name  common.String
	Param GetResultParam
}

type GetResultParam struct {
	Regex         common.String
	Phrase        common.String
	Interval      common.Integer
	Threshold     common.Float
	InSentence    bool
	Target        common.Integer
	FromEnd       bool
	DifferentRole bool
	TargetRole    common.String
	OperKeyWords  GetResultOperKeyWordList
	References    GetResultReferenceList
}

type GetResultCheckRange struct {
	Role   common.String
	Anchor GetResultAnchor
	Range  GetResultRange
}

type GetResultAnchor struct {
	AnchorCid common.String
	Location  common.String
	HitTime   common.Integer
}

type GetResultRange struct {
	From common.Integer
	To   common.Integer
}

type GetResultResultInfoList []GetResultResultInfo

func (list *GetResultResultInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetResultResultInfo)
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

type GetResultRuleHitInfoList []GetResultRuleHitInfo

func (list *GetResultRuleHitInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetResultRuleHitInfo)
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

type GetResultHandScoreIdListList []common.String

func (list *GetResultHandScoreIdListList) UnmarshalJSON(data []byte) error {
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

type GetResultConditionHitInfoList []GetResultConditionHitInfo

func (list *GetResultConditionHitInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetResultConditionHitInfo)
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

type GetResultConditionBasicInfoList []GetResultConditionBasicInfo

func (list *GetResultConditionBasicInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetResultConditionBasicInfo)
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

type GetResultHitKeyWordList []GetResultHitKeyWord

func (list *GetResultHitKeyWordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetResultHitKeyWord)
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

type GetResultHitCidList []common.String

func (list *GetResultHitCidList) UnmarshalJSON(data []byte) error {
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

type GetResultOperatorBasicInfoList []GetResultOperatorBasicInfo

func (list *GetResultOperatorBasicInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetResultOperatorBasicInfo)
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

type GetResultOperKeyWordList []common.String

func (list *GetResultOperKeyWordList) UnmarshalJSON(data []byte) error {
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

type GetResultReferenceList []common.String

func (list *GetResultReferenceList) UnmarshalJSON(data []byte) error {
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
