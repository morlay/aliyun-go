package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CheckDomainRequest struct {
	requests.RpcRequest
	DomainName string `position:"Query" name:"DomainName"`
}

func (req *CheckDomainRequest) Invoke(client *sdk.Client) (resp *CheckDomainResponse, err error) {
	req.InitWithApiInfo("Domain", "2016-05-11", "CheckDomain", "", "")
	resp = &CheckDomainResponse{}
	err = client.DoAction(req, resp)
	return
}

type CheckDomainResponse struct {
	responses.BaseResponse
	RequestId   string
	Name        string
	Avail       int
	Reason      string
	FeeCurrency string
	FeePeriod   int
	FeeFee      string
	RmbFee      string
	FeeCommand  string
}
