package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CheckDomainRequest struct {
	requests.RpcRequest
	FeeCurrency string `position:"Query" name:"FeeCurrency"`
	FeePeriod   int    `position:"Query" name:"FeePeriod"`
	DomainName  string `position:"Query" name:"DomainName"`
	FeeCommand  string `position:"Query" name:"FeeCommand"`
}

func (req *CheckDomainRequest) Invoke(client *sdk.Client) (resp *CheckDomainResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "CheckDomain", "", "")
	resp = &CheckDomainResponse{}
	err = client.DoAction(req, resp)
	return
}

type CheckDomainResponse struct {
	responses.BaseResponse
	RequestId  string
	DomainName string
	Avail      string
	Premium    string
	Reason     string
	Price      int64
}
