package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteDomainCertificateRequest struct {
	GroupId       string `position:"Query" name:"GroupId"`
	DomainName    string `position:"Query" name:"DomainName"`
	CertificateId string `position:"Query" name:"CertificateId"`
}

func (r DeleteDomainCertificateRequest) Invoke(client *sdk.Client) (response *DeleteDomainCertificateResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteDomainCertificateRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DeleteDomainCertificate", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		DeleteDomainCertificateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteDomainCertificateResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteDomainCertificateResponse struct {
	RequestId string
}
