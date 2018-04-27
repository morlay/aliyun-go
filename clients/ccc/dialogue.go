package ccc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DialogueRequest struct {
	requests.RpcRequest
	CallId        string `position:"Query" name:"CallId"`
	CallingNumber string `position:"Query" name:"CallingNumber"`
	InstanceId    string `position:"Query" name:"InstanceId"`
	CalledNumber  string `position:"Query" name:"CalledNumber"`
	ActionParams  string `position:"Query" name:"ActionParams"`
	CallType      string `position:"Query" name:"CallType"`
	ScenarioId    string `position:"Query" name:"ScenarioId"`
	TaskId        string `position:"Query" name:"TaskId"`
	Utterance     string `position:"Query" name:"Utterance"`
}

func (req *DialogueRequest) Invoke(client *sdk.Client) (resp *DialogueResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "Dialogue", "ccc", "")
	resp = &DialogueResponse{}
	err = client.DoAction(req, resp)
	return
}

type DialogueResponse struct {
	responses.BaseResponse
	RequestId      string
	Success        bool
	Code           string
	Message        string
	HttpStatusCode int
	Feedback       DialogueFeedback
}

type DialogueFeedback struct {
	Content      string
	Action       string
	ActionParams string
}
