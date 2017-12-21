package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetDomainCertificateRequest struct {
	GroupId               string `position:"Query" name:"GroupId"`
	DomainName            string `position:"Query" name:"DomainName"`
	CertificateName       string `position:"Query" name:"CertificateName"`
	CertificateBody       string `position:"Query" name:"CertificateBody"`
	CertificatePrivateKey string `position:"Query" name:"CertificatePrivateKey"`
}

func (r SetDomainCertificateRequest) Invoke(client *sdk.Client) (response *SetDomainCertificateResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetDomainCertificateRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "SetDomainCertificate", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		SetDomainCertificateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetDomainCertificateResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetDomainCertificateResponse struct {
	RequestId string
}
