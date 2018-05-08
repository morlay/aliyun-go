package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateCatalogRequest struct {
	requests.RpcRequest
	CatalogName string `position:"Query" name:"CatalogName"`
	Description string `position:"Query" name:"Description"`
	ParentId    string `position:"Query" name:"ParentId"`
}

func (req *CreateCatalogRequest) Invoke(client *sdk.Client) (resp *CreateCatalogResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "CreateCatalog", "apigateway", "")
	resp = &CreateCatalogResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateCatalogResponse struct {
	responses.BaseResponse
	RequestId common.String
	CatalogId common.String
}
