package qualitycheck

import (
	"encoding/json"

	"github.com/morlay/aliyun-go/core"
)

func (c *QualitycheckClient) UploadDataWithRules(req *UploadDataWithRulesArgs) (resp *UploadDataWithRulesResponse, err error) {
	resp = &UploadDataWithRulesResponse{}
	err = c.Request("UploadDataWithRules", req, resp)
	return
}

type UploadDataWithRulesArgs struct {
	JsonStr string
}
type UploadDataWithRulesResponse struct {
	RequestId string
	Success   core.Bool
	Code      string
	Message   string
	Data      string
}

func (c *QualitycheckClient) UploadAudioDataWithRules4Pre(req *UploadAudioDataWithRules4PreArgs) (resp *UploadAudioDataWithRules4PreResponse, err error) {
	resp = &UploadAudioDataWithRules4PreResponse{}
	err = c.Request("UploadAudioDataWithRules4Pre", req, resp)
	return
}

type UploadAudioDataWithRules4PreArgs struct {
	JsonStr string
}
type UploadAudioDataWithRules4PreResponse struct {
	RequestId string
	Success   core.Bool
	Code      string
	Message   string
	Data      string
}

func (c *QualitycheckClient) UploadRule(req *UploadRuleArgs) (resp *UploadRuleResponse, err error) {
	resp = &UploadRuleResponse{}
	err = c.Request("UploadRule", req, resp)
	return
}

type UploadRuleArgs struct {
	JsonStr string
}
type UploadRuleResponse struct {
	RequestId string
	Success   core.Bool
	Code      string
	Message   string
	Data      UploadRuleDatumList
}

type UploadRuleDatumList []string

func (list *UploadRuleDatumList) UnmarshalJSON(data []byte) error {
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

func (c *QualitycheckClient) InvalidRule(req *InvalidRuleArgs) (resp *InvalidRuleResponse, err error) {
	resp = &InvalidRuleResponse{}
	err = c.Request("InvalidRule", req, resp)
	return
}

type InvalidRuleArgs struct {
	JsonStr string
}
type InvalidRuleResponse struct {
	RequestId string
	Success   core.Bool
	Code      string
	Message   string
	Data      core.Bool
}

func (c *QualitycheckClient) UploadAudioData(req *UploadAudioDataArgs) (resp *UploadAudioDataResponse, err error) {
	resp = &UploadAudioDataResponse{}
	err = c.Request("UploadAudioData", req, resp)
	return
}

type UploadAudioDataArgs struct {
	JsonStr string
}
type UploadAudioDataResponse struct {
	RequestId string
	Success   core.Bool
	Code      string
	Message   string
	Data      string
}

func (c *QualitycheckClient) UploadAudioDataWithRules(req *UploadAudioDataWithRulesArgs) (resp *UploadAudioDataWithRulesResponse, err error) {
	resp = &UploadAudioDataWithRulesResponse{}
	err = c.Request("UploadAudioDataWithRules", req, resp)
	return
}

type UploadAudioDataWithRulesArgs struct {
	JsonStr string
}
type UploadAudioDataWithRulesResponse struct {
	RequestId string
	Success   core.Bool
	Code      string
	Message   string
	Data      string
}

func (c *QualitycheckClient) RegisterNotice(req *RegisterNoticeArgs) (resp *RegisterNoticeResponse, err error) {
	resp = &RegisterNoticeResponse{}
	err = c.Request("RegisterNotice", req, resp)
	return
}

type RegisterNoticeArgs struct {
	JsonStr string
}
type RegisterNoticeResponse struct {
	RequestId string
	Success   core.Bool
	Code      string
	Message   string
}

func (c *QualitycheckClient) GetResultCount(req *GetResultCountArgs) (resp *GetResultCountResponse, err error) {
	resp = &GetResultCountResponse{}
	err = c.Request("GetResultCount", req, resp)
	return
}

type GetResultCountArgs struct {
	JsonStr string
}
type GetResultCountResponse struct {
	RequestId string
	Success   core.Bool
	Code      string
	Message   string
	Data      int
}

func (c *QualitycheckClient) UpdateOnPurchaseSuccess(req *UpdateOnPurchaseSuccessArgs) (resp *UpdateOnPurchaseSuccessResponse, err error) {
	resp = &UpdateOnPurchaseSuccessResponse{}
	err = c.Request("UpdateOnPurchaseSuccess", req, resp)
	return
}

type UpdateOnPurchaseSuccessArgs struct {
	Data string
}
type UpdateOnPurchaseSuccessResponse struct {
	RequestId string
	Data      string
	Success   core.Bool
	Code      string
	Message   string
}

func (c *QualitycheckClient) GetScoreInfo(req *GetScoreInfoArgs) (resp *GetScoreInfoResponse, err error) {
	resp = &GetScoreInfoResponse{}
	err = c.Request("GetScoreInfo", req, resp)
	return
}

type GetScoreInfoScorePo struct {
	ScoreId    int
	ScoreName  string
	ScoreInfos GetScoreInfoScoreParamList
}

type GetScoreInfoScoreParam struct {
	ScoreNum     int
	ScoreSubId   int
	ScoreSubName string
	ScoreType    int
}
type GetScoreInfoArgs struct {
	JsonStr string
}
type GetScoreInfoResponse struct {
	RequestId string
	Success   core.Bool
	Code      string
	Message   string
	Data      GetScoreInfoScorePoList
}

type GetScoreInfoScoreParamList []GetScoreInfoScoreParam

func (list *GetScoreInfoScoreParamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetScoreInfoScoreParam)
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

type GetScoreInfoScorePoList []GetScoreInfoScorePo

func (list *GetScoreInfoScorePoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetScoreInfoScorePo)
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

func (c *QualitycheckClient) GetRule(req *GetRuleArgs) (resp *GetRuleResponse, err error) {
	resp = &GetRuleResponse{}
	err = c.Request("GetRule", req, resp)
	return
}

type GetRuleData struct {
	Rules GetRuleRuleInfoList
}

type GetRuleRuleInfo struct {
	Rid             string
	RuleLambda      string
	Name            string
	Type            int
	Status          int
	IsDelete        int
	StartTime       string
	EndTime         string
	Weight          string
	IsOnline        int
	CreateEmpid     string
	CreateTime      string
	LastUpdateTime  string
	LastUpdateEmpid string
	Comments        string
	AutoReview      int
	RuleScoreType   int
	ScoreName       string
	ScoreSubName    string
}
type GetRuleArgs struct {
	JsonStr string
}
type GetRuleResponse struct {
	RequestId string
	Success   core.Bool
	Code      string
	Message   string
	Data      GetRuleData
}

type GetRuleRuleInfoList []GetRuleRuleInfo

func (list *GetRuleRuleInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetRuleRuleInfo)
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

func (c *QualitycheckClient) UploadAudioData4Pre(req *UploadAudioData4PreArgs) (resp *UploadAudioData4PreResponse, err error) {
	resp = &UploadAudioData4PreResponse{}
	err = c.Request("UploadAudioData4Pre", req, resp)
	return
}

type UploadAudioData4PreArgs struct {
	JsonStr string
}
type UploadAudioData4PreResponse struct {
	RequestId string
	Success   core.Bool
	Code      string
	Message   string
	Data      string
}

func (c *QualitycheckClient) GetDataSetList(req *GetDataSetListArgs) (resp *GetDataSetListResponse, err error) {
	resp = &GetDataSetListResponse{}
	err = c.Request("GetDataSetList", req, resp)
	return
}

type GetDataSetListDataSet struct {
	SetId         int64
	SetName       string
	SetDomain     string
	SetRoleArn    string
	SetBucketName string
	SetFolderName string
	ChannelType   int
	CreateType    int
}
type GetDataSetListArgs struct {
	JsonStr string
}
type GetDataSetListResponse struct {
	RequestId string
	Success   core.Bool
	Code      string
	Message   string
	Count     int
	Data      GetDataSetListDataSetList
}

type GetDataSetListDataSetList []GetDataSetListDataSet

func (list *GetDataSetListDataSetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetDataSetListDataSet)
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

func (c *QualitycheckClient) GetRuleDetail(req *GetRuleDetailArgs) (resp *GetRuleDetailResponse, err error) {
	resp = &GetRuleDetailResponse{}
	err = c.Request("GetRuleDetail", req, resp)
	return
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
	InSentence    core.Bool
	Target        int
	FromEnd       core.Bool
	DifferentRole core.Bool
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
type GetRuleDetailArgs struct {
	JsonStr string
}
type GetRuleDetailResponse struct {
	RequestId string
	Success   core.Bool
	Code      string
	Message   string
	Data      GetRuleDetailData
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

func (c *QualitycheckClient) SaveReviewResult(req *SaveReviewResultArgs) (resp *SaveReviewResultResponse, err error) {
	resp = &SaveReviewResultResponse{}
	err = c.Request("SaveReviewResult", req, resp)
	return
}

type SaveReviewResultArgs struct {
	JsonStr string
}
type SaveReviewResultResponse struct {
	RequestId string
	Success   core.Bool
	Code      string
	Message   string
	Data      string
}

func (c *QualitycheckClient) GetAudioDataStatus(req *GetAudioDataStatusArgs) (resp *GetAudioDataStatusResponse, err error) {
	resp = &GetAudioDataStatusResponse{}
	err = c.Request("GetAudioDataStatus", req, resp)
	return
}

type GetAudioDataStatusTaskAsrResult struct {
	Tid        string
	StatusCode int
	StatusMsg  string
	AsrResult  GetAudioDataStatusAsrResult
}

type GetAudioDataStatusAsrResult struct {
	Asrstatus       string
	AsrStatusCode   string
	ErrorMessage    string
	SentenceResults GetAudioDataStatusSentenceResultList
}

type GetAudioDataStatusSentenceResult struct {
	BeginTime       int
	EndTime         int
	ChannelId       int
	Text            string
	EmotionValue    int
	SilenceDuration int
	SpeechRate      int
	SpeakerId       string
	AgentId         string
	ChannelKey      string
}
type GetAudioDataStatusArgs struct {
	JsonStr string
}
type GetAudioDataStatusResponse struct {
	RequestId string
	Success   core.Bool
	Code      string
	Message   string
	Count     int
	Data      GetAudioDataStatusTaskAsrResultList
}

type GetAudioDataStatusSentenceResultList []GetAudioDataStatusSentenceResult

func (list *GetAudioDataStatusSentenceResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetAudioDataStatusSentenceResult)
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

type GetAudioDataStatusTaskAsrResultList []GetAudioDataStatusTaskAsrResult

func (list *GetAudioDataStatusTaskAsrResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetAudioDataStatusTaskAsrResult)
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

func (c *QualitycheckClient) UploadData(req *UploadDataArgs) (resp *UploadDataResponse, err error) {
	resp = &UploadDataResponse{}
	err = c.Request("UploadData", req, resp)
	return
}

type UploadDataArgs struct {
	JsonStr string
}
type UploadDataResponse struct {
	RequestId string
	Success   core.Bool
	Code      string
	Message   string
	Data      string
}

func (c *QualitycheckClient) GetResult(req *GetResultArgs) (resp *GetResultResponse, err error) {
	resp = &GetResultResponse{}
	err = c.Request("GetResult", req, resp)
	return
}

type GetResultResultInfo struct {
	Tid             string
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
	InSentence    core.Bool
	Target        int
	FromEnd       core.Bool
	DifferentRole core.Bool
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
type GetResultArgs struct {
	JsonStr string
}
type GetResultResponse struct {
	RequestId string
	Success   core.Bool
	Code      string
	Message   string
	Count     int
	Data      GetResultResultInfoList
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

func (c *QualitycheckClient) GetAccAsrResult(req *GetAccAsrResultArgs) (resp *GetAccAsrResultResponse, err error) {
	resp = &GetAccAsrResultResponse{}
	err = c.Request("GetAccAsrResult", req, resp)
	return
}

type GetAccAsrResultAccAsrSentenceResult struct {
	RecordId     string
	Status       string
	StatusCode   string
	ErrorMessage string
	Results      GetAccAsrResultSentenceResultList
}

type GetAccAsrResultSentenceResult struct {
	BeginTime       int64
	EndTime         int64
	ChannelId       int
	Text            string
	EmotionValue    int
	SilenceDuration int
	SpeechRate      int
	SpeakerId       string
	AgentId         string
	ChannelKey      string
}
type GetAccAsrResultArgs struct {
	JsonStr string
}
type GetAccAsrResultResponse struct {
	RequestId string
	Success   core.Bool
	Code      string
	Message   string
	Count     int
	Data      GetAccAsrResultAccAsrSentenceResultList
}

type GetAccAsrResultSentenceResultList []GetAccAsrResultSentenceResult

func (list *GetAccAsrResultSentenceResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetAccAsrResultSentenceResult)
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

type GetAccAsrResultAccAsrSentenceResultList []GetAccAsrResultAccAsrSentenceResult

func (list *GetAccAsrResultAccAsrSentenceResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetAccAsrResultAccAsrSentenceResult)
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
