package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveTaskForSubmittingDomainNameCredentialRequest struct {
	requests.RpcRequest
	CredentialNo string `position:"Query" name:"CredentialNo"`
	SaleId       string `position:"Query" name:"SaleId"`
	Credential   string `position:"Query" name:"Credential"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	Lang         string `position:"Query" name:"Lang"`
}

func (req *SaveTaskForSubmittingDomainNameCredentialRequest) Invoke(client *sdk.Client) (resp *SaveTaskForSubmittingDomainNameCredentialResponse, err error) {
	req.InitWithApiInfo("Domain", "2016-05-11", "SaveTaskForSubmittingDomainNameCredential", "", "")
	resp = &SaveTaskForSubmittingDomainNameCredentialResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveTaskForSubmittingDomainNameCredentialResponse struct {
	responses.BaseResponse
	RequestId string
	TaskNo    string
}
