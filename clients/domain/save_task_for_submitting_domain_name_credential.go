package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveTaskForSubmittingDomainNameCredentialRequest struct {
	CredentialNo string `position:"Query" name:"CredentialNo"`
	SaleId       string `position:"Query" name:"SaleId"`
	Credential   string `position:"Query" name:"Credential"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	Lang         string `position:"Query" name:"Lang"`
}

func (r SaveTaskForSubmittingDomainNameCredentialRequest) Invoke(client *sdk.Client) (response *SaveTaskForSubmittingDomainNameCredentialResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SaveTaskForSubmittingDomainNameCredentialRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Domain", "2016-05-11", "SaveTaskForSubmittingDomainNameCredential", "", "")

	resp := struct {
		*responses.BaseResponse
		SaveTaskForSubmittingDomainNameCredentialResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SaveTaskForSubmittingDomainNameCredentialResponse

	err = client.DoAction(&req, &resp)
	return
}

type SaveTaskForSubmittingDomainNameCredentialResponse struct {
	RequestId string
	TaskNo    string
}
