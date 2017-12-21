package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RetrievalDomainNameRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
}

func (r RetrievalDomainNameRequest) Invoke(client *sdk.Client) (response *RetrievalDomainNameResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RetrievalDomainNameRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Alidns", "2015-01-09", "RetrievalDomainName", "", "")

	resp := struct {
		*responses.BaseResponse
		RetrievalDomainNameResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RetrievalDomainNameResponse

	err = client.DoAction(&req, &resp)
	return
}

type RetrievalDomainNameResponse struct {
	RequestId  string
	DomainName string
	WhoisEmail string
}
