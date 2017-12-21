package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteCatalogRequest struct {
	requests.RpcRequest
	CatalogId string `position:"Query" name:"CatalogId"`
}

func (req *DeleteCatalogRequest) Invoke(client *sdk.Client) (resp *DeleteCatalogResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DeleteCatalog", "apigateway", "")
	resp = &DeleteCatalogResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteCatalogResponse struct {
	responses.BaseResponse
	RequestId string
}
