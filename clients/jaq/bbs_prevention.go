package jaq

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type BbsPreventionRequest struct {
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
	JsToken         string `position:"Query" name:"JsToken"`
	SDKToken        string `position:"Query" name:"SDKToken"`
	ExtendData      string `position:"Query" name:"ExtendData"`
}

func (r BbsPreventionRequest) Invoke(client *sdk.Client) (response *BbsPreventionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		BbsPreventionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("jaq", "2016-11-23", "BbsPrevention", "", "")

	resp := struct {
		*responses.BaseResponse
		BbsPreventionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.BbsPreventionResponse

	err = client.DoAction(&req, &resp)
	return
}

type BbsPreventionResponse struct {
	ErrorCode int
	ErrorMsg  string
	Data      BbsPreventionData
}

type BbsPreventionData struct {
	FnalDecision     int
	EventId          string
	UserId           string
	FinalScore       int
	FinalDesc        string
	Detail           string
	CaptchaCheckData string
}
