package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest struct {
	requests.RpcRequest
	SaleId            string `position:"Query" name:"SaleId"`
	UserClientIp      string `position:"Query" name:"UserClientIp"`
	DomainName        string `position:"Query" name:"DomainName"`
	Lang              string `position:"Query" name:"Lang"`
	ContactTemplateId int64  `position:"Query" name:"ContactTemplateId"`
}

func (req *SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest) Invoke(client *sdk.Client) (resp *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse, err error) {
	req.InitWithApiInfo("Domain", "2016-05-11", "SaveTaskForSubmittingDomainNameCredentialByTemplateId", "", "")
	resp = &SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	TaskNo    string
}
