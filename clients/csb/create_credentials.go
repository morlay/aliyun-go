package csb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	Code      int
	Message   string
	RequestId string
	Data      CreateCredentialsData
}

type CreateCredentialsData struct {
	Credentials CreateCredentialsCredentials
}

type CreateCredentialsCredentials struct {
	Id                int64
	CurrentCredential CreateCredentialsCurrentCredential
}

type CreateCredentialsCurrentCredential struct {
	SecretKey string
	AccessKey string
}
