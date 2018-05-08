package domain_intl

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type TransferInRefetchWhoisEmailRequest struct {
	requests.RpcRequest
	DomainName   string `position:"Query" name:"DomainName"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

func (req *TransferInRefetchWhoisEmailRequest) Invoke(client *sdk.Client) (resp *TransferInRefetchWhoisEmailResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "TransferInRefetchWhoisEmail", "domain", "")
	resp = &TransferInRefetchWhoisEmailResponse{}
	err = client.DoAction(req, resp)
	return
}

type TransferInRefetchWhoisEmailResponse struct {
	responses.BaseResponse
	RequestId common.String
}
