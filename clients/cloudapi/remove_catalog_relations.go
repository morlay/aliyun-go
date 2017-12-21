package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RemoveCatalogRelationsRequest struct {
	CatalogId string `position:"Query" name:"CatalogId"`
}

func (r RemoveCatalogRelationsRequest) Invoke(client *sdk.Client) (response *RemoveCatalogRelationsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RemoveCatalogRelationsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "RemoveCatalogRelations", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		RemoveCatalogRelationsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RemoveCatalogRelationsResponse

	err = client.DoAction(&req, &resp)
	return
}

type RemoveCatalogRelationsResponse struct {
	RequestId string
}
