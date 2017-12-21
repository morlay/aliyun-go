package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetDomainCertificateRequest struct {
	requests.RpcRequest
	GroupId               string `position:"Query" name:"GroupId"`
	DomainName            string `position:"Query" name:"DomainName"`
	CertificateName       string `position:"Query" name:"CertificateName"`
	CertificateBody       string `position:"Query" name:"CertificateBody"`
	CertificatePrivateKey string `position:"Query" name:"CertificatePrivateKey"`
}

func (req *SetDomainCertificateRequest) Invoke(client *sdk.Client) (resp *SetDomainCertificateResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "SetDomainCertificate", "apigateway", "")
	resp = &SetDomainCertificateResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetDomainCertificateResponse struct {
	responses.BaseResponse
	RequestId string
}
