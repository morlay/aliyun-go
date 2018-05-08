package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId           common.String
	GroupId             common.String
	DomainName          common.String
	SubDomain           common.String
	DomainBindingStatus common.String
	DomainLegalStatus   common.String
	DomainRemark        common.String
}
