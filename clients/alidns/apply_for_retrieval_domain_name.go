package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ApplyForRetrievalDomainNameRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
}

func (r ApplyForRetrievalDomainNameRequest) Invoke(client *sdk.Client) (response *ApplyForRetrievalDomainNameResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ApplyForRetrievalDomainNameRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Alidns", "2015-01-09", "ApplyForRetrievalDomainName", "", "")

	resp := struct {
		*responses.BaseResponse
		ApplyForRetrievalDomainNameResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ApplyForRetrievalDomainNameResponse

	err = client.DoAction(&req, &resp)
	return
}

type ApplyForRetrievalDomainNameResponse struct {
	RequestId  string
	DomainName string
}
