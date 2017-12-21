package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest struct {
	SaleId            string `position:"Query" name:"SaleId"`
	UserClientIp      string `position:"Query" name:"UserClientIp"`
	DomainName        string `position:"Query" name:"DomainName"`
	Lang              string `position:"Query" name:"Lang"`
	ContactTemplateId int64  `position:"Query" name:"ContactTemplateId"`
}

func (r SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest) Invoke(client *sdk.Client) (response *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Domain", "2016-05-11", "SaveTaskForSubmittingDomainNameCredentialByTemplateId", "", "")

	resp := struct {
		*responses.BaseResponse
		SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse

	err = client.DoAction(&req, &resp)
	return
}

type SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse struct {
	RequestId string
	Success   bool
	TaskNo    string
}
