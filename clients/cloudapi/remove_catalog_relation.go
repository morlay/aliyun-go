package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RemoveCatalogRelationRequest struct {
	CatalogId string `position:"Query" name:"CatalogId"`
	ApiId     string `position:"Query" name:"ApiId"`
}

func (r RemoveCatalogRelationRequest) Invoke(client *sdk.Client) (response *RemoveCatalogRelationResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RemoveCatalogRelationRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "RemoveCatalogRelation", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		RemoveCatalogRelationResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RemoveCatalogRelationResponse

	err = client.DoAction(&req, &resp)
	return
}

type RemoveCatalogRelationResponse struct {
	RequestId string
}
