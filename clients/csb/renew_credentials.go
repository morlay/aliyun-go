package csb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RenewCredentialsRequest struct {
	requests.RpcRequest
	CredentialId int64 `position:"Query" name:"CredentialId"`
}

func (req *RenewCredentialsRequest) Invoke(client *sdk.Client) (resp *RenewCredentialsResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "RenewCredentials", "CSB", "")
	resp = &RenewCredentialsResponse{}
	err = client.DoAction(req, resp)
	return
}

type RenewCredentialsResponse struct {
	responses.BaseResponse
	Code      int
	Message   string
	RequestId string
	Data      RenewCredentialsData
}

type RenewCredentialsData struct {
	Credentials RenewCredentialsCredentials
}

type RenewCredentialsCredentials struct {
	GmtCreate         int64
	Id                int64
	Name              string
	UserId            string
	CurrentCredential RenewCredentialsCurrentCredential
	NewCredential     RenewCredentialsNewCredential
}

type RenewCredentialsCurrentCredential struct {
	AccessKey string
	SecretKey string
}

type RenewCredentialsNewCredential struct {
	AccessKey string
	SecretKey string
}
