package ccc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId      common.String
	Success        bool
	Code           common.String
	Message        common.String
	HttpStatusCode common.Integer
	PhoneNumber    AddPhoneNumberPhoneNumber
}

type AddPhoneNumberPhoneNumber struct {
	PhoneNumberId          common.String
	InstanceId             common.String
	Number                 common.String
	PhoneNumberDescription common.String
	TestOnly               bool
	RemainingTime          common.Integer
	AllowOutbound          bool
	Usage                  common.String
	Trunks                 common.Integer
	ContactFlow            AddPhoneNumberContactFlow
}

type AddPhoneNumberContactFlow struct {
	ContactFlowId          common.String
	InstanceId             common.String
	ContactFlowName        common.String
	ContactFlowDescription common.String
	Type                   common.String
}
