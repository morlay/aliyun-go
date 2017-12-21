package jaq

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CheckAccountAndPasswordRiskRequest struct {
	requests.RpcRequest
	CallerName   string `position:"Query" name:"CallerName"`
	AccountName  string `position:"Query" name:"AccountName"`
	PasswordHash string `position:"Query" name:"PasswordHash"`
}

func (req *CheckAccountAndPasswordRiskRequest) Invoke(client *sdk.Client) (resp *CheckAccountAndPasswordRiskResponse, err error) {
	req.InitWithApiInfo("jaq", "2016-11-23", "CheckAccountAndPasswordRisk", "", "")
	resp = &CheckAccountAndPasswordRiskResponse{}
	err = client.DoAction(req, resp)
	return
}

type CheckAccountAndPasswordRiskResponse struct {
	responses.BaseResponse
	ErrorCode int
	ErrorMsg  string
	Data      CheckAccountAndPasswordRiskData
}

type CheckAccountAndPasswordRiskData struct {
	FnalDecision     int
	EventId          string
	UserId           string
	FinalScore       int
	FinalDesc        string
	Detail           string
	CaptchaCheckData string
}
