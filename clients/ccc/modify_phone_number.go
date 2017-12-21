package ccc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyPhoneNumberRequest struct {
	ContactFlowId string `position:"Query" name:"ContactFlowId"`
	InstanceId    string `position:"Query" name:"InstanceId"`
	PhoneNumberId string `position:"Query" name:"PhoneNumberId"`
	Usage         string `position:"Query" name:"Usage"`
}

func (r ModifyPhoneNumberRequest) Invoke(client *sdk.Client) (response *ModifyPhoneNumberResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyPhoneNumberRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CCC", "2017-07-05", "ModifyPhoneNumber", "ccc", "")

	resp := struct {
		*responses.BaseResponse
		ModifyPhoneNumberResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyPhoneNumberResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyPhoneNumberResponse struct {
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
