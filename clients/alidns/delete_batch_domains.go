package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteBatchDomainsRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Domains      string `position:"Query" name:"Domains"`
}

func (r DeleteBatchDomainsRequest) Invoke(client *sdk.Client) (response *DeleteBatchDomainsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteBatchDomainsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Alidns", "2015-01-09", "DeleteBatchDomains", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteBatchDomainsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteBatchDomainsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteBatchDomainsResponse struct {
	RequestId string
	TraceId   string
}
