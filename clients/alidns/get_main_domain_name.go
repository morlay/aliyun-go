package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetMainDomainNameRequest struct {
	requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	InputString  string `position:"Query" name:"InputString"`
}

func (req *GetMainDomainNameRequest) Invoke(client *sdk.Client) (resp *GetMainDomainNameResponse, err error) {
	req.InitWithApiInfo("Alidns", "2015-01-09", "GetMainDomainName", "", "")
	resp = &GetMainDomainNameResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetMainDomainNameResponse struct {
	responses.BaseResponse
	RequestId   string
	DomainName  string
	RR          string
	DomainLevel int64
}
