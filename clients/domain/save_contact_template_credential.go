package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveContactTemplateCredentialRequest struct {
	CredentialNo      string `position:"Query" name:"CredentialNo"`
	Credential        string `position:"Query" name:"Credential"`
	UserClientIp      string `position:"Query" name:"UserClientIp"`
	Lang              string `position:"Query" name:"Lang"`
	ContactTemplateId int64  `position:"Query" name:"ContactTemplateId"`
}

func (r SaveContactTemplateCredentialRequest) Invoke(client *sdk.Client) (response *SaveContactTemplateCredentialResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SaveContactTemplateCredentialRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Domain", "2016-05-11", "SaveContactTemplateCredential", "", "")

	resp := struct {
		*responses.BaseResponse
		SaveContactTemplateCredentialResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SaveContactTemplateCredentialResponse

	err = client.DoAction(&req, &resp)
	return
}

type SaveContactTemplateCredentialResponse struct {
	RequestId string
}
