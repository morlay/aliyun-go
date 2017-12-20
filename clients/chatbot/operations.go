package chatbot

import (
	"encoding/json"
)

func (c *ChatbotClient) Chat(req *ChatArgs) (resp *ChatResponse, err error) {
	resp = &ChatResponse{}
	err = c.Request("Chat", req, resp)
	return
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
type ChatArgs struct {
	KnowledgeId string
	SenderId    string
	InstanceId  string
	SenderNick  string
	SessionId   string
	Tag         string
	Utterance   string
}
type ChatResponse struct {
	RequestId string
	SessionId string
	MessageId string
	Tag       string
	Messages  ChatMessageList
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
