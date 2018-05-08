package dcdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetDcdnDomainCertificateRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	CertType      string `position:"Query" name:"CertType"`
	SSLPub        string `position:"Query" name:"SSLPub"`
	CertName      string `position:"Query" name:"CertName"`
	SSLProtocol   string `position:"Query" name:"SSLProtocol"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Region        string `position:"Query" name:"Region"`
	SSLPri        string `position:"Query" name:"SSLPri"`
}

func (req *SetDcdnDomainCertificateRequest) Invoke(client *sdk.Client) (resp *SetDcdnDomainCertificateResponse, err error) {
	req.InitWithApiInfo("dcdn", "2018-01-15", "SetDcdnDomainCertificate", "dcdn", "")
	resp = &SetDcdnDomainCertificateResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetDcdnDomainCertificateResponse struct {
	responses.BaseResponse
	RequestId common.String
}
