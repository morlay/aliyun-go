package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveSingleTaskForDomainNameProxyServiceRequest struct {
	requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	Lang         string `position:"Query" name:"Lang"`
	Status       string `position:"Query" name:"Status"`
}

func (req *SaveSingleTaskForDomainNameProxyServiceRequest) Invoke(client *sdk.Client) (resp *SaveSingleTaskForDomainNameProxyServiceResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "SaveSingleTaskForDomainNameProxyService", "", "")
	resp = &SaveSingleTaskForDomainNameProxyServiceResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveSingleTaskForDomainNameProxyServiceResponse struct {
	responses.BaseResponse
	RequestId string
	TaskNo    string
}
