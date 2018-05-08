package domain_intl

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CheckDomainRequest struct {
	requests.RpcRequest
	DomainName string `position:"Query" name:"DomainName"`
}

func (req *CheckDomainRequest) Invoke(client *sdk.Client) (resp *CheckDomainResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "CheckDomain", "domain", "")
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
}
