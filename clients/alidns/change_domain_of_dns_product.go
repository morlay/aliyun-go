package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ChangeDomainOfDnsProductRequest struct {
	requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	InstanceId   string `position:"Query" name:"InstanceId"`
	NewDomain    string `position:"Query" name:"NewDomain"`
}

func (req *ChangeDomainOfDnsProductRequest) Invoke(client *sdk.Client) (resp *ChangeDomainOfDnsProductResponse, err error) {
	req.InitWithApiInfo("Alidns", "2015-01-09", "ChangeDomainOfDnsProduct", "", "")
	resp = &ChangeDomainOfDnsProductResponse{}
	err = client.DoAction(req, resp)
	return
}

type ChangeDomainOfDnsProductResponse struct {
	responses.BaseResponse
	RequestId      string
	OriginalDomain string
}
