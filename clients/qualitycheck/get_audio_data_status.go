package qualitycheck

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetAudioDataStatusRequest struct {
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (r GetAudioDataStatusRequest) Invoke(client *sdk.Client) (response *GetAudioDataStatusResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetAudioDataStatusRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "GetAudioDataStatus", "", "")

	resp := struct {
		*responses.BaseResponse
		GetAudioDataStatusResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetAudioDataStatusResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetAudioDataStatusResponse struct {
	RequestId string
	Success   bool
	Code      string
	Message   string
	Count     int
	Data      GetAudioDataStatusTaskAsrResultList
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
