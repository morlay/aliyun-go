package httpdns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddDomainRequest struct {
	AccountId  string `position:"Query" name:"AccountId"`
	DomainName string `position:"Query" name:"DomainName"`
}

func (r AddDomainRequest) Invoke(client *sdk.Client) (response *AddDomainResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddDomainRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Httpdns", "2016-02-01", "AddDomain", "", "")

	resp := struct {
		*responses.BaseResponse
		AddDomainResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddDomainResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddDomainResponse struct {
	RequestId  string
	DomainName string
}
