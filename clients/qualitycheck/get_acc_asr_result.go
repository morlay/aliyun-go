package qualitycheck

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId string
	Success   bool
	Code      string
	Message   string
	Count     int
	Data      GetAccAsrResultAccAsrSentenceResultList
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
