package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type PubRequest struct {
	TopicFullName  string `position:"Query" name:"TopicFullName"`
	Qos            int    `position:"Query" name:"Qos"`
	MessageContent string `position:"Query" name:"MessageContent"`
	ProductKey     string `position:"Query" name:"ProductKey"`
}

func (r PubRequest) Invoke(client *sdk.Client) (response *PubResponse, err error) {
	req := struct {
		*requests.RpcRequest
		PubRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Iot", "2017-04-20", "Pub", "", "")

	resp := struct {
		*responses.BaseResponse
		PubResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.PubResponse

	err = client.DoAction(&req, &resp)
	return
}

type PubResponse struct {
	RequestId    string
	Success      bool
	ErrorMessage string
	MessageId    string
}
