package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetDomainRequest struct {
	requests.RpcRequest
	GroupId               string `position:"Query" name:"GroupId"`
	DomainName            string `position:"Query" name:"DomainName"`
	CertificateName       string `position:"Query" name:"CertificateName"`
	CertificateBody       string `position:"Query" name:"CertificateBody"`
	CertificatePrivateKey string `position:"Query" name:"CertificatePrivateKey"`
}

func (req *SetDomainRequest) Invoke(client *sdk.Client) (resp *SetDomainResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "SetDomain", "apigateway", "")
	resp = &SetDomainResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetDomainResponse struct {
	responses.BaseResponse
	RequestId           string
	GroupId             string
	DomainName          string
	SubDomain           string
	DomainBindingStatus string
	DomainLegalStatus   string
	DomainRemark        string
}
