package csb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	Code      common.Integer
	Message   common.String
	RequestId common.String
	Data      ReplaceCredentialData
}

type ReplaceCredentialData struct {
	Credentials ReplaceCredentialCredentials
}

type ReplaceCredentialCredentials struct {
	GmtCreate         common.Long
	Id                common.Long
	Name              common.String
	UserId            common.String
	CurrentCredential ReplaceCredentialCurrentCredential
	NewCredential     ReplaceCredentialNewCredential
}

type ReplaceCredentialCurrentCredential struct {
	AccessKey common.String
	SecretKey common.String
}

type ReplaceCredentialNewCredential struct {
	AccessKey common.String
	SecretKey common.String
}
