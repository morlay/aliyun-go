package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReactivateDomainRequest struct {
	requests.RpcRequest
	GroupId    string `position:"Query" name:"GroupId"`
	DomainName string `position:"Query" name:"DomainName"`
}

func (req *ReactivateDomainRequest) Invoke(client *sdk.Client) (resp *ReactivateDomainResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "ReactivateDomain", "apigateway", "")
	resp = &ReactivateDomainResponse{}
	err = client.DoAction(req, resp)
	return
}

type ReactivateDomainResponse struct {
	responses.BaseResponse
	RequestId string
}
