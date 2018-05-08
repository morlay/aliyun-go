package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId  common.String
	DomainName common.String
	Avail      common.String
	Premium    common.String
	Reason     common.String
	Price      common.Long
}
