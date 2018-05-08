package csb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateCredentialsRequest struct {
	requests.RpcRequest
	Data  string `position:"Body" name:"Data"`
	CsbId int64  `position:"Query" name:"CsbId"`
}

func (req *CreateCredentialsRequest) Invoke(client *sdk.Client) (resp *CreateCredentialsResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "CreateCredentials", "CSB", "")
	resp = &CreateCredentialsResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateCredentialsResponse struct {
	responses.BaseResponse
	Code      common.Integer
	Message   common.String
	RequestId common.String
	Data      CreateCredentialsData
}

type CreateCredentialsData struct {
	Credentials CreateCredentialsCredentials
}

type CreateCredentialsCredentials struct {
	Id                common.Long
	CurrentCredential CreateCredentialsCurrentCredential
}

type CreateCredentialsCurrentCredential struct {
	SecretKey common.String
	AccessKey common.String
}
