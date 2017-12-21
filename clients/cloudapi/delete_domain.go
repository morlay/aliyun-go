package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteDomainRequest struct {
	GroupId    string `position:"Query" name:"GroupId"`
	DomainName string `position:"Query" name:"DomainName"`
}

func (r DeleteDomainRequest) Invoke(client *sdk.Client) (response *DeleteDomainResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteDomainRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DeleteDomain", "apigateway", "")

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
	RequestId string
}
