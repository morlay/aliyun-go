package httpdns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddDomainRequest struct {
	requests.RpcRequest
	AccountId  string `position:"Query" name:"AccountId"`
	DomainName string `position:"Query" name:"DomainName"`
}

func (req *AddDomainRequest) Invoke(client *sdk.Client) (resp *AddDomainResponse, err error) {
	req.InitWithApiInfo("Httpdns", "2016-02-01", "AddDomain", "", "")
	resp = &AddDomainResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddDomainResponse struct {
	responses.BaseResponse
	RequestId  string
	DomainName string
}
