package domain_intl

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CheckTransferInFeasibilityRequest struct {
	requests.RpcRequest
	TransferAuthorizationCode string `position:"Query" name:"TransferAuthorizationCode"`
	UserClientIp              string `position:"Query" name:"UserClientIp"`
	DomainName                string `position:"Query" name:"DomainName"`
	Lang                      string `position:"Query" name:"Lang"`
}

func (req *CheckTransferInFeasibilityRequest) Invoke(client *sdk.Client) (resp *CheckTransferInFeasibilityResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "CheckTransferInFeasibility", "domain", "")
	resp = &CheckTransferInFeasibilityResponse{}
	err = client.DoAction(req, resp)
	return
}

type CheckTransferInFeasibilityResponse struct {
	responses.BaseResponse
	RequestId   string
	CanTransfer bool
	Code        string
	Message     string
	ProductId   string
}
