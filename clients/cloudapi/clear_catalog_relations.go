package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ClearCatalogRelationsRequest struct {
	requests.RpcRequest
	CatalogId string `position:"Query" name:"CatalogId"`
}

func (req *ClearCatalogRelationsRequest) Invoke(client *sdk.Client) (resp *ClearCatalogRelationsResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "ClearCatalogRelations", "apigateway", "")
	resp = &ClearCatalogRelationsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ClearCatalogRelationsResponse struct {
	responses.BaseResponse
	RequestId common.String
}
