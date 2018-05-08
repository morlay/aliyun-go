package csb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	Code      common.Integer
	Message   common.String
	RequestId common.String
	Data      RenewCredentialsData
}

type RenewCredentialsData struct {
	Credentials RenewCredentialsCredentials
}

type RenewCredentialsCredentials struct {
	GmtCreate         common.Long
	Id                common.Long
	Name              common.String
	UserId            common.String
	CurrentCredential RenewCredentialsCurrentCredential
	NewCredential     RenewCredentialsNewCredential
}

type RenewCredentialsCurrentCredential struct {
	AccessKey common.String
	SecretKey common.String
}

type RenewCredentialsNewCredential struct {
	AccessKey common.String
	SecretKey common.String
}
