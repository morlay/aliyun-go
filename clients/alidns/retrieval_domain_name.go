package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RetrievalDomainNameRequest struct {
	requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
}

func (req *RetrievalDomainNameRequest) Invoke(client *sdk.Client) (resp *RetrievalDomainNameResponse, err error) {
	req.InitWithApiInfo("Alidns", "2015-01-09", "RetrievalDomainName", "", "")
	resp = &RetrievalDomainNameResponse{}
	err = client.DoAction(req, resp)
	return
}

type RetrievalDomainNameResponse struct {
	responses.BaseResponse
	RequestId  common.String
	DomainName common.String
	WhoisEmail common.String
}
