package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetDomainServerCertificateRequest struct {
	requests.RpcRequest
	PrivateKey              string `position:"Query" name:"PrivateKey"`
	ServerCertificateStatus string `position:"Query" name:"ServerCertificateStatus"`
	ServerCertificate       string `position:"Query" name:"ServerCertificate"`
	SecurityToken           string `position:"Query" name:"SecurityToken"`
	CertName                string `position:"Query" name:"CertName"`
	DomainName              string `position:"Query" name:"DomainName"`
	OwnerId                 int64  `position:"Query" name:"OwnerId"`
	Region                  string `position:"Query" name:"Region"`
}

func (req *SetDomainServerCertificateRequest) Invoke(client *sdk.Client) (resp *SetDomainServerCertificateResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetDomainServerCertificate", "", "")
	resp = &SetDomainServerCertificateResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetDomainServerCertificateResponse struct {
	responses.BaseResponse
	RequestId string
}
