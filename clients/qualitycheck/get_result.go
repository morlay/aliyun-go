package qualitycheck

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId string
	Success   bool
	Code      string
	Message   string
	Count     int
	Data      GetResultResultInfoList
}

type GetResultResultInfo struct {
	Tid             string
	AsrMsg          string
	Score           int
	ReviewStatus    int
	HitId           string
	Rules           GetResultRuleHitInfoList
	HandScoreIdList GetResultHandScoreIdListList
}

type GetResultRuleHitInfo struct {
	HitStatus     int
	Rid           string
	Hit           GetResultConditionHitInfoList
	ConditionInfo GetResultConditionBasicInfoList
}

type GetResultConditionHitInfo struct {
	HitKeyWords GetResultHitKeyWordList
	HitCids     GetResultHitCidList
	Phrase      GetResultPhrase
}

type GetResultHitKeyWord struct {
	Val  string
	Pid  int
	From int
	To   int
	Tid  string
}

type GetResultPhrase struct {
	Role       string
	Identity   string
	Words      string
	Begin      int64
	End        int64
	BeginTime  string
	HourMinSec string
}

type GetResultConditionBasicInfo struct {
	ConditionInfoCid string
	Lambda           string
	Operators        GetResultOperatorBasicInfoList
	CheckRange       GetResultCheckRange
}

type GetResultOperatorBasicInfo struct {
	Oid   string
	Type  string
	Name  string
	Param GetResultParam
}

type GetResultParam struct {
	Regex         string
	Phrase        string
	Interval      int
	Threshold     float32
	InSentence    bool
	Target        int
	FromEnd       bool
	DifferentRole bool
	TargetRole    string
	OperKeyWords  GetResultOperKeyWordList
	References    GetResultReferenceList
}

type GetResultCheckRange struct {
	Role   string
	Anchor GetResultAnchor
	Range  GetResultRange
}

type GetResultAnchor struct {
	AnchorCid string
	Location  string
	HitTime   int
}

type GetResultRange struct {
	From int
	To   int
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

type GetResultHandScoreIdListList []string

func (list *GetResultHandScoreIdListList) UnmarshalJSON(data []byte) error {
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

type GetResultHitCidList []string

func (list *GetResultHitCidList) UnmarshalJSON(data []byte) error {
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

type GetResultOperKeyWordList []string

func (list *GetResultOperKeyWordList) UnmarshalJSON(data []byte) error {
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

type GetResultReferenceList []string

func (list *GetResultReferenceList) UnmarshalJSON(data []byte) error {
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
