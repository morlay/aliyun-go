package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteDomainCertificateRequest struct {
	requests.RpcRequest
	GroupId       string `position:"Query" name:"GroupId"`
	DomainName    string `position:"Query" name:"DomainName"`
	CertificateId string `position:"Query" name:"CertificateId"`
}

func (req *DeleteDomainCertificateRequest) Invoke(client *sdk.Client) (resp *DeleteDomainCertificateResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DeleteDomainCertificate", "apigateway", "")
	resp = &DeleteDomainCertificateResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteDomainCertificateResponse struct {
	responses.BaseResponse
	RequestId common.String
}
