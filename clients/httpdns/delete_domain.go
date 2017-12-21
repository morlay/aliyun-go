package httpdns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteDomainRequest struct {
	requests.RpcRequest
	AccountId  string `position:"Query" name:"AccountId"`
	DomainName string `position:"Query" name:"DomainName"`
}

func (req *DeleteDomainRequest) Invoke(client *sdk.Client) (resp *DeleteDomainResponse, err error) {
	req.InitWithApiInfo("Httpdns", "2016-02-01", "DeleteDomain", "", "")
	resp = &DeleteDomainResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteDomainResponse struct {
	responses.BaseResponse
	RequestId  string
	DomainName string
}
