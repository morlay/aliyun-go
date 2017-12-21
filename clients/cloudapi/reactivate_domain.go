package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReactivateDomainRequest struct {
	GroupId    string `position:"Query" name:"GroupId"`
	DomainName string `position:"Query" name:"DomainName"`
}

func (r ReactivateDomainRequest) Invoke(client *sdk.Client) (response *ReactivateDomainResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ReactivateDomainRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "ReactivateDomain", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		ReactivateDomainResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ReactivateDomainResponse

	err = client.DoAction(&req, &resp)
	return
}

type ReactivateDomainResponse struct {
	RequestId string
}
