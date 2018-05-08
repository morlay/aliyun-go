package chatbot

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ChatRequest struct {
	requests.RpcRequest
	KnowledgeId string `position:"Query" name:"KnowledgeId"`
	SenderId    string `position:"Query" name:"SenderId"`
	InstanceId  string `position:"Query" name:"InstanceId"`
	SenderNick  string `position:"Query" name:"SenderNick"`
	SessionId   string `position:"Query" name:"SessionId"`
	Tag         string `position:"Query" name:"Tag"`
	Utterance   string `position:"Query" name:"Utterance"`
}

func (req *ChatRequest) Invoke(client *sdk.Client) (resp *ChatResponse, err error) {
	req.InitWithApiInfo("Chatbot", "2017-10-11", "Chat", "beebot", "")
	resp = &ChatResponse{}
	err = client.DoAction(req, resp)
	return
}

type ChatResponse struct {
	responses.BaseResponse
	RequestId common.String
	SessionId common.String
	MessageId common.String
	Tag       common.String
	Messages  ChatMessageList
}

type ChatMessage struct {
	Type       common.String
	Recommends ChatRecommendList
	Text       ChatText
	Knowledge  ChatKnowledge
}

type ChatRecommend struct {
	KnowledgeId  common.String
	Title        common.String
	AnswerSource common.String
}

type ChatText struct {
	Content      common.String
	AnswerSource common.String
}

type ChatKnowledge struct {
	Id           common.String
	Title        common.String
	Summary      common.String
	Content      common.String
	AnswerSource common.String
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
