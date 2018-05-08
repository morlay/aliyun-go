package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RemoveCatalogRelationRequest struct {
	requests.RpcRequest
	CatalogId string `position:"Query" name:"CatalogId"`
	ApiId     string `position:"Query" name:"ApiId"`
}

func (req *RemoveCatalogRelationRequest) Invoke(client *sdk.Client) (resp *RemoveCatalogRelationResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "RemoveCatalogRelation", "apigateway", "")
	resp = &RemoveCatalogRelationResponse{}
	err = client.DoAction(req, resp)
	return
}

type RemoveCatalogRelationResponse struct {
	responses.BaseResponse
	RequestId common.String
}
