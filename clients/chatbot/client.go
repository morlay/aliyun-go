package chatbot

import "github.com/morlay/aliyun-go/core"

func NewChatbotClient(key string, secret string, regionId string) *ChatbotClient {
	return &ChatbotClient{
		Client: core.Client{
			Endpoint:        "https://chatbot.aliyuncs.com",
			Version:         "2017-10-11",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type ChatbotClient struct {
	core.Client
}
