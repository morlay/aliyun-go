package csb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReplaceCredentialRequest struct {
	requests.RpcRequest
	CredentialId int64 `position:"Query" name:"CredentialId"`
}

func (req *ReplaceCredentialRequest) Invoke(client *sdk.Client) (resp *ReplaceCredentialResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "ReplaceCredential", "CSB", "")
	resp = &ReplaceCredentialResponse{}
	err = client.DoAction(req, resp)
	return
}

type ReplaceCredentialResponse struct {
	responses.BaseResponse
	Code      int
	Message   string
	RequestId string
	Data      ReplaceCredentialData
}

type ReplaceCredentialData struct {
	Credentials ReplaceCredentialCredentials
}

type ReplaceCredentialCredentials struct {
	GmtCreate         int64
	Id                int64
	Name              string
	UserId            string
	CurrentCredential ReplaceCredentialCurrentCredential
	NewCredential     ReplaceCredentialNewCredential
}

type ReplaceCredentialCurrentCredential struct {
	AccessKey string
	SecretKey string
}

type ReplaceCredentialNewCredential struct {
	AccessKey string
	SecretKey string
}
