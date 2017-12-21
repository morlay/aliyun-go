package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainRequest struct {
	requests.RpcRequest
	GroupId    string `position:"Query" name:"GroupId"`
	DomainName string `position:"Query" name:"DomainName"`
}

func (req *DescribeDomainRequest) Invoke(client *sdk.Client) (resp *DescribeDomainResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeDomain", "apigateway", "")
	resp = &DescribeDomainResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainResponse struct {
	responses.BaseResponse
	RequestId             string
	GroupId               string
	DomainName            string
	SubDomain             string
	CertificateId         string
	CertificateName       string
	CertificateBody       string
	CertificatePrivateKey string
	DomainBindingStatus   string
	DomainCNAMEStatus     string
	DomainLegalStatus     string
	DomainRemark          string
}
