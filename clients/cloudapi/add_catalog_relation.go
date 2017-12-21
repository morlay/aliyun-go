package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddCatalogRelationRequest struct {
	requests.RpcRequest
	CatalogId string `position:"Query" name:"CatalogId"`
	ApiId     string `position:"Query" name:"ApiId"`
}

func (req *AddCatalogRelationRequest) Invoke(client *sdk.Client) (resp *AddCatalogRelationResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "AddCatalogRelation", "apigateway", "")
	resp = &AddCatalogRelationResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddCatalogRelationResponse struct {
	responses.BaseResponse
	RequestId string
}
