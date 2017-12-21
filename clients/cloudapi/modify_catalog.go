package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyCatalogRequest struct {
	CatalogId   string `position:"Query" name:"CatalogId"`
	CatalogName string `position:"Query" name:"CatalogName"`
	Description string `position:"Query" name:"Description"`
}

func (r ModifyCatalogRequest) Invoke(client *sdk.Client) (response *ModifyCatalogResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyCatalogRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "ModifyCatalog", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		ModifyCatalogResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyCatalogResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyCatalogResponse struct {
	RequestId string
}
