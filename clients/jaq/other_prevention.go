package jaq

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OtherPreventionRequest struct {
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
	UserName        string `position:"Query" name:"UserName"`
	CompanyName     string `position:"Query" name:"CompanyName"`
	Address         string `position:"Query" name:"Address"`
	IDNumber        string `position:"Query" name:"IDNumber"`
	BankCardNumber  string `position:"Query" name:"BankCardNumber"`
	RegisterIp      string `position:"Query" name:"RegisterIp"`
	RegisterDate    int64  `position:"Query" name:"RegisterDate"`
	LoginIp         string `position:"Query" name:"LoginIp"`
	LoginDate       int64  `position:"Query" name:"LoginDate"`
	ExtendData      string `position:"Query" name:"ExtendData"`
	PasswordHash    string `position:"Query" name:"PasswordHash"`
	JsToken         string `position:"Query" name:"JsToken"`
	SDKToken        string `position:"Query" name:"SDKToken"`
}

func (req *OtherPreventionRequest) Invoke(client *sdk.Client) (resp *OtherPreventionResponse, err error) {
	req.InitWithApiInfo("jaq", "2016-11-23", "OtherPrevention", "", "")
	resp = &OtherPreventionResponse{}
	err = client.DoAction(req, resp)
	return
}

type OtherPreventionResponse struct {
	responses.BaseResponse
	ErrorCode int
	ErrorMsg  string
	Data      OtherPreventionData
}

type OtherPreventionData struct {
	FnalDecision     int
	EventId          string
	UserId           string
	FinalScore       int
	FinalDesc        string
	Detail           string
	CaptchaCheckData string
}
