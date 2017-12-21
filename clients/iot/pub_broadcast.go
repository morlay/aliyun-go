package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type PubBroadcastRequest struct {
	TopicFullName  string `position:"Query" name:"TopicFullName"`
	MessageContent string `position:"Query" name:"MessageContent"`
	ProductKey     string `position:"Query" name:"ProductKey"`
}

func (r PubBroadcastRequest) Invoke(client *sdk.Client) (response *PubBroadcastResponse, err error) {
	req := struct {
		*requests.RpcRequest
		PubBroadcastRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Iot", "2017-04-20", "PubBroadcast", "", "")

	resp := struct {
		*responses.BaseResponse
		PubBroadcastResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.PubBroadcastResponse

	err = client.DoAction(&req, &resp)
	return
}

type PubBroadcastResponse struct {
	RequestId    string
	Success      bool
	ErrorMessage string
}
