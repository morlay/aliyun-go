package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RemoveCatalogRelationsRequest struct {
	requests.RpcRequest
	CatalogId string `position:"Query" name:"CatalogId"`
}

func (req *RemoveCatalogRelationsRequest) Invoke(client *sdk.Client) (resp *RemoveCatalogRelationsResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "RemoveCatalogRelations", "apigateway", "")
	resp = &RemoveCatalogRelationsResponse{}
	err = client.DoAction(req, resp)
	return
}

type RemoveCatalogRelationsResponse struct {
	responses.BaseResponse
	RequestId string
}
