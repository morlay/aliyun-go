package ccc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddPhoneNumberRequest struct {
	requests.RpcRequest
	ContactFlowId string `position:"Query" name:"ContactFlowId"`
	InstanceId    string `position:"Query" name:"InstanceId"`
	Usage         string `position:"Query" name:"Usage"`
	PhoneNumber   string `position:"Query" name:"PhoneNumber"`
}

func (req *AddPhoneNumberRequest) Invoke(client *sdk.Client) (resp *AddPhoneNumberResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "AddPhoneNumber", "ccc", "")
	resp = &AddPhoneNumberResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddPhoneNumberResponse struct {
	responses.BaseResponse
	RequestId      string
	Success        bool
	Code           string
	Message        string
	HttpStatusCode int
	PhoneNumber    AddPhoneNumberPhoneNumber
}

type AddPhoneNumberPhoneNumber struct {
	PhoneNumberId          string
	InstanceId             string
	Number                 string
	PhoneNumberDescription string
	TestOnly               bool
	RemainingTime          int
	AllowOutbound          bool
	Usage                  string
	Trunks                 int
	ContactFlow            AddPhoneNumberContactFlow
}

type AddPhoneNumberContactFlow struct {
	ContactFlowId          string
	InstanceId             string
	ContactFlowName        string
	ContactFlowDescription string
	Type                   string
}
