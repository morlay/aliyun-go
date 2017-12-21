package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateCatalogRequest struct {
	CatalogName string `position:"Query" name:"CatalogName"`
	Description string `position:"Query" name:"Description"`
	ParentId    string `position:"Query" name:"ParentId"`
}

func (r CreateCatalogRequest) Invoke(client *sdk.Client) (response *CreateCatalogResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateCatalogRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "CreateCatalog", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		CreateCatalogResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateCatalogResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateCatalogResponse struct {
	RequestId string
	CatalogId string
}
