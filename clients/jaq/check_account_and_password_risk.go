package jaq

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CheckAccountAndPasswordRiskRequest struct {
	CallerName   string `position:"Query" name:"CallerName"`
	AccountName  string `position:"Query" name:"AccountName"`
	PasswordHash string `position:"Query" name:"PasswordHash"`
}

func (r CheckAccountAndPasswordRiskRequest) Invoke(client *sdk.Client) (response *CheckAccountAndPasswordRiskResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CheckAccountAndPasswordRiskRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("jaq", "2016-11-23", "CheckAccountAndPasswordRisk", "", "")

	resp := struct {
		*responses.BaseResponse
		CheckAccountAndPasswordRiskResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CheckAccountAndPasswordRiskResponse

	err = client.DoAction(&req, &resp)
	return
}

type CheckAccountAndPasswordRiskResponse struct {
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
