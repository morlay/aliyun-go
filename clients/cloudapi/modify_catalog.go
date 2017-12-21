package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyCatalogRequest struct {
	requests.RpcRequest
	CatalogId   string `position:"Query" name:"CatalogId"`
	CatalogName string `position:"Query" name:"CatalogName"`
	Description string `position:"Query" name:"Description"`
}

func (req *ModifyCatalogRequest) Invoke(client *sdk.Client) (resp *ModifyCatalogResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "ModifyCatalog", "apigateway", "")
	resp = &ModifyCatalogResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyCatalogResponse struct {
	responses.BaseResponse
	RequestId string
}
