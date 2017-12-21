package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ChangeDomainOfDnsProductRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	InstanceId   string `position:"Query" name:"InstanceId"`
	NewDomain    string `position:"Query" name:"NewDomain"`
}

func (r ChangeDomainOfDnsProductRequest) Invoke(client *sdk.Client) (response *ChangeDomainOfDnsProductResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ChangeDomainOfDnsProductRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Alidns", "2015-01-09", "ChangeDomainOfDnsProduct", "", "")

	resp := struct {
		*responses.BaseResponse
		ChangeDomainOfDnsProductResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ChangeDomainOfDnsProductResponse

	err = client.DoAction(&req, &resp)
	return
}

type ChangeDomainOfDnsProductResponse struct {
	RequestId      string
	OriginalDomain string
}
