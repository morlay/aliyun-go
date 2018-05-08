package jaq

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SpamRegisterPreventionRequest struct {
	requests.RpcRequest
	CallerName      string `position:"Query" name:"CallerName"`
	Ip              string `position:"Query" name:"Ip"`
	ProtocolVersion string `position:"Query" name:"ProtocolVersion"`
	Source          int    `position:"Query" name:"Source"`
	PhoneNumber     string `position:"Query" name:"PhoneNumber"`
	Email           string `position:"Query" name:"Email"`
	UserId          string `position:"Query" name:"UserId"`
	IdType          int    `position:"Query" name:"IdType"`
	CurrentUrl      string `position:"Query" name:"CurrentUrl"`
	Agent           string `position:"Query" name:"Agent"`
	Cookie          string `position:"Query" name:"Cookie"`
	SessionId       string `position:"Query" name:"SessionId"`
	MacAddress      string `position:"Query" name:"MacAddress"`
	Referer         string `position:"Query" name:"Referer"`
	NickName        string `position:"Query" name:"NickName"`
	CompanyName     string `position:"Query" name:"CompanyName"`
	Address         string `position:"Query" name:"Address"`
	IDNumber        string `position:"Query" name:"IDNumber"`
	BankCardNumber  string `position:"Query" name:"BankCardNumber"`
	JsToken         string `position:"Query" name:"JsToken"`
	SDKToken        string `position:"Query" name:"SDKToken"`
	ExtendData      string `position:"Query" name:"ExtendData"`
}

func (req *SpamRegisterPreventionRequest) Invoke(client *sdk.Client) (resp *SpamRegisterPreventionResponse, err error) {
	req.InitWithApiInfo("jaq", "2016-11-23", "SpamRegisterPrevention", "", "")
	resp = &SpamRegisterPreventionResponse{}
	err = client.DoAction(req, resp)
	return
}

type SpamRegisterPreventionResponse struct {
	responses.BaseResponse
	ErrorCode common.Integer
	ErrorMsg  common.String
	Data      SpamRegisterPreventionData
}

type SpamRegisterPreventionData struct {
	FnalDecision     common.Integer
	EventId          common.String
	UserId           common.String
	FinalScore       common.Integer
	FinalDesc        common.String
	Detail           common.String
	CaptchaCheckData common.String
}
