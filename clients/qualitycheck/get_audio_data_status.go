package qualitycheck

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetAudioDataStatusRequest struct {
	requests.RpcRequest
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (req *GetAudioDataStatusRequest) Invoke(client *sdk.Client) (resp *GetAudioDataStatusResponse, err error) {
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "GetAudioDataStatus", "", "")
	resp = &GetAudioDataStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetAudioDataStatusResponse struct {
	responses.BaseResponse
	RequestId     common.String
	Success       bool
	Code          common.String
	Message       common.String
	Count         common.Integer
	OverallStatus common.Integer
	Data          GetAudioDataStatusTaskAsrResultList
}

type GetAudioDataStatusTaskAsrResult struct {
	Tid        common.String
	StatusCode common.Integer
	StatusMsg  common.String
	AsrResult  GetAudioDataStatusAsrResult
}

type GetAudioDataStatusAsrResult struct {
	Asrstatus        common.String
	AsrStatusCode    common.String
	ErrorMessage     common.String
	Duration         common.Long
	InteractiveCount common.Integer
	SentenceResults  GetAudioDataStatusSentenceResultList
	ServiceEvStat    GetAudioDataStatusServiceEvStat
	ClientEvStat     GetAudioDataStatusClientEvStat
	ServiceSrStat    GetAudioDataStatusServiceSrStat
	ClientSrStat     GetAudioDataStatusClientSrStat
}

type GetAudioDataStatusSentenceResult struct {
	BeginTime       common.Integer
	EndTime         common.Integer
	ChannelId       common.Integer
	Text            common.String
	EmotionValue    common.Integer
	SilenceDuration common.Integer
	SpeechRate      common.Integer
	SpeakerId       common.String
	AgentId         common.String
	ChannelKey      common.String
}

type GetAudioDataStatusServiceEvStat struct {
	Srole            common.Integer
	SmaxEmotionValue common.Float
	SminEmotionValue common.Float
	SavgEmotionValue common.Float
}

type GetAudioDataStatusClientEvStat struct {
	Crole            common.Integer
	CmaxEmotionValue common.Float
	CminEmotionValue common.Float
	CavgEmotionValue common.Float
}

type GetAudioDataStatusServiceSrStat struct {
	Srole          common.Integer
	SmaxSpeechRate common.Float
	SminSpeechRate common.Float
	SavgSpeechRate common.Float
}

type GetAudioDataStatusClientSrStat struct {
	Crole          common.Integer
	CmaxSpeechRate common.Float
	CminSpeechRate common.Float
	CavgSpeechRate common.Float
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
