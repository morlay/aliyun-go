package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddCatalogRelationRequest struct {
	CatalogId string `position:"Query" name:"CatalogId"`
	ApiId     string `position:"Query" name:"ApiId"`
}

func (r AddCatalogRelationRequest) Invoke(client *sdk.Client) (response *AddCatalogRelationResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddCatalogRelationRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "AddCatalogRelation", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		AddCatalogRelationResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddCatalogRelationResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddCatalogRelationResponse struct {
	RequestId string
}
