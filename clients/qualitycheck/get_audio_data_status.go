package qualitycheck

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId     string
	Success       bool
	Code          string
	Message       string
	Count         int
	OverallStatus int
	Data          GetAudioDataStatusTaskAsrResultList
}

type GetAudioDataStatusTaskAsrResult struct {
	Tid        string
	StatusCode int
	StatusMsg  string
	AsrResult  GetAudioDataStatusAsrResult
}

type GetAudioDataStatusAsrResult struct {
	Asrstatus        string
	AsrStatusCode    string
	ErrorMessage     string
	Duration         int64
	InteractiveCount int
	SentenceResults  GetAudioDataStatusSentenceResultList
	ServiceEvStat    GetAudioDataStatusServiceEvStat
	ClientEvStat     GetAudioDataStatusClientEvStat
	ServiceSrStat    GetAudioDataStatusServiceSrStat
	ClientSrStat     GetAudioDataStatusClientSrStat
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

type GetAudioDataStatusServiceEvStat struct {
	Srole            int
	SmaxEmotionValue float32
	SminEmotionValue float32
	SavgEmotionValue float32
}

type GetAudioDataStatusClientEvStat struct {
	Crole            int
	CmaxEmotionValue float32
	CminEmotionValue float32
	CavgEmotionValue float32
}

type GetAudioDataStatusServiceSrStat struct {
	Srole          int
	SmaxSpeechRate float32
	SminSpeechRate float32
	SavgSpeechRate float32
}

type GetAudioDataStatusClientSrStat struct {
	Crole          int
	CmaxSpeechRate float32
	CminSpeechRate float32
	CavgSpeechRate float32
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
