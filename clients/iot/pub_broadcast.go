package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type PubBroadcastRequest struct {
	requests.RpcRequest
	TopicFullName  string `position:"Query" name:"TopicFullName"`
	MessageContent string `position:"Query" name:"MessageContent"`
	ProductKey     string `position:"Query" name:"ProductKey"`
}

func (req *PubBroadcastRequest) Invoke(client *sdk.Client) (resp *PubBroadcastResponse, err error) {
	req.InitWithApiInfo("Iot", "2017-04-20", "PubBroadcast", "", "")
	resp = &PubBroadcastResponse{}
	err = client.DoAction(req, resp)
	return
}

type PubBroadcastResponse struct {
	responses.BaseResponse
	RequestId    string
	Success      bool
	ErrorMessage string
}
