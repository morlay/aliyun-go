package jaq

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ActivityPreventionRequest struct {
	CallerName          string `position:"Query" name:"CallerName"`
	Ip                  string `position:"Query" name:"Ip"`
	ProtocolVersion     string `position:"Query" name:"ProtocolVersion"`
	Source              int    `position:"Query" name:"Source"`
	ActivityDescription string `position:"Query" name:"ActivityDescription"`
	ActivityId          string `position:"Query" name:"ActivityId"`
	Prize               string `position:"Query" name:"Prize"`
	PrizeType           int    `position:"Query" name:"PrizeType"`
	PhoneNumber         string `position:"Query" name:"PhoneNumber"`
	Email               string `position:"Query" name:"Email"`
	UserId              string `position:"Query" name:"UserId"`
	IdType              int    `position:"Query" name:"IdType"`
	CurrentUrl          string `position:"Query" name:"CurrentUrl"`
	Agent               string `position:"Query" name:"Agent"`
	Cookie              string `position:"Query" name:"Cookie"`
	SessionId           string `position:"Query" name:"SessionId"`
	MacAddress          string `position:"Query" name:"MacAddress"`
	Referer             string `position:"Query" name:"Referer"`
	UserName            string `position:"Query" name:"UserName"`
	CompanyName         string `position:"Query" name:"CompanyName"`
	Address             string `position:"Query" name:"Address"`
	IDNumber            string `position:"Query" name:"IDNumber"`
	BankCardNumber      string `position:"Query" name:"BankCardNumber"`
	RegisterIp          string `position:"Query" name:"RegisterIp"`
	RegisterDate        int64  `position:"Query" name:"RegisterDate"`
	ExtendData          string `position:"Query" name:"ExtendData"`
	JsToken             string `position:"Query" name:"JsToken"`
	SDKToken            string `position:"Query" name:"SDKToken"`
}

func (r ActivityPreventionRequest) Invoke(client *sdk.Client) (response *ActivityPreventionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ActivityPreventionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("jaq", "2016-11-23", "ActivityPrevention", "", "")

	resp := struct {
		*responses.BaseResponse
		ActivityPreventionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ActivityPreventionResponse

	err = client.DoAction(&req, &resp)
	return
}

type ActivityPreventionResponse struct {
	ErrorCode int
	ErrorMsg  string
	Data      ActivityPreventionData
}

type ActivityPreventionData struct {
	FnalDecision     int
	EventId          string
	UserId           string
	FinalScore       int
	FinalDesc        string
	Detail           string
	CaptchaCheckData string
}
