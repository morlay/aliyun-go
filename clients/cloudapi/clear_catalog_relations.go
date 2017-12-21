package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ClearCatalogRelationsRequest struct {
	CatalogId string `position:"Query" name:"CatalogId"`
}

func (r ClearCatalogRelationsRequest) Invoke(client *sdk.Client) (response *ClearCatalogRelationsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ClearCatalogRelationsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "ClearCatalogRelations", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		ClearCatalogRelationsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ClearCatalogRelationsResponse

	err = client.DoAction(&req, &resp)
	return
}

type ClearCatalogRelationsResponse struct {
	RequestId string
}
