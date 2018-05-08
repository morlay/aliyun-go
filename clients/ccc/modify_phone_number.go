package ccc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId      common.String
	Success        bool
	Code           common.String
	Message        common.String
	HttpStatusCode common.Integer
	PhoneNumber    ModifyPhoneNumberPhoneNumber
}

type ModifyPhoneNumberPhoneNumber struct {
	PhoneNumberId          common.String
	InstanceId             common.String
	Number                 common.String
	PhoneNumberDescription common.String
	TestOnly               bool
	RemainingTime          common.Integer
	AllowOutbound          bool
	Usage                  common.String
	Trunks                 common.Integer
	ContactFlow            ModifyPhoneNumberContactFlow
}

type ModifyPhoneNumberContactFlow struct {
	ContactFlowId          common.String
	InstanceId             common.String
	ContactFlowName        common.String
	ContactFlowDescription common.String
	Type                   common.String
}
