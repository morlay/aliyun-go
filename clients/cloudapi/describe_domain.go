package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId             common.String
	GroupId               common.String
	DomainName            common.String
	SubDomain             common.String
	CertificateId         common.String
	CertificateName       common.String
	CertificateBody       common.String
	CertificatePrivateKey common.String
	DomainBindingStatus   common.String
	DomainCNAMEStatus     common.String
	DomainLegalStatus     common.String
	DomainRemark          common.String
}
