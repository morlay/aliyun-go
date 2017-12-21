package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetMainDomainNameRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	InputString  string `position:"Query" name:"InputString"`
}

func (r GetMainDomainNameRequest) Invoke(client *sdk.Client) (response *GetMainDomainNameResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetMainDomainNameRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Alidns", "2015-01-09", "GetMainDomainName", "", "")

	resp := struct {
		*responses.BaseResponse
		GetMainDomainNameResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetMainDomainNameResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetMainDomainNameResponse struct {
	RequestId   string
	DomainName  string
	RR          string
	DomainLevel int64
}
