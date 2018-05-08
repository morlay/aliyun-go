package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ApplyForRetrievalDomainNameRequest struct {
	requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
}

func (req *ApplyForRetrievalDomainNameRequest) Invoke(client *sdk.Client) (resp *ApplyForRetrievalDomainNameResponse, err error) {
	req.InitWithApiInfo("Alidns", "2015-01-09", "ApplyForRetrievalDomainName", "", "")
	resp = &ApplyForRetrievalDomainNameResponse{}
	err = client.DoAction(req, resp)
	return
}

type ApplyForRetrievalDomainNameResponse struct {
	responses.BaseResponse
	RequestId  common.String
	DomainName common.String
}
