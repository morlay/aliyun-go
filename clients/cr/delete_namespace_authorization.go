package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteNamespaceAuthorizationRequest struct {
	requests.RoaRequest
	Namespace   string `position:"Path" name:"Namespace"`
	AuthorizeId int64  `position:"Path" name:"AuthorizeId"`
}

func (req *DeleteNamespaceAuthorizationRequest) Invoke(client *sdk.Client) (resp *DeleteNamespaceAuthorizationResponse, err error) {
	req.InitWithApiInfo("cr", "2016-06-07", "DeleteNamespaceAuthorization", "/namespace/[Namespace]/authorizations/[AuthorizeId]", "", "")
	req.Method = "DELETE"

	resp = &DeleteNamespaceAuthorizationResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteNamespaceAuthorizationResponse struct {
	responses.BaseResponse
}
