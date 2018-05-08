package qualitycheck

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetAccAsrResultRequest struct {
	requests.RpcRequest
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (req *GetAccAsrResultRequest) Invoke(client *sdk.Client) (resp *GetAccAsrResultResponse, err error) {
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "GetAccAsrResult", "", "")
	resp = &GetAccAsrResultResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetAccAsrResultResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Code      common.String
	Message   common.String
	Count     common.Integer
	Data      GetAccAsrResultAccAsrSentenceResultList
}

type GetAccAsrResultAccAsrSentenceResult struct {
	RecordId         common.String
	Status           common.String
	StatusCode       common.String
	ErrorMessage     common.String
	Duration         common.Long
	InteractiveCount common.Integer
	Results          GetAccAsrResultSentenceResultList
	ServiceEvStat    GetAccAsrResultServiceEvStat
	ClientEvStat     GetAccAsrResultClientEvStat
	ServiceSrStat    GetAccAsrResultServiceSrStat
	ClientSrStat     GetAccAsrResultClientSrStat
}

type GetAccAsrResultSentenceResult struct {
	BeginTime       common.Long
	EndTime         common.Long
	ChannelId       common.Integer
	Text            common.String
	EmotionValue    common.Integer
	SilenceDuration common.Integer
	SpeechRate      common.Integer
	SpeakerId       common.String
	AgentId         common.String
	ChannelKey      common.String
}

type GetAccAsrResultServiceEvStat struct {
	Srole            common.Integer
	SmaxEmotionValue common.Float
	SminEmotionValue common.Float
	SavgEmotionValue common.Float
}

type GetAccAsrResultClientEvStat struct {
	Crole            common.Integer
	CmaxEmotionValue common.Float
	CminEmotionValue common.Float
	CavgEmotionValue common.Float
}

type GetAccAsrResultServiceSrStat struct {
	Srole          common.Integer
	SmaxSpeechRate common.Float
	SminSpeechRate common.Float
	SavgSpeechRate common.Float
}

type GetAccAsrResultClientSrStat struct {
	Crole          common.Integer
	CmaxSpeechRate common.Float
	CminSpeechRate common.Float
	CavgSpeechRate common.Float
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
