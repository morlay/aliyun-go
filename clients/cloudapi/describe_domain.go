package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainRequest struct {
	GroupId    string `position:"Query" name:"GroupId"`
	DomainName string `position:"Query" name:"DomainName"`
}

func (r DescribeDomainRequest) Invoke(client *sdk.Client) (response *DescribeDomainResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDomainRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeDomain", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDomainResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDomainResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDomainResponse struct {
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
