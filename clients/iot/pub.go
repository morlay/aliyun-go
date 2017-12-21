package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type PubRequest struct {
	requests.RpcRequest
	TopicFullName  string `position:"Query" name:"TopicFullName"`
	Qos            int    `position:"Query" name:"Qos"`
	MessageContent string `position:"Query" name:"MessageContent"`
	ProductKey     string `position:"Query" name:"ProductKey"`
}

func (req *PubRequest) Invoke(client *sdk.Client) (resp *PubResponse, err error) {
	req.InitWithApiInfo("Iot", "2017-04-20", "Pub", "", "")
	resp = &PubResponse{}
	err = client.DoAction(req, resp)
	return
}

type PubResponse struct {
	responses.BaseResponse
	RequestId    string
	Success      bool
	ErrorMessage string
	MessageId    string
}
