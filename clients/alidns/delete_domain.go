package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteDomainRequest struct {
	requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
}

func (req *DeleteDomainRequest) Invoke(client *sdk.Client) (resp *DeleteDomainResponse, err error) {
	req.InitWithApiInfo("Alidns", "2015-01-09", "DeleteDomain", "", "")
	resp = &DeleteDomainResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteDomainResponse struct {
	responses.BaseResponse
	RequestId  string
	DomainName string
}
