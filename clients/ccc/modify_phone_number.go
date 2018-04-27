package ccc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyPhoneNumberRequest struct {
	requests.RpcRequest
	ContactFlowId string `position:"Query" name:"ContactFlowId"`
	InstanceId    string `position:"Query" name:"InstanceId"`
	PhoneNumberId string `position:"Query" name:"PhoneNumberId"`
	Usage         string `position:"Query" name:"Usage"`
}

func (req *ModifyPhoneNumberRequest) Invoke(client *sdk.Client) (resp *ModifyPhoneNumberResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "ModifyPhoneNumber", "ccc", "")
	resp = &ModifyPhoneNumberResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyPhoneNumberResponse struct {
	responses.BaseResponse
	RequestId      string
	Success        bool
	Code           string
	Message        string
	HttpStatusCode int
	PhoneNumber    ModifyPhoneNumberPhoneNumber
}

type ModifyPhoneNumberPhoneNumber struct {
	PhoneNumberId          string
	InstanceId             string
	Number                 string
	PhoneNumberDescription string
	TestOnly               bool
	RemainingTime          int
	AllowOutbound          bool
	Usage                  string
	Trunks                 int
	ContactFlow            ModifyPhoneNumberContactFlow
}

type ModifyPhoneNumberContactFlow struct {
	ContactFlowId          string
	InstanceId             string
	ContactFlowName        string
	ContactFlowDescription string
	Type                   string
}
