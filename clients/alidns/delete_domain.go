package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteDomainRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
}

func (r DeleteDomainRequest) Invoke(client *sdk.Client) (response *DeleteDomainResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteDomainRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Alidns", "2015-01-09", "DeleteDomain", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteDomainResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteDomainResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteDomainResponse struct {
	RequestId  string
	DomainName string
}
