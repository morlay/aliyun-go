package chatbot

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ChatRequest struct {
	KnowledgeId string `position:"Query" name:"KnowledgeId"`
	SenderId    string `position:"Query" name:"SenderId"`
	InstanceId  string `position:"Query" name:"InstanceId"`
	SenderNick  string `position:"Query" name:"SenderNick"`
	SessionId   string `position:"Query" name:"SessionId"`
	Tag         string `position:"Query" name:"Tag"`
	Utterance   string `position:"Query" name:"Utterance"`
}

func (r ChatRequest) Invoke(client *sdk.Client) (response *ChatResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ChatRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Chatbot", "2017-10-11", "Chat", "beebot", "")

	resp := struct {
		*responses.BaseResponse
		ChatResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ChatResponse

	err = client.DoAction(&req, &resp)
	return
}

type ChatResponse struct {
	RequestId string
	SessionId string
	MessageId string
	Tag       string
	Messages  ChatMessageList
}

type ChatMessage struct {
	Type       string
	Recommends ChatRecommendList
	Text       ChatText
	Knowledge  ChatKnowledge
}

type ChatRecommend struct {
	KnowledgeId  string
	Title        string
	AnswerSource string
}

type ChatText struct {
	Content      string
	AnswerSource string
}

type ChatKnowledge struct {
	Id           string
	Title        string
	Summary      string
	Content      string
	AnswerSource string
}

type ChatMessageList []ChatMessage

func (list *ChatMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ChatMessage)
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

type ChatRecommendList []ChatRecommend

func (list *ChatRecommendList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ChatRecommend)
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
