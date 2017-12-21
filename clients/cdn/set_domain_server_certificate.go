package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetDomainServerCertificateRequest struct {
	PrivateKey              string `position:"Query" name:"PrivateKey"`
	ServerCertificateStatus string `position:"Query" name:"ServerCertificateStatus"`
	ServerCertificate       string `position:"Query" name:"ServerCertificate"`
	SecurityToken           string `position:"Query" name:"SecurityToken"`
	CertName                string `position:"Query" name:"CertName"`
	DomainName              string `position:"Query" name:"DomainName"`
	OwnerId                 int64  `position:"Query" name:"OwnerId"`
	Region                  string `position:"Query" name:"Region"`
}

func (r SetDomainServerCertificateRequest) Invoke(client *sdk.Client) (response *SetDomainServerCertificateResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetDomainServerCertificateRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetDomainServerCertificate", "", "")

	resp := struct {
		*responses.BaseResponse
		SetDomainServerCertificateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetDomainServerCertificateResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetDomainServerCertificateResponse struct {
	RequestId string
}
