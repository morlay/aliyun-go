package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryRegistrantProfileRealNameVerificationInfoRequest struct {
	requests.RpcRequest
	FetchImage          string `position:"Query" name:"FetchImage"`
	UserClientIp        string `position:"Query" name:"UserClientIp"`
	RegistrantProfileId int64  `position:"Query" name:"RegistrantProfileId"`
	Lang                string `position:"Query" name:"Lang"`
}

func (req *QueryRegistrantProfileRealNameVerificationInfoRequest) Invoke(client *sdk.Client) (resp *QueryRegistrantProfileRealNameVerificationInfoResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "QueryRegistrantProfileRealNameVerificationInfo", "", "")
	resp = &QueryRegistrantProfileRealNameVerificationInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryRegistrantProfileRealNameVerificationInfoResponse struct {
	responses.BaseResponse
	RequestId              common.String
	SubmissionDate         common.String
	ModificationDate       common.String
	IdentityCredential     common.String
	RegistrantProfileId    common.Long
	IdentityCredentialNo   common.String
	IdentityCredentialType common.String
}
