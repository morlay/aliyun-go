package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteCatalogRequest struct {
	CatalogId string `position:"Query" name:"CatalogId"`
}

func (r DeleteCatalogRequest) Invoke(client *sdk.Client) (response *DeleteCatalogResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteCatalogRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DeleteCatalog", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		DeleteCatalogResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteCatalogResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteCatalogResponse struct {
	RequestId string
}
