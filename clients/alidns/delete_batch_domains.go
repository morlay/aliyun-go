package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteBatchDomainsRequest struct {
	requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Domains      string `position:"Query" name:"Domains"`
}

func (req *DeleteBatchDomainsRequest) Invoke(client *sdk.Client) (resp *DeleteBatchDomainsResponse, err error) {
	req.InitWithApiInfo("Alidns", "2015-01-09", "DeleteBatchDomains", "", "")
	resp = &DeleteBatchDomainsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteBatchDomainsResponse struct {
	responses.BaseResponse
	RequestId common.String
	TraceId   common.String
}
