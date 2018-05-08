package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RegistrantProfileRealNameVerificationRequest struct {
	requests.RpcRequest
	IdentityCredentialType string `position:"Query" name:"IdentityCredentialType"`
	UserClientIp           string `position:"Query" name:"UserClientIp"`
	RegistrantProfileID    int64  `position:"Query" name:"RegistrantProfileID"`
	IdentityCredential     string `position:"Body" name:"IdentityCredential"`
	Lang                   string `position:"Query" name:"Lang"`
	IdentityCredentialNo   string `position:"Query" name:"IdentityCredentialNo"`
}

func (req *RegistrantProfileRealNameVerificationRequest) Invoke(client *sdk.Client) (resp *RegistrantProfileRealNameVerificationResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "RegistrantProfileRealNameVerification", "", "")
	resp = &RegistrantProfileRealNameVerificationResponse{}
	err = client.DoAction(req, resp)
	return
}

type RegistrantProfileRealNameVerificationResponse struct {
	responses.BaseResponse
	RequestId common.String
}
