package domain_intl

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SaveSingleTaskForDomainNameProxyServiceRequest struct {
	requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	Lang         string `position:"Query" name:"Lang"`
	Status       string `position:"Query" name:"Status"`
}

func (req *SaveSingleTaskForDomainNameProxyServiceRequest) Invoke(client *sdk.Client) (resp *SaveSingleTaskForDomainNameProxyServiceResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "SaveSingleTaskForDomainNameProxyService", "domain", "")
	resp = &SaveSingleTaskForDomainNameProxyServiceResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveSingleTaskForDomainNameProxyServiceResponse struct {
	responses.BaseResponse
	RequestId common.String
	TaskNo    common.String
}
