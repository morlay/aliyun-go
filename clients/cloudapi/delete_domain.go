package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteDomainRequest struct {
	requests.RpcRequest
	GroupId    string `position:"Query" name:"GroupId"`
	DomainName string `position:"Query" name:"DomainName"`
}

func (req *DeleteDomainRequest) Invoke(client *sdk.Client) (resp *DeleteDomainResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DeleteDomain", "apigateway", "")
	resp = &DeleteDomainResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteDomainResponse struct {
	responses.BaseResponse
	RequestId string
}
